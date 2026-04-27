package middleware

import (
	"crypto/subtle"

	"github.com/gauas/config-service/utils"
	"github.com/labstack/echo/v4"
)

func (m *Middleware) Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if subtle.ConstantTimeCompare([]byte(c.Request().Header.Get("secret_key")), []byte(m.secretKey)) != 1 {
				return utils.Error(c, 401, "unauthorized")
			}
			return next(c)
		}
	}
}
