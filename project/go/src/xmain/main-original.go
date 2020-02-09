package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	ut "utils"
	"vvkv"
)

func runAppOriginal() {
	subCmd := os.Args[1]

	switch subCmd {
	case "install":
		if 3 <= len(os.Args) {
			installToDst(os.Args[2])
		} else {
			install(nil, nil)
		}
	case "token":
		println(vvkv.ReadToken())
	case "token-apply":
		get := strings.Split(os.Args[2], ",")
		put := strings.Split(os.Args[3], ",")
		info := ""
		if len(os.Args) >= 5 {
			info = strings.Join(os.Args[4:], " ")
		}
		token := VvkvClient.ApplyToken(get, put, info, time.Now().Unix()+1*365*24*60*60)
		fmt.Println(token)
	case "token-set":
		if len(os.Args) == 2 {
			log.Panicln("Please provide token.")
		}
		vvkv.WriteTokenSync(os.Args[2])
		fmt.Println("Token already written.")
	case "token-decrypt":
		var token string
		if len(os.Args) == 2 {
			log.Infoln("Decrypt the token stored in home directory.")
			token = vvkv.ReadToken()
		} else {
			token = os.Args[2]
		}
		tokenObj := VvkvClient.DecryptToken(token)
		tokenStr := ut.PrettifyJSONString(tokenObj)
		println(tokenStr)
	case "ls":
		fallthrough
	case "ll":
		var query string
		if 2 == len(os.Args) {
			query = "@official"
		} else {
			query = os.Args[2]
		}

		if strings.HasPrefix(query, "@") {
			printLs(query, "ll" == subCmd)
		} else {
			printLs("@official/"+query, "ll" == subCmd)
		}
	case "cat":
		fallthrough
	case "cat!":
		for _, v := range os.Args[2:] {
			ExecuteURIWithComplement("cat!" == subCmd, true, engineCat, []string{v})
		}
	case "run":
		fallthrough
	case "run!":
		ExecuteURIWithComplement("run!" == subCmd, true, engineAuto, os.Args[2:])
	case "update":
		ExecuteURIWithComplement(true, false, "", os.Args[2:])
	case "setaccess":
		if len(os.Args) <= 4 {
			panic("Please privide more command")
		}
		isPublic := parseAccess(os.Args[2])
		vvurl := os.Args[3]
		VvkvClient.SetPermissionBYVVURL(vvurl, isPublic)

	case "share":
		if len(os.Args) == 2 {
			println("Please provide keypath.\nx share [@org/keypath]")
			return
		}
		vvurl := os.Args[2]
		url := VvkvClient.Share(vvurl, 24*60*60)
		println("Generate url for 1 day")
		println(url)
	case "upload":
		fallthrough
	case "link":
		if len(os.Args) <= 4 {
			panic("Please privide more command")
		}

		isPublic := parseAccess(os.Args[2])
		vvurl := os.Args[3]
		fp := os.Args[4]

		var codetype string
		if len(os.Args) == 5 {
			codetype = parseExt(fp)
		} else {
			// len(os.Args) == 6
			codetype = os.Args[5]
		}

		exe := uploadCode
		if "link" == subCmd {
			exe = uploadLink
		}
		exe(fp, vvurl, isPublic, codetype)
		fmt.Printf("Command [%s]: Success!\n", subCmd)
	default:
		if ExecuteBySubCmd(subCmd, os.Args[2:]) {
			return
		}

		log.Info("Try to complete the prefix")
		if ExecuteURIWithComplement(false, true, engineAuto, os.Args[1:]) {
			return
		}

		// Error
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}
}
