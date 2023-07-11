package branch

import (
	"context"
	"go/skeleton/internal/model"
	"go/skeleton/pkg/clients/db"
)

type BranchRepo interface {
	Create(ctx context.Context, m model.Branch) error
	GetByID(ctx context.Context, ID uint) (*model.Branch, error)
}

type branchRepo struct {
	dbdget db.DBGormDelegate
}

func NewBranchRepo(dbdget db.DBGormDelegate) BranchRepo {
	return &branchRepo{
		dbdget: dbdget,
	}
}
