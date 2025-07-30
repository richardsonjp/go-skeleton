package backend_path

import (
	"context"
	"go-skeleton/internal/model"
	"go-skeleton/pkg/clients/db"
)

type BackendPathRepo interface {
	Create(ctx context.Context, m model.BackendPath) error
	GetByID(ctx context.Context, ID uint) (*model.BackendPath, error)
}

type backendPathRepo struct {
	dbdget db.DBGormDelegate
}

func NewBackendPathRepo(dbdget db.DBGormDelegate) BackendPathRepo {
	return &backendPathRepo{
		dbdget: dbdget,
	}
}
