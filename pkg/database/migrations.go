package database

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/opn-ooo/go-boilerplate/database"

	// For migrate files
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

// NewMigrate ...
func NewMigrate(db *gorm.DB) (*migrate.Migrate, error) {
	return database.Migration(db)
}
