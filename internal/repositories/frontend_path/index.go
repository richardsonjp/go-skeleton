package frontend_path

import (
	"context"
	"go-skeleton/internal/model"
	"go-skeleton/pkg/clients/db"
)

type FrontendPathRepo interface {
	Create(ctx context.Context, m model.FrontendPath) error
	GetByID(ctx context.Context, ID uint) (*model.FrontendPath, error)
}

type frontendPathRepo struct {
	dbdget db.DBGormDelegate
}

func NewFrontendPathRepo(dbdget db.DBGormDelegate) FrontendPathRepo {
	return &frontendPathRepo{
		dbdget: dbdget,
	}
}
