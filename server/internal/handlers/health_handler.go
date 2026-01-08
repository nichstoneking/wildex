package handlers

import (
	"net/http"

	"github.com/crest/wildex/server/internal/models"
	"github.com/gin-gonic/gin"
)

func HealthHandler(c *gin.Context) {
	response := models.NewSuccessResponse("Server is healthy", gin.H{
		"status": "ok",
	})
	c.JSON(http.StatusOK, response)
}

func HelloWorldHandler(c *gin.Context) {
	response := models.NewSuccessResponse("Hello, World!", nil)
	c.JSON(http.StatusOK, response)
}
