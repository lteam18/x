import { DEBUG } from "../GLOBAL";

import { compile } from "../tseval";
import { inject } from "../requireUtils";

// import { echo } from "@utilx/stream";
import { CODE_TYPE, META_TYPE } from "../typedef";
import { promises as fs } from "fs"
import { tmpdir } from "os";

type Signals =
    "SIGABRT" | "SIGALRM" | "SIGBUS" | "SIGCHLD" | "SIGCONT" | "SIGFPE" | "SIGHUP" | "SIGILL" | "SIGINT" | "SIGIO" |
    "SIGIOT" | "SIGKILL" | "SIGPIPE" | "SIGPOLL" | "SIGPROF" | "SIGPWR" | "SIGQUIT" | "SIGSEGV" | "SIGSTKFLT" |
    "SIGSTOP" | "SIGSYS" | "SIGTERM" | "SIGTRAP" | "SIGTSTP" | "SIGTTIN" | "SIGTTOU" | "SIGUNUSED" | "SIGURG" |
    "SIGUSR1" | "SIGUSR2" | "SIGVTALRM" | "SIGWINCH" | "SIGXCPU" | "SIGXFSZ" | "SIGBREAK" | "SIGLOST" | "SIGINFO";


const SIGNAL_LIST: Signals[] = [
    "SIGABRT" , "SIGALRM" , "SIGBUS" , "SIGCHLD" , "SIGCONT" , "SIGFPE" , "SIGHUP" , "SIGILL" , "SIGINT" , "SIGIO" ,
    "SIGIOT" , /*"SIGKILL" ,*/ "SIGPIPE" , "SIGPOLL" , "SIGPROF" , "SIGPWR" , "SIGQUIT" , "SIGSEGV" , "SIGSTKFLT" ,
    /*"SIGSTOP" ,*/ "SIGSYS" , "SIGTERM" , "SIGTRAP" , "SIGTSTP" , "SIGTTIN" , "SIGTTOU" , "SIGUNUSED" , "SIGURG" ,
    "SIGUSR1" , "SIGUSR2" , "SIGVTALRM" , "SIGWINCH" , "SIGXCPU" , "SIGXFSZ" , "SIGBREAK" , "SIGLOST" , "SIGINFO"
]

import cp from "child_process"

DEBUG("Before Run Class Definition")

export class Run{

    async handleContent(content: string, ...argument: string[]){
        require = await inject(content) // NOTICE: require must be follow the eval function
        process.argv = [...process.argv.slice(0, 2), ...argument]
    }

    public get(codetype: CODE_TYPE) {
        switch (codetype) {
            case "ts": 
                return this.ts
            case "js": 
                return this.js
            case "fish": 
                return this.fish
            case "py":
                return this.py
            case "bash":
            case "sh":
                return this.bash
                // case "perl":
                //     return this.perl
            case "java":
                return this.java
            case "kotlin":
                return this.kotlin
            case "jar":
                return this.jar
            case "txt":
            case "":
            default:
                return this.cat
        }
    }

    public readonly ts = (async (content: { code: Buffer }, ...argument: string[]) => {
        const content_str = content.code.toString()

        await this.handleContent(content_str, ...argument)
        try{
            const code = compile(content_str)
            DEBUG("ts before eval")
            await eval(code)
            DEBUG("after eval")
        } catch (err) {
            console.log(err)
        }
    }).bind(this)

    public readonly js = (async (content: { code: Buffer }, ...argument: string[]) => {
        const content_str = content.code.toString()

        await this.handleContent(content_str, ...argument)
        try{
            DEBUG("js before eval")
            await eval(content_str)
            DEBUG("js after eval")
            // process.exit(0)
        } catch(err) {
            console.log(err)
        }
    }).bind(this)

