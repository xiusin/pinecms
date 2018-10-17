package config

type DatabaseMysql struct {
	DbName     string `yaml:"db_name"`
	DbChatSet  string `yaml:"db_charset"`
	DbUser     string `yaml:"db_user"`
	DbPassword string `yaml:"db_pass"`
	DbServer   string `yaml:"db_server"`
	DbPort     string `yaml:"db_port"`
}

type DatabaseRedis struct {
}

type DatabaseMemcached struct {
}

type DatabaseConfig struct {
	Mysql     DatabaseMysql     `yaml:"mysql"`
	Redis     DatabaseRedis     `yaml:"redis"`
	Memcached DatabaseMemcached `yaml:"memcached"`
	Orm       Orm               `yaml:"orm"`
}

type Orm struct {
	ShowSql      bool  `yaml:"show_sql"`
	ShowExecTime bool  `yaml:"show_exec_time"`
	MaxOpenConns int64 `yaml:"max_open_conns"`
	MaxIdleConns int64 `yaml:"max_idle_conns"`
}
