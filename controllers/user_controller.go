package controllers

import (
	"HomeOps/config"
	"HomeOps/models"
	"HomeOps/utils/response"
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
		response.Error(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	response.Success(c, http.StatusOK, "Users retrieved successfully", users)
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	id := uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to hash password")
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
		response.Error(c, http.StatusConflict, result.Error.Error())
		return
	}

	response.Success(c, http.StatusCreated, "User created successfully", user)

}
