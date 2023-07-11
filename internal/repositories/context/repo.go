package context

import (
	"context"
	"go/skeleton/internal/model"
	"go/skeleton/pkg/utils/errors"
)

func (r *contextRepo) Create(ctx context.Context, m model.Context) error {
	if err := r.dbdget.Get(ctx).
		Create(&m).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *contextRepo) GetByID(ctx context.Context, ID uint) (*model.Context, error) {
	contextm := model.Context{}
	contextm.ID = ID

	query := r.dbdget.Get(ctx).Where(contextm)

	if err := query.Find(&contextm).Error; err != nil {
		return nil, err
	}

	if query.RowsAffected == 0 {
		return nil, errors.NewGenericError(errors.DATA_NOT_FOUND)
	}

	return &contextm, nil
}
