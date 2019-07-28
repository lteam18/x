import { OFFICIAL, DEBUG } from "./GLOBAL";

DEBUG("Before major import")
import { CommandApp, arg, fun, SubCmd, Handler } from "@utilx/app"

DEBUG("Before quicks")
import { token as TOKEN_STORE } from "./quicks";
DEBUG("After quicks")

import { getOrgKey, s } from "./mainUtils";

DEBUG("Before AppServiceClient in main.ts")

import { uploadAppFromFile, setAppAccess, updateApp, uploadUrl, applyToken, decryptToken } from "./quicks/AppServiceClient";
import { getCodeTypeFromUrl } from "./run/utils";

import { CODE_TYPE } from "./typedef";

DEBUG("Before ./run")

import run, { fromFetchCode, selectRunner } from "./run";

DEBUG("After all import")

import { 
    URI, KEYPATH, FILEPATH, 
    CODETYPE, REST_ARG,
    TOKEN, 
    LANGUAGES, LANGUAGES_UPDATE,
    ACCESS
} from "./main.annotation"

DEBUG("Before vvsh")

// import * as httpcache from "./http-cache"

class Vvsh extends CommandApp {

    version(){
        // import { selfUpgrade, showVersion } from "./self_update";

        DEBUG("Start version")
        const u = require("./self_update")
        return u.showVersion()
    }

    @SubCmd
    @fun.Example(
        "vvsh upgrade", 
        "If installed in npm, npm install vvsh@latest -g. If in exe mode, it will tirgger download and replace.")
    upgrade(){
        const u = require("./self_update")
        return u.selfUpgrade()
    }

    @SubCmd
    @fun.Example(
        "vvsh token", 
        "Show the token stored in ~/.vvsh/TOKEN"
    )
    @fun.Example(
        "vvsh token <TOKEN_STRING>",
        "Setting <TOKEN_STRING> as current token. Saved in ~/.vvsh/TOKEN"
    )
    token(
        @TOKEN str?: string
    ){
        if (undefined === str) {
            console.log(`TOKEN is:\n${TOKEN_STORE.read()}`)
        } else {
            TOKEN_STORE.write(str)
        }
    }

    @SubCmd
    @fun.Example(
        "vvsh token-apply", 
        "Apply a new token according to the token stored in ~/.vvsh/TOKEN"
    )
    @fun.Example(
        "vvsh token-apply 'org' 'official'",
        "Apply a new token with { get: [work, l], put: [official, l] }"
    )
    async "token-apply"(
        
        @arg.Example("Token generated for Organziation Administorator")
        @arg.Help("Add an info into the token")
        @arg.Require
        @arg.Name("info")               info: string,

        @arg.Help("get priviledge")
        @arg.Require    
        @arg.Name("get-priviledges")    get: string,

        @arg.Help("If somitted, it will equal to get-privilledges")
        @arg.Name("put-priviledges")    put?: string,
    ) {
        if (put === undefined) put = get
        const res = await applyToken(get, put, info)
        if (res === false) {
            console.log("Token apply rejected")
        } else {
            console.log(res)
        }
    }

    @SubCmd
    @fun.Example(
        "vvsh token-decrypt", 
        "Will decrypt the token stored in ~/.vvsh/TOKEN"
    )
    @fun.Example(
        "vvsh token-decrypt XXXXXX",
        "Will decrypt the token XXXXXX"
    )
    async "token-decrypt"(
        @TOKEN token?: string
    ){
        const res = await decryptToken(token)
        if (res === false) {
            console.log("Token decrypt failure")
        } else if (res === "AccessDenied") {
            console.log(res)
        } else {
            console.log(JSON.stringify(res, undefined, 4))
            console.log(`Expired in `, new Date(res.expired))
        }
    }

    @Handler
    @fun.Help("Upload command")
    @fun.Example(
        "vvsh public @vvkv/init init.py",
        "Upload init.py and stored in key '@vvkv/init', in public mode"
    )
    @fun.Example(
        "vvsh private @vvkv/init init.py",
        "Upload init.py and stored in key '@vvkv/init', in private mode"
    )
    async upload(
        @arg.Require @ACCESS      access: "public" | "private",
        @arg.Require @KEYPATH     org_key: string,

        @arg.Name("HTTP-URL or FilePaths")
        @arg.Example("https://github.com/hello.py", "./ipscan.ts", "/home/user/code/findip.js")
        @arg.Help("File path for local resource.")
        filepath?: string,

        @CODETYPE    codetype?: CODE_TYPE
        )
    {
        const [org, key] = getOrgKey(org_key)
        if (undefined === filepath) {
            return await setAppAccess(org, key, access);
        }

        if (undefined === codetype) {
            codetype = getCodeTypeFromUrl(filepath)
        }

        if (filepath.startsWith("http://") || (filepath.startsWith("https://"))) {
            return await uploadUrl(org, key, filepath, access, codetype)
        }
        return await uploadAppFromFile(org, key, filepath, codetype, access)
    }

