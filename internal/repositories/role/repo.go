package role

import (
	"context"
	"go-skeleton/internal/model"
	"go-skeleton/pkg/utils/errors"
	"math"
)

func (r *roleRepo) Create(ctx context.Context, m model.Role) (*model.Role, error) {
	if err := r.dbdget.Get(ctx).
		Create(&m).
		Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *roleRepo) Update(ctx context.Context, m model.Role, updatedFields ...string) (int64, error) {
	query := r.dbdget.Get(ctx)
	if len(updatedFields) > 0 {
		updatedFields = append(updatedFields, "updated_at")
		query = query.Select(updatedFields)
	}
	query.Updates(m)

	// execute query
	if err := query.Find(&m).Error; err != nil {
		return 0, err
	}
	return query.RowsAffected, nil
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

func (r *roleRepo) GetListRole(ctx context.Context, pagination model.Pagination) ([]*model.Role, *model.Pagination, error) {
	var roleList []*model.Role
	q := r.dbdget.Get(ctx).Model(model.Role{})
	q.Count(&pagination.TotalRows)
	err := q.Scopes(model.NewPaginate(pagination.GetLimit(), pagination.GetPage()).PaginatedResult).
		Order("role." + pagination.GetSort()).
		Find(&roleList).Error

	if err != nil {
		return nil, nil, err
	}

	totalPages := int(math.Ceil(float64(pagination.TotalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return roleList, &pagination, nil
}
