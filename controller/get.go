package controller

import (
	"net/http"

	"github.com/gauas/config-service/model"
	"github.com/gauas/config-service/packages/response"
	"github.com/gauas/config-service/utils"
	"github.com/labstack/echo/v4"
)

func (ctrl *Controller) Get(c echo.Context) error {
	req, err := utils.Bind[model.Config](c, utils.MaxBody)
	if err != nil {
		return response.NewError(http.StatusBadRequest, err.Error())
	}

	config, err := ctrl.service.GetConfig(c.Request().Context(), req)
	if err != nil {
		return response.Wrap(err)
	}

	if utils.AcceptsText(c) {
		return c.String(http.StatusOK, utils.Env(config.Config))
	}
	return c.JSON(http.StatusOK, config.Config)
}
