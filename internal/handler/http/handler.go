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
	"go.mongodb.org/mongo-driver/bson"
)

const baseDir = "./files/"

// Upload documents into agent local storage.
func (h HTTP) Upload(ctx echo.Context) error {
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

	// create a new context
	c := context.Background()

	// insert into database
	if _, er := h.DB.Collection("documents").InsertOne(c, document, nil); er != nil {
		log.Println(er)

		return echo.ErrInternalServerError
	}

	return ctx.NoContent(http.StatusOK)
}

// List returns a list of current uploads.
func (h HTTP) List(ctx echo.Context) error {
	// create context
	c := context.Background()

	// create filter
	filter := bson.D{}

	// query for documents
	cursor, err := h.DB.Collection("documents", nil).Find(c, filter, nil)
	if err != nil {
		log.Println(err)

		return echo.ErrInternalServerError
	}

	// create a docs list for fetching
	ids := make([]string, 0)

	for cursor.Next(c) {
		var tmp models.Document
		if err := cursor.Decode(&tmp); err != nil {
			log.Println(err)

			return echo.ErrInternalServerError
		}

		ids = append(ids, tmp.UUID+" "+tmp.Title)
	}

	return ctx.JSON(http.StatusOK, ids)
}

func (h HTTP) Use(ctx echo.Context) error {
	uuid := ctx.Param("uuid")

	log.Println(uuid)

	return nil
}
