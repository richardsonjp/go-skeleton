package admin

import (
	"context"
	"go-skeleton/internal/model"
)

func (r *adminRepo) Create(ctx context.Context, m *model.Admin) error {
	if err := r.dbdget.Get(ctx).
		Create(m).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *adminRepo) Update(ctx context.Context, param *model.Admin, updatedFields ...string) (int64, error) {
	query := r.dbdget.Get(ctx)
	if len(updatedFields) > 0 {
		updatedFields = append(updatedFields, "updated_at")
		query = query.Select(updatedFields)
	}

	query.Updates(param)

	// execute query
	if err := query.Find(&param).Error; err != nil {
		return 0, err
	}

	return query.RowsAffected, nil
}
