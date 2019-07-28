package ut

import (
	"io/ioutil"
	"os"
	"strings"
)

/*
WriteFileSync writeFileSync
*/
func WriteFileSync(filepath string, content string) {
	idxF, idxErr := os.Create(filepath)
	HandleError(idxErr)
	idxF.WriteString(content)
	idxF.Sync()
}

/*
ReadFile writeFileSync
*/
func ReadFile(filepath string) string {
	dat, err := ioutil.ReadFile(filepath)
	PanicError(err)
	return string(dat)
}

/*
IsFileExisted execute command
*/
func IsFileExisted(filepath string) bool {
	// println(filepath)
	// _, err := os.Stat(filepath)
	// return os.IsExist(err)
	if _, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

/*
IsLikeHTTPURL execute command
*/
func IsLikeHTTPURL(uri string) bool {
	if strings.HasPrefix(uri, "https://") {
		return true
	}
	if strings.HasPrefix(uri, "http://") {
		return true
	}
	return false
}

/*
Mkdirp a
*/
func Mkdirp(filepathList ...string) {
	for _, v := range filepathList {
		// I prefer 751, But I found linux default is 0755
		os.MkdirAll(v, 0755)
	}
}
