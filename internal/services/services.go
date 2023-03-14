package service

import (
	"go-skeleton/internal/services/admin"
)

// put handlers alias
type (
	AdminService = admin.AdminService
)

var (
	NewAdminService = admin.NewAdminService
)
