package utils

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string { return e.Message }

func RespondError(c echo.Context, err error) error {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return c.JSON(appErr.Code, echo.Map{"error": appErr.Message})
	}
	return c.JSON(http.StatusInternalServerError, echo.Map{"error": "internal server error"})
}
