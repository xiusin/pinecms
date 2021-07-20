package plugins

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
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

// https://market.topthink.com/product/389 样例
// https://h.jizhicms.cn/ 模板样例
// sql parser 模型创建
//https://github.com/GoAdminGroup/go-admin/blob/master/plugins/plugins.go
//https://github.com/tsenart/vegeta
// make plugin

type PluginIntf interface {
	Init(*pine.Application, *pine.Router, *cobra.Command) // 初始化插件
	Prefix(string)                                        // 路由前缀, 注册后修改需要重启程序

	Sign() string        // 插件的唯一标识, 需要开发者搞一个独一无二如 uuid
	Name() string        // 插件名称
	Version() string     // 版本号
	Author() string      // 开发者
	Description() string // 插件描述
	Contact() string     // 联系电话
	Course() string      // 完整教程信息, 以html/markdown方式呈现

	View() string // 配置视图json信息
	Install()     // 安装插件, 首次扫描后执行.
	Uninstall()   // 卸载后禁止访问
	Upgrade()     // 更新插件
}

// todo 可以打包子目录进行前端代码发布, 或后端源代码
// 插件端可以配置导出assets.zip

type pluginManager struct {
	sync.Mutex
	plugins        map[string]*Plug
	installPlugins []string // 已安装版本
	scannedPlugins []string
	path           string
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

func (p *pluginManager) GetLoadPlugin() []string {
	return p.installPlugins
}

func (p *pluginManager) loadPlugin() {
	p.Lock()
	defer p.Unlock()
	if len(p.plugins) == 0 {
		content, err := ioutil.ReadFile(filepath.Join(p.path, "plugins.json"))
		if err != nil {
			return
		}
		_ = json.Unmarshal(content, &p.installPlugins)
		// TODO  未安装的屏蔽, 在列表处不显示具体信息, 加一个按钮进入安装页面.
	}
}

// Download 下载插件,系统,go版本,插件版本
// TODO 怎么处理进度呢? 下载完成或下载失败
func (p *pluginManager) Download(name string) {
	if p.path == "" {
		return
	}
	pluginPath := filepath.Join(p.path, name+".so")
	if _, err := os.Stat(pluginPath); err != nil {
		pine.Logger().Warning("已经安装插件: ", err)
		return
	}

	f, err := os.OpenFile(pluginPath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}

	go func() {
		url := fmt.Sprintf("%s/%s/%s/%s.so", p.remoteDomain, runtime.GOOS, runtime.Version(), name)
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

type Plug struct {
	plug *plugin.Plugin
	pi   PluginIntf
}

var pluginMgr = &pluginManager{
	plugins:        map[string]*Plug{},
	installPlugins: []string{},
	remoteDomain:   "https://plugin.xiusin.cn",
}

func Init() {
	if runtime.GOOS == "windows" {
		return
	}
	pluginPath := config.AppConfig().PluginPath
	if len(pluginPath) > 0 {
		pluginMgr.path = pluginPath
		pluginMgr.loadPlugin() // 加载配置
	} else {
		pine.Logger().Print("没有配置PluginPath, 禁用plugin功能")
		return
	}
	scanPluginDir()
	go tickScanDir()
}

func scanPluginDir() {
	pluginMgr.Lock()
	defer pluginMgr.Unlock()
	if len(pluginMgr.installPlugins) > 0 {
		return
	}
	_ = os.Mkdir(pluginMgr.path, os.ModePerm)
	plugins, _ := filepath.Glob(pluginMgr.path + "/*.so")
	if len(plugins) > 0 {
		for _, f := range plugins {
			name := helper.UcFirst(strings.TrimSuffix(filepath.Base(f), ".so")) + "Plugin"
			if _, exist := pluginMgr.plugins[f]; !exist {
				p, err := plugin.Open(f)
				if err != nil {
					pine.Logger().Error(err)
					continue
				}
				pluginIntf, err := p.Lookup(name)
				if err != nil {
					pine.Logger().Print(name+"加载失败", err)
					continue
				}
				pluginEntity, ok := pluginIntf.(PluginIntf)
				if !ok {
					pine.Logger().Print(name, "没有实现PluginIntf接口")
					continue
				}

				pluginEntity.Init(
					di.MustGet("pine.application").(*pine.Application),
					di.MustGet("pine.backend_router_group").(*pine.Router),
					di.MustGet("pinecms.cmd.root").(*cobra.Command),
				)

				pluginMgr.plugins[f] = &Plug{plug: p, pi: pluginEntity}
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
