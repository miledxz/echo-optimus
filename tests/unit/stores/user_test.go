package stores

import (
	"golangechos/db"
	"golangechos/models"
	"golangechos/stores"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserStoreCreateSuccessCase(t *testing.T) {
	mockDB, mock := db.Mock()
	defer mockDB.Close()

	u := &models.User{
		Power:       150,
		Username: "test",
	}

	mock.NewRows([]string{"id", "username"})
	mock.ExpectQuery("INSERT INTO users (id, username) VALUES ($1, $2) RETURNING id").WithArgs(u.Power, u.Username).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	s := stores.New(mockDB)

	r, err := s.User.Create(nil, u)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), r)
	assert.NoError(t, mock.ExpectationsWereMet())
}
