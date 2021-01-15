package handlers

import (
	"github.com/omiselabs/go-boilerplate/config"
	"github.com/omiselabs/go-boilerplate/pkg/database"
)

func createHandler() HandlerInterface {
	configInstance, _ := config.LoadTestConfig()
	db, _, _ := database.GetMockDB()

	return NewHandler(
		db,
		configInstance,
	)
}
