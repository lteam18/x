package main

import (
	"os/exec"
	"strings"
	ut "utils"
)

func isHTTPURL(s string) bool {
	return strings.HasPrefix(s, "https://") || strings.HasPrefix(s, "http://")
}

func execute(args ...string) *exec.Cmd {
	cmd := args[0]
	path := args[1]

	if isHTTPURL(path) {
		cmd, err := ut.Execute("xmain", args)
		if err != nil {
			return nil
		}
		return cmd
	}

	// if local-file-path

	// if starts with @, check file

	// see if local-path, check other

	// read PREFIX files, try one by one: very slow

	c, e := ut.Execute(cmd, args[1:])
	if e != nil {
		return nil
	}
	return c
}
