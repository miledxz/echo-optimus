package handlers

import (
	"golangechos/logger"
	"golangechos/models"
	"golangechos/services"
	"golangechos/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type (
	UserHandler interface {
		CreateUser(c echo.Context) error
		GetUsers(c echo.Context) error
	}

	userHandler struct {
		services.UserService
	}
)

func New(s *services.Services) *Handlers {
	return &Handlers{
		UserHandler: &userHandler{s.User},
	}
}

func (h *userHandler) CreateUser(c echo.Context) error {

	var u *models.User

	if err := c.Bind(&u); err != nil {
		logger.Error("Mapping failure", zap.Error(err))
		return c.JSON(http.StatusBadRequest, utils.Error{Message: err.Error()})
	}

	r, err := h.UserService.CreateUser(u)
	if err != nil {
		logger.Error("Create User Failure", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, r)
}

func (h *userHandler) GetUsers(c echo.Context) error {
	r, err := h.UserService.GetUsers()

	if err != nil {
		logger.Error("Get Users Failure", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, r)
}
