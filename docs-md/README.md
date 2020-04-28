---
home: true
actionText: 进入手册
heroImage: /logo.png
actionLink: /guide/
features:
- title: 快速开发
  details: 以标签化的方式创建网站, 一杯茶的功夫即可. 只要会html就可以完全自主的搭建一个称心的网站.
- title: 自主框架
  details: 自主开发Web框架, 最少依赖打包, 可以让您的网站文件更小更易迁移. 
- title: 高性能
  details: 基于Go语言开发, 性能至上. 页面支持静态化, 让您的站点达到访问静态网站一样的性能.

footerColumn: 2
footerWrap: 
- headline: 案例
  items:
  - title: 素材网
    link: https://github.com/xiusin/pinecms.git
    details: 开发中

- headline: 其他项目
  items:
  - title: RedisDesktop
    link: https://github.com/xiusin/redis_manager
    details: Redis可视化工具
  - title: Pine
    link: https://github.com/xiusin/pine
    details: Web框架
    
footer: Apache-2.0 Licensed | Copyright © 2020-present
---


# 快速入门

## 直接下载二进制包
[GITHUB](https://github.com/xiusin/pinecms/releases)

[码云](https://gitee.net/xiusin/pinecms/releases)

## 下载源代码
使用 `git` 可用源代码:
```shell 
git clone https://github.com/xiusin/pinecms.git
```
## 编译代码
```shell
cd pinecms
go get -v -u
go build -o pinecms main.go
```

配置文件操作步骤请查看[安装](/guide/installation)

# 启动服务

```shell
./pinecms serve start

│─────────────│────────────────────────│
│ KEY (7)     │ VAL                    │
│─────────────│────────────────────────│
│ Name        │ PineCMS                │
│ Version     │ Development            │
│ Author      │ xiusin                 │
│ WebSite     │ http://pinecms.xyz/    │
│ PineVersion │ dev 0.2.1              │
│ Version     │ dev 0.1.2              │
│ GoVersion   │ go1.14                 │
│─────────────│────────────────────────│
   ___  _
  / _ \(_)__  ___
 / ___/ / _ \/ -_)
/_/  /_/_//_/\__/

Server now listening on: http://0.0.0.0:2019/
```

::: tip 提示
您可以通过执行`./pinecms example` 或 `go run main.go example`获取`演示案例`的数据库文件 、 模板文件、 素材文件。 
里面包含大部分模板标签的使用可以用来学习。
:::

# 进入后台管理
在浏览器地址栏输入: `http://localhost:2019/b/index/index` 即进入后端管理页面
- 默认账号: `admin`
- 默认密码: `123456`