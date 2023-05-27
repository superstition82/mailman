package core

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (server *Server) registerRootRoutes(g *echo.Group) {
	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! 🖖")
	})
}
