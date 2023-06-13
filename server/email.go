package server

import (
	"fmt"
	"io/ioutil"
	"mails/store"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slices"
	"gopkg.in/gomail.v2"
)

type sendEmailRequestBody struct {
	Template   int   `json:"template"`
	Sender     int   `json:"sender"`
	Recipients []int `json:"recipients"`
	BCC        []int `json:"bcc"`
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

	to := make([]string, 0)
	for _, recipientID := range body.Recipients {
		recipient, err := server.store.FindRecipient(ctx, &store.RecipientFind{
			ID: &recipientID,
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, &errorResponse{
				Message: "recipient not found",
			})
		}
		to = append(to, recipient.Email)
	}

	bcc := make([]string, 0)
	for _, recipientID := range body.BCC {
		recipient, err := server.store.FindRecipient(ctx, &store.RecipientFind{
			ID: &recipientID,
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, &errorResponse{
				Message: "recipient not found",
			})
		}
		bcc = append(bcc, recipient.Email)
	}

	resources := make([]string, 0)
	for _, resourceID := range template.ResourceIDList {
		resource, err := server.store.FindResource(ctx, &store.ResourceFind{
			ID:      &resourceID,
			GetBlob: true,
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, &errorResponse{
				Message: "resource not found",
			})
		}
		err = ioutil.WriteFile(resource.Filename, resource.Blob, 0644)
		if err != nil {
			fmt.Println("Error saving image file:", err)
			return c.String(http.StatusInternalServerError, "Failed to generate a file")
		}
		defer os.Remove(resource.Filename)
		resources = append(resources, resource.Filename)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", sender.Email)
	m.SetHeader("To", to...)
	m.SetHeader("Bcc", bcc...)
	m.SetHeader("Subject", template.Subject)
	modified, embeds, err := composeEmailHTML(template.Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}
	m.SetBody("text/html", modified)
	for _, embed := range embeds {
		if slices.Contains(resources, embed) {
			m.Embed(embed)
		}
	}

	d := gomail.NewDialer(sender.Host, sender.Port, sender.Email, sender.Password)
	if err := d.DialAndSend(m); err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: err.Error(),
		})
	}

	return c.String(
		http.StatusOK,
		fmt.Sprintf("[%s → %s (%s)] %s", sender.Email, strings.Join(to, ", "), strings.Join(bcc, ", "), template.Subject))
}

func composeEmailHTML(html string) (modified string, embeds []string, err error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println("Failed to parse HTML:", err)
		return
	}

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		if strings.HasPrefix(src, "/o/r/") {
			fileName := src[strings.LastIndex(src, "/")+1:]
			embeds = append(embeds, fileName)
			cidSrc := fmt.Sprintf("cid:%s", fileName)
			s.SetAttr("src", cidSrc)
		}
	})

	modified, err = doc.Html()
	if err != nil {
		fmt.Println("Failed to convert document back to HTML:", err)
		return
	}

	return
}
