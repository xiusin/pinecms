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

- [ ] 模板静态化功能
- [ ] 图片上传 功能修改针对ueditor的多图上传弹性输出多选或单选数据(根据组件类型)
- [ ] 添加导入数据命令

### 细化文档模型 ### 

1. 重新定义一下系统模型和独立模型(参考shuipfcms)
2. 默认字段可以设置为显示隐藏, 可以为数字类型的组件设置随机值
3. 默认字段添加title keywords description的(建立模型表的时候直接显示为不可删除状态)
4. 为字段添加排序值, 下次进入页面按照排序值调整
5. 文档模型
    - 系统模型（以附表方式关联数据，可级联管理， 表明是否以模型ID做附表方便关联数据）
    - 独立模型（单独生成数据表）    
6. 考虑模型的递归显示，如果分布在不同的表里如何组合显示。
7. 如果放到同一张表里字段数据使用json保存如何。
系统字段显示， 搜索字段展示， 页面静态化 参考水平凡cms的实现

## 第二期 ## 

- [ ] 数据模块
- [ ] plugin (模块管理)


# 相似的php项目 #
https://pro.cltphp.com/admin/index/index.html
