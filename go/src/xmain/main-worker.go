package main

import (
	"debuglog"
	"encoding/json"
	"errors"
	"fmt"
	"gh"
	"ghkv"
	"os"
	"strconv"
	"strings"
	"time"
	ut "utils"
	"vvkv"

	"github.com/urfave/cli"
)

func handleError(err error) error {
	// panic(err)
	// log.Debug(err)
	// pc, fn, line, _ := runtime.Caller(1)

	// log.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
	// check
	if debuglog.IsDebug("x") {
		panic(err)
	}
	return err
}

func writeResult(format string, args ...interface{}) {
	os.Stdout.WriteString(fmt.Sprintf(format, args...))
}

var versionCmd = cli.Command{
	Name:    "version",
	Aliases: []string{"v"},
	Action: func(context *cli.Context) error {
		writeResult("v0.0.1")
		return nil
	},
}

var installXCmd = cli.Command{
	Name:  "install-x",
	Usage: "`x install-x`",
	Action: func(context *cli.Context) error {
		args := context.Args()
		if args.Len() < 1 {
			install(nil, nil)
		} else {
			installToDst(args.Get(0))
		}
		return nil
	},
}

var installCmd = cli.Command{
	Name:    "install",
	Usage:   "`x install`",
	Aliases: []string{"install-all"},
	Action: func(context *cli.Context) error {
		args := context.Args()
		if args.Len() < 1 {
			install(nil, nil)
		} else {
			installToDst(args.Get(0))
		}
		log.Info("x installed.")
		noshexe.GetOrInstallNosh()
		log.Info("Nosh installed.")
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
		args := ctx.Args()
		if args.Len() < 3 {
			return handleError(fmt.Errorf("Missing get put data"))
		}
		get := strings.Split(ctx.Args().Get(0), ",")
		put := strings.Split(ctx.Args().Get(1), ",")
		info := ""
		if args.Len() >= 3 {
			info = ctx.Args().Get(3)
		}
		token := VvkvClient.ApplyToken(get, put, info, time.Now().Unix()+1*365*24*60*60)
		writeResult(token)
		return nil
	},
}

var tokenSetCmd = cli.Command{
	Name: "token-set",
	Action: func(ctx *cli.Context) error {
		args := ctx.Args()
		if args.Len() != 1 {
			return handleError(fmt.Errorf("require only one argument for token"))
		}
		vvkv.WriteTokenSync(ctx.Args().Get(0))
		log.Info("Token already written")
		return nil
	},
}

var tokenDecryptCmd = cli.Command{
	Name: "token-decrypt",
	Action: func(ctx *cli.Context) error {
		args := ctx.Args()
		var token string
		if args.Len() == 0 {
			log.Debug("Decrypt the token stored in home directory.")
			token = vvkv.ReadToken()
		} else {
			token = ctx.Args().Get(0)
		}
		tokenObj := VvkvClient.DecryptToken(token)
		tokenStr := ut.PrettifyJSONString(tokenObj)
		writeResult(tokenStr)

		amap := make(map[string]interface{})
		json.Unmarshal([]byte(tokenObj), &amap)
		time1 := int64(amap["expired"].(float64))

		log.Info("Expired At: " + time.Unix(time1, 0).Format("2006-01-02 03:04:05 PM"))

		return nil
	},
}

// How to add https
// Using center service to route
// Add login page
// Add encrypt once page
// Add third party authentication

var httpcmd = cli.Command{
	Name: "http",
	Action: func(ctx *cli.Context) error {

		if ctx.Args().Len() < 2 {
			log.Info("Two arguments needed: x http <port> <directory>")
			// TODO: using message less strict, at least without FATAL
			return handleError(errors.New("Two arguments needed: x http <port> <directory>"))
		}

		port, err := strconv.Atoi(ctx.Args().Get(0))
		if err != nil {
			return handleError(errors.New("Port should be number between 0 - 65535"))
		}
		dir := ctx.Args().Get(1)

		Serve(port, dir)

		return nil
	},
}

