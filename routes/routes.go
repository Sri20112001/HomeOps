package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"health": true,
	})
}

// SetupRoutes registers all application routes
func SetupRoutes(router *gin.Engine) {
	// Root level routes
	router.GET("/health", health)

	// API versioning / grouping (optional but standard practice)
	api := router.Group("/api")
	{
		UserRoutes(api)
	}
}
