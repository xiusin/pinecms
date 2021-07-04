package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
	"sync"
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
	sync.Once `yaml:"-"`
	Db        Db  `yaml:"db"`
	Orm       Orm `yaml:"orm"`
}

func (t *DbConfig) buildDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		t.Db.Conf.Username, t.Db.Conf.Password,
		t.Db.Conf.ServeIp, t.Db.Conf.Port, t.Db.Conf.Name)
}

func (t *DbConfig) CreateYaml() {
	t.Db.Dsn = t.buildDsn()
	out, _ := yaml.Marshal(t)
	err := ioutil.WriteFile(dbYml, out, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

type Orm struct {
	ShowSql      bool  `yaml:"show_sql"`
	ShowExecTime bool  `yaml:"show_exec_time"`
	MaxOpenConns int64 `yaml:"max_open_conns"`
	MaxIdleConns int64 `yaml:"max_idle_conns"`
}

var dbConfig = &DbConfig{}

func InitDB() {
	dbConfig.Do(func() {
		parseConfig(dbYml, dbConfig)
	})
}

func DBConfig() *DbConfig {
	return dbConfig
}
