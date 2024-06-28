package db

import (
	"database/sql"
	"embed"
	"golangechos/logger"
	"os"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"go.uber.org/zap"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

func New(development bool) (*sql.DB, error) {
	var uri string
	if development {
		uri = os.Getenv("DB_URI")
	} else {
		uri = os.Getenv("DB_URI")
	}

	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationsFS,
		Root:       "migrations",
	}

	if _, err := migrate.Exec(db, "postgres", migrations, migrate.Up); err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Mock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		logger.Fatal("MockDB Failure", zap.Error(err))
	}

	return db, mock
}
