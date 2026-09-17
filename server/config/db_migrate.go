package config

import (
	"HomeOps/models"
)

func MigrateDB() error {
	err := DB.AutoMigrate(&models.User{}, &models.Server{})
	return err
}
