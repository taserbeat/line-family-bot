package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taserbeat/line-family-bot/modules/models"
)

func GetVersionHandler() gin.HandlerFunc {
	return getVersion
}

func getVersion(c *gin.Context) {
	c.JSON(http.StatusOK, models.Ver)
}
