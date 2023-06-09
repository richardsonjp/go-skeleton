package role

import (
	"context"
	repos "mgw/mgw-resi/internal/repositories"
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
