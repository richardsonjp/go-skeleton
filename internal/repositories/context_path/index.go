package context_path

import (
	"context"
	"go-skeleton/internal/model"
	"go-skeleton/pkg/clients/db"
)

type ContextPathRepo interface {
	Create(ctx context.Context, m model.ContextPath) error
	GetByID(ctx context.Context, ID uint) (*model.ContextPath, error)
	GetByContextTag(ctx context.Context, roleID uint, contextTag string) ([]*PathList, error)
	DeleteByRoleID(ctx context.Context, roleID uint) error
	GetListRBAC(ctx context.Context) ([]ListRBAC, error)
	GetRBACByRoleID(ctx context.Context, roleID uint) ([]ListRBAC, error)
	CreateBackendRBAC(ctx context.Context, roleID uint, pageGroupToLabels map[string][]string) error
	CreateFrontendRBAC(ctx context.Context, roleID uint, pageGroups []string) error
}

type contextPathRepo struct {
	dbdget db.DBGormDelegate
}

func NewContextPathRepo(dbdget db.DBGormDelegate) ContextPathRepo {
	return &contextPathRepo{
		dbdget: dbdget,
	}
}
