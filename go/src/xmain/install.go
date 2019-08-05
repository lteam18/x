package main

import (
	"nosh"
	"os"
	"path"
	"runtime"
	ut "utils"
	"vvkv"
)

// Which to install
func normalizeFolder(folder *string) string {
	if nil != folder {
		return *folder
	}
	if runtime.GOOS == "windows" {
		// TODO: which env?
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
	println("Installing x to: " + newFilePath)

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

var noshexe = nosh.Create(path.Join(vvkv.XPath, "bin", "nosh"))
