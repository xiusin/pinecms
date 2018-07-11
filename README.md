# 介绍 #
iriscms 一个简单的cms框架,基础的cms管理功能.
开发不实现相关的前端功能



# 下载部署 #
```
go get -u -v github.com/lazy007/iriscms
glide i
go build main.go
./main.exe or ./main
```

# 访问 http://domain.com/b/login/index

后台管理相关的链接统统以`/b/`为前缀

# 数据库配置 #

修改`resources/configs/database.yml.dist`为`resources/configs/database.yml`, 配置自己的数据库


# 默认账号密码 #
```
username : admin
password : admin888
```

# 其他 #

现在开发的只有基本框架, 内容编辑当时想的是根据需求添加不同不页面模板扩展. 前端也没写 简单实现


# 相关扩展 #

- easyui
- github.com/kataras/iris#v10 
- github.com/go-xorm/xorm
- github.com/afocus/captcha




