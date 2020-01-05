package main

import (
	log "CommonHelper/Log"
	"fmt"
	"net/http"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/z0890412/JoinPlus/controller"
)

var Log *logrus.Entry
var bot *linebot.Client

func init() {
	Log = log.LogInit("main", "JoinPlus", logrus.DebugLevel)
	Log.Info("JoinPlus version 0.0.1")
	InitConfig()

}

func main() {
	var err error
	bot, err = linebot.New(viper.GetString("Line.ChannelSecret"), viper.GetString("Line.ChannelAccessToken"))

	if err != nil {
		Log.WithFields(logrus.Fields{
			"ChannelSecret":      viper.GetString("Line.ChannelSecret"),
			"ChannelAccessToken": viper.GetString("Line.ChannelAccessToken"),
		}).Error("new line bot error")
	}

	http.HandleFunc("/callback", controller.CallbackHandler)
	port := "8088"
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServeTLS(addr, "./static/ssl/https-server.crt", "./static/ssl/https-server.key", nil)
}
func InitConfig() {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()         // read in environment variables that match
	viper.SetEnvPrefix("gorush") // will be uppercased automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigName("config")   // name of config file (without extension)
	viper.AddConfigPath("./config") // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err == nil {
		Log.Info("Using config file:", viper.ConfigFileUsed())
	}

}
