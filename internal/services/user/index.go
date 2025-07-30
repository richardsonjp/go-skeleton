package user

import (
	"context"
	repos "go-skeleton/internal/repositories"
	"go-skeleton/internal/services/role"
	"go-skeleton/pkg/utils/transformer"
)

type UserService interface {
	CreateUser(ctx context.Context, params UserCreatePayload) error
	GetByID(ctx context.Context, userID uint) (*UserResponse, error)
	GetByEmail(ctx context.Context, email string) (*UserResponse, error)
	UpdateLastLogin(ctx context.Context, userID uint) error
	GetListUser(ctx context.Context, filter UserGetFilterPayload) (*transformer.Pagination, error)
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
