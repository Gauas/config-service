package utils

import "github.com/labstack/echo/v4"

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"model,omitempty"`
	Error  string      `json:"error,omitempty"`
}

type Context interface {
	JSON(code int, i interface{}) error
}

func OK(c echo.Context, data interface{}) error {
	return c.JSON(200, Response{Status: 200, Data: data})
}

func Created(c echo.Context, data interface{}) error {
	return c.JSON(201, Response{Status: 201, Data: data})
}

func NoContent(c echo.Context, message string) error {
	return c.JSON(200, Response{Status: 200, Data: echo.Map{"message": message}})
}

func Error(c Context, status int, message string) error {
	return c.JSON(status, Response{Status: status, Error: message})
}
