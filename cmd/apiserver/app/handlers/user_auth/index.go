package user_auth

import (
	service "mgw/mgw-resi/internal/services"
)

type UserAuthHandler struct {
	authenticationService service.AuthenticationService
}

func NewUserAuthHandler(authenticationService service.AuthenticationService) *UserAuthHandler {
	return &UserAuthHandler{
		authenticationService: authenticationService,
	}
}
