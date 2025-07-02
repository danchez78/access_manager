package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Result[T any] struct {
	Result T `json:"result"`
}

type ErrorResult struct {
	Error string `json:"error"`
}

func ReturnResult[T any](c echo.Context, res T) error {
	return c.JSON(http.StatusOK, Result[T]{Result: res})
}

func ReturnError(c echo.Context, statusCode int, err error) error {
	return c.JSON(statusCode, ErrorResult{Error: err.Error()})
}
