package admin

import (
	"context"
	"go-skeleton/internal/model"
	"go-skeleton/pkg/clients/db"
)

type AdminRepo interface {
	// Common
	Create(ctx context.Context, m *model.Admin) error
	Update(ctx context.Context, param *model.Admin, updatedFields ...string) (int64, error)
}

type adminRepo struct {
	dbdget db.DBGormDelegate
}

func NewAdminRepo(dbdget db.DBGormDelegate) AdminRepo {
	return &adminRepo{
		dbdget: dbdget,
	}
}
