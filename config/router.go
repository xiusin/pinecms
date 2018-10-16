package config

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/iris-contrib/middleware/cors"
	jwt2 "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"iriscms/application/controllers/api"
	"iriscms/application/controllers/backend"
	"iriscms/application/controllers/frontend"
	"iriscms/application/controllers/middleware"
	"net/http"
)

//利用中间件执行控制器前置操作
func registerBackendRoutes() {
	config := BaseMvc(ApplicationConfig) //会话管理器使用同一个 否则无法获取内容
	mvc.New(app).Configure(config).Party(
		ApplicationConfig.BackendRouteParty,
		middleware.ViewRequestPath(app),
		middleware.CheckAdminLoginAndAccess(sess, XOrmEngine),
		middleware.SetGlobalConfigData(XOrmEngine, redisPool),
		iris.Gzip,
	).Handle(new(backend.AdminController)).
		Handle(new(backend.LoginController)).
		Handle(new(backend.IndexController)).
		Handle(new(backend.CategoryController)).
		Handle(new(backend.ContentController)).
		Handle(new(backend.SettingController)).
		Handle(new(backend.WechatController)).
		Handle(new(backend.SystemController)).Handle(new(backend.MemberController))


	mvc.New(app).Configure(config).Party(
		"/public",
		middleware.SetGlobalConfigData(XOrmEngine, redisPool),
		InjectConfig(ApplicationConfig),
	).Handle(new(backend.PublicController))
}

func registerFrontendRoutes() {
	config := BaseMvc(ApplicationConfig)
	mvc.New(app).Configure(config).Party(
		"/",
		middleware.FrontendGlobalViewData(app),
	).Handle(new(frontend.IndexController))
}

func registerErrorRoutes() {
	err := new(backend.ErrorController)
	app.OnErrorCode(iris.StatusInternalServerError, err.ServerError)
	app.OnErrorCode(iris.StatusNotFound, err.StatusNotFound)
}

func registerApiRoutes() {
	apiParty := mvc.New(app.Party(
		"/api/v1",
		InjectConfig(ApplicationConfig),
		cors.AllowAll(),
		Jwt(),
		middleware.FrontendGlobalViewData(app),
		middleware.SetGlobalConfigData(XOrmEngine, redisPool),
	).AllowMethods(iris.MethodOptions))
	apiParty.Register(XOrmEngine, redisPool)
	apiParty.Handle(new(api.UserApiController)).Handle(new(api.WechatController)).Handle(new(api.CategoryController)).Handle(new(api.ContentController))

}

//防止相互调用先用这种不优美的方式实现
func InjectConfig(config *Application) func(ctx context.Context) {
	return func(ctx context.Context) {
		ctx.Values().Set("app.config", map[string]string{
			"uploadEngine": config.Upload.Engine,
		})
		ctx.Next()
	}
}

func Jwt() func(ctx context.Context) {
	return func(ctx context.Context) {
		jwt2.New(jwt2.Config{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte("MySecret"), nil
			},
			SigningMethod:       jwt.SigningMethodHS256,
			CredentialsOptional: true, //如果不传递默认未登录状态即可
			ErrorHandler: func(i context.Context, s string) {
				i.Header("session_time_out","timeout")
				i.StatusCode(http.StatusOK)
				i.JSON(iris.Map{})
			},
		}).Serve(ctx)
	}
}