var lsCmd = cli.Command{
	Name:    "ls",
	Aliases: []string{"ll"},
	Action: func(ctx *cli.Context) error {
		query := "@official"
		if ctx.Args().Len() >= 1 {
			query = ctx.Args().Get(0)
		}

		detailed := "ll" == ctx.Lineage()[1].Args().Get(0)

		if ghkv.IsLikeGHURL(query) {
			q := ghkv.LsQueryStr(query)
			if q == nil {
				writeResult("Nothing")
			} else {
				writeResult(*q + "\n")
			}
			return nil
		}

		if strings.HasPrefix(query, "@") {
			printLs(query, detailed)
		} else {
			printLs("@official/"+query, detailed)
		}
		return nil
	},
}

var whichCmd = cli.Command{
	Name:    "which",
	Usage:   "which [...resource-name] will return local file path of each resource",
	Aliases: []string{"which!"},
	Action: func(ctx *cli.Context) error {
		update := "which!" == ctx.Lineage()[1].Args().Get(0)
		if ctx.Args().Len() == 0 {
			cli.ShowCommandHelp(ctx, ctx.Command.Name)
			return handleError(fmt.Errorf("please provide commands"))
		}
		for _, v := range ctx.Args().Slice() {
			ExecuteURIWithComplement(update, true, engineWhich, []string{v})
		}
		return nil
	},
}

var catCmd = cli.Command{
	Name:    "cat",
	Usage:   "cat [...resource-name] will print the content of resource to stdout",
	Aliases: []string{"cat!"},
	Action: func(ctx *cli.Context) error {
		update := "cat!" == ctx.Lineage()[1].Args().Get(0)
		for _, v := range ctx.Args().Slice() {
			ExecuteURIWithComplement(update, true, engineCat, []string{v})
		}
		return nil
	},
}

var runCmd = cli.Command{
	Name:    "run",
	Aliases: []string{"run!"},
	Action: func(ctx *cli.Context) error {
		update := "run!" == ctx.Lineage()[1].Args().Get(0)
		ExecuteURIWithComplement(update, true, engineAuto, ctx.Args().Slice())
		return nil
	},
}

var cmdCmd = cli.Command{
	Name:  "cmd",
	Usage: "x --retry 1 --interval 10 cmd ping www.bing.com",
	Action: func(ctx *cli.Context) error {
		if ctx.Args().Len() == 0 {
			return handleError(fmt.Errorf("please provide commands"))
		}
		cmd := ctx.Args().Get(0)
		args := ctx.Args().Slice()[1:]
		// TODO: Consider using argument
		ut.Watchdog(ArgRetry, ArgInterval, cmd, args)
		return nil
	},
}

var execCmd = cli.Command{
	Name:  "exec",
	Usage: "x --retry 1 --interval 10 exec ping www.bing.com",
	Action: func(ctx *cli.Context) error {
		if ctx.Args().Len() == 0 {
			return handleError(fmt.Errorf("please provide commands"))
		}

		return handleError(ut.ExecReplace(ctx.Args().Get(0), ctx.Args().Slice()[1:]))
	},
}

var updateCmd = cli.Command{
	Name:  "update",
	Usage: "`x update @dryrun/work`",
	Action: func(ctx *cli.Context) error {
		ExecuteURIWithComplement(true, false, "", ctx.Args().Slice())
		return nil
	},
}

var shareCmd = cli.Command{
	Name:  "share",
	Usage: "Generate the url avaiable for 1 day: `x share @dryrun/work`",
	Action: func(ctx *cli.Context) error {
		if ctx.Args().Len() == 0 {
			return handleError(fmt.Errorf("please provice [vvurl]"))
		}
		vvurl := ctx.Args().Get(0)

		if ghkv.IsLikeGHURL(vvurl) {
			if ctx.Args().Len() != 1 {
				log.Info("Cannot control the duration of this github store resource. It might last a day long.")
			}
			codetype, url, err := ghkv.CreateGitRes(vvurl).GetCodeTypeAndTempURL()
			if err != nil {
				return err
			}

			cmd := fmt.Sprintf("x %s %s", *codetype, *url)
			writeResult(cmd)
			return nil
		}

		seconds := int64(24 * 60 * 60)
		if ctx.Args().Len() == 1 {
			log.Info("Generate url for 1 day")
		} else {
			seconds = ut.DateStr2Seconds(ctx.Args().Get(1))
		}
		url := VvkvClient.Share(vvurl, seconds)
		writeResult(url)
		return nil
	},
}

