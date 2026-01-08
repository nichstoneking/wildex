package main

import (
	"fmt"
	"log"

	"github.com/crest/wildex/server/configs"
	"github.com/crest/wildex/server/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config := configs.LoadConfig()

	if config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	routes.SetupRoutes(router)

	serverAddress := fmt.Sprintf(":%s", config.ServerPort)
	log.Printf("Starting server on port %s in %s mode...", config.ServerPort, config.Environment)

	if err := router.Run(serverAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
