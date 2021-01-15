package middlewares

import (
	"bytes"
	"fmt"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/omiselabs/go-boilerplate/config"
	"github.com/omiselabs/go-boilerplate/pkg/logger"
	"go.uber.org/zap"
)

// Logger ... Middleware for logging request and response
func Logger(zapLogger *zap.Logger, config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		requestID := c.GetString("RequestID")
		reqLogger := zapLogger.With(
			zap.String("request_id", requestID),
			zap.String("ip", c.Request.RemoteAddr),
		)

		defer reqLogger.Sync()

		if config.Log.Level == "debug" {
			var reqDump strings.Builder

			r, _ := httputil.DumpRequest(c.Request, true)
			reqDump.Write(r)

			reqLogger.Debug(
				"client request",
				zap.String("request", reqDump.String()),
			)
		}

		writer := getOrSetResponseWriter(c)

		c.Set("Logger", reqLogger)
		c.Next()

		if config.Log.Level == "debug" {
			var respDump strings.Builder

			for name, headers := range writer.Header() {
				name = strings.ToLower(name)

				for _, h := range headers {
					respDump.WriteString(fmt.Sprintf("%v: %v", name, h))
				}
			}

			respDump.WriteString("\n")
			respDump.Write(writer.Body.Bytes())

			reqLogger.Debug(
				"client response",
				zap.String("response", respDump.String()),
			)
		}

		end := time.Now().UTC()
		latency := end.Sub(start)

		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				reqLogger.Error(e)
			}
		} else {
			reqLogger.Info(path,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("user_agent", c.Request.UserAgent()),
				zap.String("time", end.Format(time.RFC3339)),
				zap.Duration("latency", latency),
			)
		}
	}
}

func getOrSetResponseWriter(c *gin.Context) *logger.ResponseWriter {
	if w, ok := c.Writer.(*logger.ResponseWriter); ok {
		return w
	}

	w := &logger.ResponseWriter{
		Body:           bytes.NewBuffer([]byte{}),
		ResponseWriter: c.Writer,
	}
	c.Writer = w

	return w
}
