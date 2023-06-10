package server

import (
	"crypto/tls"
	"fmt"
	"mails/store"
	"net/http"
	"net/smtp"
	"strconv"

	"github.com/labstack/echo/v4"
)

type createSenderRequestBody struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (server *Server) createSender(c echo.Context) error {
	ctx := c.Request().Context()

	var body createSenderRequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: "bad request",
		})
	}

	// SMTP Login Test
	err := smtpLoginTest(body.Host, body.Port, body.Email, body.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: fmt.Sprintf("failed to SMTP login: %s", err.Error()),
		})
	}

	createSenderParams := store.CreateSenderParams{
		Host:     body.Host,
		Port:     body.Port,
		Email:    body.Email,
		Password: body.Password,
	}
	sender, err := server.store.CreateSender(ctx, createSenderParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &okResponse{
		Data: sender,
	})
}

func smtpLoginTest(host string, port int, email string, password string) error {
	addr := fmt.Sprintf("%s:%d", host, port)
	auth := smtp.PlainAuth("", email, password, host)
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()

	c.StartTLS(tlsconfig)
	if err = c.Auth(auth); err != nil {
		return err
	}

	return nil
}

func (server *Server) listAllSenders(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := server.store.FindSenderList(ctx, &store.SenderFind{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &okResponse{
		Data: result,
	})
}

func (server *Server) getSender(c echo.Context) error {
	ctx := c.Request().Context()

	senderID, err := strconv.Atoi(c.Param("senderId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("senderId")))
	}
	sender, err := server.store.FindSender(ctx, &store.SenderFind{
		ID: &senderID,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &okResponse{
		Data: sender,
	})
}

func (server *Server) deleteSender(c echo.Context) error {
	ctx := c.Request().Context()

	senderID, err := strconv.Atoi(c.Param("senderId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("senderId")))
	}
	err = server.store.DeleteSender(ctx, &store.SenderDelete{
		ID: senderID,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, true)
}

type deleteBulkSenderRequestBody struct {
	Senders []int `json:"senders"`
}

func (server *Server) deleteBulkSender(c echo.Context) error {
	ctx := c.Request().Context()

	var body deleteBulkSenderRequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: "bad request",
		})
	}

	bulkDelete := &store.SenderBulkDelete{
		IDs: body.Senders,
	}
	if err := server.store.DeleteBulkSender(ctx, bulkDelete); err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: "bad request",
		})
	}

	return c.JSON(http.StatusOK, true)
}
