package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/taserbeat/line-family-bot/modules/auth"
	"github.com/taserbeat/line-family-bot/modules/env"
	"github.com/taserbeat/line-family-bot/modules/functions"
	"github.com/thoas/go-funk"
)

var client *http.Client
var bot *linebot.Client

func init() {
	var err error

	client = &http.Client{}
	bot, err = linebot.New(env.Env.LineChannelSecret, env.Env.LineChannelAccessToken, linebot.WithHTTPClient(client))
	if err != nil {
		log.Fatalln("linebotクライアントの作成に失敗", err)
	}
}

func WebhookHandler() gin.HandlerFunc {
	return webhook
}

func webhook(c *gin.Context) {
	// GWサーバーにステータスコード200をなるべく早く返す
	c.JSON(http.StatusOK, gin.H{})

	// 署名検証とイベント解析
	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		switch err {
		case linebot.ErrInvalidSignature:
			log.Println("リクエストの署名が不正", err)
		default:
			log.Println("イベントの解析に失敗", err)
		}
		return
	}

	for _, event := range events {

		if event.Type == linebot.EventTypeMessage {
			log.Printf("eventType: %s, userId: %s, groupId: %s, roomId: %s\n", event.Type, event.Source.UserID, event.Source.GroupID, event.Source.RoomID)

			// 認証
			userId := event.Source.UserID
			if !auth.AuthenticateClient.IsFamily(userId) {
				log.Println("未承認ユーザーからのアクセス", userId)
				return
			}

			if event.Source.Type == linebot.EventSourceTypeGroup {
				// グループでのトークの場合

				// ボット宛のメンションであるかを判定
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					isMentionToBot := funk.Contains(message.Mention.Mentionees, func(mentionee *linebot.Mentionee) bool {
						return mentionee.UserID == env.Env.LineChannelUserId
					})

					// ボット宛のメンションではない場合は処理を終了
					if !isMentionToBot {
						return
					}

					today, err := functions.GetToday()
					if err != nil {
						log.Println("「今日は何の日」を取得できなかった", err)
						return
					}

					nowString := time.Now().Format("1/2")

					textTitle := fmt.Sprintf("今日(%s)は%sです", nowString, today.Title)
					replyTitleMsg := linebot.NewTextMessage(textTitle)
					replyDescMsg := linebot.NewTextMessage(today.Description)
					_, err = bot.ReplyMessage(event.ReplyToken, replyTitleMsg, replyDescMsg).Do()
					if err != nil {
						log.Println("メッセージのリプライエラー", err)
						return
					}

				default:
					return
				}

			} else if event.Source.Type == linebot.EventSourceTypeUser {
				// 個人チャットの場合
				replyText := linebot.NewTextMessage("あなたのIDは " + userId + " です")
				_, err := bot.ReplyMessage(event.ReplyToken, replyText).Do()
				if err != nil {
					log.Println("メッセージのリプライエラー", err)
					return
				}
			}
		}
	}
}
