package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"strings"
	ut "utils"
	"vvkv"
)

/*
CatCmd fake cmd object
*/
var CatCmd exec.Cmd

func executeOrNil(execute bool, engine string, filepath string, args []string) bool {
	if execute {
		res, cmd := removeEnginePrefix(engine)
		if !res {
			log.Panicln("engine not right: " + engine)
		}

		switch cmd {
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
			log.Infoln("Running cmd")
			ut.Watchdog(
				ArgRetry, ArgInterval,
				args[0], ut.SliceOrEmpty(args, 1),
			)
		default:
			all := append([]string{cmd, filepath}, args...)
			log.Infof(strings.Join(all, " "))
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

/*
ExecuteVVKVURI execute command
*/
func ExecuteVVKVURI(update bool, execute bool, engine string, vvkvURI string, args []string) bool {
	filePath := vvkv.AppDirPath + "/" + vvkvURI[1:]
	fileIdxPath := vvkv.AppIdxDirPath + "/" + vvkvURI[1:]

	if update || !ut.IsFileExisted(fileIdxPath) {
		meta := VvkvClient.GetVVKVURIToLocalDisk(vvkvURI, filePath, fileIdxPath)
		if nil == meta {
			return false
		}
	}

	meta := readMeta(fileIdxPath)
	log.WithField("engine", engine).WithField("meta", meta).Infoln("Meta read")

	if engine == engineAuto {
		eng, ok := SubCmd2Runtime[meta.CodeType]
		if ok {
			engine = eng
		} else {
			engine = engineCat
		}
		log.WithField("engine", engine).WithField("codetype", meta.CodeType).Infoln("Setting engine by codetype")
	}

	// Just update the link
	if meta.IsURL {
		url := readLink(filePath).URL
		if isLikeVVURL(url) {
			return ExecuteVVKVURI(update, execute, engine, url, args)
		}

		return executeHTTPOrLocalDisk(update, execute, engine, url, args)
	}

	return executeOrNil(execute, engine, filePath, args)
}

func executeHTTPOrLocalDisk(update bool, execute bool, engine string, uri string, args []string) bool {

	// Where to place it?
	if engine == engineAuto {
		eng, ok := parseEngineByExt(uri)
		if ok {
			engine = eng
		} else {
			engine = engineCat
		}
	}

	if ut.IsLikeHTTPURL(uri) {
		return ExecuteHTTPURI(update, execute, engine, uri, args)
	}

	if ut.IsFileExisted(uri) {
		return executeOrNil(execute, engine, uri, args)
	}

	return false
}

/*
ExecuteURI execute command
*/
func ExecuteURI(update bool, execute bool, engine string, args []string) bool {

	uri := args[0]
	cmdArgs := ut.SliceOrEmpty(args, 1)
	if isLikeVVURL(uri) {
		return ExecuteVVKVURI(update, execute, engine, uri, cmdArgs)
	}

	return executeHTTPOrLocalDisk(update, execute, engine, uri, cmdArgs)
}

/*
ExecuteURIWithComplement execute command
*/
func ExecuteURIWithComplement(update bool, execute bool, engine string, args []string) bool {
	cmd := ExecuteURI(update, execute, engine, args)
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
		log.WithField("Prefix", prefix).Infof("Try Complement")
		uri := prefix + "/" + partURI
		cmd := ExecuteURI(update, execute, engine, append([]string{uri}, cmdArgs...))
		if cmd {
			return true
		}
	}

	return false
}
