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

1. 目前存在的问题 缓存采用badger, 但是2.0.2版本在windows下存在开发时存在丢失数据的情况. 如果无法搭建起来, 请留言. 撸不易
### 细化文档模型 ### 

1. 重新定义一下模型以及使用场景的复杂度
2. 字段可以设置为显示隐藏, 可以为数字类型的组件设置随机值
4. 搜索字段默认设置

## 第二期 ## 

- [ ] 数据模块 浏览次数, 搜索引擎来源, 服务器监控, 待审核信息, 会员信息等
- [ ] plugin (模块管理, 是否要动态状态 使用外部plugin的注入看性能如何)
