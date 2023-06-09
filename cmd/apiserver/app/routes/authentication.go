package routes

import (
	"github.com/gin-gonic/gin"
	"mgw/mgw-resi/cmd/apiserver/app/store"
)

func initAuthenticationRoute(group *gin.RouterGroup) {
	group.POST("/signin", store.UserAuthHandler.Signin)
	group.POST("/logout", store.MiddlewareDashboardAuth.AuthenticateUser(), store.UserAuthHandler.Logout)
}
