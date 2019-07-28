我现在可用的名字有：

nodeshell => nosh => ns/nh
typeshell => tysh => ts/th

我倾向于将typeshell理解为一种新的语言，而 nodeshell 则是一个生态。

ty，peshell的工作可以尽快完成。可以只是一个小语言前端，转化成 js。

以可视化和语言提示为目标，不妨僵硬。但是输入内容要少。

我认为，兼容javascript是一个重要内容。

目标是：

1. 设计一种新的语言，typescript的简化版本。
2. 建立language service

Integrate the logging facility

```go
func work(){
    info("Please notice")
    error("")
    warn("")
}

() => e
a = 3 // you could not giving var
a, b = 3, 4 // support this candy

command "-alh" "abc" 3
command("-alh", "abc", 3)

// work is function
command work 3
command(work, 3) => command work 3 tg 4
command(work(3)) => command work(3) tg(4)

support scala like feature
```



