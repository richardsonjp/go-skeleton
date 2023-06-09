package user

import (
	"context"
	"mgw/mgw-resi/internal/model"
	"mgw/mgw-resi/pkg/clients/db"
)

type UserRepo interface {
	Create(ctx context.Context, m model.User) error
	Update(ctx context.Context, m model.User, updatedFields ...string) (int64, error)
	GetByID(ctx context.Context, ID uint) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetListUser(ctx context.Context, pagination model.Pagination, branchID uint, searchQuery string) ([]*model.User, *model.Pagination, error)
}

type userRepo struct {
	dbdget db.DBGormDelegate
}

func NewUserRepo(dbdget db.DBGormDelegate) UserRepo {
	return &userRepo{
		dbdget: dbdget,
	}
}