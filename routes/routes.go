package routes

import (
	"HomeOps/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	response.Success(c, http.StatusOK, "pong")
}

func health(c *gin.Context) {
	response.Success(c, http.StatusOK, "Server is healthy", gin.H{
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
		AuthRoutes(api)
		MetricRoutes(api)
		ServerRoutes(api)

	}
}
