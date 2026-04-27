package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

const MaxBody int64 = 1 << 20

func IsType(c echo.Context, contentType string) bool {
	return strings.Contains(c.Request().Header.Get("Content-Type"), contentType)
}

func AcceptsText(c echo.Context) bool {
	return strings.Contains(c.Request().Header.Get("Accept"), "text/plain")
}

func JSON(c echo.Context, destination interface{}, maxBytes int64) error {
	c.Request().Body = http.MaxBytesReader(c.Response(), c.Request().Body, maxBytes)
	if err := json.NewDecoder(c.Request().Body).Decode(destination); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}
	return nil
}

func FileJSON(c echo.Context, fileField string, destination interface{}, maxBytes int64) error {
	raw, err := File(c, fileField, maxBytes)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(raw, destination); err != nil {
		return fmt.Errorf("invalid JSON in file: %w", err)
	}
	return nil
}
