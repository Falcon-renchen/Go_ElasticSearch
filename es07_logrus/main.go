package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logger := logrus.New()
	//logger.Warn("警告日志")
	//logger.Fatal("fatal")
	//logger.Info("info")
	//logger.Error("error")

	logfile, err := os.OpenFile("./es07_logrus/mylog", os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModeAppend)
	if err != nil {
		logrus.Fatal(err)
	}
	logger.Out = logfile
	logger.Info("abc")
}
