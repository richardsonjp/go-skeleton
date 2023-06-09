package middlewares

import (
	"github.com/gin-gonic/gin"
	"mgw/mgw-resi/config"
	"mgw/mgw-resi/pkg/utils/errors"
	"net/http"
)

// Generic header static API key validator middleware
func CheckHeaderStaticApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("STATIC-API-KEY")
		if token == "" {
			errors.ErrorCode(c, http.StatusUnauthorized)
			return
		}

		keys := config.Config.MiddlewareKeys
		if token != keys.StaticAPIKey {
			errors.ErrorCode(c, http.StatusUnauthorized)
		}

		c.Next()
	}
}
