package server

import (
	"log"
	"mails/store"
	"net/http"

	"github.com/labstack/echo/v4"
)

type sendEmailRequestBody struct {
	Template   int   `json:"template"`
	Sender     int   `json:"sender"`
	Recipients []int `json:"recipients"`
}

func (server *Server) sendEmail(c echo.Context) error {
	ctx := c.Request().Context()

	var body sendEmailRequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: "bad request",
		})
	}

	template, err := server.store.FindTemplate(ctx, &store.TemplateFind{
		ID: &body.Template,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: "template not found",
		})
	}

	sender, err := server.store.FindSender(ctx, &store.SenderFind{
		ID: &body.Sender,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: "sender not found",
		})
	}

	recipients := make([]*store.Recipient, 0)
	for _, recipientID := range body.Recipients {
		recipient, err := server.store.FindRecipient(ctx, &store.RecipientFind{
			ID: &recipientID,
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, &errorResponse{
				Message: "recipient not found",
			})
		}
		recipients = append(recipients, recipient)
	}

	log.Println(sender, recipients, template)

	return c.JSON(http.StatusOK, &okResponse{})
}
