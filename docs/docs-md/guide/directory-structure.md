# 目录结构 

## 目录结构 
PINECMS 遵循`约定优于配置`的原则, 目录结构如下:
```
.
├── LICENSE
├── data.db (sqlite3数据文件)
├── main.go (主入口源代码文件)
├── pinecms.exe (可执行二进制文件)
├── pinecms.sql (mysql文件)
├── resources   (素材目录)
│   ├── assets (对外模板素材)
│   ├── configs (系统配置)
│   ├── fonts (验证码字体)
│   ├── html    (静态文件目录)
│   ├── taglibs (系统内置标签目录)
│   ├── themes  (前端主题目录)
│   └── views  (后端页面目录)
├── runtime (运行时目录)
│   ├── cache.db
│   └── logs
└── src (源代码目录)
    ├── application (controller/models/middlewares目录)
    ├── common (核心组件目录)
    ├── config (配置定义目录)
    ├── router (路由注册, 前端目录在此配置)
    └── server (启动服务)
```

- `src` 目录是源代码目录, 如果下载的编译包, 不存在此目录. 
- `resources/assets` 前后端素材对外目录, 系统上传的文件也为保存在此目录下. 
- `data.db` sqlite3的文件, 可以直接使用. 

## 素材目录

PineCMS 推荐素材前后端分离防止, 目录结构如下: 

```
.
├── backend
│   ├── highlight
│   ├── lightgallery
│   ├── static
│   └── ueditor
├── frontend
│   ├── css
│   ├── images
│   ├── js
└── upload
    └── public
```

- `backend` 后端素材目录
- `frontend` 前端素材目录, 各主题可以单独保存到一个文件夹
- `upload` 上传素材目录. 数据库备份会放置在此目录下(`只有登录后端的会话才可下载备份`)