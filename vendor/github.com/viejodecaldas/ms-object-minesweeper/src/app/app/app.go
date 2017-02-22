package app

import (
	"github.com/labstack/echo"
	"net/http"
)

// OK responds with a 200 ok JSON packet.
func OK(ctx echo.Context, msg string) error {
	c := Confirmation{Message: msg}
	return ctx.JSON(http.StatusOK, c)
}

// Success responds with a 200 ok JSON packet.
func Success(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, data)
}

// Error responds with a 500 internal server error.
func Error(ctx echo.Context, err error) error {
	e := ErrorResponse{
		Status: http.StatusInternalServerError,
		Detail: err.Error(),
	}

	return ctx.JSON(http.StatusInternalServerError, e)
}
