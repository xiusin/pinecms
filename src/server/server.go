package config

import (
	"fmt"
	"github.com/xiusin/pine/cache/providers/bitcask"
	"github.com/xiusin/pine/sessions"
	cacheProvider "github.com/xiusin/pine/sessions/providers/cache"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/xiusin/pinecms/src/application/controllers/middleware/apidoc"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/application/plugins"

	"github.com/xiusin/logger"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pine/middlewares/cache304"
	request_log "github.com/xiusin/pine/middlewares/request-log"
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
	di.Set("pine.application", func(builder di.AbstractBuilder) (interface{}, error) {
		return app, nil
	}, true)
	diConfig()
	registerV2BackendRoutes()
	app.Use(request_log.RequestRecorder())
	var staticPathPrefix []string
	for _, static := range conf.Statics {
		staticPathPrefix = append(staticPathPrefix, static.Route)
	}
	app.Use(cache304.Cache304(30000*time.Second, staticPathPrefix...), middleware.CheckDatabaseBackupDownload())
}

func InitApp() {
	initApp()
	InitDB()
	plugins.Init()
}

func InitDB() {
	XOrmEngine = config.InitDB(nil)

	di.Set(XOrmEngine, func(builder di.AbstractBuilder) (i interface{}, err error) {
		return XOrmEngine, nil
	}, true)

	if config.AppConfig().Debug {
		go func() {
			err := XOrmEngine.Sync2(
				&tables.Dict{},
				&tables.DocumentModelDsl{},
				&tables.DocumentModelField{},
				&tables.DictCategory{},
				&tables.Category{},
				&tables.AdminRole{},
				&tables.Member{},
				&tables.MemberGroup{},
				&tables.Attachments{},
				&tables.AttachmentType{},
				&tables.Department{},
				&tables.Position{},
				&tables.Log{},
				&tables.RequestLog{},
				&tables.Plugin{},
				&tables.Tags{})
			if err != nil {
				pine.Logger().Error("同步表结构失败", err)
			}
		}()
	}
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
	app.GET("/admin/", func(ctx *pine.Context) {
		byts, err := ioutil.ReadFile("dist/index.html")
		if err != nil{
			_ = ctx.WriteString(err.Error())
		} else {
			_ = ctx.WriteHTMLBytes(byts)
		}
	})

	if config.AppConfig().Debug {
		app.Use(middleware.Cors())
	}

	app.Use(
		middleware.SetGlobalConfigData(),
		apidoc.New(app, nil),
		middleware.StatesViz(app),
	)

	g := app.Group(
		"/v2",
		middleware.VerifyJwtToken(),
	)
	g.Handle(new(backend.UserController), "/user").
		Handle(new(backend.AdminRoleController), "/role").
		Handle(new(backend.MenuController), "/menu").
		Handle(new(backend.LinkController), "/link").
		Handle(new(backend.LogController), "/log").
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
		Handle(new(backend.LoginController)).
		Handle(new(backend.IndexController)).
		Handle(new(backend.DatabaseController))

	app.Group("/v2/public").Handle(new(backend.PublicController))
	app.Group("/v2/api").Handle(new(backend.PublicController))

	di.Set("pine.backend_router_group", func(builder di.AbstractBuilder) (interface{}, error) {
		return g, nil
	}, true)
}

func runServe() {
	app.SetRecoverHandler(func(ctx *pine.Context) {
		if ctx.IsAjax() {
			_ = ctx.WriteJSON(pine.H{"code": http.StatusInternalServerError, "message": ctx.Msg})
		} else {
			ctx.Abort(http.StatusInternalServerError, ctx.Msg)
		}
	})
	//app.DumpRouteTable()
	app.Run(
		pine.Addr(fmt.Sprintf(":%d", conf.Port)),
		pine.WithCookieTranscoder(securecookie.New([]byte(conf.HashKey), []byte(conf.BlockKey))),
		pine.WithoutStartupLog(false),
		pine.WithServerName("pinecms.xiusin.cn"),
		pine.WithCookie(true),
	)
}

func InitCache() {
	cacheHandler = bitcask.New(int(conf.Session.Expires), conf.CacheDb, time.Minute*10)

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
		loggers.SetOutput()
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

	app.Use(func(ctx *pine.Context) {
		ctx.Set("cache", cacheHandler)
		ctx.Set("orm", XOrmEngine)
		ctx.Next()
	})
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
