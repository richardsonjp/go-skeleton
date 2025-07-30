package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-skeleton/internal/services/user"
	"go-skeleton/pkg/utils/api"
	"go-skeleton/pkg/utils/errors"
	"go-skeleton/pkg/utils/validator"
)

func (h *UserHandler) GetListUser(ctx *gin.Context) {
	filter := user.UserGetFilterPayload{}
	if err := ctx.ShouldBindWith(&filter, binding.Form); err != nil {
		errors.ErrorString(ctx, validator.GetValidatorMessage(err))
		return
	}
	if errorMessage, err := validator.Validate(filter); err != nil {
		errors.ErrorString(ctx, errorMessage)
		return
	}

	result, err := h.userService.GetListUser(ctx, filter)
	if err != nil {
		errors.E(ctx, err)
		return
	}

	ctx.JSON(200, api.Base{
		Data: result,
	})
}
