package handlers

import (
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	UserHandler
}

func SetDefault(e *echo.Echo) {
	e.GET("/health", HealthHandler)
}

func SetRouter(e *echo.Echo, h *Handlers) {
	g := e.Group("/service")

	g.GET("/users", h.GetUsers)
	g.POST("/user", h.CreateUser)
}
