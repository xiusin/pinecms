package config

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/iris-contrib/middleware/cors"
	jwt2 "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/cache"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"iriscms/application/controllers/api"
	"iriscms/application/controllers/backend"
	"iriscms/application/controllers/frontend"
	"iriscms/application/controllers/middleware"
	"time"
)

//利用中间件执行控制器前置操作
func registerBackendRoutes() {
	config := BaseMvc(ApplicationConfig) //会话管理器使用同一个 否则无法获取内容
	mvc.New(app).Configure(config).Party(
		ApplicationConfig.BackendRouteParty,
		middleware.ViewRequestPath(app),
		middleware.CheckAdminLoginAndAccess(sess, XOrmEngine),
		iris.Gzip).
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
	mvc.New(app).Configure(config).Party("/", cache.Handler(10*time.Second), middleware.FrontendGlobalViewData(app)).Handle(new(frontend.IndexController))
}

func registerErrorRoutes() {
	err := new(backend.ErrorController)
	app.OnErrorCode(iris.StatusInternalServerError, err.ServerError)
	app.OnErrorCode(iris.StatusNotFound, err.StatusNotFound)
}

func registerApiRoutes() {
	//middleToll := tollbooth.NewLimiter(100000, nil) //Api限流
	apiParty := mvc.New(app.Party("/api/v1", cors.AllowAll(), func(ctx context.Context) {
		jwt2.New(jwt2.Config{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte("MySecret"), nil
			},
			SigningMethod:       jwt.SigningMethodHS256,
			CredentialsOptional: true, //如果不传递默认未登录状态即可
		}).Serve(ctx)
	}, /*tollboothic.LimitHandler(middleToll),*/ middleware.FrontendGlobalViewData(app)).AllowMethods(iris.MethodOptions))

	apiParty.Register(XOrmEngine)

	apiParty.Handle(new(api.UserApiController))
}
