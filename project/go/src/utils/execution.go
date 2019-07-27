package ut

import (
	"bytes"
	"io"
	"os"
	"os/exec"
)

/*
Execute execute command
*/
func Execute(funcname string, args []string) *exec.Cmd {
	bout := bytes.NewBuffer(nil)
	berr := bytes.NewBuffer(nil)

	cmd := exec.Command(funcname, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = bout
	cmd.Stderr = berr

	cmd.Run()

	// TODO: using stream like API for that.
	io.Copy(os.Stdout, bout)
	io.Copy(os.Stderr, berr)

	return cmd
}
