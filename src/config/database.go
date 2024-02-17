package config

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/common/helper"
	"xorm.io/xorm/log"

	"gopkg.in/yaml.v2"
	"xorm.io/core"
	"xorm.io/xorm"
)

type Db struct {
	DbDriver string `yaml:"driver"`
	Dsn      string `yaml:"dsn"`
	DbPrefix string `yaml:"prefix"`
	Conf     dbInfo `yaml:"-" json:"-"`
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

func (t *DbConf) Initialized() bool {
	_, err := os.Stat(dbYml)
	return err == nil
}

func (t *DbConf) buildDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		t.Db.Conf.Username, t.Db.Conf.Password,
		t.Db.Conf.ServeIp, t.Db.Conf.Port, t.Db.Conf.Name)
}

func (t *DbConf) BuildYaml() error {
	t.Db.Dsn = t.buildDsn()
	out, err := yaml.Marshal(t)
	helper.PanicErr(err)
	return os.WriteFile(dbYml, out, os.ModePerm)
}

type orm struct {
	ShowSql      bool  `yaml:"show_sql"`
	ShowExecTime bool  `yaml:"show_exec_time"`
	MaxOpenConns int64 `yaml:"max_open_conns"`
	MaxIdleConns int64 `yaml:"max_idle_conns"`
}

var configure = &DbConf{}

func Orm() *xorm.Engine {
	return helper.GetORM()
}

func InitDB(conf ...*DbConf) *xorm.Engine {
	configure.Do(func() {
		if len(conf) > 0 {
			configure = conf[0]
		} else {
			parseConfig(dbYml, configure)
		}

		_orm, err := xorm.NewEngine(configure.Db.DbDriver, configure.Db.Dsn)
		helper.PanicErr(err)
		_orm.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, configure.Db.DbPrefix))
		_orm.SetLogger(log.NewSimpleLogger(helper.NewOrmLogFile(config.LogPath)))
 		helper.PanicErr(_orm.Ping())

		_orm.ShowSQL(configure.Orm.ShowSql)
		_orm.TZLocation = helper.GetLocation()
		_orm.SetMaxOpenConns(int(configure.Orm.MaxOpenConns))
		_orm.SetMaxIdleConns(int(configure.Orm.MaxIdleConns))
		configure.Engine = _orm
		helper.Inject(controllers.ServiceXorm, _orm)
		helper.Inject(controllers.ServiceTablePrefix, configure.Db.DbPrefix)
	})
	return configure.Engine
}

func DB() *DbConf {
	return configure
}
