package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	request_log "github.com/xiusin/pine/middlewares/request-log"
	"github.com/xiusin/pine/middlewares/traceid"

	"github.com/allegro/bigcache/v3"
	pine_bigcache "github.com/xiusin/pine/cache/providers/bigcache"

	"github.com/xiusin/pine/sessions"
	cacheProvider "github.com/xiusin/pine/sessions/providers/cache"
	"github.com/xiusin/pinecms/src/application/controllers/backend/wechat"
	"github.com/xiusin/pinecms/src/application/controllers/middleware/apidoc"
	"github.com/xiusin/pinecms/src/application/plugins"
	logger2 "github.com/xiusin/pinecms/src/common/logger"

	"github.com/gorilla/securecookie"
	"github.com/xiusin/logger"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pine/middlewares/cache304"
	"github.com/xiusin/pine/render/engine/jet"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/controllers/taglibs"
	"github.com/xiusin/pinecms/src/application/controllers/tplfun"

	"github.com/xiusin/pinecms/src/config"
	"github.com/xiusin/pinecms/src/router"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/controllers/middleware"
)

var (
	app *pine.Application

	cacheHandler cache.AbstractCache
	XOrmEngine   *xorm.Engine
	conf         = config.AppConfig()
	dc           = config.DBConfig()
)

func Dc() *config.DbConfig {
	return dc
}

func Ac() *config.Config {
	return conf
}

func initApp() {
	app = pine.New()
	di.Set(controllers.ServiceApplication, func(builder di.AbstractBuilder) (interface{}, error) {
		return app, nil
	}, true)
	diConfig()
	registerV2BackendRoutes()
	var staticPathPrefix []string
	for _, static := range conf.Statics {
		staticPathPrefix = append(staticPathPrefix, static.Route)
	}
	app.Use(cache304.Cache304(30000*time.Second, staticPathPrefix...), middleware.CheckDatabaseBackupDownload())
}

func InitApp() {
	initApp()
}

func InitDB() {
	XOrmEngine = config.InitDB(nil)
	di.Set(controllers.ServiceXorm, func(builder di.AbstractBuilder) (i interface{}, err error) {
		return XOrmEngine, nil
	}, true)
}

func Server() {
	InitCache()
	registerStatic()
	router.InitRouter(app)
	runServe()
}

func registerStatic() {
	for _, static := range conf.Statics {
		app.Static(static.Route, filepath.FromSlash(static.Path), 1)
	}
}

func registerV2BackendRoutes() {

	if config.AppConfig().Debug {

		app.Use(middleware.Demo())
		app.Use(middleware.Cors(app))
		app.Use(request_log.RequestRecorder(time.Millisecond * 200))
	}

	app.Use(traceid.TraceId(), middleware.Pprof(), middleware.SetGlobalConfigData(), apidoc.New(app, nil), middleware.StatesViz(app))

	g := app.Group("/v2", middleware.VerifyJwtToken(), middleware.Casbin(config.InitDB(nil), "resources/configs/rbac_models.conf"))

	router.InitModuleRouter(g, app)

	g.Handle(new(backend.UserController), "/user").
		Handle(new(backend.AdminRoleController), "/role").
		Handle(new(backend.MenuController), "/menu").
		Handle(new(backend.LinkController), "/link").
		Handle(new(backend.LogController), "/log").
		Handle(new(backend.ErrorLogController), "/errlog").
		Handle(new(backend.AssetsManagerController), "/assets").
		Handle(new(backend.AttachmentController), "/attachment").
		Handle(new(backend.AttachmentTypeController), "/attachment/type").
		Handle(new(backend.SettingController), "/setting").
		Handle(new(backend.DictCategoryController), "/dict/category").
		Handle(new(backend.DictController), "/dict").
		Handle(new(backend.DocumentController), "/model").
		Handle(new(backend.CategoryController), "/category").
		Handle(new(backend.DistrictController), "/district").
		Handle(new(backend.AdController), "/ad").
		Handle(new(backend.AdSpaceController), "/ad/space").
		Handle(new(backend.DepartmentController), "/department").
		Handle(new(backend.PositionController), "/position").
		Handle(new(backend.StatController), "/stat").
		Handle(new(backend.PluginController), "/plugin").
		Handle(new(backend.TagsController), "/tags").
		Handle(new(backend.MemberController), "/member").
		Handle(new(backend.MemberGroupController), "/member/group").
		Handle(new(backend.TableController), "/table").
		Handle(new(backend.ContentController), "/content").
		Handle(new(backend.ImSessionController)).
		Handle(new(backend.LoginController)).
		Handle(new(backend.IndexController)).
		Handle(new(backend.DatabaseController)).
		Handle(new(backend.DatabaseBackupController))

	wechat.InitRouter(app, g)

	app.Group("/v2/public").Handle(new(backend.PublicController))
	app.Group("/v2/api").Handle(new(backend.PublicController))

	di.Set(controllers.ServiceBackendRouter, func(builder di.AbstractBuilder) (interface{}, error) {
		return g, nil
	}, true)
}

