package server

import (
	"bufio"
	"fmt"
	"mails/store"
	"net/http"
	"os"
	"strconv"
	"strings"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/labstack/echo/v4"
)

type createRecipientRequestBody struct {
	Email string `json:"email"`
}

func (server *Server) createRecipient(c echo.Context) error {
	ctx := c.Request().Context()

	var body createRecipientRequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: "bad request",
		})
	}

	recipientCreate := store.RecipientCreate{
		Email:     body.Email,
		Reachable: "unknown",
	}
	recipient, err := server.store.CreateRecipient(ctx, &recipientCreate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &okResponse{
		Data: recipient,
	})
}

func (server *Server) findRecipientList(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := server.store.FindRecipientList(ctx, &store.RecipientFind{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &okResponse{
		Data: result,
	})
}

func (server *Server) findRecipient(c echo.Context) error {
	ctx := c.Request().Context()

	recipientID, err := strconv.Atoi(c.Param("recipientId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("recipientId")))
	}
	recipient, err := server.store.FindRecipient(ctx, &store.RecipientFind{
		ID: &recipientID,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &okResponse{
		Data: recipient,
	})
}

func (server *Server) deleteRecipient(c echo.Context) error {
	ctx := c.Request().Context()

	recipientID, err := strconv.Atoi(c.Param("recipientId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("recipientId")))
	}
	err = server.store.DeleteRecipient(ctx, &store.RecipientDelete{
		ID: recipientID,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, true)
}

type deleteBulkRecipientRequestBody struct {
	Recipients []int `json:"recipients"`
}

func (server *Server) deleteBulkRecipient(c echo.Context) error {
	ctx := c.Request().Context()

	var body deleteBulkRecipientRequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: "bad request",
		})
	}

	recipientBulkDelete := &store.RecipientBulkDelete{
		IDs: body.Recipients,
	}
	if err := server.store.DeleteBulkRecipient(ctx, recipientBulkDelete); err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: "bad request",
		})
	}

	return c.JSON(http.StatusOK, true)
}

var (
	verifier = emailverifier.
		NewVerifier().
		EnableSMTPCheck().
		DisableCatchAllCheck()
)

func (server *Server) validateRecipient(c echo.Context) error {
	ctx := c.Request().Context()

	recipientID, err := strconv.Atoi(c.Param("recipientId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("recipientId")))
	}
	recipient, err := server.store.FindRecipient(ctx, &store.RecipientFind{
		ID: &recipientID,
	})
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: "recipient not found",
		})
	}

	var updatedRecipient *store.Recipient
	if recipient.Reachable == "unknown" {
		verifiedResult, err := verifier.Verify(recipient.Email)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &errorResponse{
				Message: "failed to verify",
			})
		}
		updatedRecipient, err = server.store.PatchRecipient(ctx, &store.RecipientPatch{
			ID:        recipient.ID,
			Email:     &recipient.Email,
			Reachable: &verifiedResult.Reachable,
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, &errorResponse{
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, &okResponse{
			Data: updatedRecipient,
		})
	}

	return c.JSON(http.StatusOK, &okResponse{
		Data: recipient,
	})
}

func (server *Server) importRecipientsByFile(c echo.Context) error {
	ctx := c.Request().Context()

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: "bad request",
		})
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &errorResponse{
			Message: "unable to open file",
		})
	}
	defer src.Close()

	// Read the file content line by line
	scanner := bufio.NewScanner(src)
	var recipients []*store.Recipient
	for scanner.Scan() {
		recipient, err := server.store.CreateRecipient(ctx, &store.RecipientCreate{
			Email:     scanner.Text(),
			Reachable: "unknown",
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &errorResponse{
				Message: "unable to read file",
			})
		}
		recipients = append(recipients, recipient)
	}

	// Check for any scanner errors
	if err := scanner.Err(); err != nil {
		return err
	}

	// Send the file content as the response
	return c.JSON(http.StatusOK, &okResponse{
		Data: recipients,
	})
}

func (server *Server) exportRecipientsToFile(c echo.Context) error {
	ctx := c.Request().Context()

	recipients, err := server.store.FindRecipientList(ctx, &store.RecipientFind{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	file, err := os.CreateTemp("", "recipients.txt")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to generate a file")
	}
	defer os.Remove(file.Name()) // Remove the temporary file when done

	var emails []string
	for _, recipient := range recipients {
		emails = append(emails, recipient.Email)
	}

	content := strings.Join(emails, "\n")
	_, err = file.Write([]byte(content))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	c.Response().Header().Set("Content-Disposition", "attachment; filename=customers.txt")
	c.Response().Header().Set("Content-Type", "text/plain")

	return c.File(file.Name())
}
