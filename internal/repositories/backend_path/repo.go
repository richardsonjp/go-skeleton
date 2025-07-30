package backend_path

import (
	"context"
	"go-skeleton/internal/model"
	"go-skeleton/pkg/utils/errors"
)

func (r *backendPathRepo) Create(ctx context.Context, m model.BackendPath) error {
	if err := r.dbdget.Get(ctx).
		Create(&m).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *backendPathRepo) GetByID(ctx context.Context, ID uint) (*model.BackendPath, error) {
	backendPath := model.BackendPath{}
	backendPath.ID = ID

	query := r.dbdget.Get(ctx).Where(backendPath)

	if err := query.Find(&backendPath).Error; err != nil {
		return nil, err
	}

	if query.RowsAffected == 0 {
		return nil, errors.NewGenericError(errors.DATA_NOT_FOUND)
	}

	return &backendPath, nil
}
