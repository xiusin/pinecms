package config

import (
	"fmt"
	"github.com/xiusin/pinecms/src/common/helper"
	ormlogger "github.com/xiusin/pinecms/src/common/logger"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/go-xorm/xorm"
	"gopkg.in/yaml.v2"
	"xorm.io/core"
)

type Db struct {
	DbDriver string `yaml:"db_driver"`
	Dsn      string `yaml:"dsn"`
	DbPrefix string `yaml:"db_prefix"`
	Conf     DbInfo `yaml:"-"`
}

type DbInfo struct {
	ServeIp  string
	Port     string
	Username string
	Password string
	Name     string
}

func (t *DbInfo) Check() bool {
	return (strings.Trim(t.ServeIp, " ") == "" || strings.Contains(t.ServeIp, " ") ||
		strings.Trim(t.Port, " ") == "" || strings.Contains(t.Port, " ") ||
		strings.Trim(t.Username, " ") == "" || strings.Contains(t.Username, " ") ||
		strings.Trim(t.Password, " ") == "" || strings.Contains(t.Password, " ") ||
		strings.Trim(t.Name, " ") == "" || strings.Contains(t.Name, " ")) != true
}

type DbConfig struct {
	*xorm.Engine `yaml:"-"`
	sync.Once    `yaml:"-"`
	Db           Db  `yaml:"db"`
	Orm          Orm `yaml:"orm"`
}

func (t *DbConfig) Inited() bool {
	_, err := os.Stat(dbYml)
	return err == nil
}

func (t *DbConfig) buildDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		t.Db.Conf.Username, t.Db.Conf.Password,
		t.Db.Conf.ServeIp, t.Db.Conf.Port, t.Db.Conf.Name)
}

func (t *DbConfig) CreateYaml() error {
	t.Db.Dsn = t.buildDsn()
	out, err := yaml.Marshal(t)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dbYml, out, os.ModePerm)
}

type Orm struct {
	ShowSql      bool  `yaml:"show_sql"`
	ShowExecTime bool  `yaml:"show_exec_time"`
	MaxOpenConns int64 `yaml:"max_open_conns"`
	MaxIdleConns int64 `yaml:"max_idle_conns"`
}

var dbConfig = &DbConfig{}

func InitDB(conf *DbConfig) *xorm.Engine {
	dbConfig.Do(func() {
		if config != nil {
			dbConfig = conf
		} else {
			parseConfig(dbYml, dbConfig)
		}
		m, o := dbConfig.Db, dbConfig.Orm
		_orm, err := xorm.NewEngine(m.DbDriver, m.Dsn)
		if err != nil {
			panic(err.Error())
		}
		_orm.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, m.DbPrefix))
		_orm.SetLogger(ormlogger.NewIrisCmsXormLogger(helper.NewOrmLogFile(config.LogPath), core.LOG_INFO))
		if err = _orm.Ping(); err != nil {
			panic(err.Error())
		}

		_orm.ShowSQL(o.ShowSql)
		_orm.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
		_orm.ShowExecTime(o.ShowExecTime)
		_orm.SetMaxOpenConns(int(o.MaxOpenConns))
		_orm.SetMaxIdleConns(int(o.MaxIdleConns))
		dbConfig.Engine = _orm
	})
	return dbConfig.Engine
}

func DBConfig() *DbConfig {
	return dbConfig
}
