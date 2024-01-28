package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HTTP struct{}

func (h HTTP) Healthy(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}
