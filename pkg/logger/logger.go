package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitLog(flag string) {
	if flag == "production" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
	logrus.SetOutput(os.Stdout)
	level, err := logrus.ParseLevel("info")
	if err != nil {
		panic(err)
	}
	logrus.SetLevel(level)
}
