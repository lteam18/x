console.log("Wait 10 sec")

setTimeout(() => {
    console.log("Exit")
}, 10 * 1000)

process.on("SIGINT", (...args)=>{
    console.log(args)
    process.exit(0)
})
