package logger

import "github.com/sirupsen/logrus"

func Error(args ...interface{}) {
	logrus.Error(args...)
}
