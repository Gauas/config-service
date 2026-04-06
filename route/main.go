package route

import (
	"github.com/gauas/config-service/controller"
	"github.com/labstack/echo/v4"
)

type Router struct {
	server         *echo.Echo
	controller     *controller.Controller
	authMiddleware echo.MiddlewareFunc
}

func New(server *echo.Echo, controllerInstance *controller.Controller, authMiddleware echo.MiddlewareFunc) *Router {
	return &Router{
		server:         server,
		controller:     controllerInstance,
		authMiddleware: authMiddleware,
	}
}

func (r *Router) RegisterRoutes() {
	handler := newConfigHandler(r.controller)

	v1 := r.server.Group("/api/v1")

	v1.GET("/config/health", healthHandler)

	configGroup := v1.Group("/config", r.authMiddleware)
	configGroup.GET("", handler.get)
	configGroup.POST("", handler.create)
	configGroup.PUT("", handler.update)
	configGroup.DELETE("", handler.delete)
}
