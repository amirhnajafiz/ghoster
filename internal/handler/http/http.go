package http

import (
	"net/http"

	"github.com/amirhnajafiz/ghoster/pkg/logger"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type HTTP struct {
	DB     *mongo.Database
	Logger logger.Logger
}

func (h HTTP) Healthy(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

func (h HTTP) Register(app *echo.Echo) {
	app.GET("/", h.Healthy)

	api := app.Group("/api/docs")
	api.POST("/", h.Upload)
	api.GET("/", h.List)
	api.GET("/:uid", h.Use)
}
