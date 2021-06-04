package cmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/gookit/color"
	"github.com/schollz/progressbar"
	"github.com/spf13/cobra"
	"github.com/xiusin/logger"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/config"
	"regexp"
	"strconv"
	"strings"
	"time"
	"xorm.io/core"
)

/**
注意: 这个代码隔一天再看已经看不懂了!
*/

var (
	dc             = config.DBConfig()
	bar            = progressbar.New(100)
	dedeOrm        *xorm.Engine
	tableSchema    string
	pineOrm        *xorm.Engine
	tableFieldMaps = map[string]string{
		"click":    "visit_count",
		"body":     "content",
		"litpic":   "thumb",
		"writer":   "author",
		"source":   "from_url",
		"sortrank": "listorder",
	}
)
var dedeCmd = &cobra.Command{
	Use:   "dede",
	Short: "导入DEDECMS文章数据",
	Long: `
1. 导出广告位以及广告
2. 导入模型,自动分析模型字段以及映射关系(模型完全由导入决定. 只是替换一下必须字段为pinecms)
1. 根据模型创建表字段以及模型内容
2. 导入织梦文档表, 只匹配已知可对应字段 
`,
	// todo 尽最大可能保留非默认字段数据
	Run: func(cmd *cobra.Command, args []string) {
		dsn, _ := cmd.Flags().GetString("dsn")
		if len(dsn) == 0 {
			_ = cmd.Usage()
			return
		}
		initORM(dsn)
		importChannelType()
		importArcType()
	},
}

func init() {
	importCmd.AddCommand(dedeCmd)
	dedeCmd.Flags().String("dsn", "root:@tcp(127.0.0.1:3306)/dedecms?charset=utf8", "输入要导入数据的数据库连接DSN信息(只支持MYSQL)")

}

func initORM(dsn string) {
	_dedeOrm, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		logger.Error("连接DEDE数据库失败", err)
		return
	}
	preg, _ := regexp.Compile("/(.+?)\\?")
	tableSchema = strings.TrimRight(strings.TrimLeft(preg.FindString(dsn), "/"), "?")
	dedeOrm = _dedeOrm
	_orm, err := xorm.NewEngine(dc.Db.DbDriver, dc.Db.Dsn)
	if err != nil {
		logger.Error("连接PINECMS数据库失败", err)
		return
	}
	_orm.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, dc.Db.DbPrefix))
	bar.Reset()
	_ = bar.Add(5)
	pineOrm = _orm
}

type DedeXMLField struct {
	Field     string
	ItemName  string `xml:"itemname,attr"`
	Autofield int    `xml:"autofield,attr"`
	NotSend   int    `xml:"notsend,attr"`
	Type      string `xml:"type,attr"`
	Isnull    string `xml:"isnull,attr"`
	IsList    int    `xml:"islist,attr"`
	Default   string `xml:"default,attr"`
	Maxlength int    `xml:"maxlength,attr"`
}

var dedePineFieldMaps = map[string]int64{"text": 1, "textchar": 1, "multitext": 2, "htmltext": 3, "textdata": 3, "int": 9, "float": 10, "datetime": 14, "img": 11, "imgfile": 11, "media": 4, "addon": 4, "select": 5, "radio": 7, "checkbox": 8, "stepselect": 6,}

func clearDocumentModelTable(table, dslTable string) error {

	// 删除所有分类
	_, _ = pineOrm.Exec(fmt.Sprintf("DELETE FROM `%s` WHERE 1=1", table))
	_, _ = pineOrm.Exec(fmt.Sprintf("DELETE FROM `%s` WHERE mid != 0", dslTable))
	return nil
}

