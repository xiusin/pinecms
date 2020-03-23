package config

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/xiusin/logger"
	"github.com/xiusin/pine/cache/providers/redis"
	"github.com/xiusin/pinecms/src/application/controllers/taglibs"
	"io"
	"os"
	"path/filepath"

	"github.com/gorilla/securecookie"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pine/render/engine/jet"
	"github.com/xiusin/pine/render/engine/template"
	"github.com/xiusin/pine/sessions"
	cacheProvider "github.com/xiusin/pine/sessions/providers/cache"
	"github.com/xiusin/pinecms/src/application/controllers"
	"xorm.io/core"

	"github.com/xiusin/pinecms/src/config"
	"github.com/xiusin/pinecms/src/router"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/controllers/middleware"
	"github.com/xiusin/pinecms/src/common/helper"
	ormlogger "github.com/xiusin/pinecms/src/common/logger"
)

var (
	app *pine.Application

	iCache     cache.ICache
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
	//_orm.SetDefaultCacher(cacher)
	//configs := map[string]string{
	//	"conn": Cfg.RedisAddr,
	//	"key":  "default", // the collection name of redis for cache adapter.
	//}
	//ccStore := cachestore.NewRedisCache(configs)
	//ccStore.Debug = true
	//cacher := xorm.NewLRUCacher(ccStore, 99999999)

	//_orm.Sync(&tables.Advert{}, &tables.AdvertSpace{})
	XOrmEngine = _orm
}

func initApp() {
	app = pine.New()
	diConfig()
	app.Use(middleware.CheckDatabaseBackupDownload())
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
	app.Use(middleware.SetGlobalConfigData(XOrmEngine, iCache))
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

	//todo 两个内嵌缓存库居然都有问题
	//iCache = badger.New(badger.Option{TTL: int(conf.Session.Expires), Path: conf.CacheDb})
	//iCache = bbolt.New(bbolt.Option{
	//	TTL:             int(conf.Session.Expires),
	//	Path:            conf.CacheDb,
	//	Prefix:          "",
	//	CleanupInterval: 0,
	//})

	redisOpt := redis.DefaultOption()
	//redisOpt.Port = 6380
	iCache = redis.New(redisOpt)

	theme, _ := iCache.Get(controllers.CacheTheme)
	if len(theme) == 0 {
		theme = []byte("default")
	}
	conf.View.Theme = string(theme)
	di.Set(controllers.ServiceICache, func(builder di.BuilderInf) (i interface{}, err error) {
		return iCache, nil
	}, true)

	di.Set(controllers.ServiceConfig, func(builder di.BuilderInf) (i interface{}, e error) {
		return conf, nil
	}, true)

	di.Set(di.ServicePineLogger, func(builder di.BuilderInf) (i interface{}, err error) {
		loggers := logger.New()
		loggers.SetReportCaller(true, 3)
		loggers.SetLogLevel(logger.DebugLevel)
		loggers.SetOutput(io.MultiWriter(os.Stdout, &lumberjack.Logger{
			Filename: filepath.Join(conf.LogPath, "pinecms.log"),
			MaxSize:  500,
			Compress: true,
		}))
		return loggers, nil
	}, false)

	di.Set(di.ServicePineSessions, func(builder di.BuilderInf) (i interface{}, err error) {
		sess := sessions.New(cacheProvider.NewStore(iCache), &sessions.Config{
			CookieName: conf.Session.Name,
			Expires:    conf.Session.Expires,
		})
		return sess, nil
	}, true)

	htmlEngine := template.New(conf.View.BeDirname, ".html", conf.View.Reload)
	htmlEngine.AddFunc("GetInMap", controllers.GetInMap)

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

	di.Set(controllers.ServiceJetEngine, func(builder di.BuilderInf) (i interface{}, err error) {
		return jetEngine, nil
	}, true)

	di.Set(XOrmEngine, func(builder di.BuilderInf) (i interface{}, err error) {
		return XOrmEngine, nil
	}, true)

	di.Set(controllers.ServiceTablePrefix, func(builder di.BuilderInf) (i interface{}, err error) {
		return dc.Db.DbPrefix, nil
	}, true)

	app.Use(func(ctx *pine.Context) {
		ctx.Set("cache", iCache)
		ctx.Set("orm", XOrmEngine)
		ctx.Next()
	})
}
