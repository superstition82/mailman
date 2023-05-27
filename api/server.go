package api

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"probemail/db"
	"probemail/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	e  *echo.Echo
	db *gorm.DB

	Config *util.Config
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(ctx context.Context, config *util.Config) (*Server, error) {
	e := echo.New()
	e.Debug = true
	e.HideBanner = true

	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	db := db.NewDB(config)
	if err := db.Open(ctx); err != nil {
		return nil, errors.Wrap(err, "cannot open db")
	}

	s := &Server{
		e:      e,
		db:     db.DBInstance,
		Config: config,
	}

	// default ui
	embedFrontend(e)

	// default api routes
	apiGroup := e.Group("/api")
	s.registerSmtpRoutes(apiGroup)

	return s, nil
}

func (server *Server) Start(ctx context.Context) error {
	return server.e.Start(fmt.Sprintf(":%d", server.Config.Port))
}

func (server *Server) Shutdown(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Shutdown echo server
	if err := server.e.Shutdown(ctx); err != nil {
		fmt.Printf("Failed to shutdown server, error: %v\n", err)
	}

	// Close database connection
	// if err := server.db.Close(); err != nil {
	// 	fmt.Printf("failed to close database, error: %v\n", err)
	// }

	fmt.Printf("probemail stopped properly 👏\n")
}
