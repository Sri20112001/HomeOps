package routes

import (
	"HomeOps/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.GET("", controllers.GetUsers)
		// users.POST("")
		// users.GET("/:id")
	}
}
