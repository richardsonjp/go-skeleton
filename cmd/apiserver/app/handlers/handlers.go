package hds

import (
	"go-skeleton/cmd/apiserver/app/handlers/user"
	"go-skeleton/cmd/apiserver/app/handlers/user_auth"
)

// put handlers alias
type (
	UserHandler     = user.UserHandler
	UserAuthHandler = user_auth.UserAuthHandler
)

var (
	NewUserHandler     = user.NewUserHandler
	NewUserAuthHandler = user_auth.NewUserAuthHandler
)
