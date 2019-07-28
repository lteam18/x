package main

import (
	"os/exec"
	"path"
	"strings"
)

/*
SubCmd2Runtime sub
*/
var SubCmd2Runtime = make(map[string]string)

var prefix = "engine:"

/*
Engine a
*/
func Engine(lang string) string {
	return prefix + lang
}

var engineCat = Engine("cat")
var engineAuto = Engine("auto")

func init() {
	SubCmd2Runtime["python"] = Engine("python")
	SubCmd2Runtime["perl"] = Engine("perl")
	SubCmd2Runtime["bash"] = Engine("bash")
	SubCmd2Runtime["fish"] = Engine("fish")
	SubCmd2Runtime["jar"] = Engine("jar")
	SubCmd2Runtime["vvsh"] = Engine("vvsh")
	SubCmd2Runtime["node"] = Engine("node")
	SubCmd2Runtime["cat"] = Engine("cat")

	// shortcut
	SubCmd2Runtime["py"] = SubCmd2Runtime["python"]
	SubCmd2Runtime["pl"] = SubCmd2Runtime["perl"]
	SubCmd2Runtime["js"] = SubCmd2Runtime["vvsh"]
	SubCmd2Runtime["ts"] = SubCmd2Runtime["vvsh"]
	SubCmd2Runtime["sh"] = SubCmd2Runtime["bash"]
}

func removeEnginePrefix(str string) (bool, string) {
	if strings.HasPrefix(str, prefix) {
		return true, str[len(prefix):]
	}
	return false, str
}

// Return ext without dot
func parseExt(filepath string) string {
	ext := path.Ext(filepath)
	if len(ext) > 0 {
		return ext[1:]
	}
	return ""
}

/*
GuessEngineThenExecute guess
*/

/*
func GuessEngineThenExecute(update bool, args []string) *exec.Cmd {
	engine, ok := parseEngineByExt(args[0])
	if ok {
		return ExecuteURIWithComplement(update, true, engine, args)
	}
	return ExecuteURIWithComplement(update, true, engineCat, args)
}
*/

func parseEngineByExt(filepath string) (string, bool) {
	ext := parseExt(filepath)
	ret, ok := SubCmd2Runtime[ext]
	return ret, ok
}

/*
ExecuteBySubCmd execute command
*/
func ExecuteBySubCmd(subCmd string, args []string) *exec.Cmd {
	updateFirst := false
	if strings.HasSuffix(subCmd, "!") {
		subCmd = subCmd[0 : len(subCmd)-1]
		updateFirst = true
	}
	engine, ok := SubCmd2Runtime[subCmd]
	if ok {
		return ExecuteURIWithComplement(updateFirst, true, engine, args)
	}
	return nil
}
