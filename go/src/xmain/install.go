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
	log.Infoln("Installing x to: " + newFilePath)

	installToDst(newFilePath)
}

func upgrade() error {
	// Download From Github into temporary folder
	curFilePath, err := os.Executable()
	if err != nil {
		return err
	}
	return nosh.DownloadX(curFilePath)
}

func installToDst(newFilePath string) {
	curFilePath, _ := os.Executable()
	err := ut.CopyFile(curFilePath, newFilePath)
	if nil != err {
		log.Infoln("Try to install in **sudo** mode.")
		log.Panicln("Error", err)
	}
	log.Infoln("Install success.")
}

var noshexe = nosh.Create(path.Join(vvkv.XPath, "bin", "nosh"))
