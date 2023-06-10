package server

import (
	"fmt"
	"io"
	"mails/store"
	"mime"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

const (
	// The upload memory buffer is 32 MiB.
	// It should be kept low, so RAM usage doesn't get out of control.
	// This is unrelated to maximum upload size limit, which is now set through system setting.
	maxUploadBufferSizeBytes = 32 << 20
	MebiByte                 = 1024 * 1024
)

type createResourceRequestBody struct {
	// Domain specific fields
	Filename        string `json:"filename"`
	Blob            []byte `json:"-"`
	InternalPath    string `json:"internalPath"`
	ExternalLink    string `json:"externalLink"`
	Type            string `json:"type"`
	Size            int64  `json:"-"`
	DownloadToLocal bool   `json:"downloadToLocal"`
}

func (server *Server) createResource(c echo.Context) error {
	ctx := c.Request().Context()

	var body createResourceRequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Message: "bad request",
		})
	}

	createResourceParams := store.CreateResourceParams{
		Filename:     body.Filename,
		ExternalLink: body.ExternalLink,
	}
	if createResourceParams.ExternalLink != "" {
		// Only allow those external links scheme with http/https
		linkURL, err := url.Parse(createResourceParams.ExternalLink)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid external link").SetInternal(err)
		}
		if linkURL.Scheme != "http" && linkURL.Scheme != "https" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid external link scheme")
		}

		if body.DownloadToLocal {
			resp, err := http.Get(linkURL.String())
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Failed to request "+body.ExternalLink)
			}
			defer resp.Body.Close()

			blob, err := io.ReadAll(resp.Body)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Failed to read "+body.ExternalLink)
			}
			createResourceParams.Blob = blob

			mediaType, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Failed to read mime from "+body.ExternalLink)
			}
			createResourceParams.Type = mediaType

			filename := path.Base(linkURL.Path)
			if path.Ext(filename) == "" {
				extensions, _ := mime.ExtensionsByType(mediaType)
				if len(extensions) > 0 {
					filename += extensions[0]
				}
			}
			createResourceParams.Filename = filename
			createResourceParams.ExternalLink = ""
		}
	}

	resource, err := server.store.CreateResource(ctx, createResourceParams)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create resource").SetInternal(err)
	}

	return c.JSON(http.StatusOK, &okResponse{
		Data: resource,
	})
}

func (server *Server) createResourceBlob(c echo.Context) error {
	ctx := c.Request().Context()

	maxUploadSetting := "32"
	var settingMaxUploadSizeBytes int
	if settingMaxUploadSizeMiB, err := strconv.Atoi(maxUploadSetting); err == nil {
		settingMaxUploadSizeBytes = settingMaxUploadSizeMiB * MebiByte
	} else {
		log.Warn("Failed to parse max upload size", zap.Error(err))
		settingMaxUploadSizeBytes = 0
	}

	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get uploading file").SetInternal(err)
	}
	if file == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Upload file not found").SetInternal(err)
	}

	if file.Size > int64(settingMaxUploadSizeBytes) {
		message := fmt.Sprintf("File size exceeds allowed limit of %d MiB", settingMaxUploadSizeBytes/MebiByte)
		return echo.NewHTTPError(http.StatusBadRequest, message).SetInternal(err)
	}
	if err := c.Request().ParseMultipartForm(maxUploadBufferSizeBytes); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse upload data").SetInternal(err)
	}

	filetype := file.Header.Get("Content-Type")
	size := file.Size
	sourceFile, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to open file").SetInternal(err)
	}
	defer sourceFile.Close()

	fileBytes, err := io.ReadAll(sourceFile)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to read file").SetInternal(err)
	}

	// [*] save on db storage as blob
	createResourceParams := store.CreateResourceParams{
		Filename: file.Filename,
		Type:     filetype,
		Size:     size,
		Blob:     fileBytes,
	}

	resource, err := server.store.CreateResource(ctx, createResourceParams)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create resource").SetInternal(err)
	}

	return c.JSON(http.StatusOK, &okResponse{
		Data: resource,
	})
}
