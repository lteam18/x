import { TOKEN_FP } from "./constants";

import fs from "fs"
import path from "path"

export function get(token?: string) {
    if (token !== undefined) return token
    process.env.VVSH_SERVICE_TOKEN =
        undefined !== process.env.VVSH_SERVICE_TOKEN ?
            process.env.VVSH_SERVICE_TOKEN :
            read()
    return process.env.VVSH_SERVICE_TOKEN
}

export async function write(token: string) {
    await fs.promises.mkdir(path.dirname(TOKEN_FP), { recursive: true })
    return await fs.promises.writeFile(TOKEN_FP, token)
}

export function read() {
    try {
        return fs.readFileSync(TOKEN_FP).toString()
    } catch (err) {
        return undefined
    }
}