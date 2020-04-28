# 控制器
控制器目录(`src/application/controllers`)目前包括如下目录:

- `frontend`（前端）
- `backend`（后端）
- `middleware` (中间件目录)
- `taglibs` (模板标签目录)
- `tplfun` (模板函数目录)


# 控制器创建

> 任何需要控制器都需要内嵌`pine.Controller`结构体, 它会提供一些基础访问方法, 如果您不继承, 框架会抛出异常提醒. 

首先, 控制器创建需要遵循`pine`框架的相关规则: 

```go
type XXXController struct {
	pine.Controller
}

func (c *XXXController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY("/xxx/index", "Index")
}

func (c *XXXController) Index(service1 injectService) {
    
}

func (c *XXXController) Index(service1 injectService) string {
    return ""
}
```

通过如上代码创建一个支持`pine.di`注册服务注入功能的控制器.  这里 `Controller`后缀不是必须的, 只是一般`mvc`框架的惯用手段. 

到现在为止可以支持的`injectService`的服务有:
- *xorm.Engine
- cache.AbstractCache
- logger.AbstractLogger

> 如果控制器方法是包含有参数, 那么pine.di会自动查找相关注册并且注入, 在开发期间您不会得到并没有注册的服务

如果您的控制器方法包含返回值, 那么`pine`框架会自动收集返回值信息(如果您没有渲染过模板框架会自动解析返回值并且返回`json`). 

您也可以选择不使用`RegisterRoute`方法手动指定注册方法, 可以通过使用`请求方法`作为前缀框架会自动注册路由:
```go
type XXXController struct {
	pine.Controller
}

func (c *XXXController) GetIndex(service1 injectService) {
    
}

func (c *XXXController) PostIndex(service1 injectService) string {
    return ""
}
```
上面的代码表示需要注册两个路由`GET`方法的`index`和`POST`方法的`index`, 这种方法注册的路由,会被打印到启动日志里 

您在启动时会得到如下信息: 

![](https://raw.githubusercontent.com/xiusin/assets/master/20200427154826.png)


# 控制器属性

目前为止定义的请勿在控制器中定义属性, 框架每个请求都是重新实例化的结构体, 除非您有特殊需求, 否则不建议使用结构体属性作为操作依赖, 您可以通过使用全局变量的方式来替代.
比如您可能需要统计访问的功能:
```go
var visitCounter int32 // 正确
type VisitController struct {
	pine.Controller
    counter int32 // 错误
}
``` 

# 注册路由

在文件`src/server/server.go`中的`registerBackendRoutes`添加`.Handle(new(XXXController))`, 到此, 框架会自动加载并解析路由. 接下来您就可以编写模块逻辑了.

# 接口模块

如果您需要开发接口模块, 希望您定义到`api`目录下, 如果前端页面是动态渲染的, 希望`.go`作为结束路由. 如: `/search.go` 这样您就知道这个是动态地址, 不会走统一的渲染逻辑.  

# 请求日志
您在开发时需要关注请求日志, `pine`提供了`RequestRecorder`中间件, 您可以在`server.go`中引入它, 以便查看请求信息及耗时:

![](https://raw.githubusercontent.com/xiusin/assets/master/20200427164108.png)

> 如果您已经确定不再开发, 请关闭此中间件