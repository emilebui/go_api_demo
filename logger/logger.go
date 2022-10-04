package logger

import (
	"github.com/sirupsen/logrus"
	"go_api/helpers"
	"go_api/load_config"
	"os"
)

func InitLogger() *logrus.Logger {

	logger := logrus.New()
	logger.Out = os.Stdout

	if load_config.Conf.Log2File {
		filename := helpers.GetDirPath("log.txt")
		file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			logger.Infof("Failed to log to file (%v), logging to stdout instead", err)
		} else {
			logger.Out = file
		}
	}

	return logger
}

func LogInfo(log string, service string) {
	logger.WithFields(logrus.Fields{
		"App":     "GO_API_DEMO",
		"Service": service,
	}).Info(log)
}

func LogError(err error, service string) {
	logger.WithFields(logrus.Fields{
		"App":     "GO_API_DEMO",
		"Service": service,
	}).Error(err.Error())
}

var logger = InitLogger()
