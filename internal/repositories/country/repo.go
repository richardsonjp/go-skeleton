package country

import (
	"context"
	"mgw/mgw-resi/internal/model"
	"mgw/mgw-resi/pkg/utils/errors"
)

func (r *countryRepo) Create(ctx context.Context, m model.Country) error {
	if err := r.dbdget.Get(ctx).
		Create(&m).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *countryRepo) GetByName(ctx context.Context, name string) (*model.Country, error) {
	country := model.Country{}
	country.Name = name

	query := r.dbdget.Get(ctx).Where(country)

	if err := query.Find(&country).Error; err != nil {
		return nil, err
	}

	if query.RowsAffected == 0 {
		return nil, errors.NewGenericError(errors.DATA_NOT_FOUND)
	}

	return &country, nil
}
