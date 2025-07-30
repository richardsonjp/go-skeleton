package user_auth

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-skeleton/internal/services/authentication"
	"go-skeleton/pkg/utils/api"
	"go-skeleton/pkg/utils/errors"
	"go-skeleton/pkg/utils/validator"
)

func (h *UserAuthHandler) ListRBAC(ctx *gin.Context) {
	param := &authentication.ListRBAC{}
	if err := ctx.ShouldBindWith(param, binding.JSON); err != nil {
		errors.ErrorString(ctx, validator.GetValidatorMessage(err))
		return
	}

	path, err := h.authenticationService.GetListRBAC(ctx, param.RoleID)
	if err != nil {
		errors.E(ctx, err)
		return
	}

	ctx.JSON(200, api.Base{
		Data: path,
	})
}
