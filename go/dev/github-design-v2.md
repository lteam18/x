# 采用github的方式来实现

**为什么不采用gist**

采用这种设计的最大考虑是：

1. 使用一个正常的repo时，如果该repo没有`meta.x-cmd.com`, 可以通过x @gh:edwinjhlee:ownrepo/work.py来实现直接调用

# 当前设计：

每一个git repo的文件夹布局是：

`<owner>/source.x-cmd.com`是一个私有repo

```js
// 私有文件
|- meta.x-cmd.com/
  |- hi.py
  |- hi
    |- js
    |- py
|- hi.py
|- hi
    |- js
    |- py
// 共享文件
|- docs/
    |- meta.x-cmd.com/
        |- hi.share.py
    |- hi.share.py
```

## 共享机制

我们共享一个文件的机制是，将文件由`APP`和`APP_IDX`移动到`docs/APP`和`docs/APP_IDX`

1. 公开访问文件列表不对外暴露
2. 可以公开访问
3. 具有临时URL访问能力

