import { arg } from "@utilx/app"

import { selectRunner } from "./run";

export const URI = arg.combine(
    arg.Name("URI"),
    arg.Example("find-ip", "@myaccount/ip-scan", "https://xxx.github.com/a/b.js", "/var/scdripts/hi.py"),
    arg.Help(
        "It could be http like url, file path, or KEYPATH",
        "Vvsh will try to fetch resource as it is url, then filepath, then KEYPATH"
    )
)

export const KEYPATH = arg.combine(
    arg.Name("KeyPath"),
    arg.Example("find-ip", "@myaccount/ip-scan", "@myorg/ip-scan"),
    arg.Help(
        "Key in vvsh repostiroy, a typical form is @<org>/<keypath>",
        "If @<org> omited, vvsh will try completing the @<org> and requesting the resource",
        "The try sequence is according to the list in ~/.vvsh/PREFIX.",
        "Normally, if you don't modify the thie PREFIX file, it will be @official"
    )
)

export const FILEPATH = arg.combine(
    arg.Name("FilePath"), 
    arg.Example("./hello.py", "./ipscan.ts", "/home/user/code/findip.js"),
    arg.Help("File path for local resource.")
)


export const CODETYPE = arg.combine(
    arg.Name("CodeType"),
    arg.Example("ts", "js", "py", "fish", "sh", "pl", "txt", "awk"),
    arg.Help(
        "Code type, decide which vm to use to run code", 
        "ts for typescript, js for javascript, py for python",
        "It is optional, Code type will be guessed according to the FilePath suffix"
    )
)

export const TOKEN = arg.combine(
    arg.Name("token"), 
    arg.Example("[A base64 string]"), 
    arg.Help(
        "Provide, TOKEN will be stored in ~/.vvsh/TOKEN",
        "If not, the current token will be printed"
    )
)

export const REST_ARG = arg.combine(
    arg.Name("...args"), 
    arg.Help("Argument provided for the script")
)

export const LANGUAGES = arg.combine(
    arg.Name("Language"),
    arg.Choice(...Object.keys(selectRunner.fromCmd))
)

export const LANGUAGES_UPDATE = arg.combine(
    arg.Name("Language"),
    arg.Choice(
        ...Object.keys(selectRunner.fromCmd).map(e => `${e}!`)
    )
)

export const ACCESS = arg.combine(
    arg.Name("Access"),
    arg.Choice("public", "private")
)