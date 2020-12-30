package database

import (
	"errors"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"

	"github.com/omiselabs/gin-boilerplate/pkg/database"
)

func getDBConnection() *gorm.DB {
	return viper.Get("database").(*gorm.DB)
}

// Migrate ... migrate database
func Migrate(command *cobra.Command, args []string) error {
	_database := getDBConnection()
	_migrate, err := database.NewMigrate(_database)
	if err != nil {
		return err
	}

	command.Println("Migrating...")
	if err := _migrate.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

// Drop ... drop all migrations
func Drop(command *cobra.Command, args []string) error {
	_database := getDBConnection()
	_migrate, err := database.NewMigrate(_database)
	if err != nil {
		return err
	}

	command.Println("Deleting...")
	return _migrate.Drop()
}
