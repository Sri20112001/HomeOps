package utils

import (
	"time"

	"HomeOps/constants"
	"HomeOps/models"

	"github.com/golang-jwt/jwt/v5"
)

// getJWTSecret ensures AppEnv is evaluated at runtime, not during init
func getJWTSecret() []byte {
	return []byte(constants.AppEnv.JWTSecretKey)
}

func GenerateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"userId":   user.ID,
		"username": user.Username, // or user.Username, matching your models.User definition
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJWTSecret())
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(
		tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return getJWTSecret(), nil
		},
	)
}
