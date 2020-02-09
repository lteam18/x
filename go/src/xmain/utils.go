package main

import "debuglog"

var log = debuglog.Create("x")

func parseAccess(access string) bool {
	switch access {
	case "public":
		return true
	case "private":
		return false
	default:
		log.Panicln("Should be private or public")
		panic("")
	}
}
