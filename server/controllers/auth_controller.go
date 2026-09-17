package controllers

import (
	"errors"
	"net/http"

	"HomeOps/config"
	"HomeOps/models"
	"HomeOps/utils"
	"HomeOps/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput

	// 1. Validate incoming JSON payload
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// 2. Fetch the user record from the database
	var user models.User
	result := config.DB.Where("username = ?", input.Username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusUnauthorized, "Invalid email or password")
			return
		}
		response.Error(c, http.StatusInternalServerError, "Database error")
		return
	}

	// 3. Compare stored bcrypt hash with provided plain password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		response.Error(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// 4. Generate signed JWT token
	token, err := utils.GenerateToken(user)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate authentication token")
		return
	}

	// 5. Send token and sanitized user profile
	response.Success(c, http.StatusOK, "Login successful", gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username, // use user.Username if defined that way in models.User
			"email":    user.Email,
		},
	})
}

func GetMe(c *gin.Context) {
	// 1. Retrieve the value stored by auth middleware
	userIDVal, exists := c.Get("userId")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	userIDStr, ok := userIDVal.(string)
	if !ok {
		response.Error(c, http.StatusInternalServerError, "Invalid User ID type in context")
		return
	}

	// 2. Parse the string into a valid UUID
	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid UUID format")
		return
	}

	// 3. Query Postgres using GORM
	var user models.User
	if err := config.DB.WithContext(c.Request.Context()).First(&user, "id = ?", userUUID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "User not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, "Database error")
		return
	}

	// 4. Return sanitized user profile
	response.Success(c, http.StatusOK, "User retrieved successfully", gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}
