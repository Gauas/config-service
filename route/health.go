package route

import (
	"github.com/gauas/config-service/utils"
	"github.com/labstack/echo/v4"
)

func healthHandler(c echo.Context) error {
	return utils.OK(c, echo.Map{"health": "ok"})
}
