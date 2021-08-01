package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/taserbeat/line-family-bot/modules/controllers"
	"github.com/taserbeat/line-family-bot/modules/env"
)

func main() {
	router := gin.Default()

	router.GET("/", controllers.HealthCheckHandler())

	router.GET("/version", controllers.GetVersionHandler())

	router.POST("/webhook", controllers.WebhookHandler())

	// ポートの設定
	port := env.Env.Port
	if port == "" {
		port = "8080"
	}

	// サーバーのリッスン
	if err := router.Run(":" + port); err != nil {
		log.Fatalln("サーバーの起動に失敗")
	}
}
