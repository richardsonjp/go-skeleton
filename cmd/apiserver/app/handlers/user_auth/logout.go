package user_auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/skeleton/internal/services/authentication"
	"go/skeleton/pkg/utils/api"
	"go/skeleton/pkg/utils/errors"
	"net/http"
)

func (h *UserAuthHandler) Logout(ctx *gin.Context) {
	sessionContext := ctx.Value("session")
	fmt.Println(sessionContext)
	session, sessionOk := sessionContext.(authentication.SessionData)
	if !sessionOk {
		errors.ErrorCode(ctx, http.StatusUnauthorized)
		return
	}

	err := h.authenticationService.AuthenticateLogout(ctx, session)
	if err != nil {

		errors.E(ctx, err)
		return
	}

	ctx.JSON(200, api.Message{
		Message: "success",
	})
}
