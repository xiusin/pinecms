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

# 微信模块 

- 账号授权
![账号授权](./images/wechat-account.png)
  
- 会员管理
![会员管理](./images/wechat-member.png)
  
- 菜单管理
![菜单管理](./images/wechat-menu.png)
  
- 消息管理
![消息管理](./images/wechat-msg.png)
  
- 模板管理
![消息模板](./images/wechat-template.png)


# Doing
- 权限系统完善到按钮级别
- 系统内部BUG修复，图片处理
- 图片系统以及md5处理
- cms系统完善表单字段自定义配置（可自定义模型页面），搜索字段配置（不使用高级搜索表单）
- 插件公共页面配置 （.so下载，源代码下载）(尽可能不让二次开发)
- 插件允许暴露公共前端页面
- 微信插件
    - 素材管理
    - 自动回复素材功能
    - 客服消息
    
- redis manager管理插件
    
# TODO  
- github.com/gokeeptech/gktemplate
- Bleve 全文检索 （插件提供）
- hook方式支持插件钩子注入信息
- 一键CRUD
- http://8.140.114.57:7000/opsli-boot/doc.html#/opsli%202.X/%E4%BB%A3%E7%A0%81%E7%94%9F%E6%88%90%E5%99%A8-%E6%97%A5%E5%BF%97/createUsingGET
- https://github.com/lljj-x/vue-json-schema-form
- http://goframe.ele.rxthink.cn/tool/generate
- https://gitee.com/unifig/unifig-admin?_from=gitee_search

