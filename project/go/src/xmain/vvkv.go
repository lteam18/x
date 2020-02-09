package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	ut "utils"
	"vvkv"
)

var url = "https://1632295596863408.cn-shenzhen.fc.aliyuncs.com/2016-08-15/proxy/vvsh/vvkv"

/*
VvkvClient a
*/
var VvkvClient = vvkv.CreateClient(url)

/*
Meta a
*/
type Meta struct {
	CodeType string  `json:"codetype"`
	IsURL    bool    `json:"isURL"`
	URL      *string `json:"url"`
}

func upload(src string, vvurl string, isPublic bool, codetype string, isURL bool) {
	meta := &Meta{codetype, isURL, nil}
	st, _ := json.Marshal(meta)

	info := make(map[string]string)
	info["x-vvkv-m"] = string(st)
	ut.Pjson(info)
	VvkvClient.UploadByVVURL(src, vvurl, isPublic, info)
}

func uploadCode(src string, vvurl string, isPublic bool, codetype string) {
	ut.Pjson("upload code")
	upload(src, vvurl, isPublic, codetype, false)
}

type linkType struct {
	URL *string `json:"url"`
}

func readLink(src string) linkType {
	var ret linkType
	json.Unmarshal([]byte(ut.ReadFile(src)), &ret)
	return ret
}

func uploadLink(srcLink string, vvurl string, isPublic bool, codetype string) {
	temppath, temppathErr := ioutil.TempFile(os.TempDir(), "vvsh*.zip")
	ut.HandleError(temppathErr)
	defer os.Remove(temppath.Name())

	st, err := json.MarshalIndent(&linkType{&srcLink}, "", "  ")
	ut.HandleError(err)
	ut.WriteFileSync(temppath.Name(), string(st))

	upload(temppath.Name(), vvurl, isPublic, codetype, true)
}

func printLs(vvurl string, detail bool) {
	if detail {
		printLsAll(vvurl)
	} else {
		printLsSimple(vvurl)
	}
}

func printLsAll(vvurl string) {
	ret := VvkvClient.ListVVURL(vvurl)
	st, _ := json.MarshalIndent(ret, "", "  ")
	writeResult(string(st))
}

func printLsSimple(vvurl string) {
	ret := VvkvClient.ListVVURL(vvurl)
	for _, v := range ret {
		writeResult("%s\t%d\t%s\n", v.LastModified, v.Size, v.Name)
	}
}

func isLikeVVURL(vvurl string) bool {
	return strings.HasPrefix(vvurl, "@") && strings.Index(vvurl, "/") >= 0
}

type urlType struct {
	CodeType string `json:"codetype"`
	IsURL    string `json:"isURL"`
}
