package tx

import (
	"context"
	"mgw/mgw-resi/pkg/clients/db"
)

type TxRepo interface {
	Run(ctx context.Context, fn func(ctx context.Context) error) error
}

type txRepo struct {
	dbdget db.DBGormDelegate
}

func NewTxRepo(dbdget db.DBGormDelegate) TxRepo {
	return &txRepo{
		dbdget: dbdget,
	}
}
