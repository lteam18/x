# vvx的多种形态

## vvx0 + vvsh

解压：13K + 90MB
延时：2ms + 80ms

充分采用node的技术，并且复用当前的技术栈

问题：

1. node核心更新，就会触发一次更新，更新还需要特权
2. vvsh很大，解压后90MB（js + ts），跟vvx-go的6MB相比，是一个量级的差别

## vvx0 + vvx-go ( + vvsh )

解压后：13K + 6MB
延时：2ms + 10ms

这里面包含了SSL以及各种网络能力，跨系统。

## vvx0 ( + vvsh )

解压：200K - 1MB
延时：2ms - 10ms

具备全功能

我不认为这是一个好的方式：

1. 漏洞
2. 代码维护复杂度
3. 代码迭代发展的复杂度

另外，我发现，如果需要引入http以及ssl等处理，最终解压后的大小还是会到5MB（这个结论来自于curl on windows）

curl看起来是183K，是因为假设了系统本身存在某些库。

```
https://packages.ubuntu.com/xenial/curl
```


