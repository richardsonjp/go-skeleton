package middlewares

import (
	service "go-skeleton/internal/services"
)

type MiddlewareAccess struct {
}

func NewMiddlewareAccess() MiddlewareAccess {
	return MiddlewareAccess{}
}

type MiddlewareDashboardAuth struct {
	authenticationService service.AuthenticationService
}

func NewMiddlewareDashboardAuth(authenticationService service.AuthenticationService) MiddlewareDashboardAuth {
	return MiddlewareDashboardAuth{
		authenticationService: authenticationService,
	}
}
