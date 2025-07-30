package routes

import (
	"github.com/gin-gonic/gin"
	"go-skeleton/cmd/apiserver/app/store"
)

func initDashboardRoute(group *gin.RouterGroup) {
	group.GET("/user", store.UserHandler.GetListUser)
	group.POST("/user", store.UserHandler.CreateUser)
	group.GET("/user/roles", store.UserHandler.GetListRole)
	group.POST("/user/roles", store.UserHandler.CreateRole)
	group.PUT("/user/roles", store.UserHandler.UpdateRole)
	group.POST("/user/roles/rbac", store.UserAuthHandler.ListRBAC)
}
