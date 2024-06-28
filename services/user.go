package services

import (
	"golangechos/models"
	"golangechos/stores"
)

type (
	UserService interface {
		CreateUser(u *models.User) (int64, error)
		GetUsers() ([] models.User, error)
	}

	userService struct {
		stores *stores.Stores
	}
)

func(s *userService) CreateUser(u *models.User) (int64, error) {
	r, err := s.stores.User.Create(nil, u)
	return r, err
}

func(s *userService) GetUsers() ([]models.User, error) {
	r, err := s.stores.User.Get(nil)
	return r, err
}
