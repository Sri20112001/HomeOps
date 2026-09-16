package routes

import (
	"HomeOps/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	{
		auth.GET("/login", controllers.Login)
	}
}
