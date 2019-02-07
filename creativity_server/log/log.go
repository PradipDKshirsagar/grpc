package log

import (
	"fmt"

	logger "github.com/sirupsen/logrus"
)

func InfoWithFields(fn, msg string, fields logger.Fields) {
	logger.WithFields(fields).Info(fmt.Sprintf("[%s] - %s", fn, msg))
}

func Info(fn, msg string) {
	logger.Info(fmt.Sprintf("[%s] - %s", fn, msg))
}

func ErrorWithFields(fn, msg string, fields logger.Fields) {
	logger.WithFields(fields).Error(fmt.Sprintf("[%s] - %s", fn, msg))
}

func Error(fn, msg string) {
	logger.Error(fmt.Sprintf("[%s] - %s", fn, msg))
}
