package main

import (
	"encoding/json"
	"fmt"
	"ghkv"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	ut "utils"
	"vvkv"

	"github.com/sirupsen/logrus"
)

/*
CatCmd fake cmd object
*/
var CatCmd exec.Cmd

func executeOrNil(execute bool, engine string, filepath string, args []string) bool {
	if execute {

		// TODO: We should deprecate removeEnginePrefix function
		res, cmd := removeEnginePrefix(engine)
		if !res {
			log.Panicln("engine not right: " + engine)
		}

		switch cmd {
		case "which":
			fmt.Println(filepath)
		case "cat":
			ut.Cat(filepath)
		case "jar":
			// return ut.Execute("java", append([]string{"-jar", filepath}, args...))
			ut.Watchdog(
				ArgRetry, ArgInterval,
				"java", append([]string{"-jar", filepath}, args...),
			)

		case "nosh":
			noshPath := noshexe.GetOrInstallNosh()
			ut.Watchdog(
				ArgRetry, ArgInterval,
				noshPath, append([]string{filepath}, args...),
			)
		case "nosh-js":
			noshPath := noshexe.GetOrInstallNosh()
			os.Setenv("ENGINE", "js")
			ut.Watchdog(
				ArgRetry, ArgInterval,
				noshPath, append([]string{filepath}, args...),
			)
		case "nosh-ts":
			noshPath := noshexe.GetOrInstallNosh()
			os.Setenv("ENGINE", "ts")
			ut.Watchdog(
				ArgRetry, ArgInterval,
				noshPath, append([]string{filepath}, args...),
			)
		case "cmd":
			log.Debugln("Running cmd")
			ut.Watchdog(
				ArgRetry, ArgInterval,
				args[0], ut.SliceOrEmpty(args, 1),
			)
		default:
			all := append([]string{cmd, filepath}, args...)
			log.Debugln(strings.Join(all, " "))
			// return ut.Execute(cmd, append([]string{filepath}, args...))
			ut.Watchdog(
				ArgRetry, ArgInterval,
				cmd, append([]string{filepath}, args...),
			)
		}
	}
	return true
}

/*
ExecuteHTTPURI execute command
*/
func ExecuteHTTPURI(update bool, execute bool, engine string, uri string, args []string) bool {
	md5 := ut.CalMd5(uri)
	filePath := vvkv.AppFlatDirPath + "/" + md5
	return executeHTTPURIWithTargetFilePath(update, execute, engine, uri, filePath, args)
}

func executeHTTPURIWithTargetFilePath(update bool, execute bool, engine string, uri string, filePath string, args []string) bool {
	if update || !ut.IsFileExisted(filePath) {
		ut.HTTPCat(uri, filePath)
	}
	return executeOrNil(execute, engine, filePath, args)
}

func readMeta(fileIdxPath string) *Meta {
	if !ut.IsFileExisted(fileIdxPath) {
		return nil
	}

	var meta Meta
	err := json.Unmarshal([]byte(ut.ReadFile(fileIdxPath)), &meta)
	if nil != err {
		return nil
	}
	return &meta
}

func getEngineFromCodeTypeOrExt(codeType string, engine string, resurl string) string {
	if engine != engineAuto {
		return engine
	}

	if eng, ok := SubCmd2Runtime[codeType]; ok {
		return eng
	}

	if eng, ok := parseEngineByExt(resurl); ok {
		return eng
	}
	return engineCat
}

