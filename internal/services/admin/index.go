package admin

import (
	repos "go-skeleton/internal/repositories"
)

type AdminService interface {
}

type adminService struct {
	adminRepo repos.AdminRepo
}

func NewAdminService(adminRepo repos.AdminRepo) AdminService {
	return &adminService{
		adminRepo: adminRepo,
	}
}
