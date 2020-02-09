package main

import (
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

var engineWhich = Engine("which")
var engineCat = Engine("cat")
var engineAuto = Engine("auto")

func init() {

	SubCmd2Runtime["python"] = Engine("python")
	SubCmd2Runtime["perl"] = Engine("perl")
	SubCmd2Runtime["bash"] = Engine("bash")
	SubCmd2Runtime["fish"] = Engine("fish")
	SubCmd2Runtime["jar"] = Engine("jar")

	SubCmd2Runtime["nosh"] = Engine("nosh")

	SubCmd2Runtime["js"] = Engine("nosh-js")
	SubCmd2Runtime["ts"] = Engine("nosh-ts")

	SubCmd2Runtime["node"] = Engine("node")
	SubCmd2Runtime["cat"] = Engine("cat")
	SubCmd2Runtime["which"] = Engine("which")

	// shortcut
	SubCmd2Runtime["py"] = SubCmd2Runtime["python"]
	SubCmd2Runtime["pl"] = SubCmd2Runtime["perl"]

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
func ExecuteBySubCmd(subCmd string, args []string) bool {
	updateFirst := false
	if strings.HasSuffix(subCmd, "!") {
		subCmd = subCmd[0 : len(subCmd)-1]
		updateFirst = true
	}
	engine, ok := SubCmd2Runtime[subCmd]
	log.WithField("subcmd", subCmd).WithField("engine", engine).Debugln("subcmd ready")
	if ok {
		// What if subCmd is empty, just open the REPL perhaps ?
		return ExecuteURIWithComplement(updateFirst, true, engine, args)
	}
	return false
}
