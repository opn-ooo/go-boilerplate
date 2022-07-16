package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/opn-ooo/go-boilerplate/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// HandlerInterface ...
type HandlerInterface interface {
	HealthzHandler(c *gin.Context)
}

// Handler ...
type Handler struct {
	db     *gorm.DB
	config *config.Config
	logger *zap.Logger
}

// NewHandler ...
func NewHandler(
	db *gorm.DB,
	config *config.Config,
	logger *zap.Logger,
) HandlerInterface {
	return &Handler{
		db:     db,
		config: config,
		logger: logger,
	}
}
