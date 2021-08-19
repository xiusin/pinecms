package crud

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/cmd/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/xwb1989/sqlparser"

	"github.com/alecthomas/chroma/quick"
	"github.com/gookit/color"
	"github.com/xiusin/logger"
	config "github.com/xiusin/pinecms/src/server"
	"xorm.io/core"

	"github.com/spf13/cobra"
)

var cols = map[string]*core.Column{}
var topCode []string
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

var Cmd = &cobra.Command{
	Use:   "crud",
	Short: "生成模块的crud功能",
	Long: `
字段设定格式: 字段注释(:字典备注:组件名称)
如: 
状态 -  仅解析名称, 组件根据类型或字段后缀推导
状态:-1=禁用,0=待审核,1=正常 - 解析名称并设置下拉
状态:-1=禁用,0=待审核,1=正常:el-checkbox 
`,
	Run: func(cmd *cobra.Command, args []string) {
		config.InitDB()
		if !config.Ac().Debug {
			logger.SetReportCaller(false)
			logger.Print("非Debug模式，不支持 crud 命令")
			return
		}
		table, _ := cmd.Flags().GetString("table")
		force, _ := cmd.Flags().GetBool("force") // 强制创建
		onlyInfo, _ := cmd.Flags().GetBool("info")
		frontendPath, _ := cmd.Flags().GetString("fepath")
		if len(table) == 0 {
			logger.Print(color.Red.Sprint("请输入表名"))
			_ = cmd.Help()
			return
		}
		metas, err := config.XOrmEngine.DBMetas()
		if err != nil {
			pine.Logger().Error(err)
			return
		}
		var tableMata *core.Table
		for _, meta := range metas {
			if meta.Name == getTableName(table) {
				tableMata = meta
				break
			}
		}
		if tableMata == nil {
			logger.Errorf("无法获取数据表[%s]元信息", color.Red.Sprint(getTableName(table)))
			return
		}
		for _, v := range tableMata.Columns() {
			cols[v.Name] = v
		}
		controllerName, controllerPath := getControllerName(table)
		tablePath := tableDir + table + goExt
		if !force { // 非强制创建则检测文件是否存在
			for _, s := range []string{frontendPath + "/" + feModuleDir + table, controllerPath, tablePath} {
				if _, err := os.Stat(s); !os.IsNotExist(err) {
					logger.Print("已有存在: " + color.Red.Sprint(s))
					return
				}
			}
		}
		if err = genTableFileAndFrontendFile(onlyInfo, table, tableDir+table+goExt, frontendPath); err != nil {
			panic(err)
		}
		if err = genControllerFile(onlyInfo, controllerName, table, controllerPath); err != nil {
			panic(err)
		}
		byts, err := ioutil.ReadFile(routerFile)
		if err != nil {
			panic(err)
		}
		controllerNamespace := `"github.com/xiusin/pinecms/src/application/controllers/backend"`
		pineNamespace := `"github.com/xiusin/pine"`
		holder := "// holder"
		if !bytes.Contains(byts, []byte(controllerNamespace)) {
			byts = bytes.Replace(byts, []byte(pineNamespace), []byte(pineNamespace+"\r\n\t"+controllerNamespace), 1)
		}
		byts = bytes.Replace(byts, []byte(holder), []byte(holder+"\r\n\t"+`backendRouter.Handle(new(backend.`+controllerName+`), "/`+util.SnakeString(table)+`")`), 1)
		ioutil.WriteFile(routerFile, byts, os.ModePerm)
		logger.Print("创建模块文件成功, 已注册路由信息至: " + color.Green.Sprint(routerFile))
	},
}

func init() {
	Cmd.Flags().String("table", "", "数据库表名")
	Cmd.Flags().Bool("force", false, "是否强制覆盖（可能导致已有代码丢失）")
	Cmd.Flags().Bool("info", false, "是否只打印生成文件以及操作步骤")
	Cmd.Flags().String("fepath", "admin", "前端开发根目录")
}

