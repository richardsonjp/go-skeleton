package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go/skeleton/internal/services/authentication"
	"go/skeleton/internal/services/user"
	"go/skeleton/pkg/utils/api"
	"go/skeleton/pkg/utils/errors"
	"go/skeleton/pkg/utils/validator"
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
