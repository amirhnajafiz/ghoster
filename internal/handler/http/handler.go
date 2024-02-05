package http

import (
	"context"
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

const baseDir = "./files/"

func (h HTTP) Upload(ctx echo.Context) error {
	// create a new context
	c := context.Background()

	// get title and create uid for document
	title := ctx.FormValue("title")
	uid := uuid.New().String()
	path := baseDir + uid + "."

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

	path = path + file.Filename

	// open file
	src, err := file.Open()
	if err != nil {
		log.Println(err)

		return echo.ErrInternalServerError
	}

	// create local file
	dst, err := os.Create(path)
	if err != nil {
		log.Println(err)

		return echo.ErrInternalServerError
	}

	// save content
	if _, err = io.Copy(dst, src); err != nil {
		log.Println(err)

		return echo.ErrInternalServerError
	}

	_ = src.Close()
	_ = dst.Close()

	// create a new document instance
	document := models.Document{
		Title:       title,
		UUID:        uid,
		CreatedAt:   time.Now(),
		Forbidden:   false,
		StoragePath: path,
	}

	// insert into database
	if _, err := h.DB.Collection("documents").InsertOne(c, document, nil); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}

func (h HTTP) Use(ctx echo.Context) error {
	uuid := ctx.Param("uuid")

	log.Println(uuid)

	return nil
}
