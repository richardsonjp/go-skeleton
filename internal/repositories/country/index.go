package country

import (
	"context"
	"mgw/mgw-resi/internal/model"
	"mgw/mgw-resi/pkg/clients/db"
)

type CountryRepo interface {
	Create(ctx context.Context, m model.Country) error
	GetByName(ctx context.Context, name string) (*model.Country, error)
}

type countryRepo struct {
	dbdget db.DBGormDelegate
}

func NewCountryRepo(dbdget db.DBGormDelegate) CountryRepo {
	return &countryRepo{
		dbdget: dbdget,
	}
}
