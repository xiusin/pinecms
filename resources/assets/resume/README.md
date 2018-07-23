# Funday -- 个人简历模板

<a href='http://gitee.com/xiaodan_yu/resume.io/stargazers'><img src='http://gitee.com/xiaodan_yu/resume.io/badge/star.svg?theme=dark' alt='star'></img></a>
<a href='http://gitee.com/xiaodan_yu/resume.io/members'><img src='http://gitee.com/xiaodan_yu/resume.io/badge/fork.svg?theme=dark' alt='fork'></img></a>

## 名字释义

<strong>Funday</strong>，是星期八，开心Day的意思。以前很多玩笑都是说要等星期八才能实现，OK，Then it happened now!

希望每个人在使用了Funday之后在找工作的时候都是处在星期八，开开心心，快快乐乐，轻轻松松拿到心仪的多金的Offer！

## 使用方法

1. 修改`_config.yml`文件中的内容

```
# 个人名称或昵称
name: xiaoxiao
# 页面个人头像信息中地址展示信息
location: 大连
# 页面个人头像信息中公司展示信息
company: IBM
# 页面个人头像信息中职位展示信息
position: Java开发工程师
# 页面个人头像信息中GITHUB展示信息
github: https://github.com/XXXX
# 页面个人头像信息中Facebook展示信息
facebook: https://www.facebook.com/XXXX
# 页面个人头像信息中电话展示信息
phone: 1580424XXXX
# 页面个人头像信息中EMAIL展示信息
email: xxxx@xxx.com

#本项目的baseurl
baseurl: "/resume.io"
```

2. 修改个人头像信息

	修改 `_config.yml` 文件中内容

3. 修改基本信息
 
	修改 `_includes/resumer_01-basic.html` 文件中内容

4. 修改职业技能

    修改 `_includes/resumer_02-profetional.html` 文件中内容

5. 修改教育经历

    修改 `_includes/resumer_03-education.html` 文件中内容

6. 修改工作经历

    修改 `_includes/resumer_04-experience.html` 文件中内容

7. 修改获得证书

	修改 `_includes/resumer_05-certification.html` 文件中内容

8. 修改个人作品

	修改 `_includes/resumer_06-personal_project.html` 文件中内容


## 本地搭建

在本地安装[Jekyll](https://jekyllrb.com/).
然后在项目目录执行`jekyll s`命令,如下

```bash
[root@localhost ~]# jekyll s
Configuration file: C:/..../resume.io/_config.yml
            Source: C:/..../resume.io
       Destination: C:/..../resume.io/_site
 Incremental build: disabled. Enable with --incremental
      Generating...
                    done in 0.371 seconds.
  Please add the following to your Gemfile to avoid polling for changes:
    gem 'wdm', '>= 0.1.0' if Gem.win_platform?
 Auto-regeneration: enabled for 'C:/..../resume.io'
    Server address: http://127.0.0.1:4000/resume.io/
  Server running... press ctrl-c to stop.
```

然后可以访问[http://127.0.0.1:4000/resume.io/](http://127.0.0.1:4000/resume.io/)来访问本地的服务了。


## 在线预览

[非你莫属--个人简历模板](http://xiaodan_yu.gitee.io/resume.io)

## 截图预览

![01.png](http://xiaodan_yu.gitee.io/resume.io/snapshot/11_01.png)

![02.png](http://xiaodan_yu.gitee.io/resume.io/snapshot/11_02.png)

![03.png](http://xiaodan_yu.gitee.io/resume.io/snapshot/11_03.png)

![04.png](http://xiaodan_yu.gitee.io/resume.io/snapshot/11_04.png)


## 参考

本简历模板基于[Certy](http://sc.chinaz.com/moban/170307198220.htm)修改而来。保留了所有的样式，基于Jekyll重构了页面框架，并去掉了一些没有必要的内容，整合一页简历。

## 开源协议
[MIT](https://gitee.com/xiaodan_yu/resume.io/blob/master/LICENSE)
