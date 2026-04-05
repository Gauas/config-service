package controller

import (
	"net/http"

	"github.com/gauas/config-service/data"
	"github.com/gauas/config-service/utils"
	"github.com/google/uuid"
)

func (ctrl *Controller) Get(serviceName, environment string) (*data.Config, error) {
	configEntry, err := ctrl.configRepo.FindOne("service = ? AND environment = ?", serviceName, environment)
	if err != nil {
		return nil, err
	}
	if configEntry == nil {
		return nil, &utils.AppError{Code: http.StatusNotFound, Message: "config not found"}
	}
	return configEntry, nil
}

func (ctrl *Controller) Create(serviceName, environment string, payload data.JSONMap) (*data.Config, error) {
	existing, err := ctrl.configRepo.FindOne("service = ? AND environment = ?", serviceName, environment)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, &utils.AppError{Code: http.StatusConflict, Message: "config already exists"}
	}

	configEntry := &data.Config{
		ID:          uuid.New(),
		Service:     serviceName,
		Environment: environment,
		Config:      payload,
	}

	if err := ctrl.configRepo.Create(configEntry); err != nil {
		return nil, err
	}
	return configEntry, nil
}

func (ctrl *Controller) Merge(serviceName, environment string, payload data.JSONMap) (*data.Config, error) {
	configEntry, err := ctrl.configRepo.FindOne("service = ? AND environment = ?", serviceName, environment)
	if err != nil {
		return nil, err
	}
	if configEntry == nil {
		return nil, &utils.AppError{Code: http.StatusNotFound, Message: "config not found"}
	}

	for key, value := range payload {
		configEntry.Config[key] = value
	}

	if err := ctrl.configRepo.Save(configEntry); err != nil {
		return nil, err
	}
	return configEntry, nil
}

func (ctrl *Controller) Delete(serviceName, environment string) error {
	configEntry, err := ctrl.configRepo.FindOne("service = ? AND environment = ?", serviceName, environment)
	if err != nil {
		return err
	}
	if configEntry == nil {
		return &utils.AppError{Code: http.StatusNotFound, Message: "config not found"}
	}
	return ctrl.configRepo.Delete("service = ? AND environment = ?", serviceName, environment)
}
