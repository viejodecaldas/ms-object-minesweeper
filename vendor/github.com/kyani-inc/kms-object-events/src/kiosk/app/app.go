package app

import (
	"net/http"

	"github.com/labstack/echo"
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

// CustomError responds with a custom error code.
func CustomError(ctx echo.Context, code int, err error) error {
	e := ErrorResponse{
		Status: code,
		Detail: err.Error(),
	}

	return ctx.JSON(code, e)
}

// BadRequest responds with a 400 bad request error.
func BadRequest(ctx echo.Context, err error) error {
	e := ErrorResponse{
		Status: http.StatusBadRequest,
		Detail: err.Error(),
	}

	return ctx.JSON(http.StatusBadRequest, e)
}
