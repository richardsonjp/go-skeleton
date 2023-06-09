package context

import (
	"context"
	"mgw/mgw-resi/internal/model"
	"mgw/mgw-resi/pkg/clients/db"
)

type ContextRepo interface {
	Create(ctx context.Context, m model.Context) error
	GetByID(ctx context.Context, ID uint) (*model.Context, error)
}

type contextRepo struct {
	dbdget db.DBGormDelegate
}

func NewContextRepo(dbdget db.DBGormDelegate) ContextRepo {
	return &contextRepo{
		dbdget: dbdget,
	}
}
