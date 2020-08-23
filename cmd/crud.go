package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/xwb1989/sqlparser"

	"github.com/alecthomas/chroma/quick"
	"github.com/gookit/color"
	"github.com/xiusin/logger"
	config "github.com/xiusin/pinecms/src/server"
	"xorm.io/core"

	"github.com/spf13/cobra"
)

const (
	controllerDir = "src/application/backend/"
	modelDir      = "src/application/models/"
	tableDir      = modelDir + "tables/"
	controllerTpl = `package backend
import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type [ctrl] struct {
	BaseController
}

func (c *[ctrl]) Construct() {
	c.BindType = "form"
	c.Orm = pine.Make(controllers.ServiceXorm).(*xorm.Engine)
	c.Table = &tables.[table]{}
	c.Entries = []tables.[table]{}
}`
	modelTpl = `package models

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

	tableTpl = `package tables

[struct]
`
)

const theme = "vim"

var crudCmd = &cobra.Command{
	Use:   "crud",
	Short: "生成基本crud模块",
	Run: func(cmd *cobra.Command, args []string) {
		config.Bootstrap() // 方法不可放到init里，否则缓存组件阻塞
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
		modelName, modelPath := getModelName(table)
		controllerName, controllerPath := getControllerName(table)
		tablePath := tableDir + table + ".go"
		if !force && !print {
			f, err := os.Stat(modelPath)
			if !os.IsNotExist(err) && !f.IsDir() {
				logger.Print("已有存在的文件: " + modelPath)
			}
			f, err = os.Stat(controllerPath)
			if !os.IsNotExist(err) && !f.IsDir() {
				logger.Print("已有存在的文件: " + controllerPath)
			}
			f, err = os.Stat(tablePath)
			if !os.IsNotExist(err) && !f.IsDir() {
				logger.Print("已有存在的文件: " + tablePath)
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

		err = genTableFile(print, camelString(table), tableDir+table+".go", tableMata.Columns())
		if err != nil {
			logger.Error(err)
		}
	},
}

// SQLTable 表名结构体
type SQLTable struct {
	Name string
	Cols []SQLColumn
}

// SQLColumn 描述结构体
type SQLColumn struct {
	Name, Type    string
	IsPrimaryKey  bool
	IsUnique      bool
	Length        string
	EnumValues    []string
	AutoIncrement bool
	NotNull       bool
	Default       string
}

func init() {
	crudCmd.Flags().String("table", "", "数据库表名")
	crudCmd.Flags().Bool("force", false, "是否强制覆盖（可能导致已有代码丢失）")
	crudCmd.Flags().Bool("print", false, "是否只打印生成文件以及操作步骤")
	rootCmd.AddCommand(crudCmd)
}

func getModelName(tableName string) (model string, filename string) {
	model = camelString(tableName) + "Model"
	filename = modelDir + snakeString(tableName) + "_model.go"
	return
}

func getControllerName(tableName string) (controller string, filename string) {
	controller = camelString(tableName) + "Controller"
	filename = controllerDir + snakeString(tableName) + "_controller.go"
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
	content := strings.ReplaceAll(modelTpl, "[model]", modelName)
	if !print {
		err = ioutil.WriteFile(modelPath, []byte(content), os.ModePerm)
	}
	if err == nil {
		logger.Print("创建文件： " + color.Green.Sprint(modelPath))
	}
	if print {
		quick.Highlight(os.Stdout, content, "go", "terminal256", theme)
	}
	return err
}

func genControllerFile(print bool, controllerName, tableName, controllerPath string) error {
	var err error
	content := strings.ReplaceAll(controllerTpl, "[ctrl]", controllerName)
	content = strings.ReplaceAll(content, "[table]", tableName)
	if !print {
		err = ioutil.WriteFile(controllerPath, []byte(content), os.ModePerm)
	}
	if err == nil {
		logger.Print("创建文件： " + color.Green.Sprint(controllerPath))
	}
	if print {
		quick.Highlight(os.Stdout, content, "go", "terminal256", theme)
	}
	return err
}

func genTableFile(print bool, tableName, tablePath string, cols []*core.Column) error {
	//for _, _ := range cols {
	//	//fmt.Println(col.Name, 		col.Comment)
	//}
	realTableName := config.Dc().Db.DbPrefix + strings.ToLower(tableName)
	res, err := config.XOrmEngine.QueryString(`show create table ` + realTableName)

	if err != nil {
		return err
	}

	createSQL := res[0]["Create Table"]

	// 替换字段
	reg := regexp.MustCompile(`"(.+?)"\s`)

	createSQL = reg.ReplaceAllStringFunc(createSQL, func(s string) string {
		s = strings.Trim(s, `" `)
		return "`" + s + "` "
	})
	stmt, err := sqlparser.Parse(createSQL)
	var tableStruct string
	if err != nil {
		panic(err)
	}
	switch stmt := stmt.(type) {
	case *sqlparser.DDL:
		if stmt.TableSpec == nil {
			logger.Error("Canont get table spec")
			break
		}
		var table SQLTable

		var uniqueKeys []string
		var primaryKey string
		for _, ind := range stmt.TableSpec.Indexes {
			switch ind.Info.Type {
			case "primary key":
				primaryKey = ind.Columns[0].Column.String()
			case "unique key":
				uniqueKeys = append(uniqueKeys, ind.Columns[0].Column.String())
			default:
				fmt.Fprintln(os.Stderr, "unknown type ", ind.Info.Type)
			}
		}

		table.Name = stmt.NewName.Name.String()
		for _, col := range stmt.TableSpec.Columns {
			var scol SQLColumn

			scol.Name = col.Name.String()
			scol.Type = col.Type.Type
			scol.EnumValues = col.Type.EnumValues
			if col.Type.Length != nil {
				scol.Length = string(col.Type.Length.Val)
			}
			scol.AutoIncrement = bool(col.Type.Autoincrement)
			scol.NotNull = bool(col.Type.NotNull)
			if col.Type.Default != nil {
				scol.Default = string(col.Type.Default.Val)
			}
			scol.IsPrimaryKey = (col.Name.String() == primaryKey)
			for _, k := range uniqueKeys {
				if scol.Name == k {
					scol.IsUnique = true
					break
				}
			}

			table.Cols = append(table.Cols, scol)
		}

		tableStruct = table.toXorm(tableName)
	}

	if tableStruct == "" {
		return errors.New("没有生成模型内容, 请检查数据表是否正确")
	}

	content := strings.ReplaceAll(tableTpl, "[struct]", tableStruct)
	if !print {
		err = ioutil.WriteFile(tablePath, []byte(content), os.ModePerm)
	}
	if err == nil {
		logger.Print("创建文件： " + color.Green.Sprint(tablePath))
	}
	if print{
		quick.Highlight(os.Stdout, content, "go", "terminal256", theme)
	}
	return err
}

