package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.io/taserbeat/line-family-bot/modules/models"
)

func GetVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, models.Ver)
	}
}
