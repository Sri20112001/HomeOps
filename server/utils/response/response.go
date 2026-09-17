package response

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Error sends a JSON error response
func Error(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ErrorResponse{
		Success: false,
		Error:   message,
	})
}

// AbortError sends a JSON error response AND stops downstream middleware/handlers
func AbortError(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, ErrorResponse{
		Success: false,
		Error:   message,
	})
}

// Success sends a standard success response with optional data payload
func Success(c *gin.Context, statusCode int, message string, data ...any) {
	var payload any
	if len(data) > 0 {
		payload = data[0]
	}

	c.JSON(statusCode, SuccessResponse{
		Success: true,
		Message: message,
		Data:    payload,
	})
}
