package tx

import (
	"context"

	"gorm.io/gorm"
)

func setContextWithTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, "tx", tx)
}

// Run is a function that used to run the service under tx
func (r *txRepo) Run(ctx context.Context, fn func(ctx context.Context) error) error {
	tx := r.dbdget.BeginTx()
	// defer todo: klo panic rollback

	// set tx to context
	newCtx := setContextWithTx(ctx, tx)

	// execute function
	err := fn(newCtx)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
