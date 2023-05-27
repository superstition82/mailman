package server

import (
	"fmt"
	"net/http"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/labstack/echo/v4"
)

var (
	verifier = emailverifier.
		NewVerifier().
		EnableSMTPCheck().
		DisableCatchAllCheck()
)

func (server *Server) registerEmailVerifyRoutes(g *echo.Group) {
	g.POST("/verify/", func(c echo.Context) error {
		email := c.QueryParam("email")
		if email == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "email is required")
		}

		ret, err := verifier.Verify(email)
		if err != nil {
			fmt.Println("verify email address failed, error is: ", err)
		}
		if !ret.Syntax.Valid {
			fmt.Println("email address syntax is invalid")
		}

		return c.JSON(http.StatusOK, composeResponse(ret))
	})
}
