package server

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/smtp"
	"pocketmail/store"
	"strconv"

	"github.com/labstack/echo/v4"
)

type createSenderRequestBody struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createSenderOKResponse struct {
	SenderID int `json:"sender_id"`
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

	return c.JSON(http.StatusOK, &createSenderOKResponse{
		SenderID: sender.ID,
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

type listSenderOKResponse struct {
	Data []store.Sender `json:"data"`
}

func (server *Server) listSenders(c echo.Context) error {
	ctx := c.Request().Context()

	var listSenderParams store.ListSendersParams
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err == nil {
		listSenderParams.Limit = limit
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err == nil {
		listSenderParams.Offset = offset
	}

	result, err := server.store.ListSenders(ctx, listSenderParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &listSenderOKResponse{
		Data: result,
	})
}

func (server *Server) getSender(c echo.Context) error {
	ctx := c.Request().Context()

	senderID, err := strconv.Atoi(c.Param("senderId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("senderId")))
	}
	sender, err := server.store.GetSender(ctx, senderID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, sender)
}

func (server *Server) deleteSender(c echo.Context) error {
	ctx := c.Request().Context()

	senderID, err := strconv.Atoi(c.Param("senderId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("senderId")))
	}
	err = server.store.DeleteSender(ctx, senderID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, true)
}
