package vvkv

import (
	"testing"
	ut "utils"
)

func TestABC(t *testing.T) {
	t.Log("Given the nded")

	url1 := "https://1632295596863408.cn-shenzhen.fc.aliyuncs.com/2016-08-15/proxy/vvsh/vvkv"
	cli := CreateClient(url1)

	ut.Pjson(cli.GetToLocalDisk("official", "ding-ip", "./a", "./b"))
}
