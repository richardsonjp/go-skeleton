package routes

import (
	"go-skeleton/internal/middlewares"

	"github.com/gin-gonic/gin"
)

// Init initialize application routes
// DO NOT FORGET TO READ README#routes FIRST!
func Init(mode string) *gin.Engine {
	gin.SetMode(mode)
	r := gin.Default()
	r.HandleMethodNotAllowed = true

	// Setup middlewares
	r.Use(middlewares.Recovery(mode))
	r.Use(middlewares.AccessLog())
	r.Use(middlewares.LanguageAccept())
	r.Use(middlewares.Cors())

	// Setup pingpong
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Setup sample route
	sampleDashboardRouteGroup := r.Group("/dashboard")
	initSampleRoute(sampleDashboardRouteGroup)

	return r
}
