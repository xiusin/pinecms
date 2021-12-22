package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/xiusin/pinecms/src/common/helper"

	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"gopkg.in/yaml.v2"
)

const dbYml = "resources/configs/database.yml"
const appYml = "resources/configs/application.yml"

type Config struct {
	Debug   bool    `yaml:"debug"`
	Port    int64   `yaml:"port"`
	View    viewConf    `yaml:"view"`
	Session SessConf `yaml:"session"`

	LogPath     string `yaml:"log_path"`
	RuntimePath string `yaml:"runtime_path"`
	PluginPath  string `yaml:"plugin_path"`
	CacheDb     string `yaml:"cache_db"`

	Charset  string `yaml:"charset"`
	JwtKey   string `yaml:"jwtkey"`
	HashKey  string `yaml:"hashkey"`
	BlockKey string `yaml:"blockkey"`

	Upload struct {
		MaxBodySize int64  `yaml:"max_bodysize"`
		Engine      string `yaml:"engine"`
		BasePath    string `yaml:"base_path"`
	} `yaml:"upload"`

	Statics []struct {
		Route string `yaml:"route"`
		Path  string `yaml:"path"`
	} `yaml:"statics"`
}

type SessConf struct {
	Name    string        `yaml:"name"`
	Expires time.Duration `yaml:"expires"`
}

type Html struct {
	Path   string `yaml:"path"`
	Suffix string `yaml:"suffix"`
}

type viewConf struct {
	Reload    bool   `yaml:"reload"`
	FeDirname string `yaml:"fedirname"`
	BeDirname string `yaml:"bedirname"`
	Theme     string
}

var config = &Config{}

func (c *Config) init() {
	if len(c.Upload.BasePath) == 0 {
		c.Upload.BasePath = "uploads"
	}
	if len(c.RuntimePath) == 0 {
		c.RuntimePath = "runtime"
	}
	if len(c.LogPath) == 0 {
		c.LogPath = "logs"
	}
	c.LogPath = RuntimePath(c.LogPath)
	if len(c.CacheDb) == 0 {
		c.CacheDb = "cache.db"
	}
	c.CacheDb = RuntimePath(c.CacheDb)
}

func (c *Config) StaticPrefixArr() []string {
	var staticPathPrefix []string
	for _, static := range config.Statics {
		staticPathPrefix = append(staticPathPrefix, static.Route)
	}
	return staticPathPrefix
}

func IsDebug() bool {
	return config.Debug
}

func RuntimePath(path ...string) string {
	if len(path) == 0 {
		path = append(path, "")
	}
	return filepath.Join(config.RuntimePath, path[0])
}

func App() *Config {
	return config
}

func parseConfig(path string, out interface{}) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	fileContent, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(fileContent, out)
	if err != nil {
		panic(err)
	}
}

func SiteConfig() (map[string]string, error) {
	xorm, cache := helper.Orm(), helper.AbstractCache()
	var settingData = map[string]string{}
	if err := cache.GetWithUnmarshal(controllers.CacheSetting, &settingData); err != nil {
		var settings []tables.Setting
		err := xorm.Find(&settings)
		if err != nil {
			return nil, err
		}
		if len(settings) != 0 {
			for _, v := range settings {
				settingData[strings.ToUpper(v.Key)] = v.Value
			}
		}
		if err = cache.SetWithMarshal(controllers.CacheSetting, &settingData); err != nil {
			return nil, err
		}
	}
	return settingData, nil
}

func init() {
	parseConfig(appYml, config)
	config.init()
	_ = os.MkdirAll(config.LogPath, os.ModePerm)
}
