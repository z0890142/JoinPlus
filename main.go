package main

import (
	log "CommonHelper/Log"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
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
	channelSecret := viper.GetString("Line.ChannelSecret")
	channelAccessToken := viper.GetString("Line.ChannelAccessToken")
	bot, err = linebot.New(channelSecret, channelAccessToken)
	controller.SetLineBot(bot)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"ChannelSecret":      viper.GetString("Line.ChannelSecret"),
			"ChannelAccessToken": viper.GetString("Line.ChannelAccessToken"),
		}).Error("new line bot error")
	}

	http.HandleFunc("/callback", controller.CallbackHandler)
	port := "8088"
	addr := fmt.Sprintf(":%s", port)

	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static/.well-known/acme-challenge"))
	router.PathPrefix("/.well-known/acme-challenge/").Handler(http.StripPrefix("/.well-known/acme-challenge/", fs))

	http.ListenAndServeTLS(addr, "./static/ssl/bundle.crt", "./static/ssl/private.key", nil)
	// http.ListenAndServe(addr, router)
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
