package frontend_path

import (
	"context"
	"go-skeleton/internal/model"
	"go-skeleton/pkg/utils/errors"
)

func (r *frontendPathRepo) Create(ctx context.Context, m model.FrontendPath) error {
	if err := r.dbdget.Get(ctx).
		Create(&m).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *frontendPathRepo) GetByID(ctx context.Context, ID uint) (*model.FrontendPath, error) {
	frontendPath := model.FrontendPath{}
	frontendPath.ID = ID

	query := r.dbdget.Get(ctx).Where(frontendPath)

	if err := query.Find(&frontendPath).Error; err != nil {
		return nil, err
	}

	if query.RowsAffected == 0 {
		return nil, errors.NewGenericError(errors.DATA_NOT_FOUND)
	}

	return &frontendPath, nil
}
