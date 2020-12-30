package logger

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestResponseWriter_Write(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	rw := &ResponseWriter{
		ResponseWriter: c.Writer,
		Body:           bytes.NewBuffer([]byte{}),
	}

	n, err := rw.Write([]byte("test"))

	assert.Equal(t, 4, n)
	assert.Equal(t, []byte("test"), w.Body.Bytes())
	assert.Equal(t, []byte("test"), rw.Body.Bytes())
	assert.NoError(t, err)
}

func TestResponseWriter_WriteString(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	rw := &ResponseWriter{
		ResponseWriter: c.Writer,
		Body:           bytes.NewBuffer([]byte{}),
	}

	n, err := rw.WriteString("test")

	assert.Equal(t, 4, n)
	assert.Equal(t, "test", w.Body.String())
	assert.Equal(t, "test", rw.Body.String())
	assert.NoError(t, err)
}
