package stores

import (
	"database/sql"
	"golangechos/logger"
	"golangechos/models"

	"go.uber.org/zap"
)

type (
	UserStore interface {
		Create(tx *sql.Tx, user *models.User) (int64, error)
		Get(tx *sql.Tx) ([]models.User, error)
	}

	userStore struct {
		*sql.DB
	}
)

func (s *userStore) Create(tx *sql.Tx, user *models.User) (int64, error) {
	var err error

	query := "INSERT INTO users (id, username) VALUES ($1, $2) RETURNING id"
	if err != nil {
		return 0, err
	}

	var id int64

	if tx != nil {
		err = tx.QueryRow(query, user.Power, user.Username).Scan(&id)
	} else {
		err = s.QueryRow(query, user.Power, user.Username).Scan(&id)
	}

	if err != nil {
		logger.Error("Create User Failure", zap.Error(err))
	}

	return id, nil
}

func (s *userStore) Get(tx *sql.Tx) ([]models.User, error) {
	users := make([]models.User, 0)

	rows, err := s.Query("SELECT id, username from Users")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var u models.User
		err = rows.Scan(&u.Power, &u.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
