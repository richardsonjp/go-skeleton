package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-skeleton/internal/services/role"
	"go-skeleton/pkg/utils/api"
	"go-skeleton/pkg/utils/errors"
	"go-skeleton/pkg/utils/validator"
)

func (h *UserHandler) UpdateRole(ctx *gin.Context) {
	param := role.RoleUpdatePayload{}
	if err := ctx.ShouldBindWith(&param, binding.JSON); err != nil {
		errors.ErrorString(ctx, validator.GetValidatorMessage(err))
		return
	}
	if errorMessage, err := validator.Validate(param); err != nil {
		errors.ErrorString(ctx, errorMessage)
		return
	}

	err := h.roleService.UpdateRole(ctx, param)
	if err != nil {
		errors.E(ctx, err)
		return
	}

	ctx.JSON(200, api.Message{
		Message: "success",
	})
}
