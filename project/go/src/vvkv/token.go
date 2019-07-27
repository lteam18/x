package vvkv

import (
	"io/ioutil"
	ut "utils"
)

/*
ReadToken get token from vvkv folder
*/
func ReadToken() string {
	bytes, err := ioutil.ReadFile(XPath + "/TOKEN")
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

/*
WriteTokenSync get token into vvkv folder
*/
func WriteTokenSync(tokenStr string) {
	/*
		err := ioutil.WriteFile(appDir+"/TOKEN", []byte(tokenStr), 600)
		if err != nil {
			panic(err)
		}
	*/
	ut.WriteFileSync(AppDirPath+"/TOKEN", tokenStr)
}
