package server

import (
	"context"
	"fmt"
	"mails/store"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type createTemplateRequestBody struct {
	// Domain specific fields
	Subject string `json:"subject"`
	Body    string `json:"body"`

	// Related fields
	ResourceIDList []int `json:"resourceIdList"`
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

	for _, resourceID := range body.ResourceIDList {
		if _, err := server.store.UpsertTemplateResource(ctx, &store.TemplateResourceUpsert{
			TemplateID: template.ID,
			ResourceID: resourceID,
		}); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to upsert template resource").SetInternal(err)
		}
	}

	template, err = server.store.FindTemplate(ctx, &store.TemplateFind{
		ID: &template.ID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to compose template").SetInternal(err)
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

func (server *Server) findTemplate(c echo.Context) error {
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

type patchTemplateRequest struct {
	// Domain specific fields
	Subject        *string `json:"subject"`
	Body           *string `json:"body"`
	ResourceIDList []int   `json:"resource_id_list"`
}

func (server *Server) patchTemplate(c echo.Context) error {
	ctx := c.Request().Context()

	templateID, err := strconv.Atoi(c.Param("templateId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("templateId")))
	}

	template, err := server.store.FindTemplate(ctx, &store.TemplateFind{
		ID: &templateID,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	currentTs := time.Now().Unix()
	var patchTemplateRequest patchTemplateRequest
	if err := c.Bind(&patchTemplateRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Malformatted patch template request").SetInternal(err)
	}
	templatePatch := &store.TemplatePatch{
		ID:        template.ID,
		Subject:   patchTemplateRequest.Subject,
		Body:      patchTemplateRequest.Body,
		UpdatedTs: &currentTs,
	}
	updatedTemplate, err := server.store.PatchTemplate(ctx, templatePatch)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to patch template").SetInternal(err)
	}

	if patchTemplateRequest.ResourceIDList != nil {
		addedResourceIDList, removedResourceIDList := getIDListDiff(updatedTemplate.ResourceIDList, patchTemplateRequest.ResourceIDList)
		for _, resourceID := range addedResourceIDList {
			if _, err := server.store.UpsertTemplateResource(ctx, &store.TemplateResourceUpsert{
				TemplateID: updatedTemplate.ID,
				ResourceID: resourceID,
			}); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to upsert template resource").SetInternal(err)
			}
		}
		for _, resourceID := range removedResourceIDList {
			if err := server.store.DeleteTemplateResource(ctx, &store.TemplateResourceDelete{
				TemplateID: &updatedTemplate.ID,
				ResourceID: &resourceID,
			}); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete memo resource").SetInternal(err)
			}
		}
	}
	template, err = server.store.FindTemplate(ctx, &store.TemplateFind{
		ID: &templateID,
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

func getIDListDiff(oldList, newList []int) (addedList, removedList []int) {
	oldMap := map[int]bool{}
	for _, id := range oldList {
		oldMap[id] = true
	}
	newMap := map[int]bool{}
	for _, id := range newList {
		newMap[id] = true
	}
	for id := range oldMap {
		if !newMap[id] {
			removedList = append(removedList, id)
		}
	}
	for id := range newMap {
		if !oldMap[id] {
			addedList = append(addedList, id)
		}
	}
	return addedList, removedList
}

func (server *Server) deleteTemplate(c echo.Context) error {
	ctx := c.Request().Context()

	templateID, err := strconv.Atoi(c.Param("templateId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("templateId")))
	}

	err = server.deleteTemplateImpl(ctx, templateID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
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

	for _, id := range body.Templates {
		if err := server.deleteTemplateImpl(ctx, id); err != nil {
			return c.JSON(http.StatusBadRequest, "bad request")
		}
	}

	return c.JSON(http.StatusOK, true)
}

func (server *Server) deleteTemplateImpl(ctx context.Context, id int) error {
	template, err := server.store.FindTemplate(ctx, &store.TemplateFind{
		ID: &id,
	})
	if err != nil {
		return err
	}

	err = server.store.DeleteTemplate(ctx, &store.TemplateDelete{
		ID: id,
	})
	if err != nil {
		return err
	}

	// to be honestly, it should commit in one transaction..
	for _, resourceID := range template.ResourceIDList {
		err = server.store.DeleteResource(ctx, &store.ResourceDelete{
			ID: resourceID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
