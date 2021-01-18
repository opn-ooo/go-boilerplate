package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/opn-ooo/go-boilerplate/cmd"
	"github.com/opn-ooo/go-boilerplate/config"
	"github.com/opn-ooo/go-boilerplate/internal/http/app/handlers"
	appMiddlewares "github.com/opn-ooo/go-boilerplate/internal/http/app/middlewares"
	commonMiddlewares "github.com/opn-ooo/go-boilerplate/pkg/middlewares"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

// Start ...
func Start() error {
	var configInstance *config.Config
	var logger *zap.Logger

	container := cmd.BuildContainer()

	app, err := bootstrap(container)
	if err != nil {
		return err
	}

	err = container.Invoke(func(_config *config.Config, _logger *zap.Logger) {
		configInstance = _config
		logger = _logger
	})
	if err != nil {
		return err
	}

	logger.Info("starting APP server ( APIs for Users )",
		zap.String("url", configInstance.App.APP.URL),
		zap.Uint("port", configInstance.App.APP.Port),
	)

	err = app.Run(":" + strconv.FormatUint(uint64(configInstance.App.APP.Port), 10))

	return err
}

func bootstrap(container *dig.Container) (*gin.Engine, error) {
	var err error

	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	gin.DefaultWriter = os.Stderr

	app := gin.New()

	err = setupMiddlewares(app, container)
	if err == nil {
		err = setupRoutes(app, container)
	}

	return app, err
}

func setupMiddlewares(app *gin.Engine, container *dig.Container) error {
	var securityHeadersMiddleware gin.HandlerFunc
	var requestIDMiddleware gin.HandlerFunc
	var loggerMiddleware gin.HandlerFunc
	var configInstance *config.Config

	if err := container.Invoke(func(
		_handler handlers.HandlerInterface,
		_appMiddlewares appMiddlewares.Middlewares,
		_commonMiddlewares commonMiddlewares.Middlewares,
		_configInstance *config.Config,
	) {
		securityHeadersMiddleware = _commonMiddlewares.SecurityHeaders
		requestIDMiddleware = _commonMiddlewares.RequestID
		loggerMiddleware = _commonMiddlewares.Logger
		configInstance = _configInstance
	}); err != nil {
		return err
	}

	app.Use(gin.Recovery())
	app.Use(securityHeadersMiddleware)
	app.Use(requestIDMiddleware)
	app.Use(loggerMiddleware)

	return nil
}

func setupRoutes(app *gin.Engine, container *dig.Container) error {
	var handler handlers.HandlerInterface
	var requestHeaderMiddleware gin.HandlerFunc

	if err := container.Invoke(func(
		_handler handlers.HandlerInterface,
		_appMiddlewares appMiddlewares.Middlewares,
	) {
		handler = _handler
		requestHeaderMiddleware = _appMiddlewares.RequestHeaders
	}); err != nil {
		return err
	}

	app.GET("/healthz", handler.HealthzHandler)

	v1Router := app.Group("v1")
	v1Router.Use(requestHeaderMiddleware)

	return nil
}

func main() {
	err := Start()
	if err != nil {
		panic(err)
	}
}
