package ut

// 24*60*60

import (
	"regexp"
	"strconv"
)

func getFrom(sym string, target string) int {
	form, err := regexp.Compile("\\d" + sym)
	if err != nil {
		panic(err)
	}
	ret := form.FindString(target)
	if len(ret) == 0 {
		return 0
	}

	num, err := strconv.Atoi(ret[0 : len(ret)-1])
	if err != nil {
		panic(err)
	}
	return num
}

/*
DateStr2Seconds a
*/
func DateStr2Seconds(str string) int64 {
	years := getFrom("Y", str)
	months := getFrom("M", str)
	days := getFrom("D", str)

	return int64(((years*365 + months*30) + days) * 24 * 60 * 60)

}
