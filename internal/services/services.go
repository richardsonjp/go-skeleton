package service

import (
	"go-skeleton/internal/services/authentication"
	"go-skeleton/internal/services/context_path"
	"go-skeleton/internal/services/credential"
	"go-skeleton/internal/services/role"
	"go-skeleton/internal/services/user"
)

// put handlers alias
type (
	UserService           = user.UserService
	RoleService           = role.RoleService
	AuthenticationService = authentication.AuthenticationService
	ContextPathService    = context_path.ContextPathService
	CredentialService     = credential.CredentialService
)

var (
	NewUserService           = user.NewUserService
	NewRoleService           = role.NewRoleService
	NewAuthenticationService = authentication.NewAuthenticationService
	NewContextPathService    = context_path.NewContextPathService
	NewCredentialService     = credential.NewCredentialService
)
