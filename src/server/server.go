package config

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/xiusin/iriscms/src/application/controllers"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pine/cache/providers/badger"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pine/middlewares/pprof"
	request_log "github.com/xiusin/pine/middlewares/request-log"
	"github.com/xiusin/pine/render/engine/template"
	"github.com/xiusin/pine/sessions"
	cacheProvider "github.com/xiusin/pine/sessions/providers/cache"
	"net/http"
	"path/filepath"
	"strconv"
	"xorm.io/core"

	"github.com/xiusin/iriscms/src/config"
	"github.com/xiusin/iriscms/src/router"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/iriscms/src/application/controllers/backend"
	"github.com/xiusin/iriscms/src/application/controllers/middleware"
	"github.com/xiusin/iriscms/src/common/helper"
	"github.com/xiusin/iriscms/src/common/logger"
)

var (
	app *pine.Application

	iCache     cache.ICache
	XOrmEngine *xorm.Engine
	conf       = config.AppConfig()
)

func initDatabase() {
	dc := config.DBConfig()
	m, o := dc.Mysql, dc.Orm
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", m.DbUser, m.DbPassword, m.DbServer, m.DbPort, m.DbName, m.DbChatSet)
	_orm, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	_orm.SetLogger(logger.NewIrisCmsXormLogger(helper.NewOrmLogFile(conf.LogPath), core.LOG_INFO))
	err = _orm.Ping() //检测是否联通数据库
	if err != nil {
		panic(err.Error())
	}
	_orm.ShowSQL(o.ShowSql)
	_orm.ShowExecTime(o.ShowExecTime)
	_orm.SetMaxOpenConns(int(o.MaxOpenConns))
	_orm.SetMaxIdleConns(int(o.MaxIdleConns))
	XOrmEngine = _orm
	//syncTable()
}

func initApp() {
	//实例化服务器
	app = pine.New()

	app.Use(request_log.RequestRecorder())
	// todo 使用中间件无法显示内容!!
	//app.Use(debug.Recover(app))

	diConfig()
	//配置前端缓存10秒
	//app.Use(iris.Cache304(10 * time.Second))
	//配置PPROF
	if conf.Pprof.Open {
		p := pprof.New()
		app.GET(conf.Pprof.Route, p)
		app.GET(fmt.Sprintf("%s/*action", conf.Pprof.Route), p)
	}
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
	app.Static(conf.Favicon, "favicon.ico")
	for _, static := range conf.Statics {
		app.Static(static.Route, filepath.FromSlash(static.Path))
	}

}

//利用中间件执行控制器前置操作
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
		Handle(new(backend.LinkController))

	app.Group("/public").Handle(new(backend.PublicController))
}

func runServe() {
	if conf.Pprof.Open {
		go func() {
			pport := strconv.Itoa(int(conf.Pprof.Port))
			err := http.ListenAndServe(":"+pport, nil)
			if err != nil {
				pine.Logger().Error("启动pprof失败", err)
			}
		}()
	}
	app.Run(
		pine.Addr(fmt.Sprintf(":%d", conf.Port)),
		pine.WithCookieTranscoder(securecookie.New([]byte(conf.HashKey), []byte(conf.BlockKey))),
		pine.WithCharset(conf.Charset),
		pine.WithoutStartupLog(false),
		pine.WithServerName("xiusin/pinecms"),
		pine.WithAutoParseForm(true),
	)
}

// 获取mvc配置, 分离相关session
func diConfig() {
	iCache = badger.New(badger.Option{TTL: int(conf.Session.Expires), Path: conf.CacheDb})
	di.Set("cache.ICache", func(builder di.BuilderInf) (i interface{}, err error) {
		return iCache, nil
	}, true)

	di.Set(di.ServicePineSessions, func(builder di.BuilderInf) (i interface{}, err error) {
		sess := sessions.New(cacheProvider.NewStore(iCache), &sessions.Config{
			CookieName: conf.Session.Name,
			Expires:    conf.Session.Expires,
		})
		return sess, nil
	}, true)

	di.Set(di.ServicePineRender, func(builder di.BuilderInf) (i interface{}, err error) {
		view := template.New(conf.View.Path, conf.View.Reload)
		view.AddFunc("GetInMap", controllers.GetInMap)
		return view, nil
	}, true)

	di.Set(XOrmEngine, func(builder di.BuilderInf) (i interface{}, err error) {
		return XOrmEngine, nil
	}, true)

	app.Use(func(ctx *pine.Context) {
		ctx.Set("cache", iCache)
		ctx.Set("orm", XOrmEngine)
		ctx.Next()
	})
}