func createModelTable(table string, channel map[string]string) *tables.DocumentModel {
	data := tables.DocumentModel{}
	// 解析模型
	data.Id, _ = strconv.ParseInt(channel["id"], 10, 64)
	if channel["addtable"] == "dede_addonarticle" {
		data.Table = "articles"
	} else {
		data.Table = strings.Replace(channel["addtable"], "dede_addon", "", 1)
		if regexp.MustCompile("^\\d+$").MatchString(data.Table) {
			data.Table = channel["nid"]
		}
	}
	data.Name = channel["typename"]
	data.Enabled, _ = strconv.Atoi(channel["isshow"])
	data.Execed = 1
	pineOrm.Table(table).InsertOne(&data)

	return &data
}

func importChannelType() {
	var table = dc.Db.DbPrefix + "document_model"
	var dslTable = dc.Db.DbPrefix + "document_model_dsl"
	_ = clearDocumentModelTable(table, dslTable)
	channels, _ := dedeOrm.QueryString("SELECT * FROM dede_channeltype")
	pre := 70 / len(channels)
	for _, channel := range channels {
		data := createModelTable(table, channel)
		modelTableName := dc.Db.DbPrefix + data.Table
		// 迭代读取模型内容并入库. 注意映射字段
		_, _ = pineOrm.Exec("DROP TABLE  " + dc.Db.DbPrefix + data.Table)
		// 附表额外添加字段
		f := regexp.MustCompile(">[^>]+</field").ReplaceAllString(channel["fieldset"], "></field")
		// 解析附表自定义字段
		dedeFields := strings.Split(f, "\n")
		// 选出默认字段 model = 0
		var defaultFields []tables.DocumentModelDsl
		pineOrm.Where("mid = 0").Find(&defaultFields)
		dmField := &tables.DocumentModelField{Id: 11}
		defaultFields = append(defaultFields, tables.DocumentModelDsl{
			Mid:        data.Id,
			FieldType:  11,
			ListOrder:  7,
			FormName:   "缩略图",
			TableField: "thumb",
			Html:       dmField.Html,
		})
		// pinecms默认字段
		var modelFields = map[string]string{}
		for _, v := range defaultFields {
			v.Id = 0
			v.Mid = data.Id
			pineOrm.InsertOne(&v)
			modelFields[v.TableField] = ""
		}
		// 检查是否有与默认字段冲突的字段 额外字段
		autoIndex := 7
		for k, field := range dedeFields {
			if regexp.MustCompile("<field*").MatchString(field) {
				autoIndex++
				var f DedeXMLField
				err := xml.Unmarshal([]byte(field), &f)
				if err != nil {
					panic(err)
				}
				f.Field = regexp.MustCompile("<field:([^\\s]+) ").FindStringSubmatch(field)[1]
				if mf, ok := tableFieldMaps[f.Field]; ok {
					f.Field = mf
				}
				if _, ok := modelFields[f.Field]; ok { //默认字段忽略
					continue
				}
				dmField := &tables.DocumentModelField{Id: dedePineFieldMaps[f.Type],}
				pineOrm.Get(dmField)
				fieldDsl := &tables.DocumentModelDsl{
					Mid:        data.Id,
					FieldType:  dmField.Id,
					FormName:   f.ItemName,
					TableField: f.Field,
					ListOrder:  int64(k),
					Html:       dmField.Html,
					Default:    f.Default,
				}
				dataSource := f.Default
				if dmField.Id == 5 || dmField.Id == 7 || dmField.Id == 8 {
					defs := strings.Split(f.Default, ",")
					byt, _ := json.Marshal(defs)
					dataSource = string(byt)
					fieldDsl.Default = defs[0]
				}
				fieldDsl.ListOrder = int64(autoIndex)
				fieldDsl.Datasource = dataSource
				id, err := pineOrm.InsertOne(fieldDsl)
				if id == 0 {
					panic(fmt.Sprintf("插入模型失败: %d: %s", id, err))
				}
				modelFields[f.Field] = ""
			}
		}

		var fields []tables.DocumentModelDsl
		pineOrm.Table(&tables.DocumentModelDsl{}).Where("mid = ?", data.Id).Find(&fields)
		var list []*tables.DocumentModelField
		var mapList = map[int64]*tables.DocumentModelField{}
		_ = pineOrm.Find(&list)
		for _, v := range list {
			mapList[v.Id] = v
		}
		GenSQLFromSQLite3(data.Id, channel["maintable"], modelTableName, modelFields, fields, mapList)
		transDocument(modelTableName, data, modelFields, channel)
		bar.Add(pre)
		time.Sleep(time.Millisecond * 100)
	}
}

