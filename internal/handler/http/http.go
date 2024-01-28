package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type HTTP struct {
	DB *mongo.Database
}

func (h HTTP) Healthy(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

func (h HTTP) Register(app *echo.Echo) {
	app.GET("/", h.Healthy)

	api := app.Group("/api")
	api.POST("/upload", h.Upload)
	api.GET("/use/:uuid", h.Use)
}
