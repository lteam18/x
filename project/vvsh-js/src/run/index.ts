import { DEBUG } from "../GLOBAL";

import run, { Run } from "./Run"

DEBUG("Before fromFetchCode in index.ts")

import * as fromFetchCode from "./fromFetchCode"

DEBUG("After fromFetchCode in run/index.ts")

export { run, Run, fromFetchCode }
export default run

// export const fromFetchCode = {
//     auto: runFetchCode,
//     withCodeType: runWithCodetype,
//     withRunner: runWithRunner,
//     withUrl: runWithUrl,
//     withAppName: runWithAppName,
// }

import * as selectRunner from "./selectRunner"
DEBUG("After selectRunner in run/index.ts")

export { selectRunner }
