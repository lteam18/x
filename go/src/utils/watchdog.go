package ut

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
Watchdog watchdog
*/
func Watchdog(maxRetry int, retryIntervalInMS int64, cmd string, args []string) {

	failure := 0

	sw := true

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	sleepSig := make(chan bool, 1)
	go func() {
		s := <-sigs
		if s == nil {
			return
		}
		sw = false
		sleepSig <- false
	}()

	defer func() { sigs <- nil }()

	for sw {
		// log.Infoln("Runing: " + strconv.Itoa(failure))
		ret, _ := Execute(cmd, args)
		if nil == ret {
			failure++
		} else {
			// if nil == err {
			if ret.ProcessState.ExitCode() == 0 {
				failure = 0
			} else {
				// TODO: About error, we should think twice.
				// log.Error("1: ", err)
				// log.Error("2: ", ret.ProcessState.String())
				failure++
			}
		}

		if maxRetry <= 0 {
			break
		}

		if failure >= maxRetry {
			// panic("Exit If retry at a row more than " + strconv.Itoa(failure))
			break
		}

		go func() {
			time.Sleep(time.Duration(retryIntervalInMS) * time.Millisecond)
			sleepSig <- true
		}()

		<-sleepSig
	}

}
