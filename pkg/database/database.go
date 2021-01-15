package database

import (
	"context"
	"fmt"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/omiselabs/go-boilerplate/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	// Postgres driver for GORM
	"gorm.io/driver/postgres"
)

// InitDatabase ... initializes the database connection.
func InitDatabase(configInstance *config.Config) (*gorm.DB, error) {
	pgConfig := configInstance.Database.Postgres
	conn := &connection{
		host:     pgConfig.Host,
		port:     pgConfig.Port,
		dbname:   pgConfig.Name,
		user:     pgConfig.User,
		password: pgConfig.Password,
	}

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable",
		pgConfig.User, pgConfig.Password, pgConfig.Name, pgConfig.Port, pgConfig.Host)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logError(err, conn)
	}

	return db, err
}

// PingDB ...
func PingDB(db *gorm.DB) bool {
	ctx := context.Background()

	_db, err := db.DB()
	if err != nil {
		return false
	}
	conn, err := _db.Conn(ctx)
	if err != nil {
		return false
	}
	defer conn.Close()

	if _, err := conn.ExecContext(ctx, "SELECT 1"); err != nil {
		return false
	}

	return true
}

// GetMockDB ... Get mock db connection for unit testing purpose
func GetMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return gdb, mock, nil
}

func logError(err error, conn *connection) {
	conn.filterPassword()
	log.Println(err)
	log.Fatal("Cannot connect to database: Please check the connection: " + conn.string())
}
