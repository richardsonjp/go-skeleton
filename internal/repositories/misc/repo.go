package misc

import (
	"context"
	"go/skeleton/internal/model"
	"go/skeleton/pkg/utils/errors"
)

func (r *miscRepo) Create(ctx context.Context, m model.Misc) error {
	if err := r.dbdget.Get(ctx).
		Create(&m).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *miscRepo) GetByName(ctx context.Context, name string) (*model.Misc, error) {
	misc := model.Misc{}
	misc.Name = name

	query := r.dbdget.Get(ctx).Where(misc)

	if err := query.Find(&misc).Error; err != nil {
		return nil, err
	}

	if query.RowsAffected == 0 {
		return nil, errors.NewGenericError(errors.DATA_NOT_FOUND)
	}

	return &misc, nil
}
