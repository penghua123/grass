package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	customFormatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	logger := logrus.New()
	logger.Formatter = customFormatter

	rotateLogger := &lumberjack.Logger{
		Filename: "./foo.log",
	}
	logger.Out = rotateLogger
	logger.Info("logrus log to lumberjack in normal text formatter")
}
