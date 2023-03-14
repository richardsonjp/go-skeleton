package repos

import (
	"go-skeleton/internal/repositories/admin"
	"go-skeleton/internal/repositories/tx"
)

// put repos alias
type (
	AdminRepo = admin.AdminRepo
	TxRepo    = tx.TxRepo
)

var (
	NewAdminRepo = admin.NewAdminRepo
	NewTxRepo    = tx.NewTxRepo
)
