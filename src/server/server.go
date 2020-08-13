package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/xiusin/pine/cache/providers/bbolt"

	"github.com/xiusin/logger"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pine/middlewares/cache304"
	request_log "github.com/xiusin/pine/middlewares/request-log"
	"github.com/xiusin/pine/render/engine/jet"
	"github.com/xiusin/pine/render/engine/template"
	"github.com/xiusin/pine/sessions"
	cacheProvider "github.com/xiusin/pine/sessions/providers/cache"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/controllers/taglibs"
	"github.com/xiusin/pinecms/src/application/controllers/tplfun"
	"xorm.io/core"

	"github.com/xiusin/pinecms/src/config"
	"github.com/xiusin/pinecms/src/router"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/controllers/middleware"
	"github.com/xiusin/pinecms/src/common/helper"
	ormlogger "github.com/xiusin/pinecms/src/common/logger"
)

var (
	app *pine.Application

	iCache     cache.AbstractCache
	XOrmEngine *xorm.Engine
	conf       = config.AppConfig()
	dc         = config.DBConfig()
)

func Dc() *config.DbConfig {
	return dc
}

func initDatabase() {
	m, o := dc.Db, dc.Orm
	_orm, err := xorm.NewEngine(m.DbDriver, m.Dsn)
	if err != nil {
		panic(err.Error())
	}
	_orm.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, dc.Db.DbPrefix))
	_orm.SetLogger(ormlogger.NewIrisCmsXormLogger(helper.NewOrmLogFile(conf.LogPath), core.LOG_INFO))
	if err = _orm.Ping(); err != nil {
		panic(err.Error())
	}

	_orm.ShowSQL(o.ShowSql)
	_orm.ShowExecTime(o.ShowExecTime)
	_orm.SetMaxOpenConns(int(o.MaxOpenConns))
	_orm.SetMaxIdleConns(int(o.MaxIdleConns))
	XOrmEngine = _orm
}

func initApp() {
	app = pine.New()
	diConfig()
	//app.Use(middleware.Demo())
	app.Use(request_log.RequestRecorder())
	var staticPathPrefix []string
	for _, static := range conf.Statics {
		staticPathPrefix = append(staticPathPrefix, static.Route)
	}
	app.Use(cache304.Cache304(30000*time.Second, staticPathPrefix...))
	app.Use(middleware.CheckDatabaseBackupDownload())
}

func Bootstrap()  {
	initDatabase()
	initApp()
}

func Server() {
	Bootstrap()
	registerStatic()
	registerV2BackendRoutes()
	router.InitRouter(app)
	runServe()
}

func registerStatic() {
	for _, static := range conf.Statics {
		app.Static(static.Route, filepath.FromSlash(static.Path))
	}
}

func registerV2BackendRoutes() {
	app.Use(middleware.SetGlobalConfigData())
	app.Use(func(ctx *pine.Context) {
		ctx.Response.Header.Add("Vary", "Origin")
		ctx.Response.Header.Add("Vary", "Access-Control-Allow-Methods")
		ctx.Response.Header.Add("Vary", "Access-Control-Allow-Headers")
		ctx.Response.Header.Add("Vary", "Access-Control-Allow-Credentials")
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://localhost:7050")
		//ctx.Response.Header.Set("Access-Control-Allow-Headers", "X-TOKEN, Content-Type, Origin, Referer, Content-Length, Access-Control-Allow-Headers")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "*")
		if !ctx.IsOptions() {
			ctx.Next()
		}
	})
	// 解析参数
	app.Group(
		"/v2",
		middleware.VerifyJwtToken(),
	).Handle(new(backend.AdminController)).
		Handle(new(backend.LoginController)).
		Handle(new(backend.IndexController)).
		Handle(new(backend.CategoryController)).
		Handle(new(backend.ContentController)).
		Handle(new(backend.SettingController)).
		Handle(new(backend.SystemController)).
		Handle(new(backend.MemberController)).
		Handle(new(backend.DocumentController)).
		Handle(new(backend.LinkController)).
		Handle(new(backend.DatabaseController)).
		Handle(new(backend.AssetsManagerController)).
		Handle(new(backend.AdController)).
		Handle(new(backend.StatController))

	app.Group("/v2/public").Handle(new(backend.PublicController))
}

func runServe() {
	app.Run(
		pine.Addr(fmt.Sprintf(":%d", conf.Port)),
		pine.WithCookieTranscoder(securecookie.New([]byte(conf.HashKey), []byte(conf.BlockKey))), // 关闭加密cookie
		pine.WithoutStartupLog(false),
		pine.WithServerName("xiusin/pinecms"),
		pine.WithCookie(true),
	)
}

func diConfig() {
	iCache = bbolt.New(&bbolt.Option{
		TTL:  int(conf.Session.Expires),
		Path: conf.CacheDb,
	})

	//badger.New(int(conf.Session.Expires), conf.CacheDb)

	theme, _ := iCache.Get(controllers.CacheTheme)
	if len(theme) > 0 {
		conf.View.Theme = string(theme)
	}
	di.Set(controllers.ServiceICache, func(builder di.AbstractBuilder) (i interface{}, err error) {
		return iCache, nil
	}, true)

	di.Set(controllers.ServiceConfig, func(builder di.AbstractBuilder) (i interface{}, e error) {
		return conf, nil
	}, true)

	di.Set(di.ServicePineLogger, func(builder di.AbstractBuilder) (i interface{}, err error) {
		loggers := logger.New()
		loggers.SetReportCaller(true, 3)
		loggers.SetLogLevel(logger.DebugLevel)
		loggers.SetOutput(io.MultiWriter(os.Stdout))
		return loggers, nil
	}, false)

	di.Set(di.ServicePineSessions, func(builder di.AbstractBuilder) (i interface{}, err error) {
		sess := sessions.New(cacheProvider.NewStore(iCache), &sessions.Config{
			CookieName: conf.Session.Name,
			Expires:    conf.Session.Expires,
		})
		return sess, nil
	}, true)

	htmlEngine := template.New(conf.View.BeDirname, ".html", conf.View.Reload)
	htmlEngine.AddFunc("in_array", controllers.InStringArr)
	pine.RegisterViewEngine(htmlEngine)

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

	pine.RegisterViewEngine(jetEngine)

	//app.SetNotFound(func(ctx *pine.Context) {
	//	conf := di.MustGet(controllers.ServiceConfig).(*config.Config)
	//	path := filepath.Join(conf.View.Theme, "error.jet")
	//	if runtime.GOOS == "windows" {
	//		path = strings.ReplaceAll(path, "\\", "/")
	//	}
	//	ctx.Render().ViewData("msg", ctx.Msg)
	//	ctx.Render().HTML(path)
	//})

	di.Set(controllers.ServiceJetEngine, func(builder di.AbstractBuilder) (i interface{}, err error) {
		return jetEngine, nil
	}, true)

	di.Set(XOrmEngine, func(builder di.AbstractBuilder) (i interface{}, err error) {
		return XOrmEngine, nil
	}, true)

	di.Set(controllers.ServiceTablePrefix, func(builder di.AbstractBuilder) (i interface{}, err error) {
		return dc.Db.DbPrefix, nil
	}, true)

	app.Use(func(ctx *pine.Context) {
		ctx.Set("cache", iCache)
		ctx.Set("orm", XOrmEngine)
		ctx.Next()
	})
}
