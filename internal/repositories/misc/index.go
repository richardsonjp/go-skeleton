package misc

import (
	"context"
	"go/skeleton/internal/model"
	"go/skeleton/pkg/clients/db"
)

type MiscRepo interface {
	Create(ctx context.Context, m model.Misc) error
	GetByName(ctx context.Context, name string) (*model.Misc, error)
}

type miscRepo struct {
	dbdget db.DBGormDelegate
}

func NewMiscRepo(dbdget db.DBGormDelegate) MiscRepo {
	return &miscRepo{
		dbdget: dbdget,
	}
}
