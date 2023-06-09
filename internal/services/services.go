package service

import (
	"mgw/mgw-resi/internal/services/authentication"
	"mgw/mgw-resi/internal/services/branch"
	"mgw/mgw-resi/internal/services/credential"
	"mgw/mgw-resi/internal/services/role"
	"mgw/mgw-resi/internal/services/user"
)

// put services alias
type (
	AuthenticationService = authentication.AuthenticationService
	BranchService         = branch.BranchService
	CredentialService     = credential.CredentialService
	UserService           = user.UserService
	RoleService           = role.RoleService
)

var (
	NewAuthenticationService = authentication.NewAuthenticationService
	NewBranchService         = branch.NewBranchService
	NewCredentialService     = credential.NewCredentialService
	NewUserService           = user.NewUserService
	NewRoleService           = role.NewRoleService
)
