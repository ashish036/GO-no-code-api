package utils

import (
	"context"
	"os"

	logrus "github.com/sirupsen/logrus"
)

func InitializeLogger() {
	// Set logger
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.ErrorLevel)
	logrus.SetReportCaller(true)
}

func GetLogHelper(ctx context.Context, module string, statusCode int) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"module": module,
		"code":   statusCode,
	})
}
