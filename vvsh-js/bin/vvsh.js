#! /usr/bin/env node

const fork = require("child_process").fork

process.env["NODE_NO_WARNINGS"] = 1
const child = fork(`${__dirname}/../dist/main.js`, process.argv.slice(2))
