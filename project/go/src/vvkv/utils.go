package vvkv

import (
	"debuglog"
	"os/user"
	"strings"
	ut "utils"
)

var log = debuglog.Create("x")

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

/*
AppPath Pre
*/
var XPath = getHomeDir() + "/.vvx"

/*
AppDirPath Pre
*/
var AppDirPath = XPath + "/APP"

/*
AppFlatDirPath Pre
*/
var AppFlatDirPath = XPath + "/APP_FLAT"

/*
AppIdxDirPath Pre
*/
var AppIdxDirPath = XPath + "/APP_IDX"

/*
PrefixFilePath Pre
*/
var PrefixFilePath = XPath + "/PREFIX"

func init() {
	ut.Mkdirp(AppDirPath, AppFlatDirPath, AppIdxDirPath)

	if !ut.IsFileExisted(PrefixFilePath) {
		ut.WriteFileSync(PrefixFilePath, "@official\n")
	}
}

/*
Filter try to reduce filter
*/
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

/*
Map map
*/
func Map(vs []string, f func(string) string) []string {
	vsf := make([]string, len(vs))
	for idx, v := range vs {
		vsf[idx] = f(v)
	}
	return vsf
}

/*
GetDefaultPrefixes a
*/
func GetDefaultPrefixes() []string {
	str := ut.ReadFile(PrefixFilePath)
	items := strings.Split(str, "\n")

	vsf := make([]string, 0)
	for _, v := range items {
		tmpv := strings.TrimSpace(v)
		if len(tmpv) > 0 {
			vsf = append(vsf, tmpv)
		}
	}
	return vsf
}

/*
GetCacheDirPath a
*/
func GetCacheDirPath(prefix string, key string) (string, string) {
	fullPath := prefix + "/" + key
	md5 := ut.CalMd5(fullPath)
	return AppFlatDirPath + "/" + md5, fullPath
}

/*
PrefixFilePath Pre
*/

// func checkFileExists(name string) (bool, string, string) {
// 	strList := getDefaultPrefixes()
// 	for _, prefix := range strList {
// 		cacheFilePath, _ := GetCacheDirPath(prefix, name)
// 		{
// 			return true, cacheFilePath, prefix
// 		}
// 	}
// 	return false, "", ""
// }
