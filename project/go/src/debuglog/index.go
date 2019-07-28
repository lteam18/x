package debuglog

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

/*
CreateDebugLogger a
*/
func Create(name string) *logrus.Logger {

	var logger = &logrus.Logger{
		Out: os.Stderr,
		Formatter: &logrus.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "02 15:04:05", // "2006-01-02 15:04:05"
			FullTimestamp:   true,
		},
	}

	debugEnv, ok := os.LookupEnv("DEBUG")

	if ok {
		itemList := strings.Split(debugEnv, ";")
		for _, v := range itemList {
			v1 := strings.TrimSpace(v)
			if name == strings.ToLower(v1) {
				logger.SetLevel(logrus.DebugLevel)
				return logger
			}
		}
	}

	logger.SetLevel(logrus.WarnLevel)

	return logger

}
