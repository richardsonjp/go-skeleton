package handler

import (
	"mgw/mgw-resi/cmd/apiserver/app/handlers/user"
	"mgw/mgw-resi/cmd/apiserver/app/handlers/user_auth"
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
