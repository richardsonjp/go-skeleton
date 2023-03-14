package middlewares

import (
	"go-skeleton/pkg/utils/errors"

	"github.com/gin-gonic/gin"
)

// Generic header static API key validator middleware
func (m *MiddlewareAccess) CheckHeaderStaticApiKey(header *string, apiKey *string, errorCode *int) gin.HandlerFunc {
	return func(c *gin.Context) {
		retrievedApiKey := c.Request.Header.Get(*header)
		if retrievedApiKey == "" || retrievedApiKey != *apiKey {
			errors.ErrorCode(c, *errorCode)
			return
		}

		c.Next()
	}
}
