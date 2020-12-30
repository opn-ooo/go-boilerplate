package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// Middlewares ... middleware structure for DI
type Middlewares struct {
	dig.In
	SecurityHeaders gin.HandlerFunc `name:"securityHeaders"`
	RequestID       gin.HandlerFunc `name:"requestID"`
	Logger          gin.HandlerFunc `name:"logger"`
}
