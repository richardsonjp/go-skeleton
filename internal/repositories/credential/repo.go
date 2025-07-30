package credential

import (
	"context"
	"go-skeleton/internal/model"
	"go-skeleton/pkg/utils/errors"
)

func (r *credentialRepo) Create(ctx context.Context, m model.Credential) error {
	if err := r.dbdget.Get(ctx).
		Create(&m).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *credentialRepo) Update(ctx context.Context, m model.Credential, updatedFields ...string) (int64, error) {
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

func (r *credentialRepo) GetSecret(ctx context.Context, secret string) (*model.Credential, error) {
	credential := model.Credential{}
	credential.Secret = secret

	query := r.dbdget.Get(ctx).Where(credential)

	if err := query.Find(&credential).Error; err != nil {
		return nil, err
	}

	if query.RowsAffected == 0 {
		return nil, errors.NewGenericError(errors.DATA_NOT_FOUND)
	}

	return &credential, nil
}
