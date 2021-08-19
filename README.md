# 项目描述 #
PineCMS是一个GO语言开发的内容管理系统, 让您可以在短时间内以制作模板的方式搭建出来一个网站, 非开发者也能快速愉悦的使用系统. 
简单使用情况下无需关注代码逻辑. 

![](./resources/assets/backend/static/images/1.png)

# 编译部署 #

## 下载并编译 ##
 ```
 git clone https://github.com/xiusin/pinecms.git
 cd pinecms
 go build -o pinecms
```

## 配置 ##
1. 拷贝数据库文件 `data.db.dist` 并且命名为 `data.db`

2. 数据库配置
    > 导入数据库结构`resources/pinecms.sql`
    >
    >修改`resources/configs/database.yml.dist`为`resources/configs/database.yml`
    >
    > 配置数据源

2. 安装依赖
    > `go build`

3. 运行项目
    > `./pinecms serve start` 

4. 开发期间自动构建
    > `go run main.go serve dev`

6. 访问后端登陆页面
    > 访问 `http://localhost:2019/admin/` 默认账号密码 `用户名: admin 密码: 123456`

# 自动静态化路由 #
完全自动静态文件和动态路由. 更友好的SEO方式

# 主题系统 #
系统支持多种主题, 可自由切换不同风格的模板. (一般需要相同类型模型, 字段一致, 减少错误问题) 

# 内置标签系统 #
支持类似织梦的标签系统, 可以让您在不写任何代码的情况下完成网站建设. 

# 文档 #
[doc.xiusin.cn](http://doc.xiusin.cn/)

# 演示 # 
- 首页 http://pinecms.xiusin.cn/
- 后端 http://pinecms.xiusin.cn/b/login/index 
- 新版后端 http://new_pinecms.xiusin.cn/


# 新功能

### 插件系统 (doing)
支持动态插拔插件, 并注册到系统功能, 提供方便便捷的扩展功能. 
系统可以动态扫描插件目录,自动发现并可以热加载进系统.  
也可以导入第三方人员开发的扩展动态库(受限于系统和版本,后面会提供编译个版本的docker镜像)

- 插件系统界面
![插件系统界面](./images/plugin.png)


### 服务器监控 
> 常用的系统资源监控, 以及环境检测

![服务器监控界面](./images/stat.png)

### 模型配置
- 模型列表
> 模型管理界面

![模型列表界面](./images/model.png)
  
- 模型变更SQL
> 当模型变更时,会检测变更生成执行SQL, 需要手动操作

![模型变更SQL](./images/presql.png)


- 模型字段界面
![模型字段界面](./images/field_list.png)
  
- 添加字段界面
> 内置CMS常用字段, 可以设置表单显示, 列表显示, 可搜索字段, 搜索类型

![添加模型字段](./images/add_field.png)

- CRUD命令
> 允许直接从表创建crud模块, 根据表字段自动解析form, 区别于cms是此命令仅适用于开发区间 

![crud命令](./images/crud.png)

# Doing
- 重写为前后端分离
- 一键CRUD(view models controllers menus)

# TODO  
- github.com/gokeeptech/gktemplate
- github.com/traefik/yaegi `作为脚本处理`
- Bleve 全文检索
- 微信插件
- daemon 守护进程
- 插件公共页面配置 (尽可能不让二次开发)
- hook方式支持插件钩子注入信息
- http://8.140.114.57:7000/opsli-boot/doc.html#/opsli%202.X/%E4%BB%A3%E7%A0%81%E7%94%9F%E6%88%90%E5%99%A8-%E6%97%A5%E5%BF%97/createUsingGET
- https://chu1204505056.gitee.io/admin-pro/#/mall/goods
- https://editor.yanmao.cc/zh-CN
- https://github.com/textbus/textbus
- github.com/mojocn/base64Captcha 验证
- pine记录函数参数反射缓存
