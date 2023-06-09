package role

import (
	"context"
	"mgw/mgw-resi/internal/model"
	"mgw/mgw-resi/pkg/utils/errors"
)

func (r *roleRepo) Create(ctx context.Context, m model.Role) error {
	if err := r.dbdget.Get(ctx).
		Create(&m).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *roleRepo) GetByID(ctx context.Context, ID uint) (*model.Role, error) {
	role := model.Role{}
	role.ID = ID

	query := r.dbdget.Get(ctx).Where(role)

	if err := query.Find(&role).Error; err != nil {
		return nil, err
	}

	if query.RowsAffected == 0 {
		return nil, errors.NewGenericError(errors.DATA_NOT_FOUND)
	}

	return &role, nil
}
