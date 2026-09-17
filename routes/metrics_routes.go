package routes

import (
	"HomeOps/controllers"

	"github.com/gin-gonic/gin"
)

func MetricRoutes(rg *gin.RouterGroup) {
	system := rg.Group("/system")
	{
		metrics := system.Group("/metrics")
		{
			metrics.GET("", controllers.GetSystemMetrics)
		}
	}
}
