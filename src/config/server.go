package config

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/etcd-io/bbolt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/gorilla/securecookie"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/pprof"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/boltdb"
	"github.com/kataras/iris/v12/view"
	"github.com/xiusin/iriscms/src/application/controllers"
	"github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/iriscms/src/common/cache"
	"github.com/xiusin/iriscms/src/common/helper"
	"github.com/xiusin/iriscms/src/common/logger"
	"gopkg.in/yaml.v2"
)

var (
	dbYml           = "resources/configs/database.yml"
	appYml          = "resources/configs/application.yml"
	app             *iris.Application
	mvcApp          *mvc.Application
	XOrmEngine      *xorm.Engine
	sess            *sessions.Sessions
	config          *Config // config 全局配置文件对象
	sessionInitSync sync.Once
	sessCache       *boltdb.Database
	iCache          cache.ICache
)

func syncTable() {
	if err := XOrmEngine.Sync( // 同步表结构
		new(tables.IriscmsAdmin), new(tables.IriscmsAdminRole), new(tables.IriscmsAdminRolePriv),
		new(tables.IriscmsCategory), new(tables.IriscmsCategoryPriv), new(tables.IriscmsContent),
		new(tables.IriscmsLog), new(tables.IriscmsMember), new(tables.IriscmsPage),
		new(tables.IriscmsMenu), new(tables.IriscmsSetting), new(tables.IriscmsWechatMember),
		new(tables.IriscmsWechatMessageLog),
	); err != nil {
		golog.Error("同步表结构失败", err)
	}
}

func initDatabase() {
	var dc DbConfig
	parseConfig(dbYml, &dc)
	m, o := dc.Mysql, dc.Orm
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", m.DbUser, m.DbPassword, m.DbServer, m.DbPort, m.DbName, m.DbChatSet)
	_orm, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	_orm.SetLogger(logger.NewIrisCmsXormLogger(helper.NewOrmLogFile(config.LogPath)))
	err = _orm.Ping() //检测是否联通数据库
	if err != nil {
		panic(err.Error())
	}
	_orm.ShowSQL(o.ShowSql)
	_orm.ShowExecTime(o.ShowExecTime)
	_orm.SetMaxOpenConns(int(o.MaxOpenConns))
	_orm.SetMaxIdleConns(int(o.MaxIdleConns))
	XOrmEngine = _orm
	syncTable()
}

func parseConfig(path string, out interface{}) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err.Error())
	}
	fileContent, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(err.Error())
	}
	err = yaml.Unmarshal(fileContent, out)
	if err != nil {
		panic(err.Error())
	}
}

func initApp() {
	//实例化服务器
	app = iris.New()
	mvcApp = mvc.New(app).Configure(GetMvcConfig())
	//配置前端缓存10秒
	app.Use(iris.Cache304(10 * time.Second))
	//配置PPROF
	if config.Pprof.Open {
		app.Logger().Debug("pprof enabled")
		app.Get(config.Pprof.Route, pprof.New())
	}
	var viewEngine = 0
	engines := config.View.Engine
	//附加视图
	if engines.Django.Path != "" && engines.Django.Suffix != "" {
		viewEngine++
		app.Logger().Debug("注册模板引擎Django")
		app.RegisterView(view.Django(engines.Django.Path, engines.Django.Suffix).Reload(config.View.Reload)) //不缓存模板
	}
	if engines.Html.Path != "" && engines.Html.Suffix != "" {
		app.Logger().Debug("注册模板引擎Html")
		viewEngine++
		app.RegisterView(view.HTML(engines.Html.Path, engines.Html.Suffix, ).Reload(config.View.Reload))
	}
	if viewEngine == 0 {
		app.Logger().Error("请至少配置一个模板引擎")
		panic("请至少配置一个模板引擎")
	}
}

func GetConfig() *Config {
	//解析配置
	if config == nil {
		config = &Config{}
		parseConfig(appYml, config)
	}
	return config
}

