package handlers

import (
	"github.com/omiselabs/gin-boilerplate/config"
	"github.com/omiselabs/gin-boilerplate/pkg/database"
)

func createHandler() HandlerInterface {
	configInstance, _ := config.LoadTestConfig()
	db, _, _ := database.GetMockDB()

	return NewHandler(
		db,
		configInstance,
	)
}
