package main

import (
	"encoding/json"
	"os/exec"
	"strings"
	ut "utils"
	"vvkv"
)

/*
CatCmd fake cmd object
*/
var CatCmd exec.Cmd

func executeOrNil(execute bool, engine string, filepath string, args []string) *exec.Cmd {
	if execute {
		res, cmd := removeEnginePrefix(engine)
		if !res {
			log.Panicln("engine not right: " + engine)
		}

		switch cmd {
		case "cat":
			ut.Cat(filepath)
			return &CatCmd
		case "jar":
			return ut.Execute("java", append([]string{"-jar", filepath}, args...))
		default:
			return ut.Execute(cmd, append([]string{filepath}, args...))
		}
	}
	return nil
}

/*
ExecuteHTTPURI execute command
*/
func ExecuteHTTPURI(update bool, execute bool, engine string, uri string, args []string) *exec.Cmd {
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
func ExecuteVVKVURI(update bool, execute bool, engine string, vvkvURI string, args []string) *exec.Cmd {
	filePath := vvkv.AppDirPath + "/" + vvkvURI[1:]
	fileIdxPath := vvkv.AppIdxDirPath + "/" + vvkvURI[1:]

	if update || !ut.IsFileExisted(fileIdxPath) {
		meta := VvkvClient.GetVVKVURIToLocalDisk(vvkvURI, filePath, fileIdxPath)
		if nil == meta {
			return nil
		}
	}

	meta := readMeta(fileIdxPath)
	if engine == engineAuto {
		eng, ok := SubCmd2Runtime[meta.CodeType]
		if ok {
			engine = eng
		} else {
			engine = engineCat
		}
	}

	// Just update the link
	if meta.IsURL {
		return ExecuteHTTPURI(update, execute, engine, readLink(filePath).URL, args)
	}
	return executeOrNil(execute, engine, filePath, args)
}

/*
ExecuteURI execute command
*/
func ExecuteURI(update bool, execute bool, engine string, args []string) *exec.Cmd {

	uri := args[0]
	if strings.HasPrefix(uri, "@") {
		return ExecuteVVKVURI(update, execute, engine, uri, args[1:])
	}

	if engine == engineAuto {
		eng, ok := parseEngineByExt(uri)
		if ok {
			engine = eng
		} else {
			engine = engineCat
		}
	}

	if ut.IsLikeHTTPURL(uri) {
		return ExecuteHTTPURI(update, execute, engine, uri, args[1:])
	}

	if ut.IsFileExisted(uri) {
		return executeOrNil(execute, engine, args[0], args[1:])
	}

	return nil
}

/*
ExecuteURIWithComplement execute command
*/
func ExecuteURIWithComplement(update bool, execute bool, engine string, args []string) *exec.Cmd {
	cmd := ExecuteURI(update, execute, engine, args)
	if nil != cmd {
		return cmd
	}

	partURI := args[0]
	if strings.HasPrefix(partURI, "@") {
		return nil
	}

	for _, prefix := range vvkv.GetDefaultPrefixes() {
		log.WithField("Prefix", prefix).Infof("Try Complement")
		uri := prefix + "/" + partURI
		cmd := ExecuteURI(update, execute, engine, append([]string{uri}, args[1:]...))
		if nil != cmd {
			return cmd
		}
	}

	return nil
}
