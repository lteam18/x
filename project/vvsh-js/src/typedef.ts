export type CODE_TYPE = "" | "ts" | "js" | "bash" | "sh" | "java" | "jar" | "kotlin" | "txt" | "url" | "py" | "fish" | "perl" | "awk" | "r" | "json" | "xml"

export type META_TYPE = {
    codetype: CODE_TYPE,
    isURL?: true
}

export type URL_TYPE = {
    url: string
}
