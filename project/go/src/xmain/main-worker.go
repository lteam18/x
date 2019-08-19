package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	ut "utils"
	"vvkv"

	"github.com/urfave/cli"
)

var installXCmd = cli.Command{
	Name:  "install-x",
	Usage: "`x install-x`",
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			install(nil, nil)
		} else {
			installToDst(context.Args().Get(0))
		}
		return nil
	},
}

var installCmd = cli.Command{
	Name:    "install",
	Usage:   "`x install`",
	Aliases: []string{"install-all"},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			install(nil, nil)
		} else {
			installToDst(context.Args().Get(0))
		}
		println("x installed.")
		noshexe.GetOrInstallNosh()
		println("Nosh installed.")
		return nil
	},
}

var installNoshCmd = cli.Command{
	Name:  "install-nosh",
	Usage: "`x install-nosh`",
	Action: func(context *cli.Context) error {
		noshexe.GetOrInstallNosh()
		return nil
	},
}

var updateNoshCmd = cli.Command{
	Name:  "update-nosh",
	Usage: "`x update-nosh`",
	Action: func(context *cli.Context) error {
		noshexe.Upgrade()
		return nil
	},
}

var tokenCmd = cli.Command{
	Name:  "token",
	Usage: "display current token: `x token`",
	Action: func(context *cli.Context) error {
		println(vvkv.ReadToken())
		return nil
	},
}

var tokenApplyCmd = cli.Command{
	Name:      "token-apply",
	Usage:     "token-apply [get] [put]: `x token-apply username,orgname username`",
	UsageText: "`x token-apply username,orgname username`",
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 3 {
			return fmt.Errorf("Missing get put data")
		}
		get := strings.Split(ctx.Args().Get(0), ",")
		put := strings.Split(ctx.Args().Get(1), ",")
		info := ""
		if len(ctx.Args()) >= 3 {
			info = ctx.Args().Get(3)
		}
		token := VvkvClient.ApplyToken(get, put, info, time.Now().Unix()+1*365*24*60*60)
		fmt.Println(token)
		return nil
	},
}

var tokenSetCmd = cli.Command{
	Name: "token-set",
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) == 1 {
			return fmt.Errorf("please provide token")
		}
		vvkv.WriteTokenSync(ctx.Args().Get(0))
		fmt.Println("Token already written.")
		return nil
	},
}

var tokenDecryptCmd = cli.Command{
	Name: "token-decrypt",
	Action: func(ctx *cli.Context) error {
		var token string
		if len(ctx.Args()) == 0 {
			log.Infoln("Decrypt the token stored in home directory.")
			token = vvkv.ReadToken()
		} else {
			token = ctx.Args().Get(0)
		}
		tokenObj := VvkvClient.DecryptToken(token)
		tokenStr := ut.PrettifyJSONString(tokenObj)
		println(tokenStr)
		return nil
	},
}

var lsCmd = cli.Command{
	Name:    "ls",
	Aliases: []string{"ll"},
	Action: func(ctx *cli.Context) error {
		query := "@official"
		if len(ctx.Args()) >= 1 {
			query = ctx.Args().Get(0)
		}

		detailed := "ll" == ctx.Parent().Args().First()
		if strings.HasPrefix(query, "@") {
			printLs(query, detailed)
		} else {
			printLs("@official/"+query, detailed)
		}
		return nil
	},
}

var catCmd = cli.Command{
	Name:    "cat",
	Aliases: []string{"cat!"},
	Action: func(ctx *cli.Context) error {
		update := "cat!" == ctx.Parent().Args().First()
		for _, v := range ctx.Args() {
			ExecuteURIWithComplement(update, true, engineCat, []string{v})
		}
		return nil
	},
}

var runCmd = cli.Command{
	Name:    "run",
	Aliases: []string{"run!"},
	Action: func(ctx *cli.Context) error {
		update := "run!" == ctx.Parent().Args().First()
		ExecuteURIWithComplement(update, true, engineAuto, ctx.Args())
		return nil
	},
}

var cmdCmd = cli.Command{
	Name:  "cmd",
	Usage: "x --retry 1 --interval 10 cmd ping www.bing.com",
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) == 0 {
			return fmt.Errorf("please provide commands")
		}
		cmd := ctx.Args().Get(0)
		args := ctx.Args()[1:]
		// TODO: Consider using argument
		ut.Watchdog(ArgRetry, ArgInterval, cmd, args)
		return nil
	},
}