var setAccessCmd = cli.Command{
	Name:  "set-access",
	Usage: "`x set-access private @dryrun/work`",
	Action: func(ctx *cli.Context) error {
		if ctx.Args().Len() < 2 {
			return handleError(fmt.Errorf("please provice [access] [url]"))
		}
		isPublic := parseAccess(ctx.Args().Get(0))
		vvurl := ctx.Args().Get(1)

		if ghkv.IsLikeGHURL(vvurl) {
			res := ghkv.CreateGitRes(vvurl)
			return handleError(res.SetAccess(isPublic))
		}

		VvkvClient.SetPermissionBYVVURL(vvurl, isPublic)
		return nil
	},
}

var deleteCmd = cli.Command{
	Name:  "delete",
	Usage: "`x delete @dryrun/work`, or `x delete @gh/hi/py`",
	Action: func(ctx *cli.Context) error {
		if ctx.Args().Len() < 1 {
			return handleError(fmt.Errorf("please provice [access] [url]"))
		}
		vvurl := ctx.Args().Get(0)

		if ghkv.IsLikeGHURL(vvurl) {
			res := ghkv.CreateGitRes(vvurl)
			return handleError(res.Delete())
		}

		// VvkvClient.SetPermissionBYVVURL(vvurl, isPublic)
		return handleError(errors.New("Not implemented"))
	},
}

var ghTokenSetUpCmd = cli.Command{
	Name:  "gh-setup",
	Usage: "`x gh-setup <username> <GITHUB_TOKEN_HAVE_FULL_ACCESS_FOR_REPOSITORY>`",
	Action: func(ctx *cli.Context) error {
		if ctx.Args().Len() < 2 {
			return handleError(fmt.Errorf("please provice [username] [token]"))
		}

		owner := ctx.Args().Get(0)
		log.Infoln("Token written into '~/.x-cmd.com/GH_TOKEN'")
		gh.SetToken(owner, ctx.Args().Get(1))

		log.Infoln("Prepare for Repo'")

		if e := ghkv.InitGHRes(owner); e != nil {
			// TODO: Perhaps repo already initialized.
			// log.Infoln("Create Repo and init repo encountering errors")
			log.Debug(e)
		}
		return nil
	},
}

var ghTokenCmd = cli.Command{
	Name:  "gh-token",
	Usage: "`x gh-token`",
	Action: func(ctx *cli.Context) error {
		token := gh.GetToken()
		if token == nil {
			return handleError(fmt.Errorf("Token is inavailable. You could check in '~/.x-cmd.com/GH_TOKEN'"))
		}
		writeResult("%s %s", token.Owner, token.Token)
		return nil
	},
}

/*
Consider retire this command
*/
var ghCatCmd = cli.Command{
	Name:  "gh-cat",
	Usage: "x gh-cat <user> <repo> <filepath>",
	Action: func(ctx *cli.Context) error {
		if ctx.Args().Len() < 3 {
			return handleError(fmt.Errorf("please provice <user> <repo> <filepath>"))
		}
		args := ctx.Args()

		gf := &gh.File{Owner: args.Get(0), Repo: args.Get(1), Keypath: args.Get(2)}
		s, e := gf.Cat()
		if e != nil {
			return handleError(e)
		}
		writeResult(s)
		return nil
	},
}

var ghInitCmd = cli.Command{
	Name:  "gh-init",
	Usage: "`x gh-init`",
	Action: func(ctx *cli.Context) error {
		token := gh.GetToken()
		if token == nil {
			return handleError(fmt.Errorf("Token is inavailable. You should run `x gh-token-set <YOUR_GITHUB_TOKEN>` first. You could check the file in '~/.x-cmd.com/GH_TOKEN'"))
		}
		return handleError(ghkv.InitGHRes(token.Owner))
	},
}

