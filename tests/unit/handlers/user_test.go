package handlers

import (
	"golangechos/app"
	"golangechos/handlers"
	"golangechos/logger"
	"golangechos/models"
	"golangechos/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockUserService struct {
	services.UserService
	MockGetUsers func() ([]models.User, error)
}

func (m *MockUserService) GetUsers() ([]models.User, error) {
	return m.MockGetUsers()
}

func SetupUserTest() func() {
	logger.New()
	return func() {
		logger.Sync()
		logger.Delete()
	}
}

func TestGetUsersSuccessCase(t *testing.T) {
	defer SetupUserTest()()

	s := &MockUserService{
		MockGetUsers: func() ([]models.User, error) {
			var r []models.User
			return r, nil
		},
	}

	mockService := &services.Services{User: s}

	e := app.Echo()
	h := handlers.New(mockService)

	req := httptest.NewRequest(http.MethodGet, "/service/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, h.UserHandler.GetUsers(c))
	assert.Equal(t, rec.Code, http.StatusOK)
}
