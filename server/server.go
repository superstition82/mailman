package server

import (
	"context"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"

	"probemail/db"
	"probemail/server/config"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	e  *echo.Echo
	db *gorm.DB

	Config *config.Config
}

// New creates a new HTTP server and set up routing.
func New(ctx context.Context, config *config.Config) (*Server, error) {
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

	// default routes
	apiGroup := e.Group("/api")
	s.registerClientRoutes(apiGroup)
	s.registerEmailVerifyRoutes(apiGroup)

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
