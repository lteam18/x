# Goal

1. 完成下载功能
2. 加入vvkv的客户端功能
3. 执行本地的bash，node
4. vvsh-js以及js，将使用vvsh运行时
5. 对于code-type指定js, nodejs的，则与python等同例
6. vvsh-daemon：支持热启动，定时执行，远程登陆
7. vvsh-ssh

codetype = js ts sh fish py rb

# 增强功能

1. 加入gbt的文件copy-paste-share，以运行时方式实现
2. 加入自动安装设定：可以指定运行时按需安装
3. vvsh-link，指定hostname，实现转发


# vvsh的网络功能

1. 无中心的分布式
2. 有中心可加速
3. 自动共享，构建cdn


```bash
x github://edwinjhlee/

x repo  # Find

# Support it only
x @ljh/work # name/work
xi work

# Using github
x github://edwinjhlee/work.sh
xg work.sh
xg publish work work.sh


```


# xg token


```bash
xg token github://edwinjhlee <TOKEN>
xg token-sync github://edwinjhlee <TOKEN>
```

# 资费问题

1. 避免资费负担过重，每个人只能提供0.01元（10MB流量，1MB存储），100万程序员，运营1万元
2. DDoS攻击防止：必须使用某种token机制，该机制能够支持在lambda层断开，避开

