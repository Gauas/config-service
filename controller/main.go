package controller

import "github.com/gauas/config-service/data"

type Controller struct {
	Repository data.Repositories
}

func New(dataInstance *data.Data) *Controller {
	return &Controller{
		Repository: dataInstance.Repository,
	}
}
