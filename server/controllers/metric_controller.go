package controllers

import (
	"net/http"

	"HomeOps/metrics"
	"HomeOps/utils/response"

	"github.com/gin-gonic/gin"
)

func GetSystemMetrics(c *gin.Context) {
	systemInfo, err := metrics.GetSystemInfo()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to collect system metrics")
		return
	}

	response.Success(c,
		http.StatusOK, "Successfully collected the Metrics", systemInfo,
	)
}
