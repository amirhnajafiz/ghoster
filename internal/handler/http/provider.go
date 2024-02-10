package http

import (
	"context"
	"errors"
	"time"

	"github.com/amirhnajafiz/ghoster/pkg/enum"
	"github.com/amirhnajafiz/ghoster/pkg/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	ErrDocumentNotFound = errors.New("")
)

func (h HTTP) Provider(uid string) error {
	c := context.Background()

	// create filter
	filter := bson.M{"uuid": uid}

	// fetch the first object
	doc := new(models.Document)

	cursor := h.DB.Collection(h.Collection).FindOne(c, filter, nil)
	if err := cursor.Err(); err != nil {
		h.Logger.Error(err)

		return echo.ErrInternalServerError
	}

	// parse into the docs object
	if err := cursor.Decode(doc); err != nil {
		h.Logger.Error(err)

		return echo.ErrInternalServerError
	}

	// create a new worker
	w, err := h.Agent.NewWorker()
	if err != nil {
		h.Logger.Error(err)

		return echo.ErrInternalServerError
	}

	// get worker stdin and stdout
	stdin := w.GetStdin()
	stdout := w.GetStdout()

	// pass the storage path for starting the process
	stdin <- doc.StoragePath

	// get the result from the process
	result := <-stdout
	msg := result.(string)

	// dismiss the process
	stdin <- enum.CodeDismiss

	// update fields
	update := bson.D{
		{
			"$set",
			bson.D{
				{"forbidden", msg == string(enum.CodeFailure)},
				{"last_execute", time.Now()},
			},
		},
	}

	// update document
	if _, er := h.DB.Collection(h.Collection).UpdateOne(c, filter, update, nil); er != nil {
		h.Logger.Error(er)

		return echo.ErrInternalServerError
	}

	return nil
}
