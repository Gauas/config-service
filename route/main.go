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
	v1 := r.server.Group("/v1")

	v1.GET("/config/health", healthHandler)
	r.registry(v1)
}

func (r *Router) registry(v1 *echo.Group) {
	c := r.controller
	configGroup := v1.Group("/config", r.authMiddleware)
	configGroup.GET("", c.Get)
	configGroup.POST("", c.Create)
	configGroup.PUT("", c.Update)
	configGroup.DELETE("", c.Delete)
}
