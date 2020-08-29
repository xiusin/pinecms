package cmd

import (
	"bytes"
	"encoding/json"
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
	controllerDir = "src/application/controllers/backend/"
	modelDir      = "src/application/models/"
	tableDir      = modelDir + "tables/"
	feDir         = "frontend/src/pages/"
	theme         = "vim"
	controllerTpl = `package backend
import (
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
	c.SearchFields = map[string]searchFieldDsl{}
	c.Orm = pine.Make(controllers.ServiceXorm).(*xorm.Engine)
	c.Table = &tables.[table]{}
	c.Entries = &[]tables.[table]{}
}`
	modelTpl = `package models

import (
	"github.com/go-xorm/xorm"
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

	indexTsTpl = `export const schema = {
  type: 'page',

  body: {
    type: 'lib-crud',
    api: '$preset.apis.list',
    filter: '$preset.forms.filter',
    filterTogglable: true,
    perPageAvailable: [50, 100, 200],
    defaultParams: {
      size: 50,
    },
    perPageField: 'size',
    pageField: 'page',
    headerToolbar: [
      'filter-toggler',
      {
        type: 'columns-toggler',
        align: 'left',
      },
      {
        type: 'pagination',
        align: 'left',
      },
      '$preset.actions.add',
    ],
    footerToolbar: ['statistics', 'switch-per-page', 'pagination'],
    columns:  [
      [tableDSL]
      {
        type: 'operation',
        label: '操作',
        width: 60,
        limits: ['edit', 'del'],
        limitsLogic: 'or',
        buttons: ['$preset.actions.edit', '$preset.actions.del'],
      },
    ],
  },
  definitions: {
    updateControls: {
      controls: [formDSL],
    },
  },
  preset: {
    actions: {
      add: {
        limits: 'add',
        type: 'button',
        align: 'right',
        actionType: 'drawer',
        label: '添加',
        icon: 'fa fa-plus pull-left',
        size: 'sm',
        primary: true,
        drawer: {
          title: '新增',
          size: 'xl',
          body: {
            type: 'form',
            api: '$preset.apis.add',
            mode: 'normal',
            $ref: 'updateControls',
          },
        },
      },
      edit: {
        limits: 'edit',
        type: 'button',
        icon: 'fa fa-pencil',
        tooltip: '编辑',
        actionType: 'dialog',
        dialog: {
          title: '编辑',
          size: 'xl',
          body: {
            type: 'form',
            mode: 'normal',
            api: '$preset.apis.edit',
            $ref: 'updateControls',
          },
        },
      },
      del: {
        limits: 'del',
        type: 'action',
        icon: 'fa fa-times text-danger',
        actionType: 'ajax',
        tooltip: '删除',
        confirmText: '您确认要删除?',
        api: {
          $preset: 'apis.del',
        },
        messages: {
          success: '删除成功',
          failed: '删除失败',
        },
      },
    },
    forms: {
      filter: {
        controls: [
          [filterDSL]
          {
            type: 'submit',
            className: 'm-l',
            label: '搜索',
          },
        ],
      },  // 搜索
    },
  },
}
`

	presetTsTpl = `export default {
  limits: {
    $page: {
      label: '查看列表',
    },
    add: {
      label: '添加',
    },
    edit: {
      label: '编辑',
    },
    del: {
      label: '删除',
    },
  },
  apis: {
    list: {
      url: 'GET [table]/list',
      limits: '$page',
      onPreRequest: (source) => {
        const { dateRange } = source.data
        if (dateRange) {
          const arr = dateRange.split('%2C')
          source.data = {
            ...source.data,
            startDate: arr[0],
            endDate: arr[1],
          }
        }
        return source
      },
    },
    add: {
      url: 'POST [table]/add',
      limits: 'add',
    },
    edit: {
      url: 'POST [table]/edit?linkid=$id',
      limits: 'edit',
    },
    del: {
      url: 'POST [table]/delete?id=$id',
      limits: 'del',
    },
  },
}`
)

var cols = map[string]*core.Column{}

var matchSuffix = struct {
	enumRadioSuffix, setCheckboxSuffix, switchSuffix []string
	imageSuffix, fileSuffix                          []string
	editorSuffix                                     []string
	match                                            func([]string, string) bool
}{
	enumRadioSuffix:   []string{"data", "state", "status"},
	setCheckboxSuffix: []string{"data", "state", "status"},
	switchSuffix:      []string{"switch"},
	editorSuffix:      []string{"editor", "content"},
	imageSuffix:       []string{"image", "images", "avatar", "avatars", "logo", "logos", "img", "imgs", "pic", "pics"},
	fileSuffix:        []string{"file", "files"},
	match: func(suffixes []string, field string) bool {
		var matched bool
		for _, suffix := range suffixes {
			if strings.HasSuffix(field, suffix) {
				matched = true
				break
			}
		}
		return matched
	},
}

var crudCmd = &cobra.Command{
	Use:   "crud",
	Short: "生成基本crud模块",
	Run: func(cmd *cobra.Command, args []string) {
		config.Bootstrap() // 方法不可放到init里，否则缓存组件阻塞
		if !config.Ac().Debug {
			logger.SetReportCaller(false)
			logger.Print("非Debug模式，不支持 CRUD 命令")
			return
		}
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

		for _, v := range tableMata.Columns() {
			cols[v.Name] = v
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
			return
		}
		err = genControllerFile(print, controllerName, table, controllerPath)
		if err != nil {
			logger.Error(err)
			return
		}

		err = genTableFile(print, table, tableDir+table+".go")

		if err != nil {
			logger.Error(err)
			return
		}

		logger.Print("创建模块文件成功, 请将控制器注册到路由: registerV2BackendRoutes方法内")
	},
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
		quick.Highlight(logger.DefaultWriter(), content, "go", "terminal256", theme)
	}
	return err
}

func genControllerFile(print bool, controllerName, tableName, controllerPath string) error {
	var err error
	content := strings.ReplaceAll(controllerTpl, "[ctrl]", controllerName)
	content = strings.ReplaceAll(content, "[table]", camelString(tableName))
	if !print {
		err = ioutil.WriteFile(controllerPath, []byte(content), os.ModePerm)
	}
	if err == nil {
		logger.Print("创建文件： " + color.Green.Sprint(controllerPath))
	}
	if print {
		quick.Highlight(logger.DefaultWriter(), content, "go", "terminal256", theme)
	}
	return err
}

func genTableFile(print bool, tableName, tablePath string) error {
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
			scol.IsPrimaryKey = col.Name.String() == primaryKey
			for _, k := range uniqueKeys {
				if scol.Name == k {
					scol.IsUnique = true
					break
				}
			}

			table.Cols = append(table.Cols, scol)
		}

		tableStruct = table.toXorm(print, tableName)
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
	if print {
		quick.Highlight(logger.DefaultWriter(), content, "go", "terminal256", theme)
	}
	return err
}

func genFrontendFile(print bool, table string, tableDsl, formDsl, filterDSL []map[string]interface{}) {
	// 根据路由创建目录文件
	moduleFeDir := feDir + table + "/list"
	indexFile := strFirstToUpper(moduleFeDir + "/index.ts")
	presetFile := strFirstToUpper(moduleFeDir + "/preset.ts")

	// 生成curd描述文件
	data, _ := json.MarshalIndent(tableDsl, "", "\t")
	content := bytes.ReplaceAll([]byte(indexTsTpl), []byte("[tableDSL]"), append(bytes.Trim(data, "[]"), ','))
	data, _ = json.MarshalIndent(filterDSL, "", "\t")
	content = bytes.ReplaceAll(content, []byte("[filterDSL]"), append(bytes.Trim(data, "[]"), ','))
	data, _ = json.MarshalIndent(formDsl, "", "\t")
	content = bytes.ReplaceAll(content, []byte("[formDSL]"), data)
	presetContent := bytes.ReplaceAll([]byte(presetTsTpl), []byte("[table]"), []byte(table))
	if !print {
		os.RemoveAll(moduleFeDir) // 强制创建
		os.MkdirAll(moduleFeDir, os.ModePerm)
		err := ioutil.WriteFile(indexFile, content, os.ModePerm)
		if err == nil {
			logger.Print("创建文件: " + indexFile)
		}
		err = ioutil.WriteFile(presetFile, presetContent, os.ModePerm)
		if err == nil {
			logger.Print("创建文件: " + presetFile)
		}
	} else {
		logger.Print("创建文件: " + indexFile)
		quick.Highlight(logger.DefaultWriter(), string(content), "typescript", "terminal256", theme)
		logger.Print("创建文件: " + presetFile)
		quick.Highlight(logger.DefaultWriter(), string(presetContent), "typescript", "terminal256", theme)
	}
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

func strFirstToUpper(str string) string {
	temp := strings.Split(strings.ReplaceAll(str, "_", "-"), "-")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		if y != 0 {
			for i := 0; i < len(vv); i++ {
				if i == 0 {
					vv[i] -= 32
					upperStr += string(vv[i]) // + string(vv[i+1])
				} else {
					upperStr += string(vv[i])
				}
			}
		}
	}
	return temp[0] + upperStr
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

func getFieldType(fieldName, fieldType string) string {
	inputType := "text"
	switch fieldType {
	case "bigint", "int", "mediumint", "smallint", "tinyint":
		inputType = "number"
	case "set":
		inputType = "checkboxes"
	case "enum":
		inputType = "radios"
	case "decimal", "double", "float":
		inputType = "number"
	case "longtext", "text", "mediumtext", "smalltext", "tinytext":
		inputType = "textarea"
	case "year", "date", "time", "datetime", "timestamp":
		inputType = "datetime"
	}

	if matchSuffix.match(matchSuffix.imageSuffix, fieldName) {
		inputType = "image"
	} else if matchSuffix.match(matchSuffix.fileSuffix, fieldName) {
		inputType = "file"
	} else {
		if fieldType == "enum" && matchSuffix.match(matchSuffix.enumRadioSuffix, fieldName) {
			inputType = "radios"
		}
		if fieldType == "set" && matchSuffix.match(matchSuffix.setCheckboxSuffix, fieldName) {
			inputType = "checkboxes"
		}
		if strings.HasSuffix(fieldType, "text") && matchSuffix.match(matchSuffix.editorSuffix, fieldName) {
			inputType = "rich-text"
		}
	}

	return inputType
}

// 生成表单字段额外扩展字段
func getFieldFormExtra(amisType string, item map[string]interface{}) {
	col := cols[item["name"].(string)]
	if !col.DefaultIsEmpty {
		item["value"] = col.Default	// 设置字段默认值
	}
	switch amisType {
	case "radios", "checkboxes":
		var options []map[string]interface{}
		optionSets := col.SetOptions
		if amisType == "radios" {
			optionSets = col.EnumOptions
		}
		if len(optionSets) > 0 {	// 设置字段选项
			for _, v := range optionSets {
				options = append(options, map[string]interface{}{
					"label": v,
					"value": v,
				})
			}
			item["options"] = options
		}
	case "image", "file":
		if strings.HasSuffix(item["name"].(string), "s") {
			item["multiple"] = true
		}
	}
}


