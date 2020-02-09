package nosh

import (
	"fmt"
	"os"
	"runtime"
	ut "utils"
)

func detectXVersion(os string, arch string) string {

	// https://github.com/lteam18/x/releases/download/latest-dev/x-installer.darwin-amd64

	if os == "windows" {
		return fmt.Sprintf("https://github.com/lteam18/x/releases/download/latest-dev/x-installer.%s-%s.exe", os, arch)
	}

	return fmt.Sprintf("https://github.com/lteam18/x/releases/download/latest-dev/x-installer.%s-%s", os, arch)

	// if os == "linux" {
	// 	return "https://github.com/lteam18/x/releases/download/latest-dev/x.linux"
	// }

	// return "https://github.com/lteam18/x/releases/download/latest-dev/x.darwin"
}

func DownloadX(dstPath string) error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}

	fi, _ := os.Stat(exe)

	if err := ut.Download(detectXVersion(runtime.GOOS, runtime.GOARCH), dstPath); err != nil {
		return err
	}

	return os.Chmod(dstPath, fi.Mode())
}
