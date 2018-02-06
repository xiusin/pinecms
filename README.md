# 介绍 #
iriscms 一个简单的cms框架, 比较简陋, 本人go语言也不是很熟练. 是最开始学习GO的时候的一个即兴代码产品. 结合了easyui开发的后台管理框架.

# 下载部署 #
```
go get -u -v github.com/lazy007/iriscms

glide i

go build main.go

./main.exe

http://domain.com/b/login/index # 登陆后台地址

```

# 默认账号密码 #
```
username : admin
password : admin888
```

# 其他 

现在开发的只有基本框架, 内容编辑当时想的是根据需求添加不同不页面模板扩展. 前端也没写 简单实现


# 相关扩展

- easyui
- github.com/kataras/iris#v10 
- glide0.31.0
- github.com/go-xorm/xorm
- github.com/afocus/captcha






