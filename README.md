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
    > `./iris serve dev`

7. 访问后端登陆页面
    > 访问 `http://localhost:2019/b/login/index`
    > 默认账号密码 `用户名: admin 密码: admin888`

8. 使用 `Nginx` 或者 `Caddy` 反向代理到 `:2019` 或者自定义 端口即可

# 路由配置与实现 #
在`config/router.go`中按照已有配置实现相关的前后端路由,在控制器文件内务必实现`func (*XXController) BeforeActivation(b mvc.BeforeActivation)`进行路由注册, 然后实现各个方法的功能与需求.

> 目前功能比较简单, 对付简单的企业站应该是没问题. 其他的酌情自行开发 ^_^

# TODO #
- [ ] 模板静态化 
- [ ] 图片裁切,实现上传以后传入x1,x2,y1,y2
- [ ] 精简代码以及后端模块， 删除自用表. 
- [ ] 迁移数据模块
- [ ] 文档模型(存储库结构,字段描述, 数据页面根据文档模型动态展示内容) doing


# 查看性能 #
1. 访问pprof页面 
2. 使用tool命令打印图形 如: `go tool pprof -http=0.0.0.0:1234 profile`
