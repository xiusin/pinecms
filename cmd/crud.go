package cmd

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
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

var crudCmd = &cobra.Command{
	Use:   "crud",
	Short: "生成模块的crud功能",
	Run: func(cmd *cobra.Command, args []string) {
		config.InitDB()
		if !config.Ac().Debug {
			logger.SetReportCaller(false)
			logger.Print("非Debug模式，不支持 CRUD 命令")
			return
		}
		table, _ := cmd.Flags().GetString("table")
		force, _ := cmd.Flags().GetBool("force") // 强制创建
		onlyInfo, _ := cmd.Flags().GetBool("info")
		frontendPath, _ := cmd.Flags().GetString("fepath")
		if len(table) == 0 {
			logger.Print(color.Red.Sprint("请输入要创建CRUD的表名"))
			_ = cmd.Help()
			return
		}
		metas, err := config.XOrmEngine.DBMetas()
		if err != nil {
			panic(err)
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
		if !force {
			for _, s := range []string{frontendPath + "/" + feModuleDir + table, controllerPath, tablePath} {
				if _, err := os.Stat(s); !os.IsNotExist(err) {
					logger.Print("已有存在: " + color.Red.Sprint(s))
					return
				}
			}
		}
		if err = genTableFile(onlyInfo, table, tableDir+table+goExt); err != nil {
			panic(err)
		}
		if err = genControllerFile(onlyInfo, controllerName, table, controllerPath); err != nil {
			panic(err)
		}

		genFrontendFile(table, frontendPath)

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
		byts = bytes.Replace(byts, []byte(holder), []byte(holder+"\r\n\t"+`backendRouter.Handle(new(backend.`+controllerName+`), "/`+snakeString(table)+`")`), 1)
		ioutil.WriteFile(routerFile, byts, os.ModePerm)
		logger.Print("创建模块文件成功, 已注册路由信息至: " + color.Green.Sprint(routerFile))
	},
}

func init() {
	crudCmd.Flags().String("table", "", "数据库表名")
	crudCmd.Flags().Bool("force", false, "是否强制覆盖（可能导致已有代码丢失）")
	crudCmd.Flags().Bool("info", false, "是否只打印生成文件以及操作步骤")
	crudCmd.Flags().String("fepath", "admin", "前端开发根目录")
	rootCmd.AddCommand(crudCmd)
}

func getControllerName(tableName string) (controller string, filename string) {
	controller = camelString(tableName) + "Controller"
	filename = controllerDir + snakeString(tableName) + "_controller" + goExt
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
			camelString(tableName),
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

func genTableFile(print bool, tableName, tablePath string) error {
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

		tableStruct = table.toXorm(print, tableName)
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
func genFrontendFile(table, frontendPath string) {
	moduleBaseDir := frontendPath + "/" + feModuleDir + table
	files := map[string]string{
		filepath.Join(moduleBaseDir, "service", "index.ts"):  serviceTsTpl,
		filepath.Join(moduleBaseDir, "service", "router.ts"): serviceRouterTpl,
		filepath.Join(moduleBaseDir, "views", table+".vue"):  indexVueTpl,
		filepath.Join(moduleBaseDir, "index.ts"):             serviceIndexTsTpl,
	}
	for filename, content := range files {
		os.MkdirAll(filepath.Dir(filename), os.ModePerm)
		if err := os.WriteFile(filename, bytes.ReplaceAll([]byte(content), []byte("[table]"), []byte(snakeString(table))), os.ModePerm); err == nil {
			logger.Print("创建文件： " + color.Green.Sprint(filename))
		} else {
			logger.Print("创建文件", color.Red.Sprint(filename)+"失败")
		}
	}
}

// getFieldType 生成表单和列表字段类型
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
	case "datetime", "timestamp":
		inputType = "datetime"
	case "date":
		inputType = "date"
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
		if strings.Contains(fieldType, "int") && strings.Contains(fieldName, "city") {
			inputType = "city"
		}
		if strings.HasSuffix(fieldType, "text") && matchSuffix.match(matchSuffix.editorSuffix, fieldName) {
			inputType = "rich-text"
		}
	}

	return inputType
}

func parseCommentInfo(comment string) map[string]string {
	vmap := map[string]string{}
	commentInfo := strings.Split(comment, ":") // "状态:0=关闭,1=开启"
	if len(commentInfo) > 1 {
		dict := strings.Split(commentInfo[1], ",")
		for _, v := range dict {
			kv := strings.SplitN(v, "=", 2)
			if len(kv) == 2 && kv[1] != "" {
				vmap[kv[0]] = kv[1]
			}
		}
	}
	return vmap
}

// 生成表单字段额外扩展字段
func getFieldFormExtra(amisType string, item map[string]interface{}) {
	col := cols[item["name"].(string)]
	if !col.DefaultIsEmpty {
		item["value"] = col.Default // 设置字段默认值
	}
	switch amisType {
	case "radios", "checkboxes":
		// 解析备注内容
		vmap := parseCommentInfo(col.Comment)

		var options []map[string]interface{}
		optionSets := col.SetOptions
		if amisType == "radios" {
			optionSets = col.EnumOptions
		}
		if len(optionSets) > 0 { // 设置字段选项
			for setVal := range optionSets {
				label, ok := vmap[setVal]
				if !ok {
					label = setVal
				}
				options = append(options, map[string]interface{}{
					"label": label,
					"value": setVal,
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
