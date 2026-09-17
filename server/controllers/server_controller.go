package controllers

import (
	"net/http"

	"HomeOps/config"
	"HomeOps/models"
	"HomeOps/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateServerInput struct {
	Name     string `json:"name" binding:"required"`
	Hostname string `json:"hostname"`
	OS       string `json:"os"`
	Platform string `json:"platform"`
}

func CreateServer(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	userIDString, ok := userId.(string)
	if !ok || userIDString == "" {
		response.Error(c, http.StatusUnauthorized, "Invalid user identity")
		return
	}

	var input CreateServerInput

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(
			c, http.StatusBadRequest, "Invalid request body",
		)
		return
	}

	server := models.Server{
		ID:       uuid.New().String(),
		UserID:   userIDString,
		Name:     input.Name,
		Hostname: input.Hostname,
		OS:       input.OS,
		Platform: input.Platform,
	}

	if err := config.DB.Create(&server).Error; err != nil {
		response.Error(
			c, http.StatusInternalServerError, "Failed to create server",
		)
		return
	}

	response.Success(
		c, http.StatusCreated, "Server created successfully", server,
	)

}
