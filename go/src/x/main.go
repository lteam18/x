package main

import (
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "hi":
		println("hi")
	case "python":
		fallthrough
	case "bash":
		fallthrough
	case "fish":
		fallthrough
	case "js":
		fallthrough
	case "node":
		fallthrough
	case "perl":
		execute(os.Args[1:]...)
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}
}
