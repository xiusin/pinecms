package plugin

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"github.com/xiusin/pinecms/src/common/helper"
)

const outputPluginDir = "plugins"
const sourcePluginDir = "src/application/plugins"
const configName = "config.json"

const sourceTpl = `package main

import (
	"xorm.io/xorm"
	"github.com/xiusin/pine"
	"sync"
)

type [$s] struct {
	sync.Once
	orm       *xorm.Engine
	app       *pine.Application
	prefix    string
	isInstall bool
}

func (p *[$s]) Name() string {
	return ""
}

func (p *[$s]) Sign() string {
	return "[$uuid]"
}

func (p *[$s]) Version() string {
	return "dev 0.0.1"
}

func (p *[$s]) Author() string {
	return ""
}

func (p *[$s]) Description() string {
	return ""
}

func (p *[$s]) Course() string {
	return ""
}

func (p *[$s]) Contact() string {
	return ""
}

func (p *[$s]) View() string {
	return ""
}

func (p *[$s]) Install() {

}

func (p *[$s]) Prefix(prefix string) {
	p.prefix = prefix
}

func (p *[$s]) Upgrade() {
}

func (p *[$s]) Init(app *pine.Application, backend *pine.Router) {
	
}

// [$s]Plugin 导出插件可执行变量 不可删除
var [$s]Plugin = [$s]{}`

type Config struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Contact     string `json:"contact"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Page        string `json:"page"` // 页面说明
	Error       string `json:"-"`    // 记录加载信息
}

var makePluginCmd = &cobra.Command{
	Use:   "make",
	Short: "创建一个新的插件以及配置文件",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if len(name) == 0 {
			_ = cmd.Usage()
			return
		}
		var config = &Config{
			Name: name + " plugin",
			Page: "页面简介,可以是html和markdown",
		}
		pluginDir := filepath.Join(sourcePluginDir, name)
		_ = os.MkdirAll(pluginDir, os.ModePerm)

		sourcePluginPath := filepath.Join(pluginDir, name+".go")

		if _, err := os.Stat(sourcePluginPath); err == nil {
			panic(errors.New("已存在同名插件, 请换个名称"))
		}

		source := strings.ReplaceAll(strings.Replace(sourceTpl, "[$uuid]", uuid.NewV4().String(), 1), "[$s]", name)

		helper.PanicErr(os.WriteFile(sourcePluginPath, []byte(source), os.ModePerm))

		configPath := filepath.Join(pluginDir, configName)
		conf, _ := json.MarshalIndent(config, "", "    ")
		if err := os.WriteFile(configPath, conf, os.ModePerm); err != nil {
			_ = os.RemoveAll(pluginDir)
			panic(err)
		}
		fmt.Println("生成插件相关文件", sourcePluginPath, "和", configPath, "成功")
	},
}

func init() {
	makePluginCmd.Flags().String("name", "", "插件名称, 会生成同名插件模板源文件")
}
