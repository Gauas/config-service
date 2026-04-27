package middleware

import (
	"github.com/gauas/config-service/config"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

type Middleware struct {
	secretKey string
}

func New(appConfig config.Config) *Middleware {
	return &Middleware{secretKey: appConfig.SecretKey}
}

func (m *Middleware) RegisterGlobal(server *echo.Echo) {
	server.Use(echoMiddleware.Recover())
	server.Use(echoMiddleware.Logger())
	server.Use(echoMiddleware.RequestID())
}
