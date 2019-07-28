import * as path from "path"
import { promises as fs } from "fs-extra"

import * as httpcache from "./http-cache"

export async function getContentFromFileOrHttp(url: string): Promise<{ 
    code: Buffer, filepath: string
} | false>{
    if ( (url.startsWith("http://")) || (url.startsWith("https://"))){
        try{
            return httpcache.get(url)
            // return (await A.get(url)).data
        } catch (err) {
            DEBUG(err)
            return false
        }
    } else {
        DEBUG("Inside getContentFromFileOrHttp()")
        const fp = path.normalize(url)
        try {
            return {
                code: await fs.readFile(fp),
                filepath: url
            }
        } catch (err) {
            if (err.message && err.message.indexOf("no such file or directory") >= 0) {
                DEBUG("Before return false")
                return false
            }
            throw err
        }
    }
}

export async function getAllContentFromStdin(){
    let c = ""
    process.stdin.on("data", e => {
        c += e
    })
    return new Promise<string>(resolve => {
        process.stdin.on("end", () => {
            resolve(c)
        })
    })
}

import crypto from "crypto"
// import r$, { r } from "@utilx/process";
import cp from "child_process"
import { DEBUG } from "./GLOBAL";

export function getNameFromUri(uri: string){
    const m = crypto.createHash("md5")
    return m.update(uri).digest("hex")
}

export function isExecutingInBinExe(path: string = process.argv[0]){
    try {
        const ret = cp.spawnSync(`${path} verseion`)
        return (ret.status === 0)
    } catch(err) {
        return false
    }
}
