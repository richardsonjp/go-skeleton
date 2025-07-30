package routes

import (
	"github.com/gin-gonic/gin"
	"go-skeleton/cmd/apiserver/app/store"
)

func initAuthenticationRoute(group *gin.RouterGroup) {
	group.POST("/signin", store.UserAuthHandler.Signin)
	group.POST("/logout", store.MiddlewareDashboardAuth.AuthenticateUser(), store.UserAuthHandler.Logout)
	group.GET("/frontend-path", store.MiddlewareDashboardAuth.AuthenticateUser(), store.UserAuthHandler.FrontendPath)
	group.GET("/home-profile", store.MiddlewareDashboardAuth.AuthenticateUser(), store.UserAuthHandler.GetHomeProfile)
}
