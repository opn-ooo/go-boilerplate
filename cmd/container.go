package cmd

import (
	"github.com/omiselabs/go-boilerplate/config"
	appHandler "github.com/omiselabs/go-boilerplate/internal/http/app/handlers"
	appMiddlewares "github.com/omiselabs/go-boilerplate/internal/http/app/middlewares"
	"github.com/omiselabs/go-boilerplate/pkg/database"
	"github.com/omiselabs/go-boilerplate/pkg/logger"
	commonMiddlewares "github.com/omiselabs/go-boilerplate/pkg/middlewares"
	"go.uber.org/dig"
)

// BuildContainer ...
func BuildContainer() *dig.Container {
	container := dig.New()

	_ = container.Provide(config.LoadConfig)
	_ = container.Provide(logger.NewLogger)
	_ = container.Provide(database.InitDatabase)

	// Handlers for User APIs
	_ = container.Provide(appHandler.NewHandler)

	// Handler for Admin APIs

	// Middlewares for All APIs
	_ = container.Provide(commonMiddlewares.SecurityHeaders, dig.Name("securityHeaders"))
	_ = container.Provide(commonMiddlewares.RequestID, dig.Name("requestID"))
	_ = container.Provide(commonMiddlewares.Logger, dig.Name("logger"))

	// Middlewares for App APIs
	_ = container.Provide(appMiddlewares.RequestHeaders, dig.Name("requestHeaders"))

	return container
}
