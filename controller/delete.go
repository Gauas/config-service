package controller

import (
	"net/http"

	"github.com/gauas/config-service/model"
	"github.com/gauas/config-service/packages/response"
	"github.com/gauas/config-service/utils"
	"github.com/labstack/echo/v4"
)

func (ctrl *Controller) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	req, err := utils.Bind[model.Config](c, utils.MaxBody)
	if err != nil {
		return response.NewError(http.StatusBadRequest, err.Error())
	}
	
	if err := ctrl.service.DeleteConfig(ctx, req); err != nil {
		return response.Wrap(err)
	}
	return response.OK(c, "config deleted")
}
