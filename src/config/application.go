package config

import (
	"time"
)

type Config struct {
	SendMail   bool    `yaml:"sendMail"`
	Port       int64   `yaml:"port"`
	Pprof      Pprof   `yaml:"pprof"`
	View       View    `yaml:"view"`
	Session    Session `yaml:"session"`
	LogPath    string  `yaml:"log_path"`
	Statics    []struct { // 注册静态路由
		Route string `yaml:"route"`
		Path  string `yaml:"path"`
	} `yaml:"statics"`
	Charset           string `yaml:"charset"`
	HashKey           string `yaml:"hashkey"`
	BlockKey          string `yaml:"blockkey"`
	BackendRouteParty string `yaml:"backend_route_party"`
	Upload            struct {
		Engine   string `yaml:"engine"`
		BasePath string `yaml:"base_path"`
	} `yaml:"upload"`
	Redis struct {
		Host                 string `yaml:"host"`
		ConnectTimeOut       int    `yaml:"connect_timeout"`
		ReadTimeOut          int    `yaml:"read_timeout"`
		WriteTimeOut         int    `yaml:"write_timeout"`
		CacheDatabaseIndex   int    `yaml:"cache_database_index"`
		SessionDatabaseIndex int    `yaml:"session_database_index"`
		Password             string `yaml:"password"`
		MaxIdle              int    `yaml:"max_idle"`
		MaxActive            int    `yaml:"max_active"`
		IdleTimeout          int    `yaml:"idle_timeout"`
	} `yaml:"redis"`
}

type Session struct {
	Name    string        `yaml:"name"`
	Path    string        `yaml:"path"`
	Expires time.Duration `yaml:"expires"`
}
type Engine struct {
	Django Django `yaml:"django"`
	Html   Html   `yaml:"html"`
}

type Django struct {
	Path   string `yaml:"path"`
	Suffix string `yaml:"suffix"`
}

type Html struct {
	Path   string `yaml:"path"`
	Suffix string `yaml:"suffix"`
}

type Pprof struct {
	Open  bool   `yaml:"open"`
	Port  int64  `yaml:"port"`
	Route string `yaml:"route"`
}

type View struct {
	Reload bool   `yaml:"reload"`
	Engine Engine `yaml:"engine"`
}
