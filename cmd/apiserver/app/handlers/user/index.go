package user

import "go-skeleton/internal/services"

type UserHandler struct {
	userService service.UserService
	roleService service.RoleService
}

func NewUserHandler(userService service.UserService, roleService service.RoleService) *UserHandler {
	return &UserHandler{
		userService: userService,
		roleService: roleService,
	}
}
