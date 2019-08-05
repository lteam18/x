package ut

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
PanicError classical handler
*/
func PanicError(e error) {
	if e != nil {
		panic(e)
	}
}

/*
Pjson classical handler
*/
func Pjson(a interface{}) {
	st, err := json.MarshalIndent(a, "", "  ")
	PanicError(err)
	println(string(st))
}

/*
PrettifyJSONString a
*/
func PrettifyJSONString(jsonStr string) string {
	amap := make(map[string]interface{})
	json.Unmarshal([]byte(jsonStr), &amap)
	st, err := json.MarshalIndent(amap, "", "  ")
	PanicError(err)
	return string(st)
}

/*
HandleError handle error
*/
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

/*
IfThenElse handle error
*/
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

/*
CalMd5 cal md5
*/
func CalMd5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

/*
HTTPCat http cat
*/
func HTTPCat(url string, dstPath string) {

	req, _ := http.NewRequest("GET", url, nil)

	cli := &http.Client{}
	resp, _ := cli.Do(req)

	f, ferr := os.Create(dstPath)
	HandleError(ferr)
	_, copyError := io.Copy(f, resp.Body)
	HandleError(copyError)
}

/*
func FromHTTPUrl(url, dst string) {
	res, err := http.Get(url)
	ut.PanicError(err)
	f, err := os.Create(dst)
	ut.PanicError(err)
	io.Copy(f, res.Body)
}
*/

/*
SliceOrEmpty a
*/
func SliceOrEmpty(data []string, start int) []string {
	if len(data) > start {
		return data[start:]
	}
	return []string{}
}
