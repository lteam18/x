package main

import (
	"os"
	"path"
	"runtime"
	ut "utils"
)

func normalizeFolder(folder *string) string {
	if nil != folder {
		return *folder
	}
	if runtime.GOOS == "windows" {
		// TODO
		return "c:/"
	}
	return "/usr/local/bin"
}

func validateExecName(execName *string) string {
	if nil != execName {
		return *execName
	}
	return "x"
}

func normalizeExecName(execName *string) string {
	name := validateExecName(execName)
	if runtime.GOOS == "windows" {
		return name + ".exe"
	}
	return name
}

func install(targetFolder *string, execName *string) {
	newFilePath := path.Join(
		normalizeFolder(targetFolder),
		normalizeExecName(execName),
	)
	println("Installing x to: newFilePath")

	installToDst(newFilePath)
}

func installToDst(newFilePath string) {
	curFilePath, _ := os.Executable()
	err := ut.CopyFile(curFilePath, newFilePath)
	if nil != err {
		println("Try to install in **sudo** mode.")
		log.Panicln("Error", err)
	}
	println("Install success.")
}
