package main

import (
	"github.com/gauas/config-service/config"
	"github.com/gauas/config-service/controller"
	"github.com/gauas/config-service/data"
	"github.com/gauas/config-service/infra"
	"github.com/gauas/config-service/kernel"
	middleware "github.com/gauas/config-service/middlewares"
)

func main() {
	appConfig := config.New()

	infraInstance := infra.New(appConfig)
	dataInstance := data.New(infraInstance)
	controllerInstance := controller.New(dataInstance)
	middlewareInstance := middleware.New(appConfig)

	kernel.New(controllerInstance, middlewareInstance, appConfig).Start()
}