func transDocument(modelTableName string, data *tables.DocumentModel, modelFields, channel map[string]string) {
	var querySql string
	if channel["maintable"] != channel["addtable"] {
		querySql = fmt.Sprintf("SELECT *, %s.typeid as catid  FROM %s LEFT JOIN %s ON %s.id=%s.aid WHERE %s.channel=%d", channel["maintable"], channel["maintable"], channel["addtable"], channel["maintable"], channel["addtable"], channel["maintable"], data.Id)
	} else {
		querySql = fmt.Sprintf("SELECT * FROM %s WHERE channel=%d", channel["maintable"], data.Id)
	}
	archives, err := dedeOrm.QueryString(querySql)
	if err != nil {
		panic(err.Error() + ": " + querySql)
	}
	for _, archive := range archives {
		var fs []string
		var vs []interface{}
	OuterLoop:
		for field, val := range archive {
			switch field {
			case "arcrank":
				status := "1"
				if val == "-2" {
					status = "0"
				}
				field = "status"
				val = status
			case "pubdate":
				pubdate, _ := strconv.ParseInt(archive["pubdate"], 10, 64)
				if pubdate != 0 {
					val = time.Unix(pubdate, 0).Format("2006-01-02 15:04:05")
				}
				field = "pubtime"
			case "senddate":
				senddate, _ := strconv.ParseInt(archive["senddate"], 10, 64)
				if senddate != 0 {
					val = time.Unix(senddate, 0).Format("2006-01-02 15:04:05")
				}
				field = "created_time"
			default:
				for k, v := range tableFieldMaps {
					if field == k {
						field = v
						break
					}
				}
			}
			// 判断是否包含在表字段内
			if _, ok := modelFields[field]; !ok {
				continue OuterLoop
			}
			fs = append(fs, "`"+field+"`")
			vs = append(vs, val)
		}
		placeholders := strings.TrimRight(strings.Repeat("?,", len(fs)), ",")
		// 入库
		vs = append([]interface{}{"INSERT INTO `" + modelTableName + "` (" + strings.Join(fs, ", ") + ") VALUES (" + placeholders + ");"}, vs...)
		_, err := pineOrm.Exec(vs...)
		if err != nil {
			panic(err)
		}
	}
}