    @SubCmd
    @fun.Help("Upadte the cache to the latest version")
    async update(
        @arg.Require @URI     uri: string
    ){
        if (uri.startsWith("http://") || uri.startsWith("https://")) {
            const httpcache = require("./http-cache")
            return await httpcache.update(uri)
        }

        const res = await updateApp(...getOrgKey(uri))
        if (typeof res === "object") {
            process.exit(0)
        }
    }

    @Handler
    @fun.Help("list the resource with search text")
    ls(
        @arg.Require
        @arg.Name("ls | ll")
        @arg.Choice("ls", "ll")
        @arg.Help("ll is ls in DEBUG mode")     cmd: "ls" | "ll",

        // TODO: Actually is search text
        @KEYPATH @arg.Default(`@${OFFICIAL}`)   org_key: string = `@${OFFICIAL}`
    ){
        return s(org_key, cmd === "ll")
    }


    @SubCmd
    @fun.Help("Execute the stdin content in ts/js mode")
    async "-"(...args: any[]){
        const utils = require("./utils")
        await fromFetchCode.runWithString(
            await utils.getAllContentFromStdin(),
            ...args
        )
    }

    @Handler
    @fun.Help("Using cat")
    async cat(
        @arg.Require
        @arg.Name("cat?")
        @arg.Choice("cat", "cat!")
        @arg.Help("Printing the content of resource")   cmd: "cat" | "cat!",
        @arg.Require @URI                               uri: string,
        @REST_ARG                                       ...arg: any[]
    ){
        if (cmd === "cat!") {
            await this.update(uri)
        }
        return fromFetchCode.runWithRunnerAndUri(run.cat, uri, ...arg)
    }

    @Handler
    @fun.Help("Using app")
    async app(
        @arg.Require
        @arg.Name("app?")
        @arg.Choice("app", "app!")
        @arg.Help("Run the resource in app mode")       cmd: "app" | "app!",
        @arg.Require @KEYPATH                           org_key: string,
        @REST_ARG                                       ...arg: any[]
    ){
        if (cmd === "app!") {
            await updateApp(...getOrgKey(org_key))
        }
        return fromFetchCode.runWithAppName(org_key, ...arg)
    }

    @Handler
    @fun.Help("Using runner")
    async runInAllLanguage(
        @arg.Require @LANGUAGES         engineName: keyof typeof selectRunner.fromCmd,
        @arg.Require @URI               uri: string,
        @REST_ARG                       ...arg: any[]
    ) {
        const engine = selectRunner.fromCmd[engineName]
        return await fromFetchCode.runWithRunnerAndUri(engine, uri, ...arg)
    }

    @Handler
    @fun.Help("Using runner")
    async updateThenRunInAllLanguage(
        @arg.Require @LANGUAGES_UPDATE  engineName: keyof typeof selectRunner.fromCmd,
        @arg.Require @KEYPATH           org_key: string,
        @REST_ARG                       ...arg: any[]
    ) {
        const res = await updateApp(...getOrgKey(org_key))
        if (typeof res !== "object") {
            console.log("Error: ", res)
            return
        }
        const engine = selectRunner.fromCmd[engineName]
        return await fromFetchCode.runWithRunnerAndUri(engine, org_key, ...arg)
    }

    @Handler
    @fun.Help("Running the code accroding to the CODE_TYPE in metadata or path prefix")
    async run(
        @arg.Require @URI     uri: string,
        @REST_ARG             ...arg: any[]
    ) {
        DEBUG("Into run()", arg)
        const t = await fromFetchCode.runWithUrl( uri, ...arg )
        return t
    }

    @Handler
    @fun.Help("Running repl")
    async repl(
        @REST_ARG             ...arg: any[]
    ) {
         // start repl-server
        const repl3 = require("@typeshell/repl3")
        const env = require("@typeshell/env")
        repl3.repl3.main({
        ...env,
        })
        return
    }

}

(async () => {
    await CommandApp.run(new Vvsh("0.0.0"))
})()
