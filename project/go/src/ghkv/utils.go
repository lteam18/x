package ghkv

import (
	"debuglog"
	"gh"
	"strings"
)

var log = debuglog.Create("ghkv")

var CODE_REPO = "x-cmd.com"

/*
CreateGitFile c
*/
func CreateGitFile(ghurl string) *gh.File {
	owner, repo, kp := ParseGHPrefix(ghurl)
	return &gh.File{Owner: owner, Repo: repo, Keypath: kp}
}

/*
IsLikeGHURL a
*/
func IsLikeGHURL(ghuri string) bool {
	return strings.HasPrefix(ghuri, "@gh")
}

/*
ParseGHPrefix a
*/
func ParseGHPrefix(uri string) (string, string, string) {
	// @gh:edwinjhlee:repo/abc
	as := strings.SplitN(uri, "/", 2)
	var firstPart, secondPart string
	if len(as) == 2 {
		firstPart = as[0]
		secondPart = as[1]
	} else {
		firstPart = as[0]
		secondPart = ""
	}

	arrGhOnwerRepo := strings.Split(firstPart, ":")
	if len(arrGhOnwerRepo) > 2 {
		return arrGhOnwerRepo[1], arrGhOnwerRepo[2], secondPart
	}

	if len(arrGhOnwerRepo) == 2 {
		return arrGhOnwerRepo[1], CODE_REPO, secondPart
	}

	// if len(arrGhOnwerRepo) == 1 {

	tk := gh.GetToken()
	if tk == nil {
		// TODO: Put it outside, or return nil
		panic("Token is Empty")
		// log.Debug("Token is Empty")
	}

	return tk.Owner, CODE_REPO, secondPart
}
