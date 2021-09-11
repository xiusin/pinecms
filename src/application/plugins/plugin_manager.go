package plugins

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	plugin2 "github.com/xiusin/pinecms/cmd/plugin"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/config"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"plugin"
	"runtime"
	"strings"
	"sync"
	"time"
)

//https://github.com/GoAdminGroup/go-admin/blob/master/plugins/plugins.go
// TODO 废除JSON配置读取, 命令行注入不在此提供, 使用api处理
// 		允许注入到task任务管理, 使用脚本处理

const ext = ".so"

const exportedVarSuffix = "Plugin"

const jsonName = "config.json"

type PluginIntf interface {
	Init(di.AbstractBuilder) // 初始化插件
	Prefix(string)           // 路由前缀, 注册后修改需要重启程序
	Sign() string            // 插件的唯一标识, 需要开发者搞一个独一无二如 uuid
	View() string            // 配置视图json信息
	Menu(interface{}, int)   // 安装插件位置
	Install()                // 安装插件, 首次扫描后执行.
	Uninstall()              // 卸载后禁止访问
	Upgrade()                // 更新插件
}

type Plug struct {
	plug *plugin.Plugin
	pi   PluginIntf
}

var pluginMgr = &pluginManager{
	plugins:        map[string]*Plug{},
	installPlugins: map[string]struct{}{},
	scannedPlugins: map[string]*plugin2.Config{},
	remoteDomain:   "https://plugin.xiusin.cn",
	fileGlob:       "*/**" + ext,
}

// todo 可以打包子目录进行前端代码发布, 或后端源代码
// 插件端可以配置导出assets.zip
type pluginManager struct {
	sync.Mutex
	plugins        map[string]*Plug
	installPlugins map[string]struct{}        // 已经通过安装配置的插件
	scannedPlugins map[string]*plugin2.Config // 已经扫描到的插件(已安装 + 未安装)
	path           string
	fileGlob       string
	remoteDomain   string
}

//IterFn 迭代函数
type IterFn func(string, PluginIntf) error

func (p *pluginManager) Iter(fn IterFn) error {
	p.Lock()
	defer p.Unlock()
	for s, plug := range p.plugins {
		if err := fn(s, plug.pi); err != nil {
			return err
		}
	}
	return nil
}

func (p *pluginManager) GetLocalPlugin() map[string]plugin2.Config {
	p.Lock()
	defer p.Unlock()
	var scannedPlugins = map[string]plugin2.Config{}
	for s, s2 := range p.scannedPlugins { // 复制一份
		scannedPlugins[s] = *s2
	}
	return scannedPlugins
}

func (p *pluginManager) GetPluginInfo(name string) (*plugin2.Config, error) {
	conf, exist := p.scannedPlugins[name]
	if !exist {
		return conf, errors.New("插件不存在")
	}
	return conf, nil
}

func (p *pluginManager) loadPlugin() {
	orm := di.MustGet(&xorm.Engine{}).(*xorm.Engine)
	var ps []tables.Plugin
	orm.Where("enable = ?", 1).Find(&ps)
	for _, plug := range ps {
		if _, err := p.Install(plug.Path); err != nil {
			pine.Logger().Printf("安装插件%s失败: %s", plug.Path, err)
		} else {
			pine.Logger().Print("安装插件成功 ", plug.Path)
		}
	}
}

func (p *pluginManager) Reload() {
	p.Lock()
	defer p.Unlock()
	var pluginNames []string

	for name := range p.installPlugins {
		pluginNames = append(pluginNames, name)
	}
	_ = os.MkdirAll(p.path, os.ModePerm)
	jsonPath := filepath.Join(p.path, "plugins.json")
	conf, _ := json.Marshal(&pluginNames)

	ioutil.WriteFile(jsonPath, conf, os.ModePerm)
}

