package main

import (
	"log"

	"HomeOps/config"
	"HomeOps/constants"
	"HomeOps/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	constants.LoadEnv()

	// 1. MUST BE FIRST
	config.InitDB()

	// 2. ONLY THEN call AutoMigrate
	err := config.MigrateDB()
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// 3. Setup routes and run
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":" + constants.AppEnv.Port)
}
