package middleware

import (
	"crypto/subtle"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m *Middleware) Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if subtle.ConstantTimeCompare([]byte(c.Request().Header.Get("secret_key")), []byte(m.secretKey)) != 1 {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
			}
			return next(c)
		}
	}
}
