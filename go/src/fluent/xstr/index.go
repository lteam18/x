package xstr

import "fluent/list"

func c(strList ...[]string) []string {
	ret := []string{}
	for _, v := range strList {
		ret = append(ret, v...)
	}
	return ret
}

func ret() []string {
	return []string{"a"}
}

var a = c(ret(), ret(), list.Str("a", "b"))
