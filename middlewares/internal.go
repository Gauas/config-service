package middleware

import (
	"crypto/subtle"
	"net/http"

	"github.com/gauas/config-service/packages/response"
	"github.com/labstack/echo/v4"
)

func (m *Middleware) Internal() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if subtle.ConstantTimeCompare([]byte(c.Request().Header.Get("Secret-Key")), []byte(m.secretKey)) != 1 {
				return response.NewError(http.StatusUnauthorized, "unauthorized")
			}
			return next(c)
		}
	}
}
