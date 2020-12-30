package logger

import (
	"bytes"

	"github.com/gin-gonic/gin"
)

// ResponseWriter ...
type ResponseWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

// Write ...
func (w ResponseWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

// WriteString ...
func (w ResponseWriter) WriteString(s string) (int, error) {
	w.Body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
