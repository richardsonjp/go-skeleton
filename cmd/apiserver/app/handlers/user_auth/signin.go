package user_auth

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-skeleton/internal/services/authentication"
	"go-skeleton/pkg/utils/api"
	"go-skeleton/pkg/utils/errors"
	"go-skeleton/pkg/utils/validator"
)

func (h *UserAuthHandler) Signin(ctx *gin.Context) {
	param := &authentication.Signin{}
	if err := ctx.ShouldBindWith(param, binding.JSON); err != nil {
		errors.ErrorString(ctx, validator.GetValidatorMessage(err))
		return
	}
	if errorMessage, err := validator.Validate(param); err != nil {
		errors.ErrorString(ctx, errorMessage)
		return
	}

	resultData, err := h.authenticationService.AuthenticateSignin(ctx, *param)
	if err != nil {
		errors.E(ctx, err)
		return
	}

	ctx.JSON(200, api.Base{
		Data: resultData,
	})
}
