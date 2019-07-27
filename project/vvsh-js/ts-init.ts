import { ui, ui as UI } from "@typeshell/os-command";

import { npm } from "@typeshell/cmd-type";
import { r, echo } from "@typeshell/env";
import { FileEditor, re, JS } from "@utilx/string";

/*
Following content should not be uploaded to npms:
1. config file and node_modules
2. package lock
3. src code: src/test/example/mocha
4. compiled code other than dist: build/webpack/nexe-bin
5. documents: docs/dev
6. log
*/
const NPM_IGNORE = 
`.vscode
.npmignore
.gitignore

tslint.json
tsconfig.json
webpack.config.js
node_modules

package-lock.json
yarn.lock

src
test
example
mocha

build
webpack
nexe-bin

docs
dev

log
npm-debug.log
`

/*
Following content should not be uploaded:
1. node_modules
2. compiled code: build/dist/webpack/nexe-bin
3. package lock
4. log
*/
const GIT_IGNORE = 
`node_modules

build
dist
webpack
nexe-bin

package-lock.json
yarn.lock

log
npm-debug.log
`

export async function npminit(name: string){
    await npm("init", "-y").code
    const a = [
        `"ins": "function i(){ npm install $1 @types/$1 --save; }; i ",`,
        `"insd": "function i(){ npm install $1 @types/$1 --save-dev; }; i",`,
        `"lintfix": "tslint --fix --project .",`,
        `"lintfix1": "tslintlevel=1 tslint --fix --project .; exit 0",`,
        `"lintfix2": "tslintlevel=2 tslint --fix --project .; exit 0",`,
        `"tsc": "rm -rf dist; tsc && tsc -d",`,
        `"pubpub": "npm run test && npm run tsc && npm publish --access public",`,
        `"mocha": "mocha -r ts-node/register",`,
    ].map(e => `\t\t${e}`)
    const fe = FileEditor.load$("./package.json")!
    fe.locate(/"scripts"/)!.append(...a)
    fe.locate(/"test"/)!.replace2(re.STRING, JS(`npm run mocha mocha/**/*.ts`))
    fe.save$()
}

export async function prepareTypescriptEnv() {
    ui.info("Prepare Typescript Environment");
    await npm.install.save_dev("typescript", "ts-node", "ts-lint").code;
    ui.info("Install tslints");
    await npm.install.save_dev("@utilx/tslint-rules").code;

    await r(
      "node_modules/.bin/tsc", "--init",
      "-t", "ESNEXT",
      "--outDir", "./dist",
    ).code;

    const fe = FileEditor.load$("./tsconfig.json")!
    fe.locate(/^\s+}/)!.replace(/}/, "},")!.append(`
\t"include": [
\t\t"src"
\t]
`)
    fe.save$()
  }

const text = `
const path = require('path');

module.exports = {
  entry: '%%ENTRY%%',

  module: {
    rules: [{
        test: /\.tsx?$/,
        use: 'ts-loader'
        // configFile: "./tsconfig.json"
    }]
  },
  resolve: {
    extensions: ['.tsx', '.ts', '.js']
  },
  output: {
    filename: '%%OUTPUT_FILENAME%%',
    path: path.resolve(__dirname, 'webpack')
  },

  mode: "development",
//   mode: "production",
  target: "node",
  node: {
    __dirname: false,
  }
};
`;

export function initWebpackConfig(entry: string, outputFilename: string) {
    return text
        .replace("%%ENTRY%%", entry)
        .replace("%%OUTPUT_FILENAME%%", outputFilename);
}

export async function prepareEnv(pts: Set<PACKAGE_TYPE>) {
    const STEP = 4;

    UI.info(`(1/${STEP}): PackageJson`);
    await npminit("main");

    UI.info(`(2/${STEP}): Typescript`);
    await prepareTypescriptEnv();

    if (! pts.has("lib")) {
        UI.info(`(3/${STEP}): Prepared webpack`);
        const webpack = initWebpackConfig("./src/main.ts", "main.js")
        echo(webpack).tee("webpack.config.js")
    } else {
        UI.info(`(3/${STEP}): Prepared webpack SKIPPED`);
    }

    UI.info(`(4/${STEP}): Write .gitignore`);
    echo(GIT_IGNORE).tee("./.gitignore")

    UI.info(`(4/${STEP}): Write .npmignore`);
    echo(NPM_IGNORE).tee("./.npmignore")
}

export async function installPackages(pts: Set<PACKAGE_TYPE>) {

    UI.info("Install axios");
    await npm.install("axios").code;

    if (pts.has("koa")) {
        UI.info("Install koa pack");
        await npm.install(
            "koa", "@types/koa",
            "koa-router", "@types/koa-router",
            "koa-body",
        ).code;
    }

    if (pts.has("command")) {
        UI.info("Install os-command");
        await npm.install("git+ssh://git@github.com/edwinjhlee/os-command").code;
    }

}

type PACKAGE_TYPE = "command" | "koa" | "lib";
const candidateList = ["command", "koa", "lib"];
const candidateSet = new Set(candidateList);

export async function main(...pt: PACKAGE_TYPE[]) {

    for (const p of pt) {
        if (!candidateSet.has(p)) {
            UI.error(`${p} not inside the candidate list: ${candidateList}`);
            process.exit(1);
        }
    }

    const pts = new Set(pt);

    UI.info("Initialize Typescript Project");

    UI.h1("Prepare environemt");
    await prepareEnv(pts);

    UI.h1("Install Packages");
    await installPackages(pts);
}

main(... process.argv.slice(2) as any[]);

