package server

import (
	"mails/store"
	"net/http"

	"github.com/labstack/echo/v4"
)

type createTemplateRequestBody struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func (server *Server) createTemplate(c echo.Context) error {
	ctx := c.Request().Context()

	var body createTemplateRequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: "bad request",
		})
	}

	createTemplateParams := store.CreateTemplateParams{
		Subject: body.Subject,
		Body:    body.Body,
	}
	template, err := server.store.CreateTemplate(ctx, createTemplateParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &okResponse{
		Data: template,
	})
}

func (server *Server) listAllTemplates(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := server.store.ListAllTemplates(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &okResponse{
		Data: result,
	})
}
