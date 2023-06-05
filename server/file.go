package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (server *Server) upload(c echo.Context) error {
	return c.JSON(http.StatusOK, "upload")
}
