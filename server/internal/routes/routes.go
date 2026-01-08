package routes

import (
	"github.com/crest/wildex/server/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/health", handlers.HealthHandler)

	api := router.Group("/api")
	{
		api.GET("/hello", handlers.HelloWorldHandler)
	}
}
