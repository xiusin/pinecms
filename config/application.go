package config

import (
	"time"
)

type Application struct {
	Port    int64 `yaml:"port"`
	Pprof   Pprof  `yaml:"pprof"`
	View    View `yaml:"view"`
	Session Session  `yaml:"session"`
	LogPath string `yaml:"log_path"`
	Charset string `yaml:"charset"`
	HashKey string `yaml:"hashkey"`
	BlockKey string `yaml:"blockkey"`
}

type Session struct {
	Name    string `yaml:"name"`
	Expires time.Duration `yaml:"expires"`
}
type Engine  struct {
	Django Django `yaml:"django"`
	Html   Html `yaml:"html"`
}

type Django struct {
	Path   string `yaml:"path"`
	Suffix string `yaml:"suffix"`
}

type Html  struct {
	Path   string `yaml:"path"`
	Suffix string `yaml:"suffix"`
}

type Pprof  struct {
	Open  bool `yaml:"open"`
	Port  int64 `yaml:"port"`
	Route string `yaml:"route"`
}

type View  struct {
	Reload bool `yaml:"reload"`
	Engine Engine `yaml:"engine"`
}

