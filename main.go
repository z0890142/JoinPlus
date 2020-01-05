package main

import (
	"strings"
	log "CommonHelper/Log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	Log := log.LogInit("main", "JoinPlus", logrus.DebugLevel)
	Log.Info("JoinPlus version 0.0.1")
	InitConfig()

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