var updateCmd = cli.Command{
	Name:  "update",
	Usage: "`x update @dryrun/work`",
	Action: func(ctx *cli.Context) error {
		ExecuteURIWithComplement(true, false, "", ctx.Args())
		return nil
	},
}

var shareCmd = cli.Command{
	Name:  "share",
	Usage: "Generate the url avaiable for 1 day: `x share @dryrun/work`",
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) == 0 {
			return fmt.Errorf("please provice [vvurl]")
		}
		vvurl := ctx.Args().Get(0)
		url := VvkvClient.Share(vvurl, 24*60*60)
		println("Generate url for 1 day")
		println(url)
		return nil
	},
}

var setAccessCmd = cli.Command{
	Name:  "setaccess",
	Usage: "`x setaccess private @dryrun/work`",
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 2 {
			return fmt.Errorf("please provice [access] [url]")
		}
		isPublic := parseAccess(ctx.Args().Get(0))
		vvurl := ctx.Args().Get(1)
		VvkvClient.SetPermissionBYVVURL(vvurl, isPublic)
		return nil
	},
}

var uploadCmd = cli.Command{
	Name:    "upload",
	Aliases: []string{"link"},
	Usage:   "x upload private @dryrun/work hi.js",
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 3 {
			panic("Please privide more command")
		}

		isPublic := parseAccess(ctx.Args().Get(0))
		vvurl := ctx.Args().Get(1)
		fp := ctx.Args().Get(2)

		var codetype string
		if len(ctx.Args()) == 3 {
			codetype = parseExt(fp)
		} else {
			codetype = ctx.Args().Get(3)
		}

		subCmd := ctx.Parent().Args().First()
		exe := uploadCode
		if "link" == subCmd {
			exe = uploadLink
		}
		exe(fp, vvurl, isPublic, codetype)
		fmt.Printf("Command [%s]: Success!\n", subCmd)
		return nil
	},
}

var app = cli.NewApp()

/*
ArgInterval arg for interval
*/
var ArgInterval int64

/*
ArgRetry arg for retry
*/
var ArgRetry int

func init() {
	app.Name = "x"
	app.Usage = `The program x is :
	- executor for web resources
	- easily share code and resource with your friends
	`

	cli.AppHelpTemplate = `
{{.Name}} - {{.Usage}}
USAGE:
	{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
	{{if len .Authors}}
AUTHOR:
	{{range .Authors}}{{ . }}{{end}}
	{{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
	{{range .VisibleFlags}}{{.}}
	{{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
	{{.Copyright}}
	{{end}}{{if .Version}}
VERSION:
	{{.Version}}
	{{end}}
`

	app.Commands = []cli.Command{
		tokenCmd,
		tokenApplyCmd,
		tokenSetCmd,
		tokenDecryptCmd,
		lsCmd,
		catCmd,
		runCmd,
		cmdCmd,
		updateCmd,
		shareCmd,
		setAccessCmd,
		uploadCmd,
		installCmd,
		installXCmd,
		installNoshCmd,
		updateNoshCmd,
	}

	/*
		app.Flags = []cli.Flag{
			cli.Int64Flag{
				Name:        "interval",
				Value:       3000,
				Usage:       "Interval",
				Destination: &ArgInterval,
			},
			cli.IntFlag{
				Name:        "retry",
				Value:       0,
				Usage:       "max retry times",
				Destination: &ArgRetry,
			},
		}
	*/

	app.Action = func(ctx *cli.Context) error {

		if len(ctx.Args()) == 0 {
			cli.ShowAppHelp(ctx)
			return nil
		}

		subCmd := ctx.Args().Get(0)
		args := ut.SliceOrEmpty(ctx.Args(), 1)

		if ExecuteBySubCmd(subCmd, args) {
			return nil
		}

		log.Info("Try to complete the prefix")
		if ExecuteURIWithComplement(false, true, engineAuto, ctx.Args()) {
			return nil
		}

		fmt.Printf("%q is not a resource name.\n", os.Args[1])

		// Error
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
		// TODO: show help
		return nil
	}
}

func runApp() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
