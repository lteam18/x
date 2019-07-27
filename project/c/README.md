# vvx-c

作为一个加强模块，目标要提升程序启动速度；

1. 对于本地文件
2. 对于已经缓存的vvuri代码
3. 对于热启动时间

# How

1. 首先解析为本地文件，尝试阅读
2. 如果是vvuri，尝试读取本地cache
3. 配合热启动模式，将启动速度控制在一个极低的程度

其他情况，本来需要一定的额外时间，我们无法优化，因此采用vvx-go来实现：

1. 没有本地cache，需要网络下载
2. 需要补全的vvuri
3. 其他命令

# 支持命令

```bash
vvx bash test.bash
vvx cat @official/hi.py @official/ding-ip
vvx python @official/hi.py
```

# 调研 libcurl

libcurl库，是跨系统，实现http client
测试引入之后的size，以及启动时间
