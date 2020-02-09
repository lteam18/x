package gh

import (
	"encoding/json"
	"io/ioutil"
	ut "utils"
	"vvkv"
)

/*
GHTokenFileContent g
*/
type GHTokenFileContent struct {
	Owner string `json:"owner"`
	Token string `json:"token"`
}

/*
SetToken g
*/
func SetToken(owner, token string) {
	ut.WriteFileSync(vvkv.XPath+"/GH_TOKEN", ut.JS(&GHTokenFileContent{
		Owner: owner,
		Token: token,
	}))
}

func readToken() *GHTokenFileContent {
	ret := &GHTokenFileContent{}
	dat, err := ioutil.ReadFile(vvkv.XPath + "/GH_TOKEN")
	if err != nil {
		return nil
	}
	json.Unmarshal(dat, &ret)
	return ret
}

var ghInitTokenSingleton = readToken()

/*
GetToken a
*/
func GetToken() *GHTokenFileContent {
	return ghInitTokenSingleton
}
