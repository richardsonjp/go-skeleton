package user_auth

import (
	"github.com/gin-gonic/gin"
	"go-skeleton/internal/services/authentication"
	"go-skeleton/pkg/utils/api"
	"go-skeleton/pkg/utils/errors"
	"net/http"
)

func (h *UserAuthHandler) FrontendPath(ctx *gin.Context) {
	sessionContext := ctx.Value("session")
	session, sessionOk := sessionContext.(authentication.SessionData)
	if !sessionOk {
		errors.ErrorCode(ctx, http.StatusUnauthorized)
		return
	}

	path, err := h.authenticationService.FrontendPath(ctx, session)
	if err != nil {
		errors.E(ctx, err)
		return
	}

	ctx.JSON(200, api.Base{
		Data: path,
	})
}