func (p *pluginManager) Install(filename string) (PluginIntf, error) {
	p.Lock()
	defer p.Unlock()
	name := helper.UcFirst(strings.TrimSuffix(filepath.Base(filename), ext)) + exportedVarSuffix
	if _, exist := p.scannedPlugins[filename]; !exist {
		return nil, fmt.Errorf("插件%s不存在", filename)
	}
	plug, err := plugin.Open(filename)
	if err != nil {
		delete(pluginMgr.scannedPlugins, filename)
		return nil, err
	}
	pluginIntf, err := plug.Lookup(name)
	if err != nil {
		delete(pluginMgr.scannedPlugins, filename)
		return nil, err
	}
	pluginEntity, ok := pluginIntf.(PluginIntf)
	if !ok {
		delete(pluginMgr.scannedPlugins, filename)
		return nil, errors.New(filename + "没有实现PluginIntf接口")
	}
	pluginEntity.Init(di.GetDefaultDI())
	pluginMgr.plugins[filename] = &Plug{plug: plug, pi: pluginEntity}
	pluginMgr.installPlugins[filename] = struct{}{} // 记录安装
	return pluginEntity, nil
}

// Download 下载插件,系统,go版本,插件版本
// TODO 怎么处理进度呢? 下载完成或下载失败
func (p *pluginManager) Download(name string) {
	if len(p.path) == 0 {
		return
	}
	pluginPath := filepath.Join(p.path, name+ext)
	if _, err := os.Stat(pluginPath); err != nil {
		return
	}

	f, err := os.OpenFile(pluginPath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}

	go func() {
		// 需要下载zip包
		url := fmt.Sprintf("%s/%s/%s/%s%s", p.remoteDomain, runtime.GOOS, runtime.Version(), name, ".tar.gz")
		client := &http.Client{}
		client.Timeout = time.Second * 60 * 10
		resp, err := client.Get(url)
		if err != nil {
			pine.Logger().Error("下载插件"+name+"失败", err)
			return
		}
		defer resp.Body.Close()

		_, err = io.Copy(f, resp.Body)

		if err != nil {
			pine.Logger().Error("保存插件"+pluginPath+"失败", err)
		} else {
			// todo 解压缩 .gzip
		}

	}()
}

func (p *pluginManager) Uninstall(name string) {
	p.Lock()
	defer p.Unlock()
	intf, exist := p.plugins[name]
	if !exist {
		return
	}
	intf.pi.Uninstall() //TODO 先由插件内部阻止路由访问
	intf.plug = nil
	intf.pi = nil

	delete(p.plugins, name)
}

func Init() {
	if runtime.GOOS == "windows" {
		return
	}
	pluginPath := config.AppConfig().PluginPath
	if len(pluginPath) > 0 {
		pluginMgr.path = pluginPath
		_ = os.Mkdir(pluginMgr.path, os.ModePerm)
	} else {
		pine.Logger().Warning("没有配置PluginPath, 禁用plugin功能")
		return
	}
	scanPluginDir()
	go tickScanDir()
	pluginMgr.loadPlugin()
}

func scanPluginDir() {
	if plugins, _ := filepath.Glob(filepath.Join(pluginMgr.path, pluginMgr.fileGlob)); len(plugins) > 0 {
		for _, f := range plugins {
			conf := &plugin2.Config{}
			confJsonPath := filepath.Join(filepath.Dir(f), jsonName)
			confContent, err := ioutil.ReadFile(confJsonPath)
			if err == nil {
				if err := json.Unmarshal(confContent, conf); err != nil {
					pine.Logger().Print("解析描述文件失败", confJsonPath)
				} else {
					pluginMgr.scannedPlugins[f] = conf
				}
			} else {
				pine.Logger().Print("无法扫描到文件", confJsonPath)
			}
		}
	}
}

func tickScanDir() {
	for range time.Tick(time.Second * 10) {
		scanPluginDir()
	}
}

func PluginMgr() *pluginManager {
	return pluginMgr
}
