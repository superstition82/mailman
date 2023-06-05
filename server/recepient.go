package server

import (
	"fmt"
	"net/http"
	"pocketmail/store"
	"strconv"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/labstack/echo/v4"
)

type createRecepientRequestBody struct {
	Email string `json:"email"`
}

func (server *Server) createRecepient(c echo.Context) error {
	ctx := c.Request().Context()

	var body createRecepientRequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: "bad request",
		})
	}

	createRecepientParams := store.CreateRecepientParams{
		Email:     body.Email,
		Reachable: "unknown",
	}
	recepient, err := server.store.CreateRecepient(ctx, createRecepientParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &okResponse{
		Data: recepient,
	})
}

type listRecepientOKResponse struct {
	Data []store.Recepient `json:"data"`
}

func (server *Server) listRecepients(c echo.Context) error {
	ctx := c.Request().Context()

	var listRecepientParams store.ListRecepientsParams
	if limit, err := strconv.Atoi(c.QueryParam("limit")); err == nil {
		listRecepientParams.Limit = limit
	}
	if offset, err := strconv.Atoi(c.QueryParam("offset")); err == nil {
		listRecepientParams.Offset = offset
	}

	result, err := server.store.ListRecepients(ctx, listRecepientParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &listRecepientOKResponse{
		Data: result,
	})
}

func (server *Server) getRecepient(c echo.Context) error {
	ctx := c.Request().Context()

	recepientID, err := strconv.Atoi(c.Param("recepientId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("recepientId")))
	}
	recepient, err := server.store.GetRecepient(ctx, recepientID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &okResponse{
		Data: recepient,
	})
}

func (server *Server) deleteRecepient(c echo.Context) error {
	ctx := c.Request().Context()

	recepientID, err := strconv.Atoi(c.Param("recepientId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("recepientId")))
	}
	err = server.store.DeleteRecepient(ctx, recepientID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, true)
}

var (
	verifier = emailverifier.
		NewVerifier().
		EnableSMTPCheck()
)

func (server *Server) validateRecepient(c echo.Context) error {
	ctx := c.Request().Context()
	recepientID, err := strconv.Atoi(c.Param("recepientId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("recepientId")))
	}

	recepient, err := server.store.GetRecepient(ctx, recepientID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: "bad request",
		})
	}

	var updatedRecepient store.Recepient
	if recepient.Reachable == "unknown" {
		verifiedResult, err := verifier.Verify(recepient.Email)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &errorResponse{
				Message: "bad request",
			})
		}

		updatedRecepient, err = server.store.UpdateRecepient(ctx, store.UpdateRecepientParams{
			ID:        recepient.ID,
			Email:     recepient.Email,
			Reachable: verifiedResult.Reachable,
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, &errorResponse{
				Message: err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, &okResponse{
		Data: updatedRecepient,
	})
}
