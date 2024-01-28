package faas

import "github.com/labstack/echo/v4"

type Handler struct {
}

func (h Handler) Execute(ctx echo.Context) error {
	return nil
}
