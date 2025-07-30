package repos

import (
	"go-skeleton/internal/repositories/backend_path"
	"go-skeleton/internal/repositories/context_path"
	"go-skeleton/internal/repositories/credential"
	"go-skeleton/internal/repositories/frontend_path"
	"go-skeleton/internal/repositories/role"
	"go-skeleton/internal/repositories/tx"
	"go-skeleton/internal/repositories/user"
)

// put repos alias
type (
	UserRepo         = user.UserRepo
	RoleRepo         = role.RoleRepo
	CredentialRepo   = credential.CredentialRepo
	BackendPathRepo  = backend_path.BackendPathRepo
	FrontendPathRepo = frontend_path.FrontendPathRepo
	ContextPathRepo  = context_path.ContextPathRepo

	TxRepo = tx.TxRepo
)

var (
	NewUserRepo         = user.NewUserRepo
	NewRoleRepo         = role.NewRoleRepo
	NewCredentialRepo   = credential.NewCredentialRepo
	NewBackendPathRepo  = backend_path.NewBackendPathRepo
	NewFrontendPathRepo = frontend_path.NewFrontendPathRepo
	NewContextPathRepo  = context_path.NewContextPathRepo

	NewTxRepo = tx.NewTxRepo
)
