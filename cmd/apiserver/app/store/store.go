package store

import (
	handler "go-skeleton/cmd/apiserver/app/handlers"
	"go-skeleton/config"
	"go-skeleton/internal/middlewares"
	repos "go-skeleton/internal/repositories"
	"go-skeleton/internal/services"
	"go-skeleton/pkg/clients/db"
	"go-skeleton/pkg/clients/redis"
)

var (
	// handlers
	UserAuthHandler *handler.UserAuthHandler
	UserHandler     *handler.UserHandler
	// services
	AuthenticationService service.AuthenticationService
	ContextPathService    service.ContextPathService
	RoleService           service.RoleService
	CredentialService     service.CredentialService
	UserService           service.UserService

	// repos
	CredentialRepo repos.CredentialRepo
	RoleRepo       repos.RoleRepo
	UserRepo       repos.UserRepo

	ContextPathRepo  repos.ContextPathRepo
	FrontendPathRepo repos.FrontendPathRepo
	BackendPathRepo  repos.BackendPathRepo

	TxRepo repos.TxRepo

	//middleware
	MiddlewareAccess        middlewares.MiddlewareAccess
	MiddlewareDashboardAuth middlewares.MiddlewareDashboardAuth
)

// Init application global variable with single instance
func InitDI() {
	// setup resources
	dbdget := db.NewDBdelegate(config.Config.DB.Debug)
	dbdget.Init()
	redisDel := redis.NewRedisDel()
	redisDel.Init()
	//nullHTTP := http.NewHTTPClient("", nil)

	// repos
	TxRepo = repos.NewTxRepo(dbdget)

	ContextPathRepo = repos.NewContextPathRepo(dbdget)
	FrontendPathRepo = repos.NewFrontendPathRepo(dbdget)
	BackendPathRepo = repos.NewBackendPathRepo(dbdget)

	CredentialRepo = repos.NewCredentialRepo(dbdget)
	RoleRepo = repos.NewRoleRepo(dbdget)
	UserRepo = repos.NewUserRepo(dbdget)

	// services
	ContextPathService = service.NewContextPathService(ContextPathRepo)
	RoleService = service.NewRoleService(TxRepo, RoleRepo, ContextPathService)
	CredentialService = service.NewCredentialService(CredentialRepo)
	UserService = service.NewUserService(UserRepo, RoleService)
	AuthenticationService = service.NewAuthenticationService(TxRepo, UserService, CredentialService, ContextPathService, RoleService)

	// handlers
	UserHandler = handler.NewUserHandler(UserService, RoleService)
	UserAuthHandler = handler.NewUserAuthHandler(AuthenticationService)

	// middleware
	MiddlewareAccess = middlewares.NewMiddlewareAccess()
	MiddlewareDashboardAuth = middlewares.NewMiddlewareDashboardAuth(AuthenticationService)
}
