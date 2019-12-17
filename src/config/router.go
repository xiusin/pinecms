package config

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/iris-contrib/middleware/cors"
	jwt2 "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"github.com/xiusin/iriscms/src/application/controllers/api"
	"github.com/xiusin/iriscms/src/application/controllers/backend"
	"github.com/xiusin/iriscms/src/application/controllers/frontend"
	"github.com/xiusin/iriscms/src/application/controllers/middleware"
)

//利用中间件执行控制器前置操作
func registerBackendRoutes() {
	mvcApp.Party(
		config.BackendRouteParty,
		middleware.ViewRequestPath(app, config.LogPath),
		middleware.CheckAdminLoginAndAccess(sess, XOrmEngine),
		middleware.SetGlobalConfigData(XOrmEngine),
		iris.Gzip,
	).Handle(new(backend.AdminController)).
		Handle(new(backend.LoginController)).
		Handle(new(backend.IndexController)).
		Handle(new(backend.CategoryController)).
		Handle(new(backend.ContentController)).
		Handle(new(backend.SettingController)).
		Handle(new(backend.SystemController)).
		Handle(new(backend.MemberController))

	mvcApp.Party(
		"/public",
		middleware.SetGlobalConfigData(XOrmEngine),
		InjectConfig(),
	).Handle(new(backend.PublicController))
}

func registerFrontendRoutes() {
	mvcApp.Party(
		"/",
		middleware.FrontendGlobalViewData(XOrmEngine),
	).Handle(new(frontend.IndexController))
}

func registerErrorRoutes() {
	err := new(backend.ErrorController)
	app.OnErrorCode(iris.StatusInternalServerError, err.ServerError)
	app.OnErrorCode(iris.StatusNotFound, err.StatusNotFound)
}

func registerApiRoutes() {
	apiParty := mvc.New(app.Party("/api/v1", InjectConfig(), cors.AllowAll(), Jwt(), middleware.SetGlobalConfigData(XOrmEngine)).AllowMethods(iris.MethodOptions))
	apiParty.Register(XOrmEngine)
	apiParty.Handle(new(api.UserApiController)).Handle(new(api.WechatController)).Handle(new(api.CategoryController)).Handle(new(api.ContentController))

}

//防止相互调用先用这种不优美的方式实现
func InjectConfig() func(ctx context.Context) {
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
			ErrorHandler: func(c context.Context, err error) {
				c.Header("session_time_out", "timeout")
				c.StatusCode(http.StatusOK)
				_, _ = c.JSON(iris.Map{"Msg": err.Error()})
			},
		}).Serve(ctx)
	}
}
