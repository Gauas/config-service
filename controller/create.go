package controller

import (
	"net/http"

	"github.com/gauas/config-service/model"
	"github.com/gauas/config-service/packages/response"
	"github.com/gauas/config-service/utils"
	"github.com/labstack/echo/v4"
)

func (ctrl *Controller) Create(c echo.Context) error {
	ctx := c.Request().Context()

	var configMap model.JSONMap
	if utils.IsType(c, "multipart/form-data") {
		if err := utils.FileJSON(c, "file", &configMap, utils.MaxBody); err != nil {
			return response.NewError(http.StatusBadRequest, err.Error())
		}
	} else {
		if err := utils.JSON(c, &configMap, utils.MaxBody); err != nil {
			return response.NewError(http.StatusBadRequest, err.Error())
		}
	}

	req := &model.Config{
		Service:     c.QueryParam("service"),
		Environment: c.QueryParam("environment"),
		Config:      configMap,
	}

	config, err := ctrl.service.NewConfig(ctx, req)
	if err != nil {
		return response.Wrap(err)
	}

	return response.OK(c, config)
}
