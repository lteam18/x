指定preload网页

1. 某定义命令
2. 例如


```typescript

//+helper:latest

// More facility about work
//+https://edwinjhlee.github.io/bash

//+https://abc/asdfa/asdfaf/asdf
import { docker, apt, run } from "command-helper"

//+https://abc/adfas/asdfas/asdfa
docker.run("")

```

--------

这个问题不像是我想象中那么简单。
我们需要解决的问题是，如何有效嵌入module

1. 参考nexe的方案
2. 

--------

# cache 机制

##  指定一个对应列表下载位置 apt-list

nodeshell app update

nodeshell app ts-init

1> app found in app cache
2> app found in app

1> app found in app cache
1.1> running

2> app not found in app cache
2.1> app not found in app cache list.


2.2> app found in app cache list
2.2.1> auto update list
2.2.1.1> network failed
2.2.1.2> running

nodeshell js
nodeshell ts
nodeshell 

## Work

ns list
ns update
ns upgrade [:app-name] | [url/http] | [filepath]
ns run [:app-name] | [url/http] | [filepath]
ns ts [:app-name] | [url/http] | [filepath]
ns js [:app-name] | [url/http] | [filepath]


nodeshell index list => ns ls
nodeshell index update => ns up

nodeshell app [url | file]
nodeshell js [url | file]
nodeshell ts [url | file]

nodeshell alias abc url
nodeshell alias

nodeshell self-update
> 提示采用的权限

## 自由安装包



