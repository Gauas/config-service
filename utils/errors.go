package utils

import (
	"errors"
	"net/http"
)

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string { return e.Message }

func RespondError(c Context, err error) error {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return Error(c, appErr.Code, appErr.Message)
	}
	return Error(c, http.StatusInternalServerError, "internal server error")
}