func (t *SQLTable) toXorm(tableName string) string {
	var str strings.Builder
	str.WriteString(fmt.Sprintf("type %s struct {\n", tableName))
	for _, col := range t.Cols {
		str.WriteRune('\t')
		str.WriteString(camelString(col.Name))

		var goType string
		switch col.Type {
		case "varchar", "text", "enum", "char", "longtext":
			goType = "string"
		case "int", "bigint":
			goType = "int64"
		case "tinyint":
			goType = "int"
		case "double", "float":
			goType = "float64"
		case "date", "datetime", "time", "timestamp":
			goType = "time.Time"
		case "blob":
			goType = "[]byte"
		default:
			panic(col.Name + " 是一个未知类型")
		}
		str.WriteString(" " + goType)
		str.WriteString(" `xorm:\"")

		// Type
		str.WriteString(col.Type)

		// Bracketed type metadata
		if len(col.EnumValues) > 0 {
			str.WriteRune('(')
			for i, en := range col.EnumValues {
				str.WriteString(en)
				if i != len(col.EnumValues)-1 {
					str.WriteRune(',')
				}
			}
			str.WriteRune(')')
		} else if len(col.Length) > 0 {
			str.WriteString("(" + col.Length + ")")
		}

		if col.AutoIncrement {
			str.WriteString(" autoincr")
		}
		if col.NotNull {
			str.WriteString(" not null")
		}
		if len(col.Default) > 0 {
			str.WriteString(" default '" + col.Default + "'")
		}
		if col.IsPrimaryKey {
			str.WriteString(" pk")
		}
		if col.IsUnique {
			str.WriteString(" unique")
		}
		str.WriteString(" '" + col.Name + "'")

		// close variable tag
		str.WriteString("\"`\n")
	}
	str.WriteString("}")
	return str.String()
}
