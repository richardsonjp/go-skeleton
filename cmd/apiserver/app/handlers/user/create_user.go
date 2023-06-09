package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"mgw/mgw-resi/internal/services/authentication"
	"mgw/mgw-resi/internal/services/user"
	"mgw/mgw-resi/pkg/utils/api"
	"mgw/mgw-resi/pkg/utils/errors"
	"mgw/mgw-resi/pkg/utils/validator"
	"net/http"
)

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	sessionContext := ctx.Value("session")
	session, sessionOk := sessionContext.(authentication.SessionData)
	if !sessionOk {
		errors.ErrorCode(ctx, http.StatusUnauthorized)
		return
	}

	param := &user.UserCreatePayload{}
	if err := ctx.ShouldBindWith(param, binding.JSON); err != nil {
		errors.ErrorString(ctx, validator.GetValidatorMessage(err))
		return
	}

	err := h.userService.CreateUser(ctx, session.User.RoleID, *param, session.Branch.ID)
	if err != nil {
		errors.E(ctx, err)
		return
	}

	ctx.JSON(200, api.Message{
		Message: "success",
	})
}
