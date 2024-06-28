package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Echo service is available")
}