var ghRepoDeleteCmd = cli.Command{
	Name:  "gh-repo-delete",
	Usage: "`x gh-repo-delete`",
	Action: func(ctx *cli.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return handleError(fmt.Errorf("please provice <user> <repo>"))
		}
		return gh.DeleteRepo(args.Get(0), args.Get(1))
	},
}

var uploadCmd = cli.Command{
	Name:    "upload",
	Aliases: []string{"link"},
	Usage:   "x upload private @dryrun/work hi.js",
	Action: func(ctx *cli.Context) error {
		if ctx.Args().Len() < 3 {
			// panic("Please privide more command")
			return handleError(fmt.Errorf("Please privide more command: \nx upload <private|public> <full-url> <filepath>; \nx link <private|public> <full-url> <target-url>"))
		}

		isPublic := parseAccess(ctx.Args().Get(0))
		vvurl := ctx.Args().Get(1)
		fp := ctx.Args().Get(2)

		var codetype string
		if ctx.Args().Len() == 3 {
			codetype = parseExt(fp)
		} else {
			codetype = ctx.Args().Get(3)
		}

		subCmd := ctx.Lineage()[1].Args().Get(0)
		if ghkv.IsLikeGHURL(vvurl) {
			res := ghkv.CreateGitRes(vvurl)
			exe := res.UploadFile // ghkv.UploadGHRes
			if "link" == subCmd {
				exe = res.UploadURL
			}
			e := exe(fp, isPublic, codetype)
			if e != nil {
				return handleError(e)
			}
		} else {
			exe := uploadCode
			if "link" == subCmd {
				exe = uploadLink
			}
			exe(fp, vvurl, isPublic, codetype)
		}

		log.Warnf("Command [%s]: Success!\n", subCmd)
		return nil
	},
}

var testCmd = cli.Command{
	Name:  "test",
	Usage: "`x test`",
	Action: func(ctx *cli.Context) error {
		// fmt.Println(gh.EnablePages("edwinjhlee", CODE_REPO))
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

	app.Commands = []*cli.Command{

		&versionCmd,

		&tokenCmd,
		&tokenApplyCmd,
		&tokenSetCmd,
		&tokenDecryptCmd,
		&lsCmd,
		&whichCmd,
		&catCmd,
		&runCmd,
		&cmdCmd,
		&execCmd,

		&updateCmd,
		&shareCmd,
		&setAccessCmd,
		&uploadCmd,
		&deleteCmd,

		&installCmd,
		&installXCmd,
		&installNoshCmd,
		&updateNoshCmd,

		&ghTokenCmd,
		&ghTokenSetUpCmd,
		&ghCatCmd,
		&ghInitCmd,
		&ghRepoDeleteCmd,

		&httpcmd,

		&testCmd,
	}

	app.Flags = []cli.Flag{
		&cli.Int64Flag{
			Name:        "interval, i",
			Value:       3000,
			Usage:       "Interval",
			Destination: &ArgInterval,
		},
		&cli.IntFlag{
			Name:        "retry, r",
			Value:       0,
			Usage:       "max retry times",
			Destination: &ArgRetry,
		},
	}

	app.Action = func(ctx *cli.Context) error {

		if ctx.Args().Len() == 0 {
			cli.ShowAppHelp(ctx)
			return nil
		}

		subCmd := ctx.Args().Get(0)
		args := ut.SliceOrEmpty(ctx.Args().Slice(), 1)

		if ExecuteBySubCmd(subCmd, args) {
			return nil
		}

		log.Debug("Try to complete with subcmd")
		if ExecuteURIWithComplement(false, true, engineAuto, ctx.Args().Slice()) {
			return nil
		}

		// log.Error("%q is not a resource name.\n", os.Args[1])

		// Error
		log.Error("%q is not valid command or resrouce name.\n", os.Args[1])
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
