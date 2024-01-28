package http

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

type HTTP struct{}

func (h HTTP) Upload(ctx echo.Context) error {
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

	src.Close()
	dst.Close()

	return ctx.NoContent(http.StatusOK)
}

func (h HTTP) Use(ctx echo.Context) error {
	return nil
}
