package ut

import (
	"errors"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
)

/*
Execute execute command
*/
func Execute(funcname string, args []string) (*exec.Cmd, error) {
	// bout := bytes.NewBuffer(nil)
	// berr := bytes.NewBuffer(nil)
	// io.Copy(os.Stdout, bout)
	// io.Copy(os.Stderr, berr)

	// ctxt, cancel := context.WithTimeout(context.Background())
	// defer cancel()

	cmd := exec.Command(funcname, args...)
	// cmd := exec.CommandContext(context.Background(), funcname, args...)

	if nil == cmd {
		return nil, errors.New("exec: Return nil pointer. Should not happenned")
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	sigs := make(chan os.Signal, 1)
	// done := make(chan bool, 1)
	cmdErrChan := make(chan error, 1)

	signal.Notify(sigs)

	go func() {
		for true {
			sig := <-sigs
			if sig == nil {
				return
			}

			cmd.Process.Signal(sig)
			if sig == syscall.SIGINT {
				// Wait and force kill
				pid := strconv.Itoa(cmd.Process.Pid)
				cmdErrChan <- errors.New("Recv SIGINT. Exit now. Subprocess PID is " + pid)
				println("Recv SIGINT. Exit now. Subprocess PID is " + pid)
				// fmt.Errorf("Recv SIGINT. Exit now. Subprocess PID is %s", pid)
				return
			}

			if sig == syscall.SIGTERM {
				// Wait and force kill
				pid := strconv.Itoa(cmd.Process.Pid)
				cmdErrChan <- errors.New("Recv SIGTERM. Exit nowSubprocess PID is " + pid)
				println("Recv SIGTERM. Exit nowSubprocess PID is " + pid)
				// fmt.Errorf("Recv SIGINT. Exit now. Subprocess PID is %s", pid)
				return
			}
		}
	}()
	defer func() { sigs <- nil }()

	// go func() {
	// 	cmd.Wait()
	// 	sigs <- nil
	// }()

	go func() {
		cmdErrChan <- cmd.Run()
	}()
	// err := cmd.Wait()
	// println(err)

	err := <-cmdErrChan

	return cmd, err
}
