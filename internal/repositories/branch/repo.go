package branch

import (
	"context"
	"go/skeleton/internal/model"
	"go/skeleton/pkg/utils/errors"
)

func (r *branchRepo) Create(ctx context.Context, m model.Branch) error {
	if err := r.dbdget.Get(ctx).
		Create(&m).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *branchRepo) GetByID(ctx context.Context, ID uint) (*model.Branch, error) {
	branch := model.Branch{}
	branch.ID = ID

	query := r.dbdget.Get(ctx).Where(branch)

	if err := query.Find(&branch).Error; err != nil {
		return nil, err
	}

	if query.RowsAffected == 0 {
		return nil, errors.NewGenericError(errors.DATA_NOT_FOUND)
	}

	return &branch, nil
}
