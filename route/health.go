package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
}
