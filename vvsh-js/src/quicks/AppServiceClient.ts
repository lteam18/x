import { DEBUG } from "../GLOBAL";

import fs from "fs"

import path from "path"

// const url = "https://1632295596863408.cn-shenzhen.fc.aliyuncs.com/2016-08-15/proxy/vvsh/manager"

import { getAppFP, getAppIdxFP } from "./constants"
import * as TOKEN from "./token"

// import * as xs from "@utilx/stream"

DEBUG("Before vvkv")
import { Client } from "vvkv"
DEBUG("After vvkv")

// const URL = "https://1632295596863408.cn-hongkong.fc.aliyuncs.com/2016-08-15/proxy/infra/vvkv"
const URL =
    "https://1632295596863408.cn-shenzhen.fc.aliyuncs.com/2016-08-15/proxy/vvsh/vvkv";

function KV(token?: string) {
    return new Client(URL, token)
}

async function _updateApp(org: string, key: string, token?: string) {
    const res = await KV(token).getAsStream<META_TYPE>(org, key)
    if (res === "AccessDenied") {
        console.log(
            "Access denied. Please contact the Administrator and upgrade the token."
        );
        return res
    }

    if (res === "NoSuchKey") {
        // console.log(`Error ${res}: ${org} ${key}`);
        return res;
    }

    if (res === false) {
        console.log("UnknownError", res);
        return res;
    }

    const fp = getAppFP(org, key);
    await fs.promises.mkdir(path.dirname(fp), { recursive: true })

    // OPT
    const xs = require("@utilx/stream")
    const s = xs.enhanceReadable(res.data).unzip1()

    const ss = s.tee(fp)
    const ret = await s.dump() as Buffer
    await ss.untilEnded() // Wait until written in disk.

    const idx_fp = getAppIdxFP(org, key)
    await fs.promises.mkdir(path.dirname(idx_fp), { recursive: true })
    await xs.echo(JSON.stringify(res.meta)).tee(idx_fp).untilEnded()

    return { meta: res.meta, code: ret, filepath: fp }
}



export async function setAppAccess(
    org: string, key: string,
    access: "private" | "public",
    token?: string
) {
    const res = await KV(TOKEN.get(token)).setAccess(org, key, access)

    switch (res) {
        case "AccessDenied":
        case "NoSuchKey":
            console.log(`Upload failure: ${res}`);
            DEBUG(`uploadApp failure: ${res}`);
            return res;
        case false:
            console.log("Upload failure: Unknown");
            DEBUG(`uploadApp failure: ${res}`);
            return res
    }

    console.log(`setAccess success: ${org} ${key}\nCurrent Access: ${access}`);
    return true;
}

export async function search(
    org: string, key: string,
    DEBUG = false
) {
    const res = await KV(TOKEN.get()).list(org, key)

    switch (res) {
        case "AccessDenied":
            console.log(`Search failure: ${res}`);
            return res;
        case false:
            console.log("Search failure: Unknown");
            return res
    }

    if (DEBUG) {
        console.log(JSON.stringify(res, undefined, 2))
    } else {
        for (let e of res) console.log(`${e.lastModified}\t${e.size}\t${e.name}`)
    }
    
    return res;
}

import { CODE_TYPE, META_TYPE, URL_TYPE } from "../typedef"

export async function uploadUrl(
    org: string, key: string, url: string, access: "private" | "public",
    codetype: CODE_TYPE, token?: string
) {
    const xs = require("@utilx/stream")
    const buffer = await xs.echo(JSON.stringify({ url } as URL_TYPE)).zip1().dump()
    const meta = { codetype, isURL: true } as META_TYPE

    const res = await KV(TOKEN.get(token)).put(org, key, buffer, access, meta)

    if (res === "AccessDenied") {
        console.log("Upload failure: AccessDenied: ", org, key)
        DEBUG(`uploadApp failure: ${res}`)
        return
    }

    if (res === false) {
        console.log("Upload failure: Unknown");
        DEBUG(`uploadApp failure: ${res}`);
        return;
    }

    console.log(`Upload success: @${org}/${key} with [codetype=${codetype}]\n`)
    console.log(`Try:\n  vvsh app @${org}/${key}\n  vvsh @${org}/${key}`)
    return true
}

export async function uploadApp(
    org: string, key: string, content: undefined | Buffer,
    codetype: CODE_TYPE, access: "private" | "public",
    token?: string
) {
    const res = await KV(token).put(org, key, content || Buffer.from(""), access, { codetype })
    if (res === "AccessDenied") {
        console.log("Upload failure: AccessDenied")
        DEBUG(`uploadApp failure: ${res}`)
        return
    }

    if (res === false) {
        console.log("Upload failure: Unknown");
        DEBUG(`uploadApp failure: ${res}`);
        return;
    }

    console.log(`Upload success: @${org}/${key} with [codetype=${codetype}]`)
    console.log(`Located in ${getAppFP(org, key)}\n`)
    console.log(`Try:\n  vvsh app @${org}/${key}\n  vvsh @${org}/${key}`)
    return true
}

// If filepath is undefined, only update app attributes
export async function uploadAppFromFile(
    org: string, key: string, filepath: string | undefined,
    codetype: CODE_TYPE, access: "private" | "public",
    token?: string
) {
    try {
        const xs = require("@utilx/stream")
        const buf = filepath === undefined ? filepath : await xs.cat(filepath).zip1().dump()
        return await uploadApp(org, key, buf, codetype, access, TOKEN.get(token))
    } catch (err) {
        console.log("Upload failure:", err.response ? err.response.data : err)
        return false
    }
}

export async function updateApp(org: string, key: string, token?: string) {
    const ret = await _updateApp(org, key, TOKEN.get(token))
    if (ret !== false) {
        console.log(`App update to latest: ${org} ${key}`)
    } else {
        console.log("Update failure")
    }
    return ret
}

export async function fetchApp(org: string, key: string, token?: string) {
    try {
        // const xs = require("@utilx/stream")
        // TODO: what if meta is corrupted?

        const filepath = getAppFP(org, key)
        const meta_content = (await fs.promises.readFile(getAppIdxFP(org, key)) ).toString()
        const code_content = (await fs.promises.readFile(filepath) )

        return {
            meta: JSON.parse(meta_content) as META_TYPE,
            code: code_content,
            filepath
        }
    } catch (err) {
        return await _updateApp(org, key, TOKEN.get(token))
    }
}

export async function applyToken(
    get: string,
    put: string,
    info: string
){
    const handle = (e: string) => e.split(",").map(e1 => e1.trim())
    return await KV(TOKEN.get()).encrypt(handle(get), handle(put), info)
}

export async function decryptToken(
    targetToken: string = TOKEN.get() || ""
){
    return await KV(TOKEN.get()).decrypt(targetToken)
}


