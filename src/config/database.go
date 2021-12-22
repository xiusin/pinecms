package config

import (
	"fmt"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/common/helper"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"
	"xorm.io/xorm/log"

	"gopkg.in/yaml.v2"
	"xorm.io/core"
	"xorm.io/xorm"
)

type Db struct {
	DbDriver string `yaml:"driver"`
	Dsn      string `yaml:"dsn"`
	DbPrefix string `yaml:"prefix"`
	Conf     dbInfo `yaml:"-"`
}

type dbInfo struct {
	ServeIp  string
	Port     string
	Username string
	Password string
	Name     string
}

type redisConf struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Password    string `json:"password"`
	Index       int    `json:"index"`
	MaxActive   int    `json:"max_active"`
	MaxIdle     int    `json:"max_idle"`
	IdleTimeout int    `json:"idle_timeout"`
}

func (t *dbInfo) Check() bool {
	return (strings.Trim(t.ServeIp, " ") == "" || strings.Contains(t.ServeIp, " ") ||
		strings.Trim(t.Port, " ") == "" || strings.Contains(t.Port, " ") ||
		strings.Trim(t.Username, " ") == "" || strings.Contains(t.Username, " ") ||
		strings.Trim(t.Password, " ") == "" || strings.Contains(t.Password, " ") ||
		strings.Trim(t.Name, " ") == "" || strings.Contains(t.Name, " ")) != true
}

type DbConf struct {
	*xorm.Engine `yaml:"-"`
	sync.Once    `yaml:"-"`
	Db           Db        `yaml:"db"`
	Orm          orm       `yaml:"orm"`
	Redis        redisConf `yaml:"redis"`
}

func (t *DbConf) Inited() bool {
	_, err := os.Stat(dbYml)
	return err == nil
}

func (t *DbConf) buildDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		t.Db.Conf.Username, t.Db.Conf.Password,
		t.Db.Conf.ServeIp, t.Db.Conf.Port, t.Db.Conf.Name)
}

func (t *DbConf) CreateYaml() error {
	t.Db.Dsn = t.buildDsn()
	out, err := yaml.Marshal(t)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dbYml, out, os.ModePerm)
}

type orm struct {
	ShowSql      bool  `yaml:"show_sql"`
	ShowExecTime bool  `yaml:"show_exec_time"`
	MaxOpenConns int64 `yaml:"max_open_conns"`
	MaxIdleConns int64 `yaml:"max_idle_conns"`
}

var dbConfig = &DbConf{}

func Orm() *xorm.Engine {
	return di.MustGet(controllers.ServiceXorm).(*xorm.Engine)
}

func InitDB(conf ...*DbConf) *xorm.Engine {
	dbConfig.Do(func() {
		if len(conf) > 0 {
			dbConfig = conf[0]
		} else {
			parseConfig(dbYml, dbConfig)
		}
		m, o := dbConfig.Db, dbConfig.Orm
		_orm, err := xorm.NewEngine(m.DbDriver, m.Dsn)
		if err != nil {
			panic(err.Error())
		}
		_orm.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, m.DbPrefix))
		_orm.SetLogger(log.NewSimpleLogger(helper.NewOrmLogFile(config.LogPath)))
		if err = _orm.Ping(); err != nil {
			panic(err.Error())
		}
		_orm.ShowSQL(o.ShowSql)
		_orm.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
		_orm.SetMaxOpenConns(int(o.MaxOpenConns))
		_orm.SetMaxIdleConns(int(o.MaxIdleConns))
		dbConfig.Engine = _orm
		di.Set(controllers.ServiceXorm, func(builder di.AbstractBuilder) (i interface{}, err error) {
			return _orm, nil
		}, true)
		helper.Inject(controllers.ServiceTablePrefix, dbConfig.Db.DbPrefix)
	})
	return dbConfig.Engine
}

func DB() *DbConf {
	return dbConfig
}

func Redis() redisConf {
	return dbConfig.Redis
}
