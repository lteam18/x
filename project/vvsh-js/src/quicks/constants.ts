export const ROOT_DIR = `${process.env.HOME}/.vvsh`
export const APP_DIR = `${ROOT_DIR}/APP`
export const APP_IDX_DIR = `${ROOT_DIR}/APP_IDX`
export const TOKEN_FP = `${ROOT_DIR}/TOKEN`
export const PREFIX_FP = `${ROOT_DIR}/PREFIX`

import fs from "fs"

fs.mkdirSync(APP_DIR, { recursive: true })
fs.mkdirSync(APP_IDX_DIR, { recursive: true })


export function getAppFP(org: string, key: string) {
    return `${APP_DIR}/${org}/${key}`
}

export function getAppIdxFP(org: string, key: string) {
    return `${APP_IDX_DIR}/${org}/${key}`
}
