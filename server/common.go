package server

import (
	"mails/common"
	"net/http"

	"github.com/labstack/echo/v4"
)

type okResponse struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func defaultGetRequestSkipper(c echo.Context) bool {
	return c.Request().Method == http.MethodGet
}

func defaultAPIRequestSkipper(c echo.Context) bool {
	path := c.Path()
	return common.HasPrefixes(path, "/api")
}
