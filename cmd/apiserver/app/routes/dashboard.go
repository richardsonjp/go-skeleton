package routes

import (
	"github.com/gin-gonic/gin"
	"mgw/mgw-resi/cmd/apiserver/app/store"
)

func initDashboardRoute(group *gin.RouterGroup) {
	group.POST("/user/register", store.UserHandler.CreateUser)
}
