package config

import (
	"iriscms/application/controllers/backend"
	"iriscms/application/controllers/frontend"
	"iriscms/application/controllers/middleware"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

//利用中间件执行控制器前置操作
func registerBackendRoutes() {
	config := BaseMvc(ApplicationConfig) //会话管理器使用同一个 否则无法获取内容
	mvc.New(app).Configure(config).Party(
		ApplicationConfig.BackendRouteParty,
		middleware.ViewRequestPath(app),
		middleware.CheckAdminLoginAndAccess(sess, XOrmEngine)).
		Handle(new(backend.AdminController)).
		Handle(new(backend.LoginController)).
		Handle(new(backend.IndexController)).
		Handle(new(backend.CategoryController)).
		Handle(new(backend.ContentController)).
		Handle(new(backend.SettingController)).
		Handle(new(backend.SystemController))

	mvc.New(app).Configure(config).Party("/public").Handle(new(backend.PublicController))
}

func registerFrontendRoutes() {
	config := BaseMvc(ApplicationConfig)
	mvc.New(app).Configure(config).Party("/").Handle(new(frontend.IndexController))
}

func registerErrorRoutes() {
	err := new(backend.ErrorController)
	app.OnErrorCode(iris.StatusInternalServerError, err.ServerError)
	app.OnErrorCode(iris.StatusNotFound, err.StatusNotFound)
}
