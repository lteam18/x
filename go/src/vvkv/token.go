package vvkv

import (
	"io/ioutil"
	ut "utils"
)

/*
ReadToken get token from vvkv folder
*/
func ReadToken() string {
	path := XPath + "/TOKEN"
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		// panic(err)
		// TODO: add debug log
		log.WithField("path", path).Debugln("No token found in the local environment.")
		return ""
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
	ut.WriteFileSync(XPath+"/TOKEN", tokenStr)
}
