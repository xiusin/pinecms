# 项目描述 #


# 下载部署 #

1. 下载源码
    ```
    git clone https://github.com/xiusin/pinecms.git
    cd pinecms
    ```
2. 数据库配置
    > 导入数据库结构`resources/pinecms.sql`
    >
    >修改`resources/configs/database.yml.dist`为`resources/configs/database.yml`
    

3. 修改go.mod (由于框架和cms是联动开发,直接本地映射了)

    删除掉 `replace github.com/xiusin/pine => ../pine`

4. 安装依赖
    > `go get -v`

5. 运行项目
    > `./pinecms serve start` 

6. 开发期间自动构建
    > `go run main.go serve dev`

7. 访问后端登陆页面
    > 访问 `http://localhost:2019/b/login/index`
    > 默认账号密码 `用户名: admin 密码: 123456`

# 系统关键描述 #

1. 目前存在的问题 缓存采用badger, 但是2.0.2版本在windows下存在开发时存在丢失数据的情况. 如果无法搭建起来, 请留言. 撸不易

## 文档模型 ## 

1. 支持完全自主定义模型, 模型会默认添加一些公共字段, 且不可删除
2. 支持隐藏字段, 不显示在发布表单上
3. 支持后端列表字段设置, 可以自定义设置字段显示关闭. 并且设置字段自己的formatter

## 前端 ##

完全兼容静态文件和动态路由. 
1. 规则具体依托后台为每个分类设置的静态目录参数
    - 分类资源生成静态地址下的index.html index_{page}.html 
    - 关于文档 根据分类设置的静态地址, 在其目录下生成 {aid}.html
    - 单页面和分类表现一致只存在index.html
    - 如果没有设置, 单页面分类会定义为page_{tid}, 文档分类会定义为 {mode_table}_{tid}
2. 支持不同的模板主题,可以自由切换.
3. 系统内置类dede标签, 可以让织梦用户的使用成本降到最低. 


# 需求 #

1. 重新定义一下模型以及使用场景的复杂度
2. 搜索字段默认设置
3. 优化前端标签, 可以参考多种CMS的实现, 聚合优点
	fmt.Println("直接通过Nil反射类型", reflect.TypeOf((*logger.AbstractLogger)(nil)).Elem())

# 第二期 #
- [ ] 数据模块 浏览次数, 搜索引擎来源, 服务器监控, 待审核信息, 会员信息等
- [ ] plugin (模块管理, 是否要动态状态 使用外部plugin的注入看性能如何)