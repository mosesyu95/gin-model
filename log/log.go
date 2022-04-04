package log

import (
	"gin-model/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

var Log *logrus.Logger

func Init() { // 初始化log的函数
	f, err := os.OpenFile(config.Config.Log.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("Open log file error, err: ", err.Error())
	}
	level, err := logrus.ParseLevel(viper.GetString("logger.level"))
	if err != nil {
		level = logrus.WarnLevel
		log.Println("Load log level failed . set default level \"Warnning\" ", err.Error())
	}
	log := logrus.Logger{
		Formatter: &logrus.JSONFormatter{},
		Out:       f,
		Level:     level,
	}
	Log = &log
	return
}

func GetLog() *logrus.Logger {
	return Log
}
