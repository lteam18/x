package nosh

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	ut "utils"
	"vvkv"
)

var noshDownloadPath = detect(runtime.GOOS)

func detect(os string) string {
	if os == "windows" {
		return "https://github.com/lteam18/nosh/releases/download/latest-dev/nosh.windows.zip"
	}

	if os == "linux" {
		return "https://github.com/lteam18/nosh/releases/download/latest-dev/nosh.linux.zip"
	}

	return "https://github.com/lteam18/nosh/releases/download/latest-dev/nosh.osx.zip"
}

func installNosh(noshPath string) {
	resp, _ := http.Get(noshDownloadPath)

	temppath, temppathErr := ioutil.TempFile(os.TempDir(), "nosh*.zip")
	ut.HandleError(temppathErr)
	_, err := io.Copy(temppath, resp.Body)
	ut.HandleError(err)

	vvkv.UnzipOneFile(temppath.Name(), noshPath)
	defer os.Remove(temppath.Name())
}
