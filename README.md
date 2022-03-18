
# 项目描述 #

PineCMS是一个GO语言开发的内容管理系统, 让您可以在短时间内以制作模板的方式搭建出来一个网站, 非开发者也能快速愉悦地使用系统.
简单使用情况下无需关注代码逻辑.

> 前端仓库: <https://github.com/xiusin/pinecms-web.git> <br/>
> 目前程序正在持续迭代开发中，提供开发期间的数据库文件`resources/pinecms.sql`, 后期数据库结构变更比较大。<br>

> 推荐一个自己写的Redis管理工具, 支持Web端: <https://github.com/xiusin/redis-web-manager.git>

# 编译部署 #

## 下载并编译 ##

---

```markdown
 git clone https://github.com/xiusin/pinecms.git
 cd pinecms
 go build -o pinecms
```

---

## 配置 ##

1. 执行数据链接生成命令: `./pinecms serve install`

2. 数据库配置
    > 导入数据库结构`resources/pinecms.sql`
    >
    >修改`resources/configs/database.yml.dist`为`resources/configs/database.yml`
    >
    > 配置数据源

3. 安装依赖
    > `go build`

4. 运行项目
    > `./pinecms serve start`

5. 开发期间自动构建
    > `go run main.go serve dev`

6. 访问后端登陆页面
    > 访问 `http://localhost:2019/admin/` 默认账号密码 `用户名: admin 密码: 123456`

# 内置模块 #

1. 用户管理：用于维护管理系统的用户，常规信息的维护与账号设置。
2. 角色管理：角色菜单管理与权限分配、设置角色所拥有的菜单权限。
3. 菜单管理：配置系统菜单，操作权限，按钮权限标识等。
4. 职级管理：主要管理用户担任的职级。
5. 岗位管理：主要管理用户担任的岗位。
6. 部门管理：主要管理系统组织架构，对组织架构进行统一管理维护。
7. 操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
8. 字典管理：对系统中常用的较为固定的数据进行统一维护。
9. 配置管理：对系统的常规配置信息进行维护，网站配置管理功能进行统一维护。
10. 城市管理：统一对全国行政区划进行维护，对其他模块提供行政区划数据支撑。
11. 友链管理：对系统友情链接、合作伙伴等相关外链进行集成维护管理的模块。
12. 个人中心：主要是对当前登录用户的个人信息进行便捷修改的功能。
13. 广告管理：主要对各终端的广告数据进行管理维护。
14. 站点栏目：主要对大型系统网站等栏目进行划分和维护的模块。
15. 会员管理：对各终端注册的会员进行统一的查询与管理的模块。
16. CMS管理: 可以系统设置不同的模型数据并添加不同的逻辑. 支持多主题。
17. 插件管理: 可以扩展/下载第三方开发的软件包。
18. 微信管理: 可以管理素材，会员，信息，自动回复等。

# 文档 #

