import { fetchApp } from "./quicks/AppServiceClient"

import { PREFIX_FP } from "./quicks/constants";
import fs from "fs";
import { getContentFromFileOrHttp } from "./utils";
import { DEBUG, OFFICIAL } from "./GLOBAL";
import { META_TYPE } from "./typedef";

function parse(appName: string): [string, string] {
    if (appName.startsWith("@")) appName = appName.slice(1)
    const idx = appName.indexOf("/")
    return [appName.slice(0, idx), appName.slice(idx + 1)]
}

import * as httpcache from "./http-cache"

async function fetchAppFromCacheOrWeb(org: string, key: string){
    const res = await fetchApp(org, key, undefined)
    if (typeof res === "object") {
        if (res.meta.isURL === true){
            const obj = JSON.parse(res.code.toString()) as { url: string }
            const code2 = await httpcache.get(obj.url)
            if (code2 === false) {
                return false
            }
            return {
                ...res,
                ...code2
            }
        }
    }
    return res
}

export async function byAlias(appName: string) {

    if (appName.startsWith("@")) {
        const [org, key] = parse(appName)
        return fetchAppFromCacheOrWeb(org, key)
    }

    if (fs.existsSync(PREFIX_FP) === false){
        DEBUG(`Prefix file does NOT exists. Creating in: ${PREFIX_FP}`)
        // await xs.echo().tee(PREFIX_FP).untilEnded()
        await fs.promises.writeFile(PREFIX_FP, `@${OFFICIAL}`)
    }

    const s: string = (await fs.promises.readFile(PREFIX_FP)).toString() // await xs.cat(PREFIX_FP).dumps();
    const candidates = s.split("\n").filter(e => e.trim().length !== 0)

    for (let c of candidates) {
        c = c.endsWith('/')? c.slice(0, c.length-1): c
        if (c.startsWith("@")) {
            DEBUG("Trying Organziation: ", c)
            const res = await fetchAppFromCacheOrWeb(c.slice(1), appName)
            // DEBUG("RESULT: \n", JSON.stringify(res, undefined, 2))
            if (res !== "NoSuchKey") return res
        } else {
            DEBUG("Trying HTTP or Local: ", c)
            const res = await getContentFromFileOrHttp(`${c}/${appName}`)
            if (res !== false) return res
        }
    }
    return false
    
}


export async function byUri(uri: string): Promise<false | "AccessDenied" | "NoSuchKey" | {
    meta?: META_TYPE
    code: Buffer;
    filepath: string;
}> {
    const ret = await getContentFromFileOrHttp(uri)
    if (ret === false) {
        return await byAlias(uri)
    }
    return ret
}
