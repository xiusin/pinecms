package cmd

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/xiusin/logger"
	config "github.com/xiusin/pinecms/src/server"
	"io/ioutil"
	"os"
	"strings"
	"xorm.io/core"

	"github.com/spf13/cobra"
)

const (
	CONTROLLER_DIR = "src/application/backend/"
	MODEL_DIR      = "src/application/models/"
	TABLE_DIR      = MODEL_DIR + "tables/"
	CONTROLLER_TPL = `package backend
import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
)

type [ctrl] struct {
	BaseController
}

func (c *[ctrl]) RegisterRoute(b pine.IRouterWrapper) {
	b.GET("/[table]/list", "List")
	b.POST("/[table]/add", "Add")
	b.ANY("/[table]/edit", "Edit")
	b.ANY("/[table]/order", "Order")
	b.ANY("/[table]/delete", "Delete")
}`
	MODEL_TPL = `package models

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
)

type [model] struct {
	orm *xorm.Engine
}

func New[model]() *[model] {
	return &[model]{orm: di.MustGet("*xorm.Engine").(*xorm.Engine)}
}`
)

var crudCmd = &cobra.Command{
	Use:   "crud",
	Short: "生成基本crud模块",
	Run: func(cmd *cobra.Command, args []string) {

		table, _ := cmd.Flags().GetString("table")
		force, _ := cmd.Flags().GetBool("force")
		print, _ := cmd.Flags().GetBool("print")
		if table == "" {
			cmd.Help()
			return
		}
		metas, _ := config.XOrmEngine.DBMetas()
		var tableMata *core.Table
		for _, meta := range metas {
			if meta.Name == getTableName(table) {
				tableMata = meta
				break
			}
		}
		if tableMata == nil {
			logger.Errorf("无法获取数据表[%s]元信息", getTableName(table))
			return
		}
		// 表字段
		//cols := tableMata.Columns()
		modelName, modelPath := getModelName(table)
		controllerName, controllerPath := getControllerName(table)

		if !force && !print {
			f, err := os.Stat(modelPath)
			if !os.IsNotExist(err) && !f.IsDir() {
				logger.Print("已有存在的文件: " + modelPath)
			}
			f, err = os.Stat(controllerPath)
			if !os.IsNotExist(err) && !f.IsDir() {
				logger.Print("已有存在的文件: " + controllerPath)
			}
		}
		err := genModelFile(print, modelName, modelPath)
		if err != nil {
			logger.Error(err)
		}
		err = genControllerFile(print, controllerName, table, controllerPath)
		if err != nil {
			logger.Error(err)
		}
	},
}

func init() {
	crudCmd.Flags().String("table", "", "数据库表名")
	crudCmd.Flags().Bool("force", false, "是否强制覆盖（可能导致已有代码丢失）")
	crudCmd.Flags().Bool("print", false, "是否只打印生成文件以及操作步骤")
	rootCmd.AddCommand(crudCmd)
	config.Bootstrap()
}

func getModelName(tableName string) (model string, filename string) {
	model = camelString(tableName) + "Model"
	filename = MODEL_DIR + snakeString(tableName) + "_model.go"
	return
}

func getControllerName(tableName string) (controller string, filename string) {
	controller = camelString(tableName) + "Controller"
	filename = CONTROLLER_DIR + snakeString(tableName) + "_controller.go"
	return
}

func getTableName(table string) string {
	prefix := config.Dc().Db.DbPrefix
	if strings.HasPrefix(table, prefix) {
		return table
	}
	return prefix + table
}

func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

func camelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

func genModelFile(print bool, modelName, modelPath string) error {
	var err error
	content := strings.ReplaceAll(MODEL_TPL, "[model]", modelName)
	if !print {
		err = ioutil.WriteFile(modelPath, []byte(content), os.ModePerm)
	}
	if err == nil {
		logger.Print("创建文件： " + color.Green.Sprint(modelPath))
	}
	if print {
		fmt.Println(content)
	}
	return err
}

func genControllerFile(print bool, controllerName, tableName, controllerPath string) error {
	var err error
	content := strings.ReplaceAll(CONTROLLER_TPL, "[ctrl]", controllerName)
	content = strings.ReplaceAll(content, "[table]", tableName)
	if !print {
		err = ioutil.WriteFile(controllerPath, []byte(content), os.ModePerm)
	}
	if err == nil {
		logger.Print("创建文件： " + color.Green.Sprint(controllerPath))
	}
	if print {
		fmt.Println(content)
	}
	return err
}

func genTableFile(print bool, tableName, tablePath string) error  {
	var err error

	if print {
		fmt.Println(content)
	}
	return err
}
