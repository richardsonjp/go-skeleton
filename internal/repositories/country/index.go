package country

import (
	"context"
	"go/skeleton/internal/model"
	"go/skeleton/pkg/clients/db"
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
