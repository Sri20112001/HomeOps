package routes

import (
	"HomeOps/controllers"
	"HomeOps/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		// auth.POST("/register", controllers.Register)
	}

	// Protected endpoints
	protected := rg.Group("/auth").Use(middleware.AuthRequired())
	{
		protected.GET("/me", controllers.GetMe)
	}
}
