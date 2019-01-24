# vblog-web 博客前端

项目预览地址：https://blog.xiaoxinfq.com/

#### 项目缘由
- 转眼硕士毕业工作已有三年多时间了，一直以来都想搭建一个自己的博客系统，分享自己的工作经历、总结工作中用到的技术、积累的工作经验技能等；
- 在网上找资料查了好久，完整基于Java语言开发的并不多，少有的大部分还是基于jsp、freemarker、velocity模板引擎这种开发的，我个人实在不喜欢这种前后端耦合在一起开发的方式，而且工作中一直都是前后端分离的开发方式；
- 由于工作中用过 vue+element-ui 做过管理系统，故而在今年六月份终于下决心搞一套自己的、基于当前最新技术栈、前后端分离的博客系统。

#### 项目介绍
- VBlog 是一款基于最新技术开发的多人在线、简洁的博客系统；
- vblog-web是该博客系统的前端页面代码；
- vblog-api是该博客系统的后端API接口服务，代码详见https://gitee.com/seu-lfh/vblog.git。

#### 技术架构
- 采用Vue2.5、ElementUI、mavon-editor、vuex、axios等框架


#### 项目特点
- 友好的代码结构及注释，便于阅读及二次开发
- 实现前后端分离，通过token进行数据交互，前端再也不用关注后端技术
- 页面交互使用Vue2.x，极大的提高了开发效率
- 引入quartz定时任务，可动态完成任务的添加、修改、删除、暂停、恢复及日志查看等功能
- 引入Hibernate Validator校验框架，轻松实现后端校验
- 引入swagger文档支持，方便编写API接口文档
<br>

#### 使用说明
1. npm install；
2. npm run dev；


#### 参与贡献

1. Fork 本项目
2. 新建 Feat_xxx 分支
3. 提交代码
4. 新建 Pull Request


