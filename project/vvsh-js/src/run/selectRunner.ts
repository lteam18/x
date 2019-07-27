import run from "./Run"

export const fromSuffix = {
    "js":   run.js,
    "ts":   run.ts,
    "py":   run.py,

    "bash": run.bash,
    "sh":   run.bash,
    "fish": run.fish,

    "awk":  run.awk,
    "txt":  run.cat,
    
    "java": run.java,
    "jar":  run.jar,
    "kts":  run.kotlin
}

export const fromCmd = {
    ...fromSuffix,
    "javascript":   run.js,
    "typescript":   run.ts,
    "python":       run.py,
    "text":         run.cat,
    "cat":          run.cat,
    "kotlin":       run.kotlin
}
