package server

import (
	"fmt"
	"mails/store"
	"net/http"
	"strconv"

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

	createTemplateParams := store.TemplateCreate{
		Subject: body.Subject,
		Body:    body.Body,
	}
	template, err := server.store.CreateTemplate(ctx, &createTemplateParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &okResponse{
		Data: template,
	})
}

func (server *Server) findTemplateList(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := server.store.FindTemplateList(ctx, &store.TemplateFind{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &okResponse{
		Data: result,
	})
}

func (server *Server) getTemplate(c echo.Context) error {
	ctx := c.Request().Context()

	templateId, err := strconv.Atoi(c.Param("templateId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("templateId")))
	}
	template, err := server.store.FindTemplate(ctx, &store.TemplateFind{
		ID: &templateId,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &okResponse{
		Data: template,
	})
}

type updateTemplateRequestBody struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func (server *Server) updateTemplate(c echo.Context) error {
	ctx := c.Request().Context()

	var body updateTemplateRequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: "bad request",
		})
	}
	templateId, err := strconv.Atoi(c.Param("templateId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("templateId")))
	}
	template, err := server.store.FindTemplate(ctx, &store.TemplateFind{
		ID: &templateId,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	//TODO: better way to solve this?
	if body.Subject == "" {
		body.Subject = template.Subject
	}

	if body.Body == "" {
		body.Body = template.Body
	}

	templatePatch := store.TemplatePatch{
		ID:      templateId,
		Subject: &body.Subject,
		Body:    &body.Body,
	}
	template, err = server.store.PatchTemplate(ctx, &templatePatch)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &okResponse{
		Data: template,
	})
}

func (server *Server) deleteTemplate(c echo.Context) error {
	ctx := c.Request().Context()

	templateId, err := strconv.Atoi(c.Param("templateId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("templateId")))
	}
	err = server.store.DeleteTemplate(ctx, &store.TemplateDelete{
		ID: templateId,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, true)
}

type deleteBulkTemplateRequestBody struct {
	Templates []int `json:"templates"`
}

func (server *Server) deleteBulkTemplate(c echo.Context) error {
	ctx := c.Request().Context()

	var body deleteBulkTemplateRequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: "bad request",
		})
	}

	if err := server.store.DeleteBulkTemplate(ctx, body.Templates); err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: "bad request",
		})
	}

	return c.JSON(http.StatusOK, true)
}
