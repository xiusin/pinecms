# 下载部署 #

1. 下载源码
    > `go get -u -v github.com/xiusin/pinecms`

2. 数据库配置
    > 导入数据库结构`resources/pinecms.sql`
    > 修改`resources/configs/database.yml.dist`为`resources/configs/database.yml`

4. 安装依赖
    > `go get -v`

5. 运行项目
    > `./iris serve start` 

6. 开发期间自动构建
    > `go run main.go serve dev`

7. 访问后端登陆页面
    > 访问 `http://localhost:2019/b/login/index`
    > 默认账号密码 `用户名: admin 密码: admin888`

# TODO #

## 第一期 ##

- [ ] 模板静态化功能可以使用nginx代理静态资源目录, 找不到资源的转发给pinecms生成静态文件
- [ ] 添加导入数据命令

### 细化文档模型 ### 

1. 重新定义一下系统模型和独立模型(参考shuipfcms)
2. 默认字段可以设置为显示隐藏, 可以为数字类型的组件设置随机值
4. 搜索字段默认设置

## 第二期 ## 

- [ ] 数据模块 浏览次数, 搜索引擎来源, 服务器监控, 待审核信息, 会员信息等
- [ ] plugin (模块管理, 是否要动态状态 使用外部plugin的注入看性能如何)


# 相似的php项目 #
https://pro.cltphp.com/admin/index/index.html
