package database

import (
	"embed"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"gorm.io/gorm"
)

//go:embed migrations/*.sql
var migrations embed.FS

func Migration(db *gorm.DB) (*migrate.Migrate, error) {

	source, err := httpfs.New(http.FS(migrations), "migrations")

	_db, err := db.DB()
	if err != nil {
		return nil, err
	}
	dbDriver, err := postgres.WithInstance(_db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	return migrate.NewWithInstance("httpfs", source, "postgres", dbDriver)

}
