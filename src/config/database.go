package config

type Db struct {
	DbDriver string `yaml:"db_driver"`
	Dsn      string `yaml:"dsn"`
	DbPrefix string `yaml:"db_prefix"`
}

type DbConfig struct {
	Db  Db  `yaml:"db"`
	Orm Orm `yaml:"orm"`
}

type Orm struct {
	ShowSql      bool  `yaml:"show_sql"`
	ShowExecTime bool  `yaml:"show_exec_time"`
	MaxOpenConns int64 `yaml:"max_open_conns"`
	MaxIdleConns int64 `yaml:"max_idle_conns"`
}

var dbConfig *DbConfig

func init() {
	dbConfig = &DbConfig{}
	parseConfig(dbYml, dbConfig)
}

func DBConfig() *DbConfig {
	return dbConfig
}
