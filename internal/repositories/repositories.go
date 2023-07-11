package repos

import (
	"go/skeleton/internal/repositories/branch"
	"go/skeleton/internal/repositories/context"
	"go/skeleton/internal/repositories/country"
	"go/skeleton/internal/repositories/credential"
	"go/skeleton/internal/repositories/misc"
	"go/skeleton/internal/repositories/page"
	"go/skeleton/internal/repositories/role"
	"go/skeleton/internal/repositories/tx"
	"go/skeleton/internal/repositories/user"
)

// put repos alias
type (
	BranchRepo     = branch.BranchRepo
	ContextRepo    = context.ContextRepo
	CountryRepo    = country.CountryRepo
	CredentialRepo = credential.CredentialRepo
	MiscRepo       = misc.MiscRepo
	PageRepo       = page.PageRepo
	RoleRepo       = role.RoleRepo
	UserRepo       = user.UserRepo

	TxRepo = tx.TxRepo
)

var (
	NewBranchRepo     = branch.NewBranchRepo
	NewContextRepo    = context.NewContextRepo
	NewCountryRepo    = country.NewCountryRepo
	NewCredentialRepo = credential.NewCredentialRepo
	NewMiscRepo       = misc.NewMiscRepo
	NewPageRepo       = page.NewPageRepo
	NewRoleRepo       = role.NewRoleRepo
	NewUserRepo       = user.NewUserRepo

	NewTxRepo = tx.NewTxRepo
)
