# vvsh

## vvsh究竟与npx有什么差别：

表面上，两者相同，能够通过简单命令，诸如`vvsh xxx`就能启动脚本。拥有一个上传代码的社区，能够支持一键运维。

事实上，我们还能做到：

### 提供更简单的开机自启的能力，以及类似crontab的能力

vvsh-daemon将是其中极其关键的设计

### 作为ssh的一个替代

vvsh-daemon能够启动一个安全的，基于https的加密反向代理连接
这个代理将会直接连接服务，本地设定口令，然后远端通过服务器进行加密认证。

我们将会进行轻微收费：每个用户0.99美元/月，默认5个连接；中转流量，1GB/月；超出部分，1元/G，否则按照1KB/s来运作。

1. 正常情况，服务器自动协商端口，进行通讯
2. 按照需要，提供反向服务，按流量来计算
3. 可以考虑提供打洞技术，这个用来提高用户的满意度

估计每个服务器能承受10K连接，能够服务1K个用户，每个用户大约10美元/年，10000美元/年，一个服务器一年就600美元/年
10美元 * 10万用户 = 100万美元
管理100台弹性服务器，即可；分散在不同区域。

后面主要就是：
1. 服务器管理
2. 手机app的开发
3. 云桌面，替代windows/linux/macos/bsd/android，发行基于linux内核云桌面

核心技术就是wasm

这里面最关键的，就是要快，比开源要快；并且事先实现将大量已有应用的绑定，形成强大的利益联盟。
以生态来构建，后面不断开源技术，并且通过开发者分成来吸引高质量的开发者。
最后，将大部分非核心代码都以开源社区模式运营，持续降低开发成本以及招聘和训练成本（每年举行hackathon，来吸引优秀开发者了解我们的系统，从而降低招聘和训练成本）。



### ansible的替代

替代puppet，ansible，用于多台服务器的管理

### WebIDE与远程桌面功能

1. 缓存类似tmux的远程进程 - 小广告
2. 可以启动web桌面以及ide，并作基本的管理 - 另外收费10美元/年；
3. Android/IOS app管理，提供mobile版本，可以一键启动命令脚本
4. 能够作为入口，售卖云服务器
5. WebIDE升级，用来作为SaaS开发/部署，高性能计算入口，软件测试入口，各种软件的部署，例如项目管理以及测试管理，等等
6. 用这个功能来重塑开源生态，提供有价值开源

### 加速器：github/npm/apt，用来改善中国区的开发情况，

用加速镜像来加快开发速度；按照流量来计算 1元/1GB；等流量降到0.1元/GB，意味着1个月1元即可。
1. 采用bt共享的方法

vvsh proxy git
vvsh proxy apt
vvsh proxy npm

2. 指定节点的方法

1. github加速器，采用一个命令，来加速github，支持windows/linux/mac
2. npm/apt加速器

只支持github/npm/apt的流量
