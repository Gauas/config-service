package route

import (
	"github.com/gauas/config-service/packages/response"
	"github.com/labstack/echo/v4"
)

func healthHandler(c echo.Context) error {
	return response.OK(c, echo.Map{"health": "ok"})
}
