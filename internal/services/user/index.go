package user

import (
	"context"
	repos "go/skeleton/internal/repositories"
	"go/skeleton/internal/services/role"
)

type UserService interface {
	CreateUser(ctx context.Context, accessorRole uint, params UserCreatePayload, branchID uint) error
	GetByID(ctx context.Context, userID uint) (*UserResponse, error)
	GetByEmail(ctx context.Context, email string) (*UserResponse, error)
	UpdateLastLogin(ctx context.Context, userID uint) error
}

type userService struct {
	userRepo    repos.UserRepo
	roleService role.RoleService
}

func NewUserService(userRepo repos.UserRepo, roleService role.RoleService) UserService {
	return &userService{
		userRepo:    userRepo,
		roleService: roleService,
	}
}
