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
	config     config.Config
}

func New(controllerInstance *controller.Controller, middlewareInstance *middleware.Middleware, configInstance config.Config) *Kernel {
	return &Kernel{
		controller: controllerInstance,
		middleware: middlewareInstance,
		config:  configInstance,
	}
}

func (k *Kernel) Start() {
	server := echo.New()
	server.HideBanner = true

	k.middleware.RegisterGlobal(server)

	router := route.New(server, k.controller, k.middleware.Auth())
	router.RegisterRoutes()

	addr := fmt.Sprintf(":%s", k.config.Port)
	log.Printf("config-service listening on %s", addr)

	if err := server.Start(addr); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
