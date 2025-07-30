package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-skeleton/internal/model/enum"
	"go-skeleton/internal/services/authentication"
	"go-skeleton/internal/services/context_path"
	constUtil "go-skeleton/pkg/utils/constant"
	"go-skeleton/pkg/utils/errors"
	timeutil "go-skeleton/pkg/utils/time"
	"strings"
	"time"
)

// Generic header static API key validator middleware
func (m *MiddlewareDashboardAuth) AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		if bearerToken == "" {
			errors.ErrorCode(c, errors.CLIENT_AUTH_ERROR, errors.SetDetails([]constUtil.ErrorDetails{
				{Key: "Authorization", Value: "Authorization Header not found"},
			}))
			return
		}

		token := strings.Split(bearerToken, " ")
		if len(token) != 2 {
			if token[0] != "Bearer" {
				errors.ErrorCode(c, errors.CLIENT_AUTH_ERROR, errors.SetDetails([]constUtil.ErrorDetails{
					{Key: "Authorization", Value: "Bearer not found"},
				}))
				return
			} else {
				errors.ErrorCode(c, errors.CLIENT_AUTH_ERROR, errors.SetDetails([]constUtil.ErrorDetails{
					{Key: "Authorization", Value: "Token not found"},
				}))
				return
			}
		}

		SID := token[1]
		var session *authentication.SessionData
		session, err := m.authenticationService.GetSession(c, SID)
		if err != nil {
			genericError := err.(*errors.GenericError)
			errors.ErrorCode(c, genericError.GetCode(), errors.SetDetails([]constUtil.ErrorDetails{
				{Key: "Authorization", Value: "Session expired or not found"},
			}))
			return
		} else {
			if session.User.Status != enum.ACTIVE.String() {
				errors.ErrorCode(c, errors.CLIENT_AUTH_ERROR, errors.SetDetails([]constUtil.ErrorDetails{
					{Key: "Authorization", Value: "User is not active"},
				}))
				return
			}
			loc, _ := time.LoadLocation("Asia/Jakarta")
			expiredAt, _ := time.ParseInLocation("2006-01-02 15:04:05", session.Credential.ExpiredAt, loc)
			buffer := 1 * time.Hour
			if expiredAt.Sub(timeutil.Now()) < buffer {
				if err := m.authenticationService.ExtendSession(c, SID); err != nil {
					genericError := err.(*errors.GenericError)
					errors.ErrorCode(c, genericError.GetCode(), errors.SetDetails([]constUtil.ErrorDetails{
						{Key: "Authorization", Value: "Session extension got error"},
					}))
				}
				session.Credential.ExpiredAt = timeutil.StrFormat(expiredAt.Add(1 * time.Hour))
			}
			c.Set("session", *session)
		}

		if err := m.validateAuthorizedAPI(c, session.BackendPath); err != nil {
			genericError := err.(*errors.GenericError)
			errors.ErrorCode(c, genericError.GetCode(), errors.SetDetails([]constUtil.ErrorDetails{
				{Key: "Authorization", Value: "You don't have authorization to view this page"},
			}))
		}
		c.Next()
	}
}

func (m *MiddlewareDashboardAuth) validateAuthorizedAPI(c *gin.Context, backendPath []context_path.BackendPath) error {
	for _, v := range backendPath {
		if v.Name == c.FullPath() && v.Method == c.Request.Method {
			return nil
		}
	}
	return errors.NewGenericError(errors.FORBIDDEN)
}