    public readonly tsOrJS = (async(content: { code: Buffer }, ...argument: string[]) => {
        const content_str = content.code.toString()

        await this.handleContent(content_str, ...argument)

        let code: string | null = null
        try{
            code = compile(content_str)
        } catch (err) {
            console.log(err)
        }

        try{
            if (null === code) {
                await eval(content_str)
            } else {
                await eval(code)
            }
        } catch (err) {
            console.log(err)
        }

    }).bind(this)

    public readonly py = this.exec.bind(this, "python")
    public readonly perl = this.exec.bind(this, "perl")

    public readonly bash = this.exec.bind(this, "bash")
    public readonly fish = this.exec.bind(this, "fish")
    public readonly csh = this.exec.bind(this, "csh")
    public readonly zsh = this.exec.bind(this, "zsh")

    public readonly haskell = this.exec.bind(this, "haskell")
    public readonly agda = this.exec.bind(this, "agda")

    public readonly awk = this.exec.bind(this, "awk")

    public readonly java = this.exec.bind(this, "java")
    public readonly jar = this.exec.bind(this, ["java", "-jar"])
    public readonly kotlin = this.exec.bind(this, ["kotlinc", "-script"])
    
    // Consider using Go as starter
    // If we do sth about download binary, we will use Go as a crossplatform starter.
    // Kotlin, Java, Scala, Sed, Julia, Haskell, Agda

    // TODO: go

    static parseRunningCmd(runningCmd: string | string[]){
        if (Array.isArray(runningCmd)) {
            return { cmd: runningCmd[0], args: runningCmd.slice(1) }
        } else {
            return { cmd: runningCmd, args: [] }
        }
    }

    public async exec(runningCmd: string | string[], content: { meta?: META_TYPE, code: Buffer, filepath?: string }, ...argument: string[]){
        if (content.filepath) {
            return this.execFile(runningCmd, content.filepath, ...argument)
        } else {
            return this.execContent(runningCmd, content.code, ...argument)
        }
    }

    public async execContent(runningCmd: string | string[], content: Buffer, ...argument: string[]){
        const cmd = Run.parseRunningCmd(runningCmd)
        let tmppath = `${tmpdir()}/${Math.random()}`
        if (cmd.cmd === "kotlinc") tmppath += ".kts" // Exception

        DEBUG("execContent", runningCmd, tmppath)
        try {
            await fs.writeFile(tmppath, content, { encoding: "binary"} )
            await this._execFile(runningCmd, tmppath, ...argument)
        } catch(err) {
            console.log(err)
        }
        finally {
            try{
                await fs.unlink(tmppath)
            } catch (err) {
                console.log(err)
            }
            process.exit(1)
        }
    }

    public async execFile(runningCmd: string | string[], filepath: string, ...argument: string[]) {
        let code = 0
        try {
            code = await this._execFile(runningCmd, filepath, ...argument)
        } finally {
            process.exit(code)
        }
    }

    public async _execFile(runningCmd: string | string[], filepath: string, ...argument: string[]){

        const cmd = Run.parseRunningCmd(runningCmd)

        // import { ProcessInstance } from "@utilx/process";
        const xproc = require("@utilx/process")

        try {
            DEBUG(cmd.cmd, [...cmd.args, filepath, ...argument])
            const pi = xproc.ProcessInstance.createWithChildProcess(
                cp.spawn(cmd.cmd, [...cmd.args, filepath, ...argument], { stdio: "inherit" }),
                "Running"
            )
            // Propagate All Signals
            SIGNAL_LIST.forEach(e => process.on(e, () => pi.child_process.kill(e)))
            return await pi.code
        } catch(err) {
            const msg = (err as Error).message
            if (/spawn\s.+\sENOENT/.test(msg)) {
                console.log(`Runtime Not Exists: ${cmd}\nPlease install first before run.`)
            } else {
                console.log(err)
            }
            return 1
        }
    }

    public readonly cat = async (content: { code: Buffer }, ...argument: string[]) => {
        DEBUG("Using cat")
        console.log(content.code.toString())
    }
}

const run = new Run()
export default run
