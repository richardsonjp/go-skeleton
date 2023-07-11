package service

import (
	"go/skeleton/internal/services/authentication"
	"go/skeleton/internal/services/branch"
	"go/skeleton/internal/services/credential"
	"go/skeleton/internal/services/role"
	"go/skeleton/internal/services/user"
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
