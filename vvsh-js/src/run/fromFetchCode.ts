import { DEBUG } from "../GLOBAL";

import { META_TYPE } from "../typedef";
import { getCodeTypeFromUrl } from "./utils"

DEBUG("Before ./fetchCode")
import * as fetchCode from "../fetchCode";
DEBUG("Before ./Run")
import run from "./Run"
DEBUG("After ./Run")

async function handleAccessDeniedAnNoSuchKey<T>(
    res: Promise<"AccessDenied" | "NoSuchKey" | false | T>
): Promise<T> {
    const r = await res
    if (r === "AccessDenied") {
        console.log("Accessed Denied")
        process.exit(1)
        throw new Error() // For stupid typescript
    }

    if (r === "NoSuchKey") {
        console.log("Error NoSuchKey")
        process.exit(1)
        throw new Error() // For stupid typescript
    }

    if (r === false) {
        console.log("Unknow error")
        process.exit(1)
        throw new Error() // For stupid typescript
    }

    return r
}

export async function runWithAppName(
    appName: string,
    ...argument: string[]
) {
    const res0 = fetchCode.byAlias(appName)
    const res = await handleAccessDeniedAnNoSuchKey(res0)
    return await runWithCodeType(
        res, 
        ...argument
    )
}

export async function runWithUrl(
    url: string,
    ...argument: string[]
) {
    const res0 = fetchCode.byUri(url)
    const res = await handleAccessDeniedAnNoSuchKey(res0);

    if ("meta" in res) {
        DEBUG("Before withCodeType")
        return runWithCodeType(res, ...argument)
    }

    return await runWithCodeType({
        ...res,
        meta: { codetype: getCodeTypeFromUrl(url) }
    }, ...argument)

}

export async function runWithString(
    code: string,
    ...argument: string[]
) {
    await run.tsOrJS({ code: Buffer.from(code) }, ...argument)
    return true
}

export async function runWithCodeType(
    code: { meta?: META_TYPE, code: Buffer, filepath?: string },
    ...argument: string[]
) {
    if (undefined === code.meta) {
        await run.tsOrJS(code, ...argument)
        return code
    }

    const runner = run.get(code.meta.codetype)
    if (runner === undefined) {
        await run.tsOrJS(code, ...argument)
        return true
    }

    await runner(code, ...argument)
    return true
}

export async function runWithRunnerAndUri(
    runner: typeof run.perl,
    uri: string,
    ...argument: string[]
){
    const res0 = fetchCode.byUri(uri)
    const res = await handleAccessDeniedAnNoSuchKey(res0)
    return await runner(res, ...argument)
}
