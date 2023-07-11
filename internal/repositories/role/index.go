package role

import (
	"context"
	"go/skeleton/internal/model"
	"go/skeleton/pkg/clients/db"
)

type RoleRepo interface {
	Create(ctx context.Context, m model.Role) error
	GetByID(ctx context.Context, ID uint) (*model.Role, error)
}

type roleRepo struct {
	dbdget db.DBGormDelegate
}

func NewRoleRepo(dbdget db.DBGormDelegate) RoleRepo {
	return &roleRepo{
		dbdget: dbdget,
	}
}
