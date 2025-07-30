package role

import (
	"context"
	"go-skeleton/internal/model"
	"go-skeleton/pkg/clients/db"
)

type RoleRepo interface {
	Create(ctx context.Context, m model.Role) (*model.Role, error)
	Update(ctx context.Context, m model.Role, updatedFields ...string) (int64, error)
	GetByID(ctx context.Context, ID uint) (*model.Role, error)
	GetListRole(ctx context.Context, pagination model.Pagination) ([]*model.Role, *model.Pagination, error)
}

type roleRepo struct {
	dbdget db.DBGormDelegate
}

func NewRoleRepo(dbdget db.DBGormDelegate) RoleRepo {
	return &roleRepo{
		dbdget: dbdget,
	}
}
