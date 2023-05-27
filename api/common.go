package api

import (
	"net/http"
	"probemail/util"

	"github.com/labstack/echo/v4"
)

type response struct {
	Data any `json:"data"`
}

func composeResponse(data any) response {
	return response{
		Data: data,
	}
}

func defaultGetRequestSkipper(c echo.Context) bool {
	return c.Request().Method == http.MethodGet
}

func defaultAPIRequestSkipper(c echo.Context) bool {
	path := c.Path()
	return util.HasPrefixes(path, "/api")
}
