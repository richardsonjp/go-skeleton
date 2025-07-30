package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-skeleton/internal/services/role"
	"go-skeleton/pkg/utils/api"
	"go-skeleton/pkg/utils/errors"
	"go-skeleton/pkg/utils/validator"
)

func (h *UserHandler) GetListRole(ctx *gin.Context) {
	filter := role.RoleGetFilterByPayload{}
	if err := ctx.ShouldBindWith(&filter, binding.Form); err != nil {
		errors.ErrorString(ctx, validator.GetValidatorMessage(err))
		return
	}
	if errorMessage, err := validator.Validate(filter); err != nil {
		errors.ErrorString(ctx, errorMessage)
		return
	}
	result, err := h.roleService.GetListRole(ctx, filter)
	if err != nil {
		errors.E(ctx, err)
		return
	}
	ctx.JSON(200, api.Base{
		Data: result,
	})
}
