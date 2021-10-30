package plugins

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/valyala/fasthttp"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	plugin2 "github.com/xiusin/pinecms/cmd/plugin"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/config"
)

const ext = ".so"

const exportedVarSuffix = "Plugin"

const jsonName = "config.json"

type PluginIntf interface {
	Init(di.AbstractBuilder)         // 初始化插件
	Sign() string                    // 插件的唯一标识, 需要开发者搞一个独一无二如 uuid
	View() string                    // 配置视图json信息
	Menu(interface{}, int)           // 安装插件位置
	Install()                        // 安装插件, 首次扫描后执行.
	IsInstall() bool                 // 安装状态
	Uninstall()                      // 卸载后禁止访问
	Upgrade()                        // 更新插件
	SetStatus(bool)                  // 设置插件状态
	Status() bool                    // 获取状态
	GetController() pine.IController // 控制器返回
	Prefix() string                  // 路由前缀
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
	for s, s2 := range p.scannedPlugins {
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
			pine.Logger().Printf("启用插件%s失败 %s", plug.Name, err.Error())
		} else {
			pine.Logger().Print("启用插件成功", plug.Name)
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
		return nil, errors.New(filename + "插件未实现PluginIntf接口")
	}
	pluginEntity.Init(di.GetDefaultDI())
	p.registerRouter(pluginEntity)
	pluginMgr.plugins[filename] = &Plug{plug: plug, pi: pluginEntity}
	pluginMgr.installPlugins[filename] = struct{}{}
	return pluginEntity, nil
}

// registerRouter 注册路由, 添加中间件拦截非正常状态
func (p *pluginManager) registerRouter(plug PluginIntf) {
	group := di.MustGet(controllers.ServiceBackendRouter).(*pine.Router)
	group.Use(func(ctx *pine.Context) {
		if !plug.IsInstall() {
			ctx.Abort(fasthttp.StatusNotFound)
		} else if !plug.Status() {
			disableMsg := "插件功能已禁用, 不可访问"
			if ctx.IsAjax() {
				helper.Ajax(disableMsg, 1, ctx)
			} else {
				ctx.Abort(fasthttp.StatusForbidden, disableMsg)
			}
		} else {
			ctx.Next()
		}
	})
	group.Handle(plug.GetController(), plug.Prefix())
	pine.Logger().Print("[plugin:task] 注册路由分组:" + plug.Prefix() + "成功")
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
		url := fmt.Sprintf("%s/%s/%s/%s/%s%s", p.remoteDomain, runtime.GOOS, runtime.GOARCH, runtime.Version(), name, ".tar.gz")
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
			pine.Logger().Error("保存插件到"+pluginPath+"失败", err)
		} else {
			tarExecPath, err := exec.LookPath("tar")
			if err != nil {
				pine.Logger().Error("无法查找到tar命令, 请手动解压" + pluginPath)
				return
			}
			cmd := exec.Command(tarExecPath, "-zxvf", pluginPath)
			cmd.Dir = p.path
			if err := cmd.Run(); err != nil {
				pine.Logger().Error("解压" + pluginPath + "失败, 请手动解压")
			} else {
				pine.Logger().Print("解压" + pluginPath + "成功, 等待程序扫描加载")
			}
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
	intf.pi.Uninstall()
	intf.plug = nil
	intf.pi = nil

	delete(p.plugins, name)
}

func Init() {
	if helper.IsWindows() {
		pine.Logger().Warning("windows 不支持plugin功能")
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
			conf := plugin2.Config{}
			jsonPath := filepath.Join(filepath.Dir(f), jsonName)
			content, err := ioutil.ReadFile(jsonPath)
			if err == nil {
				if err := json.Unmarshal(content, &conf); err != nil {
					pine.Logger().Warning("解析文件"+jsonPath+"失败", err.Error())
				} else {
					pluginMgr.scannedPlugins[f] = &conf
				}
			} else {
				pine.Logger().Warning("获取文件"+jsonPath+"内容异常", err.Error())
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
