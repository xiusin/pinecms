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

type Config struct {
	Debug   bool    `yaml:"debug"`
	Favicon string  `yaml:"favicon"`
	Port    int64   `yaml:"port"`
	View    View    `yaml:"view"`
	Session Session `yaml:"session"`
	LogPath string  `yaml:"log_path"`
	CacheDb string  `yaml:"cache_db"`
	Statics []struct {
		Route string `yaml:"route"`
		Path  string `yaml:"path"`
	} `yaml:"statics"`
	Charset           string `yaml:"charset"`
	JwtKey            string `yaml:"jwtkey"`
	HashKey           string `yaml:"hashkey"`
	BlockKey          string `yaml:"blockkey"`
	BackendRouteParty string `yaml:"backend_route_party"`
	PluginPath        string `yaml:"plugin_path"`

	Upload struct {
		MaxBodySize int64  `yaml:"max_bodysize"`
		Engine      string `yaml:"engine"`
		BasePath    string `yaml:"base_path"`
	} `yaml:"upload"`
}

type Session struct {
	Name    string        `yaml:"name"`
	Expires time.Duration `yaml:"expires"`
}

type Html struct {
	Path   string `yaml:"path"`
	Suffix string `yaml:"suffix"`
}

type View struct {
	Reload    bool   `yaml:"reload"`
	FeDirname string `yaml:"fedirname"`
	BeDirname string `yaml:"bedirname"`
	Theme     string
}

const dbYml = "resources/configs/database.yml"
const appYml = "resources/configs/application.yml"

var config = &Config{}

func init() {
	parseConfig(appYml, config)
	_ = os.MkdirAll(config.LogPath, os.ModePerm)
}

func AppConfig() *Config {
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
	xorm, cache := helper.XormEngine(), helper.AbstractCache()
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
