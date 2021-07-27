package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.io/taserbeat/line-family-bot/modules/env"
)

func WebhookHandler() gin.HandlerFunc {
	return webhook
}

func webhook(c *gin.Context) {
	// GWサーバーにステータスコード200をなるべく早く返す
	c.JSON(http.StatusOK, gin.H{})

	// TODO: 署名検証処理

	// TODO: イベント解析
	client := &http.Client{}
	bot, err := linebot.New(env.Env.LineChannelSecret, env.Env.LineChannelAccessToken, linebot.WithHTTPClient(client))
	if err != nil {
		log.Println("linebotのクライアント作成に失敗", err)
		return
	}

	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		log.Println("イベントの解析に失敗", err)
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			log.Printf("eventType: %s, userId: %s, groupId: %s, roomId: %s\n", event.Type, event.Source.UserID, event.Source.GroupID, event.Source.RoomID)
		}
	}
}
