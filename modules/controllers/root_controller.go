package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler() gin.HandlerFunc {
	return healthCheck
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
