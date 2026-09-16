package config

import (
	"fmt"
	"log"
	"time"

	"HomeOps/constants"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type ConnectionConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
	SSLMode  string
	TimeZone string
}

func (c ConnectionConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLMode, c.TimeZone,
	)
}

func InitDB() {
	var err error

	cfg := ConnectionConfig{
		Host:     constants.AppEnv.DBHost,
		User:     constants.AppEnv.DBUser,
		Password: constants.AppEnv.DBPassword,
		DBName:   constants.AppEnv.DBName,
		Port:     constants.AppEnv.DBPort,
		SSLMode:  constants.AppEnv.DBSSLMode,
		TimeZone: constants.AppEnv.DBTimeZone,
	}

	DB, err = gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database handle: %v", err)
	}

	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetMaxOpenConns(2)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("Successfully connected to PostgresSql!")

}
