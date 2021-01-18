package handlers

import (
	"github.com/opn-ooo/go-boilerplate/config"
	"github.com/opn-ooo/go-boilerplate/pkg/database"
)

func createHandler() HandlerInterface {
	configInstance, _ := config.LoadTestConfig()
	db, _, _ := database.GetMockDB()

	return NewHandler(
		db,
		configInstance,
	)
}
