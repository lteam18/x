import { DEBUG } from "./GLOBAL";

import { isExecutingInBinExe } from "./utils";

import os from "os"

DEBUG("Before env")

import { network, cat } from "@typeshell/env"

import path from "path"

import r$, { r } from "@utilx/process"


const ROOT_URL = "https://github.com/lteam18/auto/releases/download/v0"

export function getDownloadUrl(){
    const pl = os.platform()
    switch (pl) {
        case "linux": return `${ROOT_URL}/vvsh.linux.exe.zip`
        case "darwin": return `${ROOT_URL}/vvsh.mac.exe.zip`
        case "win32": return `${ROOT_URL}/vvsh.windows.exe.zip`
        default:
            throw new Error("Platform not supported.")
    }
}

export const info = console.log.bind(console)

export async function upgrade(mylocation: string){
    
    const download_url = getDownloadUrl()

    // TODO: generate temporary filepath
    const tg = path.join(os.tmpdir(), "___TMP_NODE_SHELL")

    // OPT
    const fs = require("fs-extra")

    info("Start downloading to filepath: " + tg)
    await network.download(download_url).unzip1().tee(tg).untilEnded()
    info("Download finished. " + tg)
    await fs.chmod(tg, "755")
    const ret = await r(`${tg}`, "version").get()
    // TODO: not good.
    if (("code" in ret) && (ret.code === 0)) {
        info(`Moving file from ${tg} to ${mylocation}`)
        await fs.move(tg, mylocation, { overwrite: true })
    } else {
        info(`File downloaded is broken. Remove the unused files.`)
        await fs.remove(tg)
    }
}

export async function selfUpgrade(){
    if (! isExecutingInBinExe()){
        info(`Detected it is installed in npm.`)
        info(`cmd: npm upgrade -g vvsh@latest`)
        const p = r("npm upgrade -g vvsh@latest")
        const ret = await p.get()
        if (("code" in ret) && (ret.code === 0)) {
            console.log("vvsh upgrade success.")
        }
        return
    }
    info(`Detected it is installed as a standlable binary.`)
    await upgrade(process.argv[0])
}

export async function showVersion(){
    try {
        const v = JSON.parse(await cat(__dirname + "/../package.json").dumps())["version"]
        info(`Version ${v}. \nDesigned for LTeam. All rights reserved by Junhao Li.`)
    } catch (err) {
        info("Error", err)
    }
}

const nix = `If you are using linux or macos, please using following command:
curl https://edwinjhlee.github.io/vvsh/install.sh | sudo bash
`

const win = `If you are using windows, please download from https://github.com/lteam18/auto/releases/download/v0/vvsh.win.exe`

export function printInstallCommand(){
    info(nix + "\n" + win)
}

