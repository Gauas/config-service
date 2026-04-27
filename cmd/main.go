package main

import (
	"log"

	"github.com/gauas/config-service/config"
	"github.com/gauas/config-service/controller"
	"github.com/gauas/config-service/infra"
	"github.com/gauas/config-service/kernel"
	middleware "github.com/gauas/config-service/middlewares"
	"github.com/gauas/config-service/model"
	"github.com/gauas/config-service/repository"
	"github.com/gauas/config-service/service"
)

func main() {
	configInstance := config.New()

	infraInstance := infra.New(configInstance)

	if err := model.Migrate(infraInstance.DB); err != nil {
		log.Fatal(err)
	}

	repositoryInstance := repository.New(infraInstance.DB)

	serviceInstance := service.New(repositoryInstance)

	controllerInstance := controller.New(serviceInstance)

	middlewareInstance := middleware.New(configInstance)

	kernel.New(controllerInstance, middlewareInstance, configInstance).Start()
}
