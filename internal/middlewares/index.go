package middlewares

import (
	service "go/skeleton/internal/services"
)

type MiddlewareAccess struct {
}

func NewMiddlewareAccess() MiddlewareAccess {
	return MiddlewareAccess{}
}

type MiddlewareDashboardAuth struct {
	authenticationService service.AuthenticationService
	branchService         service.BranchService
}

func NewMiddlewareDashboardAuth(authenticationService service.AuthenticationService, branchService service.BranchService) MiddlewareDashboardAuth {
	return MiddlewareDashboardAuth{
		authenticationService: authenticationService,
		branchService:         branchService,
	}
}
