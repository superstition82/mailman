package server

import (
	"github.com/labstack/echo/v4"
)

func (server *Server) registerClientRoutes(g *echo.Group) {
	g.POST("/client", func(c echo.Context) error {
		return nil
	})
}
