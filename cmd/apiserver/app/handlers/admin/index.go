package admin

import (
	srvs "go-skeleton/internal/services"
)

type AdminHandler struct {
	adminService srvs.AdminService
}

func NewAdminHandler(adminService srvs.AdminService) *AdminHandler {
	return &AdminHandler{
		adminService: adminService,
	}
}
