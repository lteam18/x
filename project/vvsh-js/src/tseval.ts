import { join, resolve } from "path";

import { writeFileSync, unlinkSync } from "fs";

const EVAL_FILENAME = `[eval].ts`
const EVAL_PATH = join(process.cwd(), EVAL_FILENAME)
const EVAL_INSTANCE = { input: '', output: '', version: 0, lines: 0 }

const option = {
    compilerOptions: {
        "target": "esnext",
        "module": "commonjs",
        "esModuleInterop": true
    },
    skipProject: true
}

export function tseval(code: string, module: string){

    const tmp = join(process.cwd(), "nodeshell." +  (10000000 + Math.random() * 10000000) + ".ts")

    writeFileSync(tmp, code)

    try {
        // Prepend `ts-node` arguments to CLI for child processes.
        process.execArgv.unshift(
            __filename,
            ...process.argv.slice(2, process.argv.length - process.argv.length)
        );

        // import { register } from "ts-node";
        const tsnode = require("ts-node")
        const service = tsnode.register(); // should use option

        process.argv = [process.argv[1], tmp]
            // .concat(process.argv.length ? resolve(process.cwd(), process.argv[2]) : [])
            .concat(process.argv.slice(3));

        const compile_output = service.compile(code, EVAL_FILENAME, 0)
        // console.log(compile_output)

        // (Module as any)._preloadModules("os-command", "fs")

        const Module = require("module");
        Module.runMain();
    } finally {
        unlinkSync(tmp)
    }

}

export function compile(code: string){

    const tsnode = require("ts-node")
    const service = tsnode.register({
        ignoreDiagnostics: [2307, 2580],
        ...option
    })
    const compile_output = service.compile(code, EVAL_FILENAME, 0)
    return compile_output
}

