# x命令

## 安装

```bash
x install       # 安装，默认全部
x install-all   # 安装全部
x install-x     # 安装x
x install-nosh  # 安装nosh
x update-nosh   # 重新安装nosh
```

## 上传脚本

```bash
x upload public @username/ip ip.js js
```

## 执行的脚本

nosh

```bash
x @username/ip
x nosh @username/ip
x nosh ip.ts
```

python

```bash
x python ip.py
```

cat: 查看资源

```bash
x cat ip
```

执行其他脚本，请查看后面的章节

## 执行命令

使用watchdog功能

```bash
x --retry 1 --interval 10 cmd ping www.bing.com
```
