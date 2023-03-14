package hds

import (
	"go-skeleton/cmd/apiserver/app/handlers/admin"
)

// put handlers alias
type (
	AdminHandler = admin.AdminHandler
)

var (
	NewAdminHandler = admin.NewAdminHandler
)
