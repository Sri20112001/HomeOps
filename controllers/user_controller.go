package controllers

import (
	"HomeOps/config"
	"HomeOps/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func GetUsers(c *gin.Context) {
	var users []models.User

	result := config.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(
		http.StatusOK, gin.H{
			"data": users,
		},
	)
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"error": "Failed to hash password",
			},
		)
		return
	}

	user := models.User{
		ID:       id.String(),
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(
			http.StatusConflict, gin.H{
				"error": result.Error.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated, gin.H{
			"data": user,
		},
	)

}
