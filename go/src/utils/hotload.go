package ut

import (
	"os"
	"os/exec"
	"syscall"
)

/*
ExecReplace a
*/
func ExecReplace(cmd string, args []string) error {
	binary, lookErr := exec.LookPath(cmd)
	if lookErr != nil {
		panic(lookErr)
	}

	return syscall.Exec(binary, args, os.Environ())
}

// func server() {
// 	dir := vvkv.XPath + "/workers"
// 	ut.Mkdirp(dir)
//  ut.WriteFileSync(dir + "/1")
// 	c, e := ut.Execute(os.Args[0], []string{"worker"})
// }

// func worker(cmd string, args []string) {
// 	dir := vvkv.XPath + "/workers"
// 	ut.Mkdirp(dir)
// 	ut.WriteFileSync(dir + "/1")
// 	c, e := ut.Execute(os.Args[0], []string{"worker"})
// }

// func send(){

// }