func Server() {
	GetConfig()
	//初始化数据库ORM
	initDatabase()
	// 初始化APP
	initApp()
	//注册静态资源路由
	registerStatic()
	//配置异常拦截
	CatchError()
	//注册错误路由
	registerErrorRoutes()
	//注册后端路由
	registerBackendRoutes()
	////注册前端路由
	registerFrontendRoutes()
	//注册API路由
	registerApiRoutes()
	//构建并且运行应用
	runServe()
}

func registerStatic() {
	for _, static := range config.Statics {
		app.HandleDir(static.Route, filepath.FromSlash(static.Path))
		app.Logger().Info("注册静态资源: ", static.Route, "  -->  ", static.Path)
	}

}

func runServe() {
	if config.Pprof.Open {
		go func() {
			pport := strconv.Itoa(int(config.Pprof.Port))
			err := http.ListenAndServe(":"+pport, nil)
			if err != nil {
				app.Logger().Error("启动pprof失败", err)
			}
		}()
	}
	port := strconv.Itoa(int(config.Port))
	_ = app.Run(iris.Addr(":"+port),
		iris.WithCharset(config.Charset),
		iris.WithoutBanner,
		iris.WithOptimizations,
		iris.WithPostMaxMemory(config.Upload.MaxBodySize<<20),
	)
}

// 获取mvc配置, 分离相关session
func GetMvcConfig() func(app *mvc.Application) {
	sessionInitSync.Do(func() {
		hashKey, blockKey := []byte(config.HashKey), []byte(config.BlockKey)
		sec, ssc := securecookie.New(hashKey, blockKey), config.Session
		sess = sessions.New(sessions.Config{Cookie: ssc.Name, Encode: sec.Encode, Decode: sec.Decode, Expires: ssc.Expires * time.Second,})
		db, err := bbolt.Open(ssc.Path, os.FileMode(0750), &bbolt.Options{Timeout: 20 * time.Second})
		if err != nil {
			app.Logger().Error("打开缓存文件失败", err)
			panic(err)
		}
		sessCache, err = boltdb.NewFromDB(db, "sessions")
		if err != nil {
			app.Logger().Error("创建session缓存失败", err)
			panic(err)
		}
		iCache = cache.New(db, string(controllers.WebSiteCacheBucket))
		sess.UseDatabase(sessCache)
		iris.RegisterOnInterrupt(func() {
			if err := sessCache.Close(); err != nil {
				app.Logger().Error("关闭cache失败", err)
			}
		})
	})
	return func(m *mvc.Application) {
		// 注册依赖服务, 内部可以用反射类型方式获取
		m.Register(sess.Start, XOrmEngine, iCache)
	}
}

func getRequestLogs(ctx context.Context) string {
	var status, ip, method, path string
	status = strconv.Itoa(ctx.GetStatusCode())
	path = ctx.Path()
	method = ctx.Method()
	ip = ctx.RemoteAddr()
	return fmt.Sprintf("%v %s %s %s", status, path, method, ip)
}

func CatchError() {
	app.Use(func(ctx context.Context) {
		defer func() {
			if err := recover(); err != nil {
				if ctx.IsStopped() {
					return
				}
				var stacktrace string
				for i := 1; ; i++ {
					_, f, l, got := runtime.Caller(i)
					if !got {
						break
					}
					stacktrace += fmt.Sprintf("%s:%d\n", f, l)
				}
				logMessage := fmt.Sprintf("Recovered from a route's Handler('%s')\n", ctx.HandlerName())
				logMessage += fmt.Sprintf("At Request: %s\n", getRequestLogs(ctx))
				logMessage += fmt.Sprintf("Trace: %s\n", err)
				logMessage += fmt.Sprintf("\n%s", stacktrace)
				app.Logger().Error(logMessage)
				ctx.StatusCode(500)
				ctx.StopExecution()
			}
		}()
		ctx.Next()
	})
}
