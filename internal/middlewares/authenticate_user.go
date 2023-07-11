package middlewares

import (
	"github.com/gin-gonic/gin"
	"go/skeleton/internal/services/authentication"
	"go/skeleton/pkg/utils/errors"
	"net/http"
	"strings"
)

// Generic header static API key validator middleware
func (m *MiddlewareDashboardAuth) AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		if bearerToken == "" {
			errors.ErrorCode(c, http.StatusUnauthorized)
			return
		}

		token := strings.Split(bearerToken, " ")
		if len(token) != 2 {
			if token[0] != "Bearer" {
				errors.ErrorString(c, "Bearer not found")
				return
			} else {
				errors.ErrorString(c, "Token not found")
				return
			}
		}

		SID := token[1]
		var session *authentication.SessionData
		session, err := m.authenticationService.GetSession(c, SID)
		if err != nil {
			errors.ErrorCode(c, http.StatusUnauthorized)
			return
		} else {
			c.Set("session", *session)
		}

		c.Next()
	}
}
