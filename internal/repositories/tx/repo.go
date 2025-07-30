package tx

import (
	"context"
	"gorm.io/gorm"
)

func setContextWithTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, "tx", tx)
}

func getContextTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("tx").(*gorm.DB); ok && tx != nil {
		return tx
	}
	return nil
}

// Run is a function that used to run the service under tx
func (r *txRepo) Run(ctx context.Context, fn func(ctx context.Context) error) error {
	// Check if there's already an active transaction in the context
	existingTx, ok := ctx.Value("tx").(*gorm.DB)
	if ok && existingTx != nil {
		// If a transaction exists, just pass it to the function
		return fn(ctx)
	}

	// No existing transaction, start a new one
	tx := r.dbdget.BeginTx()

	defer func() {
		if p := recover(); p != nil {
			r.dbdget.Rollback(tx) // Rollback on panic
			panic(p)              // Re-throw the panic after rollback
		}
	}()

	newCtx := setContextWithTx(ctx, tx)

	// Execute the user-defined function
	err := fn(newCtx)
	if err != nil {
		// Ensure rollback on any error
		r.dbdget.Rollback(tx)
		return err
	}

	// Commit only if no errors occurred
	if err := r.dbdget.Commit(tx); err != nil {
		r.dbdget.Rollback(tx) // Safeguard rollback if commit fails
		return err
	}

	return nil
}
