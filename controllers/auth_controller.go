package controllers

import (
	"errors"
	"net/http"

	"HomeOps/config"
	"HomeOps/models"
	"HomeOps/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginInput struct {
	username string `json:"username" binding:"required,username"`
	password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput

	// 1. Validate incoming JSON payload
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Fetch the user record from the database
	var user models.User
	result := config.DB.Where("username = ?").First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// 3. Compare stored bcrypt hash with provided plain password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// 4. Generate signed JWT token
	token, err := utils.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate authentication token"})
		return
	}

	// 5. Send token and sanitized user profile
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username, // use user.Username if defined that way in models.User
			"email":    user.Email,
		},
	})
}
