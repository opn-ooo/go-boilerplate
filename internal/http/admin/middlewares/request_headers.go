package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/omiselabs/go-boilerplate/internal/http/admin/responses"
)

// RequestHeaders ... Parse Request Headers
func RequestHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {

		platform := strings.ToLower(c.GetHeader("X-OPN-Platform"))
		osVersion := strings.ToLower(c.GetHeader("X-OPN-OS-Version"))
		deviceName := strings.ToLower(c.GetHeader("X-OPN-Device-Name"))
		clientVersion := strings.ToLower(c.GetHeader("X-OPN-Client-Version"))

		if platform == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewInvalidHeaderErrorStatus("Platform ( X-OPN-Platform ) should be ios or android"))
			return
		}
		if osVersion == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewInvalidHeaderErrorStatus("OS Version( X-OPN-OS-Version ) is empty"))
			return
		}
		if deviceName == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewInvalidHeaderErrorStatus("Device Name( X-OPN-Device-Name ) is empty"))
			return
		}
		if clientVersion == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewInvalidHeaderErrorStatus("Client Version( X-OPN-Client-Version ) is empty"))
			return
		}

		c.Set("Platform", platform)
		c.Set("OsVersion", osVersion)
		c.Set("DeviceName", deviceName)
		c.Set("ClientVersion", clientVersion)

		c.Next()
	}
}
