package plugin

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

const outputPluginDir = "plugins"
const sourcePluginDir = "src/application/plugins"

const sourceTpl = `package main

import (
	"github.com/go-xorm/xorm"
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



var makePluginCmd = &cobra.Command{
	Use:   "make",
	Short: "创建一个新的插件",
	Run: func(cmd *cobra.Command, args []string) {
		os.Mkdir(sourcePluginDir, os.ModePerm)

		name,_ := cmd.Flags().GetString("name")
		if len(name) == 0 {
			cmd.Usage()
			return
		}


		sourcePluginPath := filepath.Join(sourcePluginDir, name, name + ".go")

		if _, err := os.Stat(sourcePluginPath); err == nil {
			panic(errors.New("已存在同名插件, 请换个名称"))
		}

		os.Mkdir(filepath.Dir(sourcePluginPath), os.ModePerm)

		uuidCode := uuid.NewV4().String()

		source := strings.ReplaceAll(strings.Replace(sourceTpl, "[$uuid]", uuidCode, 1), "[$s]", name)


		err := ioutil.WriteFile(sourcePluginPath, []byte(source), os.ModePerm)
		if err != nil {
			panic(err)
		}
		fmt.Println("生成插件", sourcePluginPath, "源代码文件成功")
	},
}


func init() {
	makePluginCmd.Flags().String("name", "","插件名称, 会生成同名插件模板源文件")
}
