package ut

import "encoding/json"

/*
JS a
*/
func JS(a interface{}) string {
	st, err := json.MarshalIndent(a, "", "  ")
	HandleError(err)
	return string(st)
}
