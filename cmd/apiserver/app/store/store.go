package store

import (
	handler "go/skeleton/cmd/apiserver/app/handlers"
	"go/skeleton/config"
	"go/skeleton/internal/middlewares"
	repos "go/skeleton/internal/repositories"
	service "go/skeleton/internal/services"
	"go/skeleton/pkg/clients/db"
)

var (
	// handlers
	UserAuthHandler *handler.UserAuthHandler
	UserHandler     *handler.UserHandler

	// services
	AuthenticationService service.AuthenticationService
	BranchService         service.BranchService
	RoleService           service.RoleService
	CredentialService     service.CredentialService
	UserService           service.UserService

	// repos
	BranchRepo     repos.BranchRepo
	ContextRepo    repos.ContextRepo
	CountryRepo    repos.CountryRepo
	CredentialRepo repos.CredentialRepo
	MiscRepo       repos.MiscRepo
	PageRepo       repos.PageRepo
	RoleRepo       repos.RoleRepo
	UserRepo       repos.UserRepo

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
	//redisDel := redis.NewRedisDel()
	//redisDel.Init()

	// repos
	TxRepo = repos.NewTxRepo(dbdget)
	BranchRepo = repos.NewBranchRepo(dbdget)
	ContextRepo = repos.NewContextRepo(dbdget)
	CountryRepo = repos.NewCountryRepo(dbdget)
	CredentialRepo = repos.NewCredentialRepo(dbdget)
	MiscRepo = repos.NewMiscRepo(dbdget)
	PageRepo = repos.NewPageRepo(dbdget)
	RoleRepo = repos.NewRoleRepo(dbdget)
	UserRepo = repos.NewUserRepo(dbdget)

	// services
	BranchService = service.NewBranchService(BranchRepo)
	RoleService = service.NewRoleService(RoleRepo)
	CredentialService = service.NewCredentialService(CredentialRepo)
	UserService = service.NewUserService(UserRepo, RoleService)
	AuthenticationService = service.NewAuthenticationService(TxRepo, UserService, CredentialService, BranchService)

	// handlers
	UserAuthHandler = handler.NewUserAuthHandler(AuthenticationService)
	UserHandler = handler.NewUserHandler(UserService)

	// middlewares
	MiddlewareAccess = middlewares.NewMiddlewareAccess()
	MiddlewareDashboardAuth = middlewares.NewMiddlewareDashboardAuth(AuthenticationService, BranchService)
}
