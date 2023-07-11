package role

import (
	"context"
	repos "go/skeleton/internal/repositories"
)

type RoleService interface {
	CreateRole(ctx context.Context, name string) error
	GetByID(ctx context.Context, ID uint) (*RoleResponse, error)
}

type roleService struct {
	roleRepo repos.RoleRepo
}

func NewRoleService(roleRepo repos.RoleRepo) RoleService {
	return &roleService{
		roleRepo: roleRepo,
	}
}
