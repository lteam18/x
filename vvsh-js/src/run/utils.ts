import { CODE_TYPE } from "../typedef";

import { DEBUG } from "../GLOBAL";

export function _getCodeTypeFromUrl(name: string): CODE_TYPE {
    if (name.endsWith(".ts")) return "ts"
    if (name.endsWith(".js")) return "js"
    if (name.endsWith(".py")) return "py"

    if (name.endsWith(".bash")) return "bash"
    if (name.endsWith(".sh")) return "sh"
    if (name.endsWith(".fish")) return "fish"

    if (name.endsWith(".awk")) return "awk"
    if (name.endsWith(".txt")) return "txt"

    if (name.endsWith(".java")) return "java"
    if (name.endsWith(".jar")) return "jar"
    if (name.endsWith(".kts")) return "kotlin"

    return ""
}

export function getCodeTypeFromUrl(name: string): CODE_TYPE {
    const res = _getCodeTypeFromUrl(name)
    DEBUG(`Guessing Code type from Name: ${name}\nType is: ${res}`)
    return res
}