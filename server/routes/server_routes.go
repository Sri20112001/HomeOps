package routes

import (
	"HomeOps/controllers"
	"HomeOps/middleware"

	"github.com/gin-gonic/gin"
)

func ServerRoutes(rg *gin.RouterGroup) {
	servers := rg.Group("/server")

	servers.Use(middleware.AuthRequired())
	servers.POST("", controllers.CreateServer)
}
