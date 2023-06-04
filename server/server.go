package server

import (
	"context"
	"database/sql"
	"fmt"
	"pocketmail/server/config"
	"pocketmail/store"
	"pocketmail/store/db"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
)

type Server struct {
	e  *echo.Echo
	db *sql.DB

	ID     string
	config *config.Config
	store  *store.Store
}

func NewServer(ctx context.Context, config *config.Config) (*Server, error) {
	e := echo.New()
	e.Debug = true
	e.HideBanner = true
	e.HidePort = true

	db := db.NewDB(config)
	if err := db.Open(ctx); err != nil {
		return nil, errors.Wrap(err, "cannot open db")
	}

	s := &Server{
		e:      e,
		db:     db.DBInstance,
		config: config,
	}
	storeInstance := store.New(db.DBInstance, config)
	s.store = storeInstance

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339}",` +
			`"method":"${method}","uri":"${uri}",` +
			`"status":${status},"error":"${error}"}` + "\n",
	}))
	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		ErrorMessage: "Request timeout",
		Timeout:      30 * time.Second,
	}))

	// frontend embedding with FS
	embedFrontend(e)

	apiGroup := e.Group("/api")

	// recepient group
	recepientGroup := apiGroup.Group("/recepient")
	recepientGroup.POST("", s.createRecepient)
	recepientGroup.GET("", s.listRecepients)
	recepientGroup.GET("/:recepientId", s.getRecepient)
	recepientGroup.DELETE("/:recepientId", s.deleteRecepient)
	recepientGroup.POST("/:recepientId/verification", s.validateRecepient)

	// sender group
	senderGroup := apiGroup.Group("/sender")
	senderGroup.POST("", s.createSender)
	senderGroup.GET("", s.listSenders)
	senderGroup.GET("/:senderId", s.getSender)
	senderGroup.DELETE("/:senderId", s.deleteSender)

	return s, nil
}

func (s *Server) Start(ctx context.Context) error {
	return s.e.Start(fmt.Sprintf(":%d", s.config.Port))
}

func (s *Server) Shutdown(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Shutdown echo server
	if err := s.e.Shutdown(ctx); err != nil {
		fmt.Printf("failed to shutdown server, error: %v\n", err)
	}

	// Close database connection
	if err := s.db.Close(); err != nil {
		fmt.Printf("failed to close database, error: %v\n", err)
	}

	fmt.Printf("memos stopped properly\n")
}
