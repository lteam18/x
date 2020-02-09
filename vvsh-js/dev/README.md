# 使用手册

一切从这里开始：

```bash
curl -s https://edwinjhlee.github.io/cmd-type/helloworld.js | node - "arg1" "arg2" "arg3"
```

我们一般细化将代码放到一个公开可用的位置，然后采用上述命令实现启动。

但实际上，我们遇到如下三个问题：

1. 不是所有系统都有curl；但是有node了，curl的功能是能实现的
1. js代码如果涉及到某些库，需要webpack成一个js文件，使用者难以进行二次修改；
2. 从脚本的可运维角度，typescript文件直接运行，可读性更好
2. 需要在当地安装特定nodejs，可能是某个特定版本，这样势必影响部署系统的node环境。

我们这个项目尝试解决：

1. 固定问题
    - 可直接加载ts文件，当然，js文件也支持
    - 我们将这个nodejs环境打包成单个可执行文件，与系统node环境平行
    - 以后考虑加载类似coffeescript之类的脚本
3. 额外需求：
    - 脚本别名 以及 运维系统
        - 我们觉得可以将一些常用的脚本放在一个启动列表内，然后通过标识符来快速启动，这样就不需要借助外部力量了。
        - 可登陆。设计一个简单的邮件运维系统进行登陆设置。利用github作分发
        - 该系统的赢利采用主页广告赢利，控制篇幅，资金用于资助服务器
        - 私有repo以及加密分发需要资金
    - 脚本加密
        - 可以设定点对点加密，脚本采用公钥加密
    - 自动更新
    - 内置库的自动更新以及引用新版本：问题搁置
        - 做一个自动更新 - 最新库，在import上加入注释
    - 增加更多的内置库：
        - 减轻webpack打包
        - 增加更多的功能：例如puppeteer
        - 成为更多程序的客户端，例如mongo，redis，以及其他数据库

理想情况下，采用该脚本能达到如下目标：

```bash
nosh :ts-init
nosh :get-ip

nosh :mongo/dbman
nosh :mongo/work
nosh org mongo
nosh :dbman

nosh https://org.github.io/exmaple-script.ts
nosh local_script.ts
```



## 安装

### 脚本

Linux/Mac用户： 

If you are using linux or mac, please using following command:

```bash
curl https://edwinjhlee.github.io/vvsh/install.sh | sudo bash
```

Windows:

```doc
...
```

### 直接下载


https://github.com/lteam18/auto/releases/download/v0/nodeweb.linux.exe
https://github.com/lteam18/auto/releases/download/v0/nodeweb.mac.exe
https://github.com/lteam18/auto/releases/download/v0/nodeweb.win.exe


# 开发手册

## 增强功能

1. js: 默认可读，默认由os-command包
2. ts: 内置ts-loader，需要在前面备注这是tsshell，会直接翻译


## 需求列表：

### 基本功能：

1. 自动从os-command库获取最新的代码；可以将代码缓存到`/tmp`，但需要校验
2. 输出手动更新脚本的地址
3. os-command只保留最小化的功能代码
4. 将命令行代码以远程require方法提供
5. 官方更新command-helper最新版


### 社区生态：

1. 自动获取scripts列表
2. 提供scripts加载
3. 版本规划，确保版本能保证这些script的运作正常


