package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"probemail/db/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CreateSmtpRequest struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Server) registerSmtpRoutes(g *echo.Group) {
	g.POST("/smtp", func(c echo.Context) error {
		createSmtpRequest := &CreateSmtpRequest{}
		if err := json.NewDecoder(c.Request().Body).Decode(createSmtpRequest); err != nil {
			return echo.
				NewHTTPError(http.StatusBadRequest, "Malformatted request").
				SetInternal(err)
		}

		smtp := models.SMTP{
			Host:     createSmtpRequest.Host,
			Port:     createSmtpRequest.Port,
			Email:    createSmtpRequest.Email,
			Password: createSmtpRequest.Password,
		}
		s.db.Create(smtp)

		return c.JSON(http.StatusOK, smtp)
	})

	g.GET("/smtp/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.NoContent(http.StatusPreconditionFailed)
		}

		var smtp models.SMTP
		if err := s.db.Where("id = ?", id).First(&smtp).Error; err != nil {
			return c.NoContent(http.StatusNotFound)
		}

		fmt.Println(smtp)

		return c.JSON(http.StatusOK, smtp)
	})

	g.DELETE("/smtp/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.NoContent(http.StatusPreconditionFailed)
		}

		fmt.Println(id)

		var smtp models.SMTP
		if err := s.db.Where("id = ?", id).First(&smtp).Error; err != nil {
			return c.NoContent(http.StatusNotFound)
		}

		s.db.Delete(&smtp)

		return c.NoContent(http.StatusOK)
	})
}
