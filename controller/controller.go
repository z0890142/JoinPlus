package controller

import (
	log "CommonHelper/Log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
	"github.com/z0890412/JoinPlus/service"
)

var bot *linebot.Client

func SetLineBot(_bot *linebot.Client) {
	bot = _bot
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	Log := log.LogInit("CallbackHandler", "JoinPlus", logrus.DebugLevel)
	events, err := bot.ParseRequest(r)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error": err,
		}).Error("lineBot.ParseRequest error")
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				// quota, err := bot.GetMessageQuota().Do()
				// if err != nil {
				// 	Log.Error("Quota err:", err)
				// }

				placeID := service.GetPlaceID(message.Text, Log)
				if placeID == "-99" {
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("輸入地點錯誤")).Do(); err != nil {
						Log.Info(err)
					}
				}

				lat, lng := service.GetPlaceDetail(placeID, Log)
			}
		}
	}

}
