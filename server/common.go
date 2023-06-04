package server

import (
	"net/http"
	"pocketmail/common"

	"github.com/labstack/echo/v4"
)

func defaultGetRequestSkipper(c echo.Context) bool {
	return c.Request().Method == http.MethodGet
}

func defaultAPIRequestSkipper(c echo.Context) bool {
	path := c.Path()
	return common.HasPrefixes(path, "/api")
}
