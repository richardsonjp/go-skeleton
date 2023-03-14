package store

import (
	hdrs "go-skeleton/cmd/apiserver/app/handlers"
	"go-skeleton/config"
	"go-skeleton/internal/middlewares"
	repos "go-skeleton/internal/repositories"
	srvs "go-skeleton/internal/services"
	"go-skeleton/pkg/clients/db"
	"go-skeleton/pkg/clients/redis"
)

var (
	// hdrs
	AdminHandler *hdrs.AdminHandler

	// srvs
	AdminService srvs.AdminService

	// repos
	AdminRepo repos.AdminRepo
	TxRepo    repos.TxRepo

	//middleware
	MiddlewareAccess middlewares.MiddlewareAccess
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
	AdminRepo = repos.NewAdminRepo(dbdget)
	TxRepo = repos.NewTxRepo(dbdget)

	// services
	AdminService = srvs.NewAdminService(AdminRepo)

	// hdrs
	AdminHandler = hdrs.NewAdminHandler(AdminService)

	// middleware
	MiddlewareAccess = middlewares.NewMiddlewareAccess(redisDel)
}
