package middlewares

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// SecurityHeaders ... Add common security headers
func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.NewV4().String()
		c.Set("RequestID", requestID)
		c.Writer.Header().Set("X-Request-Id", requestID)
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		c.Writer.Header().Set("X-UA-Compatible", "chrome=1")
		c.Writer.Header().Set("Content-Security-Policy", "default-src 'none'")
		c.Next()
	}
}
