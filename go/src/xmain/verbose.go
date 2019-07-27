package main

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// var Debug = log.New(os.Stderr, "vvx: ", log.LstdFlags)

/*
Log a
*/
var Log = &logrus.Logger{
	Out: os.Stderr,
	Formatter: &logrus.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "02 15:04:05", // "2006-01-02 15:04:05"
		FullTimestamp:   true,
	},
}

func init() {

	debugEnv, ok := os.LookupEnv("DEBUG")

	if ok {
		itemList := strings.Split(debugEnv, ";")
		for _, v := range itemList {
			v1 := strings.TrimSpace(v)
			if "x" == strings.ToLower(v1) {
				Log.SetLevel(logrus.DebugLevel)
				return
			}
		}
	}

	Log.SetLevel(logrus.WarnLevel)
}