func getControllerName(tableName string) (controller string, filename string) {
	controller = util.CamelString(tableName) + "Controller"
	filename = controllerDir + util.SnakeString(tableName) + "_controller" + goExt
	return
}

func getTableName(table string) string {
	prefix := config.Dc().Db.DbPrefix
	if strings.HasPrefix(table, prefix) {
		return table
	}
	return prefix + table
}

func genControllerFile(print bool, controllerName, tableName, controllerPath string) error {
	var err error
	content := strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(controllerTpl, "[ctrl]", controllerName),
			"[table]",
			util.SnakeString(tableName),
		),
		"[searchFieldDsl]",
		searchFieldDsl,
	)
	if !print {
		err = ioutil.WriteFile(controllerPath, []byte(content), os.ModePerm)
	}
	if err == nil {
		logger.Print("创建文件： " + color.Green.Sprint(controllerPath))
	}
	if print {
		_ = quick.Highlight(logger.DefaultWriter(), content, "go", "terminal256", theme)
	}
	return err
}

func genTableFileAndFrontendFile(print bool, tableName, tablePath, frontendPath string) error {
	realTableName := config.Dc().Db.DbPrefix + strings.ToLower(tableName)
	res, err := config.XOrmEngine.QueryString(`SHOW CREATE TABLE ` + realTableName)
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
			logger.Error("Can't get table spec")
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
				logger.Warning("未知类型 ", ind.Info.Type)
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

		tableStruct = table.toXorm(print, tableName, frontendPath)
	}

	if len(tableStruct) == 0 {
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
		_ = quick.Highlight(logger.DefaultWriter(), content, "go", "terminal256", theme)
	}
	return err
}

// genFrontendFile 生成前端模块
func genFrontendFile(table, frontendPath string, tableDsl, formDsl, filterDsl []map[string]interface{}) {
	//fmt.Println(tableDsl)
	//data, _ := json.Marshal(tableDsl)
	//fmt.Println(string(data))
	//fmt.Println(filterDsl)

	moduleBaseDir := filepath.Join(frontendPath, feModuleDir+table)
	files := map[string]string{
		filepath.Join(moduleBaseDir, "service", "index.ts"):  serviceTsTpl,
		filepath.Join(moduleBaseDir, "service", "router.ts"): serviceRouterTpl,
		filepath.Join(moduleBaseDir, "views", table+".vue"):  indexVueTpl,
		filepath.Join(moduleBaseDir, "index.ts"):             serviceIndexTsTpl,
	}
	for filename, content := range files {
		os.MkdirAll(filepath.Dir(filename), os.ModePerm)
		if strings.Contains(filename, ".vue") {
			data, _ := json.MarshalIndent(formDsl, "\t\t\t", "\t")
			content = string(bytes.ReplaceAll([]byte(content), []byte("[formDSL]"), data))

			data, _ = json.MarshalIndent(tableDsl, "\t\t\t", "\t")
			content = string(bytes.ReplaceAll([]byte(content), []byte("[tableDSL]"), data))
		}
		if err := os.WriteFile(filename, bytes.ReplaceAll([]byte(content), []byte("[table]"), []byte(util.SnakeString(table))), os.ModePerm); err == nil {
			logger.Print("创建文件： " + color.Green.Sprint(filename))
		} else {
			logger.Print("创建文件", color.Red.Sprint(filename)+"失败")
		}
	}
}