func importArcType() {
	pineOrm.Where("1=1").Delete(&tables.Category{})
	dedeQuerySql := `SELECT * FROM dede_arctype`
	arctypes, err := dedeOrm.QueryString(dedeQuerySql)
	if len(arctypes) == 0 {
		logger.Error("读取织梦分类数据为空", err)
		return
	}
	placeholders := strings.TrimRight(strings.Repeat("?,", 15), ",")
	for _, arctype := range arctypes {
		ismenu := 0
		if arctype["ishidden"] == "0" {
			ismenu = 1
		}
		pineDir, url := "", ""
		if arctype["ispart"] == "2" {
			url = arctype["typedir"]
		} else {
			typedir := strings.Split(arctype["typedir"], "/")
			l := len(typedir)
			pineDir = strings.Join(typedir[l-1:], "")
		}

		arctype["templist"] = regexp.MustCompile("[^/]+/(.+)?.html?").ReplaceAllString(arctype["templist"], "$1.jet")
		arctype["temparticle"] = regexp.MustCompile("[^/]+/(.+)?.html?").ReplaceAllString(arctype["temparticle"], "$1.jet")

		res, err := pineOrm.Exec("INSERT INTO `"+dc.Db.DbPrefix+"category` ("+
			"catid, type, parentid, topid, model_id, catname, keywords, description, url, listorder, dir, thumb, ismenu"+
			",list_tpl, detail_tpl) VALUES ("+placeholders+");",
			arctype["id"], arctype["ispart"],
			arctype["reid"], arctype["topid"], arctype["channeltype"], arctype["typename"],
			arctype["keywords"], arctype["description"], url, arctype["sortrank"],
			pineDir, "", ismenu, arctype["templist"], arctype["temparticle"], )
		if err != nil {
			panic(err)
		}
		if lid, _ := res.LastInsertId(); lid == 0 {
			panic("插入分类数据失败")
		}

		if arctype["ispart"] == "1" || len(arctype["content"]) > 0 {
			tid, _ := strconv.Atoi(arctype["id"])
			pineOrm.InsertOne(&tables.Page{
				Catid:       int64(tid),
				Title:       arctype["typename"],
				Keywords:    arctype["keywords"],
				Description: arctype["description"],
				Content:     arctype["content"],
				Updatetime:  time.Now().In(helper.GetLocation()).Unix(),
			})
		}
	}
	bar.Reset()
	bar.Add(100)
	fmt.Println(color.Green.Sprint(`

SUCCESS!

1. 导入织梦数据成功, 请修改对应模板文件到pinecms
2. 配置模型字段属性和显隐性
3. 使用标签开发模板
4. Enjoy! 😃

%s`, color.Red.Sprint("注意: 导入不保证完全正确,建议进行模型设置(固化字段被设置为text类型)")))
}

func importAd() {
	// 导入广告
	dedeOrm.Table("dede_")
}

// 检测时也需要把cms之间的映射字段辅助添加上
var excludeDedeFields = []string{"id", "typeid", "typeid2", "flag", "ismake", "arcrank", "channel", "click", "title", "color", "litpic", "pubdate", "senddate", "mid", "keywords", "lastpost", "goodpost", "badpost", "voteid", "notpost", "description", "filename", "dutyadmin", "tackid", "mtype", "weight",}

func needSkip(field string) bool {
	for _, v := range excludeDedeFields {
		if v == field {
			return true
		}
	}
	for k := range tableFieldMaps {
		if k == field {
			return true
		}
	}
	return false
}

// 生成SQL 传入模型ID
func GenSQL(tableName string, hasFields map[string]string, fields []tables.DocumentModelDsl, fieldTypes map[int64]*tables.DocumentModelField) {
	var existsFields []map[string]string
	querySQL := ""
	existsFields = append(existsFields, ExtraFields...)
	querySQL += "CREATE TABLE `" + tableName + "` ( \n"
	querySQL += fmt.Sprintf("\t`%s` %s %s %s %s %s,\n", "id", "int", "NOT NULL", "", "auto_increment", `COMMENT "ID自增字段"`)

	for _, field := range fields {
		colType, ok := SqlFieldTypeMap[fieldTypes[field.FieldType].Type]
		if !ok {
			colType = fieldTypes[field.FieldType].Type
		}
		querySQL += fmt.Sprintf("\t`%s` %s %s %s %s %s,\n", field.TableField, strings.ToUpper(colType), "", "", "", `COMMENT "`+field.FormName+`"`)
	}

	for _, f := range existsFields {
		var notNull = ""
		if f["IS_NULLABLE"] == "NO" {
			notNull = "NOT NULL"
		}
		var defaultVal = ""
		if f["COLUMN_DEFAULT"] != "" {
			defaultVal = "DEFAULT '" + f["COLUMN_DEFAULT"] + "'"
		}
		querySQL += fmt.Sprintf("\t`%s` %s %s %s %s %s,\n", f["COLUMN_NAME"], strings.ToUpper(f["COLUMN_TYPE"]), notNull, defaultVal, f["EXTRA"], `COMMENT "`+f["COLUMN_COMMENT"]+`"`)
	}
	querySQL += "\tPRIMARY KEY (`id`) USING BTREE) ENGINE=InnoDB DEFAULT CHARSET=utf8;"
	querySQL = regexp.MustCompile(" +").ReplaceAllString(querySQL, " ")
	pineOrm.Exec(querySQL)
}

