package http

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/amirhnajafiz/ghoster/pkg/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h HTTP) Upload(ctx echo.Context) error {
	title := ctx.FormValue("title")
	user := ctx.FormValue("user")
	uid := uuid.New().String()
	now := time.Now()

	// get file from form data
	file, err := ctx.FormFile("project")
	if err != nil {
		log.Println(err)

		return echo.ErrBadRequest
	}

	// check file extension
	parts := strings.Split(file.Filename, ".")
	if parts[len(parts)-1] != "zip" {
		return echo.ErrBadRequest
	}

	// open file
	src, err := file.Open()
	if err != nil {
		log.Println(err)

		return echo.ErrInternalServerError
	}

	// create local file
	dst, err := os.Create(file.Filename)
	if err != nil {
		log.Println(err)

		return echo.ErrInternalServerError
	}

	// save content
	if _, err = io.Copy(dst, src); err != nil {
		log.Println(err)

		return echo.ErrInternalServerError
	}

	document := models.Document{
		Title:       title,
		User:        user,
		UUID:        uid,
		CreatedAt:   now,
		Forbidden:   false,
		StoragePath: dst.Name(),
	}

	_ = src.Close()
	_ = dst.Close()

	if _, err := h.DB.Collection("").InsertOne(nil, document, nil); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}

func (h HTTP) Use(ctx echo.Context) error {
	uuid := ctx.Param("uuid")

	log.Println(uuid)

	return nil
}
