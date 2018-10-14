package config

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
	"io/ioutil"
	"iriscms/common/helper"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql" // 初始化 Mysql 驱动
	"github.com/go-xorm/xorm"
	"github.com/gorilla/securecookie"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/pprof"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	redisSessionDB "github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/view"
	"gopkg.in/yaml.v2"
)

var app *iris.Application
var XOrmEngine *xorm.Engine // XOrmEngine 全局 Xorm 引擎对象
var databaseYml = "resources/configs/database.yml"
var applicationYml = "resources/configs/application.yml"
var sess *sessions.Sessions
var ApplicationConfig *Application // ApplicationConfig 全局配置文件对象
var sessionInitSync sync.Once
var redisPool *redis.Pool

func initDatabase() {
	dbconfig := new(DatabaseConfig)
	//解析数据库配置项
	parseConfig(databaseYml, dbconfig)
	dsn := dbconfig.Mysql.DbUser + ":" +
		dbconfig.Mysql.DbPassword + "@tcp(" +
		dbconfig.Mysql.DbServer + ":" +
		dbconfig.Mysql.DbPort + ")/" +
		dbconfig.Mysql.DbName + "?charset=" +
		dbconfig.Mysql.DbChatSet
	_orm, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	err = _orm.Ping() //检测是否联通数据库
	if err != nil {
		panic(err.Error())
	}
	XOrmEngine = _orm
	//XOrmEngine.Logger().SetLevel(core.LOG_DEBUG | core.LOG_ERR | core.LOG_WARNING)
	XOrmEngine.ShowSQL(dbconfig.Orm.ShowSql)
	XOrmEngine.ShowExecTime(dbconfig.Orm.ShowExecTime)
	XOrmEngine.SetMaxOpenConns(int(dbconfig.Orm.MaxOpenConns))
	XOrmEngine.SetMaxIdleConns(int(dbconfig.Orm.MaxIdleConns))
}

func initRedis() {
	redisPool = &redis.Pool{
		MaxIdle:     ApplicationConfig.Redis.MaxIdle,
		MaxActive:   ApplicationConfig.Redis.MaxActive,
		IdleTimeout: time.Duration(ApplicationConfig.Redis.IdleTimeout) * time.Second, //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ApplicationConfig.Redis.Host,
				redis.DialConnectTimeout(time.Second*time.Duration(ApplicationConfig.Redis.ConnectTimeOut)),
				redis.DialReadTimeout(time.Second*time.Duration(ApplicationConfig.Redis.ReadTimeOut)),
				redis.DialWriteTimeout(time.Second*time.Duration(ApplicationConfig.Redis.WriteTimeOut)),
				redis.DialDatabase(ApplicationConfig.Redis.CacheDatabaseIndex),
				redis.DialPassword(ApplicationConfig.Redis.Password),
			)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
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

func StartApplication() {
	//初始化数据库ORM
	initDatabase()
	ApplicationConfig = new(Application)
	parseConfig(applicationYml, ApplicationConfig) //解析配置
	initRedis()
	//实例化服务器
	app = iris.New()
	//配置前端缓存10秒
	app.Use(iris.Cache304(10 * time.Second))
	//配置PPROF
	if ApplicationConfig.Pprof.Open {
		app.Get(ApplicationConfig.Pprof.Route, pprof.New())
	}
	var viewEngine int64 = 0
	//附加视图
	if ApplicationConfig.View.Engine.Django.Path != "" && ApplicationConfig.View.Engine.Django.Suffix != "" {
		viewEngine = viewEngine + 1
		app.RegisterView(
			view.Django(
				ApplicationConfig.View.Engine.Django.Path,
				ApplicationConfig.View.Engine.Django.Suffix,
			).Reload(ApplicationConfig.View.Reload),
		) //不缓存模板
	}
	if ApplicationConfig.View.Engine.Html.Path != "" && ApplicationConfig.View.Engine.Html.Suffix != "" {
		viewEngine = viewEngine + 1
		app.RegisterView(view.HTML(
			ApplicationConfig.View.Engine.Html.Path,
			ApplicationConfig.View.Engine.Html.Suffix,
		).Reload(ApplicationConfig.View.Reload))

	}
	if viewEngine == 0 {
		panic("请至少配置一个模板引擎")
	}
	//注册静态资源路由
	registerStatic()
	//配置异常拦截
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
				ctx.Application().Logger().Warn(logMessage)
				go helper.SendEmail("系统发生异常", logMessage, []string{"826466266@qq.com"}, ctx.Values().Get("setting").(map[string]string))
				ctx.StatusCode(500)
				ctx.StopExecution()
			}
		}()
		ctx.Next()
	})
	//日志
	app.Use(logger.New())
	//注册错误路由
	registerErrorRoutes()
	//注册后端路由
	registerBackendRoutes()
	////注册前端路由
	registerFrontendRoutes()
	//注册API路由
	registerApiRoutes()
	//构建并且运行应用
	runServe(ApplicationConfig)
}

func registerStatic() {
	app.StaticWeb("/upload", filepath.FromSlash("./resources/assets/upload"))
	app.StaticWeb("/assets", filepath.FromSlash("./resources/assets"))
	app.StaticWeb("/frontend", filepath.FromSlash("./resources/assets/frontend"))
	app.StaticWeb("/backend", filepath.FromSlash("./resources/assets/backend"))
	app.StaticWeb("/resume", filepath.FromSlash("./resources/assets/resume"))
}

func runServe(config *Application) {
	if config.Pprof.Open {
		go func() {
			pport := strconv.Itoa(int(config.Pprof.Port))
			err := http.ListenAndServe("0.0.0.0:"+pport, nil)
			if err != nil {
				log.Println(err.Error())
			}
		}()
	}
	port := strconv.Itoa(int(config.Port))
	app.Run(iris.Addr(":"+port), iris.WithCharset(config.Charset))
}

// BaseMvc 构造 mvc基础,分离相关session
func BaseMvc(config *Application) func(app *mvc.Application) {
	sessionInitSync.Do(func() {
		hashKey := []byte(config.HashKey)
		blockKey := []byte(config.BlockKey)
		secureCookie := securecookie.New(hashKey, blockKey)
		sess = sessions.New(sessions.Config{
			Cookie:  config.Session.Name,
			Encode:  secureCookie.Encode,
			Decode:  secureCookie.Decode,
			Expires: config.Session.Expires * time.Second,
		})
		db := redisSessionDB.New(service.Config{
			Network:     "tcp",
			Addr:        ApplicationConfig.Redis.Host,
			Password:    ApplicationConfig.Redis.Password,
			Database:    strconv.Itoa(ApplicationConfig.Redis.SessionDatabaseIndex),
			MaxIdle:     ApplicationConfig.Redis.MaxIdle,
			MaxActive:   ApplicationConfig.Redis.MaxActive,
			IdleTimeout: time.Duration(ApplicationConfig.Redis.IdleTimeout) * time.Second,
		})

		iris.RegisterOnInterrupt(func() { //关机时候关闭数据
			db.Close()
			redisPool.Close()
		})

		sess.UseDatabase(db)
	})
	return func(app *mvc.Application) {
		app.Register(sess.Start, XOrmEngine, redisPool)
	}
}

func getRequestLogs(ctx context.Context) string {
	var status, ip, method, path string
	status = strconv.Itoa(ctx.GetStatusCode())
	path = ctx.Path()
	method = ctx.Method()
	ip = ctx.RemoteAddr()
	// the date should be logged by iris' Logger, so we skip them
	return fmt.Sprintf("%v %s %s %s", status, path, method, ip)
}