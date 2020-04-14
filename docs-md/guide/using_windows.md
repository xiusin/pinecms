---
title: 在 Windows 中运行 PINECMS
---

# 在 Windows 中运行 PINECMS #


::: warning 注意
如果需要使用`SQLITE`作为数据库， 需要安装gcc的命令。 这里推荐`mattn/go-sqlite3`推荐的包[TDM-GCC](https://sourceforge.net/projects/tdm-gcc/), 就此可以支持`CGO` 相关的包。 
:::

## 1. 直接下载编译好的安装包 ##
如果您不打算对代码进行任何修改， 可以直接[下载](https://github.com/xiusin/pinecms)编译好的安装包， 达到一键部署的目的。 

## 2. 修改配置文件 ##
### 配置数据库信息 ###
请复制`$project/resources/configs/database.yml.dist` 并且命名为：`database.yml`

在复制好的文件内修改如下配置:
```yaml
db:
  #dsn: "debian-sys-maint:3Av7BK8pUOaxn0Or@tcp(127.0.0.1:3306)/pinecms?charset=utf8"  mysql链接方式
  dsn: "./data.db"    # sqlite的链接方式
  db_prefix: "pinecms_" 
  db_driver: "sqlite3"  # 设置数据库驱动 sqlite3 或者 mysql

orm: # 配置orm信息
  show_sql: true
  show_exec_time: true
  max_open_conns: 10
  max_idle_conns: 10
```

### 配置application.yml ###
```yaml
port: 2019  #运行端口号

view:
  reload: true　#非开发时期，关闭reload提升性能
  fedirname: "./resources/themes/" #前端模板目录地址
  bedirname: "./resources/views/"　#后端模板地址

session:　#session相关
  name: "gosessionid"
  expires: 0

cache_db: "./runtime/cache.db"  # 全局缓存数据保存地址 (bbolt)
log_path: "./runtime/logs/" #日志目录，　包括运行日志和数据库日志

favicon: "./resources/assets/favicon.ico"　
charset: "UTF-8"
hashkey: "the-big-and-secret-fash-key-here"  # bug AES only supports key sizes of 16, 24 or 32 bytes.
blockkey: "lot-secret-of-characters-big-too"
backend_route_party: "/b"

max_bodysize: 32 # MB

upload:
  engine: "oss" #oss or file
  base_path: "assets" # 基本路径

statics:
  - {route: "/assets", path: "./resources/assets"}
  - {route: "/upload", path: "./resources/assets/upload"}
```

## 运行系统 ##
```shell 
pinecms.exe serve start
```
运行后显示如下信息(关于个人信息部分可以选择删除)：
```shell 
│─────────────│────────────────────────│
│ KEY (7)     │ VAL                    │
│─────────────│────────────────────────│
│ Name        │ xiusin                 │
│ Version     │ Development            │
│ Author      │ xiusin                 │
│ WebSite     │ http://www.xiusin.com/ │
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

至此, 我们到这里算是已经完全启动了服务. 

## 进入后台 ##

在浏览器上输入地址: `http://localhost:2019/b/login/index` 然后登录界面管理自己的网站内容. 

## 其他 ##
这里并没有像其他CMS一样启动install 安装配置目录.