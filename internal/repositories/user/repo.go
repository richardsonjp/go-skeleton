package user

import (
	"context"
	"go/skeleton/internal/model"
	"go/skeleton/pkg/utils/errors"
	"math"
	"strings"
)

func (r *userRepo) Create(ctx context.Context, m model.User) error {
	if err := r.dbdget.Get(ctx).
		Create(&m).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepo) Update(ctx context.Context, m model.User, updatedFields ...string) (int64, error) {
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

func (r *userRepo) GetByID(ctx context.Context, ID uint) (*model.User, error) {
	user := model.User{}
	user.ID = ID

	query := r.dbdget.Get(ctx).Where(user)

	if err := query.Find(&user).Error; err != nil {
		return nil, err
	}

	if query.RowsAffected == 0 {
		return nil, errors.NewGenericError(errors.DATA_NOT_FOUND)
	}

	return &user, nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	user := model.User{}
	user.Email = email

	query := r.dbdget.Get(ctx).Where("email = ?", email)

	if err := query.Find(&user).Error; err != nil {
		return nil, err
	}

	if query.RowsAffected == 0 {
		return nil, errors.NewGenericError(errors.DATA_NOT_FOUND)
	}

	return &user, nil
}

func (r *userRepo) GetListUser(ctx context.Context, pagination model.Pagination, branchID uint, searchQuery string) ([]*model.User, *model.Pagination, error) {
	var user []*model.User
	query := r.dbdget.Get(ctx).Model(user)

	// search
	if len(strings.Trim(searchQuery, " ")) > 0 {
		searchValue := "%" + searchQuery + "%"
		searchCondition := `(
			name LIKE ?
			OR email LIKE ?
			OR phone_number LIKE ?
		)`
		query.
			Where(searchCondition, searchValue, searchValue, searchValue)
	}

	query.Find(&user, model.User{BranchID: branchID})
	query.Count(&pagination.TotalRows)
	err := query.Scopes(model.NewPaginate(pagination.GetLimit(), pagination.GetPage()).PaginatedResult).Order(pagination.GetSort()).Find(&user).Error
	if err != nil {
		return nil, nil, err
	}

	totalPages := int(math.Ceil(float64(pagination.TotalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return user, &pagination, nil
}