func GenSQLFromSQLite3(mid int64, dedeMainTable, tableName string, hasFields map[string]string, fields []tables.DocumentModelDsl, fieldTypes map[int64]*tables.DocumentModelField) {
	querySQL := ""
	querySQL += "CREATE TABLE `" + tableName + "` ( \n"
	var createFields []string
	createFields = append(createFields, fmt.Sprintf("\t`%s` %s %s %s %s", "id", "INTEGER", "NOT NULL", "PRIMARY KEY AUTOINCREMENT", ""))
	hasFields["id"] = ""
	mainTableFields, _ := dedeOrm.QueryString("select * from information_schema.columns where TABLE_NAME='" + dedeMainTable + "' and  table_schema = '" + tableSchema + "'")
	autoIndex := 70
	for _, mtf := range mainTableFields {
		if _, ok := hasFields[mtf["COLUMN_NAME"]]; ok {
			continue
		}
		if !needSkip(mtf["COLUMN_NAME"]) { // 手动添加的字段需要拷贝数据
			var dmf = &tables.DocumentModelField{}
			switch mtf["DATA_TYPE"] {
			case "char", "varchar": // 使用textbox
				dmf.Id = 1
			case "int", "tinyint", "bigint", "mediumint", "integer": // 使用 int
				dmf.Id = 9
			case "float", "double", "decimal", "real": //使用float
				dmf.Id = 10
			default: //使用text
				dmf.Id = 2
			}
			dedeOrm.Get(&dmf)
			autoIndex++
			dsl := tables.DocumentModelDsl{
				Mid:        mid,
				FieldType:  dmf.Id,
				FormName:   mtf["COLUMN_NAME"],
				TableField: mtf["COLUMN_NAME"],
				Html:       dmf.Html,
				ListOrder:  int64(autoIndex),
			}
			pineOrm.InsertOne(&dsl)
			fields = append(fields, dsl)
			hasFields[mtf["COLUMN_NAME"]] = ""
		}
	}

	// 根据现有的模型定义创建表
	for _, field := range fields {
		colType, ok := SqlLite3FieldTypeMap[fieldTypes[field.FieldType].Type]
		if !ok {
			colType = fieldTypes[field.FieldType].Type
		}
		createFields = append(createFields, strings.Trim(fmt.Sprintf("\t`%s` %s %s %s %s", field.TableField, strings.ToUpper(colType), "", "", ""), " "))
	}

	// 内置字段固定添加
	for _, f := range ExtraFields {
		if _, ok := hasFields[f["COLUMN_NAME"]]; ok {
			continue
		}
		var notNull = ""
		if f["IS_NULLABLE"] == "NO" {
			notNull = "NOT NULL"
		}
		var defaultVal = ""
		if f["COLUMN_DEFAULT"] != "" {
			defaultVal = "DEFAULT " + f["COLUMN_DEFAULT"]
		}
		if f["COLUMN_TYPE"] == "int" {
			f["COLUMN_TYPE"] = "INTEGER"
		}
		createFields = append(createFields, strings.Trim(fmt.Sprintf("\t`%s` %s %s %s %s", f["COLUMN_NAME"], strings.ToUpper(f["COLUMN_TYPE"]), notNull, f["EXTRA"], defaultVal), " "))
		hasFields[f["COLUMN_NAME"]] = ""
	}
	querySQL += strings.Join(createFields, ", \n")
	querySQL += "\n);"
	querySQL = regexp.MustCompile(" +").ReplaceAllString(querySQL, " ")
	_, err := pineOrm.Exec(querySQL)
	if err != nil {
		panic(err.Error() + ":" + dedeMainTable)
	}
}
