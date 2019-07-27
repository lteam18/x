import { OFFICIAL, DEBUG } from "./GLOBAL";

DEBUG("Before AppServiceClient in mainUtils")

import { search } from "./quicks/AppServiceClient";

DEBUG("After AppServiceClient in mainUtils")

export function getOrgKey(appName: string): [string, string] {
    if (appName.startsWith("@")) {
        const idx = appName.indexOf("/")
        if (idx < 0) throw new Error(`Expect to be <@org>/<key>, but get: ${appName}`)
        return [appName.slice(1, idx), appName.slice(idx+1)]
    }
    return [OFFICIAL, appName]
}

export async function s(searchText: string, DEBUG = false){
    if (! searchText.startsWith("@")) {
        searchText = `@${OFFICIAL}/${searchText}`
    }
    if (searchText.indexOf('/') < 0) searchText += "/"
    const [org, prefix] = getOrgKey(searchText)
    search(org, prefix, DEBUG)
}

export function arg<T extends string>(idx: number, value?: T){
    const arg = process.argv[idx]
    if (undefined === arg) {    
        if (undefined === value) {
            throw new Error(`Expect ${idx+1}th argument `)
        }
        return value as T
    }
    return arg as T
}
