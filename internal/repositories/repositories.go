package repos

import (
	"mgw/mgw-resi/internal/repositories/branch"
	"mgw/mgw-resi/internal/repositories/context"
	"mgw/mgw-resi/internal/repositories/country"
	"mgw/mgw-resi/internal/repositories/credential"
	"mgw/mgw-resi/internal/repositories/misc"
	"mgw/mgw-resi/internal/repositories/page"
	"mgw/mgw-resi/internal/repositories/role"
	"mgw/mgw-resi/internal/repositories/tx"
	"mgw/mgw-resi/internal/repositories/user"
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
