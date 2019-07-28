console.log("I am rising.")


async function main(astr){
    console.log(astr)
}

main("hello")

const osc = require("os-command")
// console.log(osc)
console.log(osc.network.getIPV4Address())

//+use cmd-type https://edwinjhlee.github.io/cmd-type/cmd-type.js
const cmd_type = require("cmd-type")
cmd_type.helloworld()

