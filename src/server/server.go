package config

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/natefinch/lumberjack"
	"github.com/xiusin/logger"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pine/cache/providers/bbolt"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pine/middlewares/cache304"
	request_log "github.com/xiusin/pine/middlewares/request-log"
	"github.com/xiusin/pine/render/engine/jet"
	"github.com/xiusin/pine/render/engine/template"
	"github.com/xiusin/pine/sessions"
	cacheProvider "github.com/xiusin/pine/sessions/providers/cache"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/controllers/taglibs"
	"io"
	"os"
	"path/filepath"
	"time"
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

	// todo 使用xorm缓存适配器
	XOrmEngine = _orm
}

func initApp() {
	app = pine.New()
	//app.StartPProf()
	diConfig()
	app.Use(request_log.RequestRecorder())
	var staticPathPrefix []string
	for _, static := range conf.Statics {
		staticPathPrefix = append(staticPathPrefix, static.Route)
	}
	app.Use(cache304.Cache304(30000 * time.Second, staticPathPrefix...))
	app.Use(middleware.CheckDatabaseBackupDownload())
	//app.Use(middleware.Demo())

	//pine.RegisterErrorCodeHandler(http.StatusNotFound, func(ctx *pine.Context) {
	//	// 不同状态码可以显示不同的内容
	//	// 自定义页面
	//	ctx.WriteString("404 NotFround")
	//})
}

func Server() {
	initDatabase()
	initApp()
	registerStatic()
	registerBackendRoutes()
	router.InitRouter(app)
	runServe()
}

func registerStatic() {
	for _, static := range conf.Statics {
		app.Static(static.Route, filepath.FromSlash(static.Path))
	}
}

func registerBackendRoutes() {
	app.Use(middleware.SetGlobalConfigData())
	app.Group(
		conf.BackendRouteParty,
		middleware.CheckAdminLoginAndAccess(XOrmEngine, iCache),
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
		Handle(new(backend.AttachmentController)).
		Handle(new(backend.AdController))

	app.Group("/public").Handle(new(backend.PublicController))
}

func runServe() {
	app.Run(
		pine.Addr(fmt.Sprintf(":%d", conf.Port)),
		pine.WithCookieTranscoder(securecookie.New([]byte(conf.HashKey), []byte(conf.BlockKey))),
		pine.WithCharset(conf.Charset),
		pine.WithoutStartupLog(false),
		pine.WithServerName("xiusin/pinecms"),
		pine.WithAutoParseForm(true),
	)
}

func diConfig() {

	//iCache = badger.New(badger.Option{TTL: int(conf.Session.Expires), Path: conf.CacheDb})
	iCache = bbolt.New(bbolt.Option{
		TTL:             int(conf.Session.Expires),
		Path:            conf.CacheDb,
	})
	//redisOpt := redis.DefaultOption()
	//if runtime.GOOS != "darwin" {
	//	redisOpt.Port = 6380
	//}
	//iCache = redis.New(redisOpt)

	theme, _ := iCache.Get(controllers.CacheTheme)
	if len(theme) == 0 {
		theme = []byte("default")
	}
	conf.View.Theme = string(theme)
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
		loggers.SetOutput(io.MultiWriter(os.Stdout, &lumberjack.Logger{
			Filename: filepath.Join(conf.LogPath, "pinecms.log"),
			MaxSize:  5000,
			Compress: true,
		}))
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
	htmlEngine.AddFunc("GetInMap", controllers.GetInMap)
	htmlEngine.AddFunc("in_array", controllers.InStringArr)
	pine.RegisterViewEngine(htmlEngine)

	jetEngine := jet.New(conf.View.FeDirname, ".jet", conf.View.Reload)
	jetEngine.AddPath("./resources/taglibs/")
	jetEngine.AddGlobalFunc("flink", taglibs.Flink)
	jetEngine.AddGlobalFunc("type", taglibs.Type)
	jetEngine.AddGlobalFunc("ad", taglibs.Ad)
	jetEngine.AddGlobalFunc("channel", taglibs.Channel)
	jetEngine.AddGlobalFunc("channelartlist", taglibs.ChannelArtList)
	jetEngine.AddGlobalFunc("artlist", taglibs.ArcList)
	jetEngine.AddGlobalFunc("query", taglibs.Query)

	pine.RegisterViewEngine(jetEngine)

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