[doc.xiusin.cn](http://doc.xiusin.cn/)

# 演示 #

- <http://pinecms.xiusin.cn/admin/>

# 新功能 #

## ApiDoc 管理插件 ##

内置集成接口ApiDoc插件, 支持从请求参数到响应结果的自动生成`略微侵入代码: 需要设置分组名称以及接口名称`. 可以在管理界面修改参数以及全局参数, 接口db可以随意迁移, 支持直接调试请求.

> 接口文档仓库: <https://github.com/xiusin/pinecms-apidoc-ui.git>

<table>
	<tr>
        <td><img src="https://raw.githubusercontent.com/xiusin/pinecms/react-pinecms/images/apidoc-detail.png"/></td>
        <td><img src="https://raw.githubusercontent.com/xiusin/pinecms/react-pinecms/images/apidoc-debug.png"/></td>
    </tr>	 
</table>


## 插件系统 ##

支持动态插拔插件, 并注册到系统功能, 提供方便便捷的扩展功能.
系统可以动态扫描插件目录,自动发现并可以热加载进系统.  
也可以导入第三方人员开发的扩展动态库(受限于系统和版本,后面会提供编译个版本的docker镜像)

## 系统截图
<table>
    <tr>
        <td><img src="https://raw.githubusercontent.com/xiusin/pinecms/react-pinecms/images/plugin.png"/></td>
        <td><img src="https://raw.githubusercontent.com/xiusin/pinecms/react-pinecms/images/stat.png"/></td>
        <td><img src="https://raw.githubusercontent.com/xiusin/pinecms/react-pinecms/images/model.png"/></td>
        <td><img src="https://raw.githubusercontent.com/xiusin/pinecms/react-pinecms/images/presql.png"/></td>
    </tr>
    <tr>
        <td><img src="https://raw.githubusercontent.com/xiusin/pinecms/react-pinecms/images/field_list.png"/></td>
        <td><img src="https://raw.githubusercontent.com/xiusin/pinecms/react-pinecms/images/add_field.png"/></td>
        <td><img src="https://raw.githubusercontent.com/xiusin/pinecms/react-pinecms/images/crud.png"/></td>
        <td><img src="https://raw.githubusercontent.com/xiusin/pinecms/react-pinecms/images/wechat-account.png"/></td>
    </tr>
	<tr>
        <td><img src="https://raw.githubusercontent.com/xiusin/pinecms/react-pinecms/images/wechat-member.png"/></td>
        <td><img src="https://raw.githubusercontent.com/xiusin/pinecms/react-pinecms/images/wechat-menu.png"/></td>
        <td><img src="https://raw.githubusercontent.com/xiusin/pinecms/react-pinecms/images/wechat-template.png"/></td>
        <td><img src="https://raw.githubusercontent.com/xiusin/pinecms/react-pinecms/images/wechat-member.png"/></td>
    </tr>	 
	<tr>
        <td><img src="https://raw.githubusercontent.com/xiusin/pinecms/react-pinecms/images/wechat-menu.png"/></td>
    </tr>	 
</table>


# Doing #

- 权限系统完善到按钮级别
- 系统内部BUG修复
- cms系统完善表单字段自定义配置（可自定义模型页面），搜索字段配置（不使用高级搜索表单）
- 插件公共页面配置 （.so下载，源代码下载）
- 微信插件
  - 素材管理
  - 自动回复素材功能
  - 客服消息

# TODO #

- Bleve 全文检索 （插件提供）
- 页面编辑器: <https://github.com/lljj-x/vue-json-schema-form>
- <http://goframe.ele.rxthink.cn/tool/generate>
- 菜单表增加perms字段， 可以配置权限标识，可以配置节点类型为权限。
- <http://relation-graph.com/#/demo/scene-network> 组织架构显示
- 动态构建表单: <https://eddyzhang1986.github.io/antd-jsonschema-form/>
- <http://fundemo.funadmin.com/2KmvVJA8dU.php/index/index.html>
- <https://github.com/wangyuan389/mall-cook> 低代码商城搭建平台
- <http://fast-crud.docmirror.cn/element/#/crud/basis/value-change>
- https://www.npmjs.com/package/dc-search-table
- https://github.com/huzhushan/vue3-pro-table

[comment]: <> (https://www.dowebok.com/demo/6918/)
  
<!-- 参考CMS: http://demo2.wooadmin.cn/run -->
<!-- http://pigx.pig4cloud.com/#/mp/wxaccountfans/index -->
<!-- 热门语言卡片 -->
<!-- https://github-readme-stats.vercel.app/api/top-langs/?username=xiusin&layout=compact -->

<!-- 统计卡片 -->
<!-- https://github-readme-stats.vercel.app/api?username=xiusin&show_icons=true&theme=radical -->

<!-- https://github-readme-streak-stats.herokuapp.com/?user=xiusin&theme=monokai-metallian&hide_border=true -->

<!-- https://github.com/ashutosh00710/github-readme-activity-graph -->
<!-- <a href="https://github.com/ashutosh00710/github-readme-activity-graph"><img alt="xiusin's Activity Graph" src="https://activity-graph.herokuapp.com/graph?username=xiusin&bg_color=1F222E&color=F8D866&line=F85D7F&point=FFFFFF&hide_border=true&theme=xcode&custom_title=提交日志" /></a> -->
