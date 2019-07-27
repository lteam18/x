import * as path from "path"
import { promises as fs } from "fs"
import fs$ from "fs"

import { DEBUG } from "../GLOBAL"

export function http_to_base64(http_url: string){
    return Buffer.from(http_url).toString("base64")
}

import os from "os"

fs$.mkdirSync(path.join(os.tmpdir(), ".vvsh", "url_cache"), { recursive: true })

export function getURLCachePath(http_url: string){
    return path.join(os.tmpdir(), ".vvsh", "url_cache", http_to_base64(http_url))
}

import A from "axios"

export async function update(url: string) {
    try {
        const code = (await A.get(url, { responseType: "arraybuffer" })).data as Buffer
        const filepath = getURLCachePath(url)
        await fs.writeFile(filepath, code)
        return {
            code: code,
            filepath
        }
    } catch(err) {
        DEBUG(err)
        return false
    }
}

export async function getFromCache(url: string) {
    const p = getURLCachePath(url)
    try {
        return {
            code: await fs.readFile(p),
            filepath: p
        }
    } catch (err) {
        DEBUG(err)
        return false
    }
}

export async function get(http_url: string) {
    const res = await getFromCache(http_url)
    if (false === res) {
        return update(http_url)
    }

    DEBUG("Using cache")
    return res
}
