package ut

import (
	"io/ioutil"
	"os"
)

/*
Cat facility
*/
func Cat(filepathList ...string) {
	for _, filepath := range filepathList {
		dat, err := ioutil.ReadFile(filepath)
		PanicError(err)
		// fmt.Print(string(dat))
		os.Stdout.Write(dat)
	}
}
