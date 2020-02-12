# 下载部署 #

1. 下载源码
    > `go get -u -v github.com/xiusin/iriscms`

2. 数据库配置
    > 导入数据库结构`resources/iriscms.sql`
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
- [ ] 模板静态化
- [ ] 图片裁切,实现上传以后传入x1,x2,y1,y2
- [ ] 迁移数据模块
- [ ] 文档模型(存储库结构,字段描述, 数据页面根据文档模型动态展示内容, 以及相关的权限管理)，系统字段显示
- [ ] plugin
- [ ] 广告位设计， 可以先前端埋点， 然后自动注册， 后台更新对应点的图片内容
- [ ] 数据模块

# bug #
- [ ] dev无法监听yml的注册
- [ ] easyui的模型必填单选框如何实现多字段支持

# 关于文档模型 #

1. 考虑模型的递归显示，如果分布在不同的表里如何组合显示。
2. 如果放到同一张表里字段数据使用json保存如何。

# 相似的php项目 #
https://pro.cltphp.com/admin/index/index.html