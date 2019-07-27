package main

import (
	"fluent/list"
	"fmt"
	"strings"
)

// Actual code

func ci(name string, description string) string {
	return fmt.Sprintf("\t%s\t%s", name, description)
}

func displayHelpInfo() {
	info := list.Str(
		"usage: x <command> [<args>]",
		"The most commonly used vvx commands are: ",
		ci("cat", "show files"),
		ci("list", "list files"),
		ci("fish", "run fish file"),
		ci("token", "run node file"),
		ci("token-decrypt", "run node file"),
	)

	fmt.Println(strings.Join(info, "\n"))
}

//

type argumentItem struct {
	Name        string
	Description []string
	Example     []string
	Type        string
}

type helpCmdItem struct {
	Name        string
	Description []string
	Examples    []string
	Alias       []string
	Args        []argumentItem
}

/*
Help a
*/
type Help struct {
	Description []string
	Cmds        []helpCmdItem
}

var arg = argumentItem{
	"",
	[]string{""},
	[]string{""},
	"1",
}

var helpcmd = helpCmdItem{
	"Hello",
	[]string{""},
	[]string{""},
	[]string{""},
	[]argumentItem{
		arg,
	},
}

var help = Help{
	[]string{
		"Right",
	},
	[]helpCmdItem{
		helpcmd,
	},
}
