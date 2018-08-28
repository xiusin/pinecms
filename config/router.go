package config

import (
	"iriscms/application/controllers/backend"
	"iriscms/application/controllers/frontend"
	"iriscms/application/controllers/middleware"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/context"
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"

	"github.com/iris-contrib/middleware/cors"
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
	mvc.New(app).Configure(config).Party("/", middleware.FrontendGlobalViewData(app)).Handle(new(frontend.IndexController))

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"*"},
		AllowCredentials: true,
	})

	frontendApp := mvc.New(app).Configure(config).Party("/api", crs, func(ctx context.Context) {
		jwtmiddleware.New(jwtmiddleware.Config{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte("MySecret"), nil
			},
			SigningMethod:       jwt.SigningMethodHS256,
			CredentialsOptional: true, //如果不传递默认未登录状态即可
		}).Serve(ctx)
	}, middleware.FrontendGlobalViewData(app)).Handle(new(frontend.ApiController))
	frontendApp.Router.AllowMethods(iris.MethodOptions)

}

func registerErrorRoutes() {
	err := new(backend.ErrorController)
	app.OnErrorCode(iris.StatusInternalServerError, err.ServerError)
	app.OnErrorCode(iris.StatusNotFound, err.StatusNotFound)
}
