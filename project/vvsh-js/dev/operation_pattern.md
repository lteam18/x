

- 运维步骤
    - 采用lambda
    - 三月份开始内测运营
    - 六月份进入公测运营
    - 九月份中英文发布
- 运维资金
    - buy me a coffee, paypal账号, 支付宝, 微信
    - 组织以及企业收费，paypal账号, 支付宝, 微信
    - 点对点加密 - 最高安全性
    - 页面广告 - 尽量不要浪费时间，外包

# 与npm的关系

## 世界观差异

实际上，瞄准的就是 `npm + node` 作为一个整体的差异。
它需要是一个稳定的发布，集合全功能的体系，执行python的社区哲学：
1. 重量级
2. 只有一种最好的解决方案

加上一个对于优质内容运营的关键：
1. 保证可维护性，即使牺牲一部分的简洁性
2. 书写维护性通过一个vim风格的提示器来快速提醒，确保代码的优雅
3. 采用typeshell这门语言来解决这个问题

## 竞争关系

1. 例如，我做了一个功能，这个功能可以实现自动同步；
    - 可以写成npm的package；生成bin/，然后全局可用
    - 可以写成nosh的形式，单文件，以url可运行

本质上，npm可以发展成这样一个环境，但是，有两个问题：
1. npm并没有集成node，这样意味着，node的升级
2. npm的global package跟node的版本相关联，运维应用的运行环境跟系统开发环境应用是两个不一样的概念，应该要分离。

万一，npm决定要作nodeshell的事情，那么：
1. 采用npm以及node
    1. 安装：npm instal ubermensch -g
    2. 运行：ubermensch
    3. ubermensch的依赖会导致 npm 的更新，另外，我每次发布ubermensch，要重新安装
2. 采用nosh
    1. 运行：nosh :ubermensch
    2. 相当于自动安装，并且执行ubermensch；ubermensch的依赖并不影响环境

在一个新环境，需要安装的步骤：

1. 安装node，nvm的多环境会让事情更复杂；
2. 脚本下载nosh，只需一个exe以及home_dir，就可使用

如果要完成nosh的事情，意味着node要改变社区哲学：

1. 采用重量级的封装 - 如同python一样，集成大量packages
2. 独裁，一种功能只有一种api提供
3. 编写cookbook，提供quick access

只有一个社区，叫做 @typeshell，取名自一门还没发布的语言

## 统一：nosh作为后发者，需要

### nodeshell可以去读npm的包

typeshell as a language: a language for the future
nodeshell as a facility: an environment with node and npm

nosh & yesh

### nodeshell可以作为npm的包进行安装nosh & yesh

然后整体使用

### nodeshell可以



