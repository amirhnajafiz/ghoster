package http

import (
	"github.com/labstack/echo/v4"
)

type HTTP struct{}

func (h HTTP) Use(ctx echo.Context) error {
	return nil
}
