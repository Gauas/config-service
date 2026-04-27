package controller

import (
	"net/http"

	"github.com/gauas/config-service/model"
	"github.com/gauas/config-service/utils"
	"github.com/labstack/echo/v4"
)

func (ctrl *Controller) Update(c echo.Context) error {
	ctx := c.Request().Context()

	req, err := utils.Bind[model.Config](c, utils.MaxBody)
	if err != nil {
		return utils.Error(c, http.StatusBadRequest, err.Error())
	}
	if utils.IsType(c, "multipart/form-model") {
		if err := utils.FileJSON(c, "file", &req.Config, utils.MaxBody); err != nil {
			return utils.Error(c, http.StatusBadRequest, err.Error())
		}
	}

	config, err := ctrl.service.UpdateConfig(ctx, req)
	if err != nil {
		return utils.RespondError(c, err)
	}

	return utils.OK(c, config)
}
