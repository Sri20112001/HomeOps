package middleware

import (
	"HomeOps/constants"
	"HomeOps/utils/response"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.AbortError(c, http.StatusUnauthorized, "Authorization header is missing")
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			response.AbortError(c, http.StatusUnauthorized, "Authorization header format must be 'Bearer <token>'")
			return
		}

		tokenString := strings.TrimSpace(parts[1])
		secret := constants.AppEnv.JWTSecretKey
		if secret == "" {
			response.AbortError(c, http.StatusInternalServerError, "Server configuration error: missing JWT secret")
			return
		}
		jwtSecret := []byte(secret)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			response.AbortError(c, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.AbortError(c, http.StatusUnauthorized, "Invalid token claims")
			return
		}

		// Support both standard "sub" and custom "userId"
		var userID string
		if id, ok := claims["userId"].(string); ok {
			userID = id
		} else if sub, ok := claims["sub"].(string); ok {
			userID = sub
		}

		if userID == "" {
			response.AbortError(c, http.StatusUnauthorized, "User identity missing from token")
			return
		}

		c.Set("userId", userID)
		c.Next()
	}
}
