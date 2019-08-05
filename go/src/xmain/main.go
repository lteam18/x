package main

import (
	"debuglog"
	"os"
	"strings"
)

var log = debuglog.Create("x")

func parseAccess(access string) bool {
	switch access {
	case "public":
		return true
	case "private":
		return false
	default:
		log.Panicln("Should be private or public")
		panic("")
	}
}

var cmdName = os.Args[0]

func main() {

	if len(os.Args) == 1 {
		if strings.Index(os.Args[0], "x-installer") >= 0 {
			install(nil, nil)
			return
		}
	}

	runApp()
	// runAppOriginal()

	// getOrInstallNosh()
}
