package http

import (
	"fmt"
	"net/http"

	"github.com/amirhnajafiz/ghoster/internal/agent"
	"github.com/amirhnajafiz/ghoster/pkg/logger"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type HTTP struct {
	DB         *mongo.Database
	Agent      *agent.Agent
	Logger     logger.Logger
	Collection string
}

func (h HTTP) Healthy(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

func (h HTTP) Register(port int) {
	app := echo.New()

	app.GET("/", h.Healthy)

	api := app.Group("/api/docs")
	api.POST("/", h.Upload)
	api.GET("/", h.List)
	api.GET("/:uid", h.Use)

	if err := app.Start(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
}
