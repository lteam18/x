import repl3 from "@typeshell/repl3"

import { getAllContentFromStdin, getContentFromFileOrHttp } from "./utils";

import { showVersion, selfUpgrade } from "./self_update";
import { uploadAppFromFile, updateApp, setAppAccess, uploadUrl, search } from "./quicks/AppServiceClient";

import { token } from "./quicks";
import run, { fromFetchCode } from "./run";

import * as env from "@typeshell/env"

import { getCodeTypeFromUrl } from "./run/utils";
import { OFFICIAL } from "./GLOBAL";
import { getOrgKey, arg, s } from "./mainUtils";
import { CODE_TYPE } from "./typedef";

export async function uploadURL(
    org: string, key: string, 
    access: "public" | "private", 
){
    if (process.argv.length === 5) {
        const url = arg(4)
        const code_type = getCodeTypeFromUrl(url)
        return await uploadUrl(org, key, url, access, code_type) 
    }
    return await uploadUrl(org, key, arg(5), access, arg<CODE_TYPE>(4))
}

export async function uploadFromFile(access: "public" | "private", appName: string){
    const [org, key] = getOrgKey(appName)
    if (process.argv.length === 6) {
        const filepath = process.argv[5]
        const codetype = process.argv[4] as any

        return await uploadAppFromFile(org, key, filepath, codetype, access)
    }
    if (process.argv.length === 5) {
        const filepath = process.argv[4]
        const codetype = getCodeTypeFromUrl(filepath)
        return await uploadAppFromFile(
            org, key, filepath,
            codetype, access
        )
    }

    if (process.argv.length === 4) {
        return await setAppAccess(org, key, access);
    }
    throw new Error(`Unexpected argument length: ${arguments}`)
}


async function main(){
    if (2 === process.argv.length) {
        // start repl-server
        repl3.main(env)
        return
    }
    
    let runner: typeof run.cat | undefined = undefined
    const setRunner = (r: typeof runner) => {
        if (undefined === runner) {
            runner = r
        }
    }
    const getRunner = () => runner

    const cmd = process.argv[2]
    switch (cmd) {
        case "version":
            return await showVersion()
        case "upgrade":
            return await selfUpgrade()
        case "token":
            if (process.argv.length === 3) {
                console.log(`TOKEN is:\n${token.read()}`)
            } else {
                token.write(process.argv[3])
            }
            return

        // arg0 vvsh --public <key> <js | ts> [filename]
        case "public":
        case "private":
            // If not data provide change access
            return await uploadFromFile(cmd, arg(3))
        case "update":
            return await updateApp(...getOrgKey(arg(3)))
        
        case "public-url":
        case "private-url":
            const [org, key] = getOrgKey(arg(3))
            return await uploadURL(org, key, cmd === "public-url" ? "public" : "private")

        case "ls":
        case "ll":
            return s(arg(3, `@${OFFICIAL}`), cmd === "ll")

        case "-": 
            return await fromFetchCode.runWithCodeType(
                { code: Buffer.from(await getAllContentFromStdin()) },
                ...process.argv.slice(4)
            )
        case "app!":
            await updateApp(...getOrgKey(arg(3)))
            // fallthrough
        case "app":
            return await fromFetchCode.runWithAppName(arg(3), ...process.argv[4])

        case "cat":     setRunner(run.cat)
        case "py":
        case "python":  setRunner(run.py)
        case "perl":    setRunner(run.perl)
        case "sh":
        case "bash":    setRunner(run.bash)
        case "fish":    setRunner(run.fish)
        
        case "awk":     setRunner(run.awk)

        case "haskell": setRunner(run.haskell)
        case "agda": setRunner(run.agda)

        case "js":      setRunner(run.js)
        case "ts":      setRunner(run.ts)
            const runner = getRunner()
            return await fromFetchCode.runWithRunnerAndUri(
                runner || run.cat, 
                arg(3),
                ...process.argv.slice(4)
            )
        default:
            return await fromFetchCode.runWithUrl(
                arg(2),
                ...process.argv.slice(3)
            )
    }
    
}


main().catch(e => console.log(e))

