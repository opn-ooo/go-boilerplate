package database

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	// For migrate files
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

// NewMigrate ...
func NewMigrate(db *gorm.DB) (*migrate.Migrate, error) {
	_db, err := db.DB()
	if err != nil {
		return nil, err
	}
	dbDriver, err := postgres.WithInstance(_db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance("file://database/migrations", "postgres", dbDriver)
	if err != nil {
		return nil, err
	}

	return m, nil
}
