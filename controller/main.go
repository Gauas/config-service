package controller

import "github.com/gauas/config-service/data"

type Controller struct {
	configRepo data.Repository[data.Config]
}

func New(dataInstance *data.Data) *Controller {
	return &Controller{
		configRepo: dataInstance.NewConfigRepo(),
	}
}
