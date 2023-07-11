package handler

import (
	"go/skeleton/cmd/apiserver/app/handlers/user"
	"go/skeleton/cmd/apiserver/app/handlers/user_auth"
)

// put handlers alias
type (
	UserAuthHandler = user_auth.UserAuthHandler
	UserHandler     = user.UserHandler
)

var (
	NewUserAuthHandler = user_auth.NewUserAuthHandler
	NewUserHandler     = user.NewUserHandler
)
