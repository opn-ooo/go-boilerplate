package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/opn-ooo/go-boilerplate/config"
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
}

// NewHandler ...
func NewHandler(
	db *gorm.DB,
	config *config.Config,
// {{ START HANDLER ARGUMENTS }}
// {{ END HANDLER ARGUMENTS }}
) HandlerInterface {
	return &Handler{
		db:     db,
		config: config,
// {{ START HANDLER }}
// {{ END HANDLER }}
	}
}