func runServe() {
	//pine.RegisterErrorCodeHandler(http.StatusInternalServerError, func(ctx *pine.Context) {
	//	if ctx.IsAjax() {
	//		_ = ctx.WriteJSON(pine.H{"code": http.StatusInternalServerError, "message": ctx.Msg})
	//	} else {
	//		ctx.Abort(http.StatusInternalServerError, ctx.Msg)
	//	}
	//})
	go plugins.Init()
	app.Run(
		pine.Addr(fmt.Sprintf(":%d", conf.Port)),
		pine.WithCookieTranscoder(securecookie.New([]byte(conf.HashKey), []byte(conf.BlockKey))),
		//pine.WithoutStartupLog(true),
		pine.WithServerName("xiusin/pinecms"),
		pine.WithCookie(true),
	)
}

func InitCache() {
	cacheHandler = pine_bigcache.New(bigcache.DefaultConfig(24 * time.Hour))

	theme, _ := cacheHandler.Get(controllers.CacheTheme)
	if len(theme) > 0 {
		conf.View.Theme = string(theme)
	}
	di.Set(controllers.ServiceICache, func(builder di.AbstractBuilder) (i interface{}, err error) {
		return cacheHandler, nil
	}, true)

	di.Set(di.ServicePineSessions, func(builder di.AbstractBuilder) (i interface{}, err error) {
		sess := sessions.New(cacheProvider.NewStore(cacheHandler), &sessions.Config{
			CookieName: conf.Session.Name,
			Expires:    conf.Session.Expires,
		})
		return sess, nil
	}, true)
}

func diConfig() {
	di.Set(controllers.ServiceConfig, func(builder di.AbstractBuilder) (i interface{}, e error) {
		return conf, nil
	}, true)

	di.Set(di.ServicePineLogger, func(builder di.AbstractBuilder) (i interface{}, err error) {
		loggers := logger.New()
		writer := logger2.NewPineCmsLogger(XOrmEngine, 10)
		pine.RegisterOnInterrupt(func() {
			writer.Close()
		})
		loggers.SetOutput(io.MultiWriter(os.Stdout, writer))
		logger.SetDefault(loggers)
		loggers.SetReportCaller(true, 3)
		if config.AppConfig().Debug {
			loggers.SetLogLevel(logger.DebugLevel)
		} else {
			loggers.SetLogLevel(logger.ErrorLevel)
		}
		return loggers, nil
	}, false)

	pine.RegisterViewEngine(getJetEngine())

	di.Set(controllers.ServiceTablePrefix, func(builder di.AbstractBuilder) (i interface{}, err error) {
		return dc.Db.DbPrefix, nil
	}, true)
}

func getJetEngine() *jet.PineJet {
	jetEngine := jet.New(conf.View.FeDirname, ".jet", conf.View.Reload)
	jetEngine.AddPath("./resources/taglibs/")

	jetEngine.AddGlobalFunc("flink", taglibs.Flink)
	jetEngine.AddGlobalFunc("type", taglibs.Type)
	jetEngine.AddGlobalFunc("adlist", taglibs.AdList)
	jetEngine.AddGlobalFunc("myad", taglibs.MyAd)
	jetEngine.AddGlobalFunc("channel", taglibs.Channel)
	jetEngine.AddGlobalFunc("channelartlist", taglibs.ChannelArtList)
	jetEngine.AddGlobalFunc("artlist", taglibs.ArcList)
	jetEngine.AddGlobalFunc("pagelist", taglibs.PageList)
	jetEngine.AddGlobalFunc("list", taglibs.List)
	jetEngine.AddGlobalFunc("query", taglibs.Query)
	jetEngine.AddGlobalFunc("likearticle", taglibs.LikeArticle)
	jetEngine.AddGlobalFunc("hotwords", taglibs.HotWords)
	jetEngine.AddGlobalFunc("tags", taglibs.Tags)
	jetEngine.AddGlobalFunc("position", taglibs.Position)
	jetEngine.AddGlobalFunc("toptype", taglibs.TopType)
	jetEngine.AddGlobalFunc("format_time", tplfun.FormatTime)
	jetEngine.AddGlobalFunc("cn_substr", tplfun.CnSubstr)
	jetEngine.AddGlobalFunc("GetDateTimeMK", tplfun.GetDateTimeMK)
	jetEngine.AddGlobalFunc("MyDate", tplfun.MyDate)

	di.Set(controllers.ServiceJetEngine, func(builder di.AbstractBuilder) (i interface{}, err error) {
		return jetEngine, nil
	}, true)

	return jetEngine
}
