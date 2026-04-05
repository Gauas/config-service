package route

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gauas/config-service/controller"
	"github.com/gauas/config-service/data"
	"github.com/gauas/config-service/utils"
	"github.com/labstack/echo/v4"
)

const maxBodyBytes = 1 << 20

type configHandler struct {
	controller *controller.Controller
}

func newConfigHandler(controllerInstance *controller.Controller) *configHandler {
	return &configHandler{controller: controllerInstance}
}

func (h *configHandler) get(c echo.Context) error {
	serviceName := c.QueryParam("service")
	if serviceName == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "service is required"})
	}

	environment := c.QueryParam("env")
	if environment == "" {
		environment = "default"
	}

	configEntry, err := h.controller.Get(serviceName, environment)
	if err != nil {
		return utils.RespondError(c, err)
	}

	if strings.Contains(c.Request().Header.Get("Accept"), "text/plain") {
		return c.String(http.StatusOK, utils.ToEnvFormat(configEntry.Config))
	}

	return c.JSON(http.StatusOK, configEntry.Config)
}

func (h *configHandler) create(c echo.Context) error {
	serviceName, environment, payload, err := parseRequest(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := utils.ValidateFlatMap(payload); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	created, err := h.controller.Create(serviceName, environment, payload)
	if err != nil {
		return utils.RespondError(c, err)
	}

	return c.JSON(http.StatusCreated, created)
}

func (h *configHandler) update(c echo.Context) error {
	serviceName, environment, payload, err := parseRequest(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := utils.ValidateFlatMap(payload); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	updated, err := h.controller.Merge(serviceName, environment, payload)
	if err != nil {
		return utils.RespondError(c, err)
	}

	return c.JSON(http.StatusOK, updated)
}

func (h *configHandler) delete(c echo.Context) error {
	var body struct {
		Service     string `json:"service"`
		Environment string `json:"environment"`
	}

	c.Request().Body = http.MaxBytesReader(c.Response(), c.Request().Body, maxBodyBytes)

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	if body.Service == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "service is required"})
	}

	if body.Environment == "" {
		body.Environment = "default"
	}

	if err := h.controller.Delete(body.Service, body.Environment); err != nil {
		return utils.RespondError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func parseRequest(c echo.Context) (serviceName, environment string, payload data.JSONMap, err error) {
	if strings.Contains(c.Request().Header.Get("Content-Type"), "multipart/form-data") {
		return parseMultipart(c)
	}
	return parseJSON(c)
}

func parseJSON(c echo.Context) (serviceName, environment string, payload data.JSONMap, err error) {
	var body struct {
		Service string       `json:"service"`
		Env     string       `json:"env"`
		Config  data.JSONMap `json:"config"`
	}

	if decodeErr := json.NewDecoder(http.MaxBytesReader(c.Response(), c.Request().Body, maxBodyBytes)).Decode(&body); decodeErr != nil {
		err = fmt.Errorf("invalid JSON: %w", decodeErr)
		return
	}

	if body.Service == "" {
		err = fmt.Errorf("service is required")
		return
	}
	if body.Config == nil {
		err = fmt.Errorf("config is required")
		return
	}

	serviceName = body.Service
	environment = body.Env
	if environment == "" {
		environment = "default"
	}
	payload = body.Config
	return
}

func parseMultipart(c echo.Context) (serviceName, environment string, payload data.JSONMap, err error) {
	serviceName = c.FormValue("service")
	if serviceName == "" {
		err = fmt.Errorf("service is required")
		return
	}

	environment = c.FormValue("env")
	if environment == "" {
		environment = "default"
	}

	file, fileErr := c.FormFile("file")
	if fileErr != nil {
		err = fmt.Errorf("file is required: %w", fileErr)
		return
	}

	if file.Size > maxBodyBytes {
		err = fmt.Errorf("file exceeds 1 MB limit")
		return
	}

	src, openErr := file.Open()
	if openErr != nil {
		err = fmt.Errorf("cannot open file: %w", openErr)
		return
	}
	defer src.Close()

	raw, readErr := io.ReadAll(io.LimitReader(src, maxBodyBytes))
	if readErr != nil {
		err = fmt.Errorf("cannot read file: %w", readErr)
		return
	}

	if jsonErr := json.Unmarshal(raw, &payload); jsonErr != nil {
		err = fmt.Errorf("invalid JSON in file: %w", jsonErr)
	}
	return
}
