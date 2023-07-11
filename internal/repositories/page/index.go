package page

import (
	"context"
	"go/skeleton/internal/model"
	"go/skeleton/pkg/clients/db"
)

type PageRepo interface {
	Create(ctx context.Context, m model.Page) error
	GetByID(ctx context.Context, ID uint) (*model.Page, error)
}

type pageRepo struct {
	dbdget db.DBGormDelegate
}

func NewPageRepo(dbdget db.DBGormDelegate) PageRepo {
	return &pageRepo{
		dbdget: dbdget,
	}
}