/*
ExecuteVVKVURI execute command
*/
func ExecuteVVKVURI(update bool, execute bool, engine string, vvkvURI string, args []string) bool {

	rewriteToURL := func(prefix, targetURLRoot, targetFilePathRoot string) *bool {
		if strings.HasPrefix(vvkvURI, prefix) {
			rest := vvkvURI[len(prefix):]
			httpURL := targetURLRoot + rest
			targetFilePath := targetFilePathRoot + rest
			// ret := executeHTTPOrLocalDisk(update, execute, engine, httpURL, cmdArgs)
			engine := guessEngineUsingURI(engine, httpURL)
			ut.Mkdirp(filepath.Dir(targetFilePath))
			ret := executeHTTPURIWithTargetFilePath(update, execute, engine, httpURL, targetFilePath, args)
			return &ret
		}
		return nil
	}

	xBashFPRoot := vvkv.XPath + "/x-bash/"
	xBashURLRoot := "https://x-bash.github.io/"

	if res := rewriteToURL("@official/bash/", xBashURLRoot, xBashFPRoot); res != nil {
		return *res
	}

	if res := rewriteToURL("@bash/", xBashURLRoot, xBashFPRoot); res != nil {
		return *res
	}

	filePath := vvkv.AppDirPath + "/" + vvkvURI[1:]
	fileIdxPath := vvkv.AppIdxDirPath + "/" + vvkvURI[1:]

	if update || !ut.IsFileExisted(fileIdxPath) {
		meta := VvkvClient.GetVVKVURIToLocalDisk(vvkvURI, filePath, fileIdxPath)
		if nil == meta {
			return false
		}
	}

	meta := readMeta(fileIdxPath)
	log.WithField("engine", engine).WithField("meta", meta).Debug("Meta read")
	engine = getEngineFromCodeTypeOrExt(meta.CodeType, engine, vvkvURI)

	return ExecuteWithMeta(meta.IsURL, update, execute, engine, filePath, args)
}

/*
ExecuteGHURI e
*/
func ExecuteGHURI(update bool, execute bool, engine string, ghURI string, args []string) bool {
	log.WithFields(logrus.Fields{
		"ghURI":  ghURI,
		"update": update,
		"engine": engine,
		"args":   args,
	}).Debug("ExecuteGHURI")

	res := ghkv.CreateGitRes(ghURI)

	meta, filepath, err := res.Retrieve(update)
	if err != nil {
		log.Warn(err)
		return false
	}

	engine = getEngineFromCodeTypeOrExt(meta.CodeType, engine, ghURI)
	return ExecuteWithMeta(meta.IsURL, update, execute, engine, *filepath, args)
}

func guessEngineUsingURI(engine string, uri string) string {
	// Where to place it?
	if engine == engineAuto {
		eng, ok := parseEngineByExt(uri)
		if ok {
			return eng
		}
		return engineCat
	}
	return engine
}

func executeHTTPOrLocalDisk(update bool, execute bool, engine string, uri string, args []string) bool {
	engine = guessEngineUsingURI(engine, uri)

	if ut.IsLikeHTTPURL(uri) {
		return ExecuteHTTPURI(update, execute, engine, uri, args)
	}

	if ut.IsFileExisted(uri) {
		return executeOrNil(execute, engine, uri, args)
	}

	return false
}

/*
ExecuteWithMeta execute command
*/
func ExecuteWithMeta(isURL bool, update bool, execute bool, engine string, filePath string, args []string) bool {
	// Jump
	if isURL {
		link := readLink(filePath)
		if link.URL != nil {
			return ExecuteURI(update, execute, engine, *link.URL, args)
		}
		log.Errorln("Expect there is url field")
		return false
	}

	return executeOrNil(execute, engine, filePath, args)
}

/*
ExecuteURI execute command
*/
func ExecuteURI(update bool, execute bool, engine string, uri string, args []string) bool {
	// @gh:edwinjhlee/work
	if ghkv.IsLikeGHURL(uri) {
		return ExecuteGHURI(update, execute, engine, uri, args)
	}

	if isLikeVVURL(uri) {
		return ExecuteVVKVURI(update, execute, engine, uri, args)
	}

	return executeHTTPOrLocalDisk(update, execute, engine, uri, args)
}

/*
ExecuteURIWithComplement execute command
*/
func ExecuteURIWithComplement(update bool, execute bool, engine string, args []string) bool {
	cmd := ExecuteURI(update, execute, engine, args[0], ut.SliceOrEmpty(args, 1))
	if cmd {
		return true
	}

	partURI := args[0]
	// if strings.HasPrefix(partURI, "@") {
	if isLikeVVURL(partURI) {
		// it means it fails on ExecuteURI
		return false
	}

	cmdArgs := ut.SliceOrEmpty(args, 1)
	for _, prefix := range vvkv.GetDefaultPrefixes() {
		log.WithField("Prefix", prefix).Debug("Try Complement")
		uri := prefix + "/" + partURI
		cmd := ExecuteURI(update, execute, engine, uri, cmdArgs)
		if cmd {
			return true
		}
	}

	return false
}
