import debug from "debug"

export const VERBOSE = debug("vvsh")

export const OFFICIAL = "official"

export const START = Date.now()

// export class TimeObj extends Object{
//     toString(){
//         console.log("Not invoked")
//         return (Date.now() - START).toString()
//     }
// }

// const time = new TimeObj()

export function DEBUG(...msg: any[]){
    if (VERBOSE.enabled) VERBOSE(Date.now() - START, ...msg)
}
