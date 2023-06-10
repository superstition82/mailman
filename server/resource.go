package server

import (
	"bytes"
	"fmt"
	"io"
	"mails/store"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

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

func (server *Server) downloadResource(c echo.Context) error {
	ctx := c.Request().Context()
	resourceID, err := strconv.Atoi(c.Param("resourceId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("resourceId"))).SetInternal(err)
	}

	resource, err := server.store.FindResource(ctx, &store.ResourceFind{
		ID:      &resourceID,
		GetBlob: true,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to find resource by ID: %v", resourceID)).SetInternal(err)
	}

	blob := resource.Blob
	if resource.InternalPath != "" {
		resourcePath := resource.InternalPath
		src, err := os.Open(resourcePath)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to open the local resource: %s", resourcePath)).SetInternal(err)
		}
		defer src.Close()
		blob, err = io.ReadAll(src)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to read the local resource: %s", resourcePath)).SetInternal(err)
		}
	}

	c.Response().Writer.Header().Set(echo.HeaderCacheControl, "max-age=31536000, immutable")
	c.Response().Writer.Header().Set(echo.HeaderContentSecurityPolicy, "default-src 'self'")
	resourceType := strings.ToLower(resource.Type)
	if strings.HasPrefix(resourceType, "text") {
		resourceType = echo.MIMETextPlainCharsetUTF8
	} else if strings.HasPrefix(resourceType, "video") || strings.HasPrefix(resourceType, "audio") {
		http.ServeContent(c.Response(), c.Request(), resource.Filename, time.Unix(resource.UpdatedTs, 0), bytes.NewReader(blob))
		return nil
	}

	return c.Stream(http.StatusOK, resourceType, bytes.NewReader(blob))
}

func (server *Server) findResourceList(c echo.Context) error {
	ctx := c.Request().Context()

	resourceFind := &store.ResourceFind{}
	if limit, err := strconv.Atoi(c.QueryParam("limit")); err == nil {
		resourceFind.Limit = &limit
	}
	if offset, err := strconv.Atoi(c.QueryParam("offset")); err == nil {
		resourceFind.Offset = &offset
	}

	list, err := server.store.FindResourceList(ctx, resourceFind)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch resource list").SetInternal(err)
	}
	return c.JSON(http.StatusOK, &okResponse{
		Data: list,
	})
}

func (server *Server) deleteResource(c echo.Context) error {
	ctx := c.Request().Context()
	resourceID, err := strconv.Atoi(c.Param("resourceId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("resourceId"))).SetInternal(err)
	}

	resource, err := server.store.FindResource(ctx, &store.ResourceFind{
		ID: &resourceID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to find resource").SetInternal(err)
	}
	if resource.InternalPath != "" {
		if err := os.Remove(resource.InternalPath); err != nil {
			log.Warn(fmt.Sprintf("failed to delete local file with path %s", resource.InternalPath), zap.Error(err))
		}
	}

	if err := server.store.DeleteResource(ctx, &store.ResourceDelete{ID: resourceID}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete resource").SetInternal(err)
	}

	return c.JSON(http.StatusOK, true)
}
