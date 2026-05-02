package kernel

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gauas/config-service/config"
	"github.com/gauas/config-service/controller"
	middleware "github.com/gauas/config-service/middlewares"
	"github.com/gauas/config-service/packages/response"
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
		config:     configInstance,
	}
}

func (k *Kernel) Start() {
	server := echo.New()
	server.HideBanner = true
	server.HTTPErrorHandler = func(err error, c echo.Context) {
		var e *response.Error
		if errors.As(err, &e) {
			_ = c.JSON(e.Code, response.Response{Status: e.Code, Error: e.Message})
			return
		}
		var httpErr *echo.HTTPError
		if errors.As(err, &httpErr) {
			code := httpErr.Code
			_ = c.JSON(code, response.Response{Status: code, Error: fmt.Sprintf("%v", httpErr.Message)})
			return
		}
		_ = c.JSON(http.StatusInternalServerError, response.Response{Status: http.StatusInternalServerError, Error: "internal server error"})
	}

	k.middleware.RegisterGlobal(server)

	router := route.New(server, k.controller, k.middleware.Internal())
	router.RegisterRoutes()

	addr := fmt.Sprintf(":%s", k.config.Port)
	log.Printf("config-service listening on %s", addr)

	if err := server.Start(addr); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