// getLabelAndFieldTypeAndProps 解析生成基础结构
func getLabelAndFieldTypeAndProps(col SQLColumn, xormCol *core.Column) (labelName, inputType string, props map[string]interface{}) {
	inputType, fieldType, fieldName, props := "el-input", col.Type, col.Name, map[string]interface{}{"size": "mini", "is_number": false, "is_float": false}

	if matchSuffix.match(matchSuffix.imageSuffix, fieldName) {
		fieldType = "image"
	} else if matchSuffix.match(matchSuffix.fileSuffix, fieldName) {
		fieldType = "file"
	} else {
		if fieldType == "enum" && matchSuffix.match(matchSuffix.enumRadioSuffix, fieldName) {
			fieldType = "radios"
		}
		if fieldType == "set" && matchSuffix.match(matchSuffix.setCheckboxSuffix, fieldName) {
			fieldType = "checkboxes"
		}
		if strings.Contains(fieldType, "int") && strings.Contains(fieldName, "switch") {
			fieldType = "switch"
		}
		if strings.Contains(fieldType, "int") && strings.Contains(fieldName, "city") {
			fieldType = "city"
		}
		if strings.HasSuffix(fieldType, "text") && matchSuffix.match(matchSuffix.editorSuffix, fieldName) {
			fieldType = "rich-text"
		}
	}

	switch fieldType {
	case "bigint", "int", "mediumint", "smallint", "tinyint":
		inputType = "el-input-number"
		props["controls-position"] = "right"
		props["step-strictly"] = true
		props["step"] = 1
		props["is_number"] = true
	case "set", "checkboxes":
		inputType = "cms-checkbox"
	case "switch":
		inputType = "el-switch"
		props["active-value"] = 1
		props["inactive-value"] = 0
		props["is_number"] = true
	case "enum", "radios":
		inputType = "cms-radio"
	case "decimal", "double", "float":
		inputType = "el-input-number"
		props["controls-position"] = "right"
		props["step-strictly"] = true
		props["precision"] = 2
		props["step"] = 0.01
		props["is_number"] = true
		props["is_float"] = true
	case "longtext", "text", "mediumtext", "smalltext", "tinytext", "rich-text":
		inputType = "cl-editor-quill"
		props["height"] = 350
		props["width"] = "100%"
	case "datetime", "timestamp":
		inputType = "el-date-picker"
		props["type"] = "datetime"
	case "date":
		inputType = "el-date-picker"
		props["type"] = "date"
	case "image":
		inputType = "cl-upload-space"
		props["text"] = "请选择图片"
		if fieldName[len(fieldName)-1:] == "s" {
			props["multiple"] = true
		}
		props["accept"] = ".jpg,.png,.jpeg,.bmp,.gif"
		props["listType"] = "picture-card"
		props["icon"] = "el-icon-picture"
		props["size"] = [2]int{45, 45}
		props["drag"] = true
	case "file":
		inputType = "cl-upload-space"
		props["text"] = "请选择附件"
		props["accept"] = "*"
		props["list-type"] = "text"
		props["drag"] = true
		if fieldName[len(fieldName)-1:] == "s" {
			props["multiple"] = true
		}
	}

	prop, options, comp := parseCommentInfo(col.Name, xormCol.Comment, props["is_number"].(bool))
	labelName = prop

	if len(comp) > 0 {
		inputType = comp // 指定组件
	}

	if len(options) > 0 {
		props["options"] = options
	}

	if inputType == "el-date-picker" { // 检测是否展示为区间
		props["type"] = "daterange"
	}

	return
}

// 返回默认值
func parseCommentInfo(colName, comment string, isNumber bool) (name string, options []map[string]interface{}, customComponent string) {
	if len(comment) == 0 {
		name = colName
		return
	}
	options = []map[string]interface{}{}
	commentInfo := strings.Split(comment, ":") // "状态:0=关闭,1=开启"
	name = commentInfo[0]
	if len(commentInfo) > 1 {
		dict := strings.Split(commentInfo[1], ",")
		for _, v := range dict {
			kv := strings.SplitN(v, "=", 2)
			if len(kv) == 2 && kv[1] != "" {
				opt := map[string]interface{}{
					"key":   strings.TrimSpace(kv[0]),
					"label": kv[1],
				}
				if isNumber {
					opt["key"], _ = strconv.Atoi(opt["key"].(string))
				}
				options = append(options, opt)
			}
		}
	}
	if len(commentInfo) > 2 {
		customComponent = commentInfo[2]
	}
	return
}
