package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-skeleton/config"
	"go-skeleton/pkg/utils/errors"
)

// Generic header static API key validator middleware
func CheckHeaderStaticApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("STATIC-API-KEY")
		if token == "" {
			errors.ErrorCode(c, errors.TOKEN_UNAUTHORIZED)
			return
		}

		keys := config.Config.MiddlewareKeys
		if token != keys.StaticAPIKey {
			errors.ErrorCode(c, errors.TOKEN_UNAUTHORIZED)
		}

		c.Next()
	}
}
