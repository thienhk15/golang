package utils

import (
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

func ShowErrorLogs(errData error) {
	if errData != nil {
		file, line := Caller(2)
		logrus.SetFormatter(&logrus.TextFormatter{})
		logrus.Errorf("%s[%d] %s", filepath.Base(file), line, errData.Error())
	}
}

func ShowInfoLogs(infoWantToLog string) {
	if infoWantToLog != "" {
		file, line := Caller(2)
		logrus.SetFormatter(&logrus.TextFormatter{})
		logrus.Infof("%s[%d] %s", filepath.Base(file), line, infoWantToLog)
	}
}

func Caller(level int) (string, int) {
	_, file, line, _ := runtime.Caller(level)
	return file, line
}
