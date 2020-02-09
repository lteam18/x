package main

import (
	"os"
	"strings"
)

var cmdName = os.Args[0]

func main() {
	if len(os.Args) == 1 {
		if strings.Index(os.Args[0], "x-installer") >= 0 {
			install(nil, nil)
			return
		}
	}

	// fmt.Println(getClient())

	runApp()
	// runAppOriginal()

	// getOrInstallNosh()
}
