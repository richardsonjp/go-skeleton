package page

import (
	"context"
	"mgw/mgw-resi/internal/model"
	"mgw/mgw-resi/pkg/utils/errors"
)

func (r *pageRepo) Create(ctx context.Context, m model.Page) error {
	if err := r.dbdget.Get(ctx).
		Create(&m).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *pageRepo) GetByID(ctx context.Context, ID uint) (*model.Page, error) {
	page := model.Page{}
	page.ID = ID

	query := r.dbdget.Get(ctx).Where(page)

	if err := query.Find(&page).Error; err != nil {
		return nil, err
	}

	if query.RowsAffected == 0 {
		return nil, errors.NewGenericError(errors.DATA_NOT_FOUND)
	}

	return &page, nil
}
