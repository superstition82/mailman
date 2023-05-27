package core

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"

	_db "probemail/db"
	"probemail/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	e  *echo.Echo
	db *sql.DB

	Config *util.Config
	Store  *_db.Store
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(ctx context.Context, config *util.Config) (*Server, error) {
	e := echo.New()
	e.Debug = true
	e.HideBanner = true
	e.HidePort = true

	db := _db.NewDB(config)
	if err := db.Open(ctx); err != nil {
		return nil, errors.Wrap(err, "cannot open db")
	}

	s := &Server{
		e:      e,
		db:     db.DBInstance,
		Config: config,
	}
	storeInstance := _db.NewStore(db.DBInstance, config)
	s.Store = storeInstance

	e.Use(middleware.Gzip())

	e.Use(middleware.CORS())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339}",` +
			`"method":"${method}","uri":"${uri}",` +
			`"status":${status},"error":"${error}"}` + "\n",
	}))

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		ErrorMessage: "Request timeout",
		Timeout:      30 * time.Second,
	}))

	rootGroup := e.Group("")
	s.registerRootRoutes(rootGroup)

	emailVerifyGroup := e.Group("/verify")
	s.registerEmailVerifyRoutes(emailVerifyGroup)

	return s, nil
}

func (s *Server) Start(ctx context.Context) error {
	return s.e.Start(fmt.Sprintf(":%d", s.Config.Port))
}

func (s *Server) Shutdown(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Shutdown echo server
	if err := s.e.Shutdown(ctx); err != nil {
		fmt.Printf("Failed to shutdown server, error: %v\n", err)
	}

	// Close database connection
	if err := s.db.Close(); err != nil {
		fmt.Printf("failed to close database, error: %v\n", err)
	}

	fmt.Printf("probemail stopped properly 👏\n")
}
