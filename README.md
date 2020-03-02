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

## 第一期 ##

- [ ] 图片裁切,实现上传以后传入x1,x2,y1,y2
- [ ] 模板静态化功能
- [ ] 前端模板标签开发（Jet）

### 细化文档模型 ### 

- [ ] 文档模型
    - 系统模型（以附表方式关联数据，可级联管理， 表明是否以模型ID做附表方便关联数据）
    - 独立模型（单独生成数据表）

系统字段显示， 搜索字段展示， 页面静态化

## 第二期 ## 

- [ ] 广告位设计， 可以先前端埋点， 然后自动注册， 后台更新对应点的图片内容
- [ ] 数据模块
- [ ] plugin (模块管理)

# 关于文档模型 #

1. 考虑模型的递归显示，如果分布在不同的表里如何组合显示。
2. 如果放到同一张表里字段数据使用json保存如何。

# 相似的php项目 #
https://pro.cltphp.com/admin/index/index.html
