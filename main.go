package main

import (
	"log"
	"net/http"

	"HomeOps/config"
	"HomeOps/models"
	"HomeOps/routes"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {
	// 1. MUST BE FIRST
	config.InitDB()

	// 2. ONLY THEN call AutoMigrate
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// 3. Setup routes and run
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":3500")
}
