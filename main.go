package main

import (
	"github.com/gin-gonic/gin"
	"github.io/taserbeat/line-family-bot/modules/controllers"
	"github.io/taserbeat/line-family-bot/modules/env"
)

func main() {
	router := gin.Default()

	router.GET("/", controllers.HealthCheck())

	router.GET("/version", controllers.GetVersion())

	// ポートの設定
	port := env.Env.Port
	if port == "" {
		port = "8080"
	}

	// サーバーのリッスン
	if err := router.Run(":" + port); err != nil {
		panic(err)
	}
}
