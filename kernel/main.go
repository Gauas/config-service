package kernel

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gauas/config-service/config"
	"github.com/gauas/config-service/controller"
	middleware "github.com/gauas/config-service/middlewares"
	"github.com/gauas/config-service/route"
	"github.com/labstack/echo/v4"
)

type Kernel struct {
	controller *controller.Controller
	middleware *middleware.Middleware
	appConfig  config.AppConfig
}

func New(controllerInstance *controller.Controller, middlewareInstance *middleware.Middleware, appConfig config.AppConfig) *Kernel {
	return &Kernel{
		controller: controllerInstance,
		middleware: middlewareInstance,
		appConfig:  appConfig,
	}
}

func (k *Kernel) Start() {
	server := echo.New()
	server.HideBanner = true

	k.middleware.RegisterGlobal(server)

	routerInstance := route.New(server, k.controller, k.middleware.Auth())
	routerInstance.RegisterRoutes()

	addr := fmt.Sprintf(":%s", k.appConfig.Port)
	log.Printf("config-service listening on %s", addr)

	if err := server.Start(addr); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
