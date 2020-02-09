package debuglog

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var debugComponentSet = make(map[string]bool)

func init() {
	debugEnv, _ := os.LookupEnv("DEBUG")
	for _, v := range strings.Split(debugEnv, ":") {
		v1 := strings.TrimSpace(v)
		v1 = strings.ToLower(v1)
		debugComponentSet[v1] = true
	}
}

/*
IsDebug g
*/
func IsDebug(com string) bool {
	_, ok := debugComponentSet[com]
	return ok
}

/*
Create a
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

	if IsDebug(name) {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
	return logger
}
