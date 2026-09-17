package constants

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ENV struct {
	JWTSecretKey string
	Port         string

	// Database fields
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     int
	DBSSLMode  string
	DBTimeZone string
}

var AppEnv *ENV

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Note: .env file not found, using system environment variables")
	}

	// Parse DB_PORT from string to int
	portInt, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		portInt = 5432
	}

	AppEnv = &ENV{
		JWTSecretKey: getEnv("JWT_SECRET_KEY", "default-secret"),
		Port:         getEnv("PORT", "3500"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBPassword:   getEnv("DB_PASSWORD", ""),
		DBName:       getEnv("DB_NAME", "homeops"),
		DBPort:       portInt,
		DBSSLMode:    getEnv("DB_SSLMODE", "disable"),
		DBTimeZone:   getEnv("DB_TIMEZONE", "UTC"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return fallback
}
