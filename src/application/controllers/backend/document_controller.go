package backend

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/config"
	"html/template"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/xiusin/pinecms/src/application/controllers"

	"github.com/go-xorm/xorm"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

/**
1. 文档模型管理
*/
type DocumentController struct {
	pine.Controller
}

type ModelForm struct {
	ID                string `form:"id" json:"id"`
	intID             int64
	Enabled           string   `form:"enabled" json:"enabled"`
	ModelType         string   `form:"type" json:"type"`
	Name              string   `form:"name" json:"name"`
	Table             string   `form:"table" json:"table"`
	FeTplIndex        string   `form:"tpl_index" json:"tpl_index"`
	FeTplList         string   `form:"tpl_list" json:"tpl_list"`
	FeTplDetail       string   `form:"tpl_detail" json:"tpl_detail"`
	FieldID           []string `form:"field_id" json:"field_id"`
	fieldID           []uint
	FieldDataSource   []string `form:"field_datasource" json:"field_datasource"`
	FieldField        []string `form:"field_field" json:"field_field"`
	FieldSort         []string `form:"field_sort" json:"field_sort"`
	FieldHtml         []string `form:"field_html" json:"field_html"`
	FieldName         []string `form:"field_name" json:"field_name"`
	FieldRequired     []string `form:"field_required" json:"field_required"`
	FieldRequiredTips []string `form:"field_required_tips" json:"field_required_tips"`
	FieldValidator    []string `form:"field_validator" json:"field_validator"`
	FieldDefault      []string `form:"field_default" json:"field_default"`
	FieldType         []string `form:"field_type" json:"field_type"`
	fieldType         []int64
}

var extraFields = []map[string]string{
	{
		"COLUMN_NAME":    "catid",
		"EXTRA":          "",
		"COLUMN_TYPE":    "int",
		"IS_NULLABLE":    "NO",
		"COLUMN_COMMENT": "所属栏目ID",
		"COLUMN_DEFAULT": "0",
	},
	{
		"COLUMN_NAME":    "mid",
		"EXTRA":          "",
		"COLUMN_TYPE":    "int",
		"IS_NULLABLE":    "NO",
		"COLUMN_COMMENT": "模型ID",
		"COLUMN_DEFAULT": "0",
	},
	{
		"COLUMN_NAME":    "created_time",
		"EXTRA":          "",
		"COLUMN_TYPE":    "datetime",
		"IS_NULLABLE":    "YES",
		"COLUMN_COMMENT": "",
		"COLUMN_DEFAULT": "",
	},
	{
		"COLUMN_NAME":    "updated_time",
		"EXTRA":          "",
		"COLUMN_TYPE":    "datetime",
		"IS_NULLABLE":    "YES",
		"COLUMN_COMMENT": "",
		"COLUMN_DEFAULT": "",
	},
	{
		"COLUMN_NAME":    "deleted_time",
		"EXTRA":          "",
		"COLUMN_TYPE":    "datetime",
		"IS_NULLABLE":    "YES",
		"COLUMN_COMMENT": "",
		"COLUMN_DEFAULT": "",
	},
}

func (c *DocumentController) RegisterRoute(b pine.IRouterWrapper) {
	b.GET("/model/list", "ModelList")
	b.ANY("/model/add", "ModelAdd")
	b.ANY("/model/edit", "ModelEdit")
	b.ANY("/model/delete", "ModelDelete")
	b.ANY("/model/list-field-show", "ModelFieldShowInListPage")
	if config.DBConfig().Db.DbDriver == "mysql" {
		b.ANY("/model/gen-sql", "GenSQL")
	} else {
		b.ANY("/model/gen-sql", "GenSQLFromSQLite3")
	}
	b.ANY("/model/preview-page", "PreviewPage")
}

func (c *DocumentController) ModelFieldShowInListPage(icache cache.AbstractCache) {
	mid, _ := c.Ctx().GetInt64("mid")
	if mid < 1 {
		return
	}
	model := models.NewDocumentModel().GetByID(mid)
	if model == nil || model.Id < 1 {
		return
	}
	fields := models.NewDocumentFieldDslModel().GetList(mid)
	var showInPage = map[string]controllers.FieldShowInPageList{}
	if c.Ctx().IsPost() {
		postDatas := c.Ctx().PostData()
		for _, field := range fields {
			_, showExists := postDatas["show_"+field.TableField]
			search := 0
			if _, exists := postDatas["search_"+field.TableField]; exists {
				ssearch := postDatas["search_"+field.TableField][0]
				search, _ = strconv.Atoi(ssearch)
			}
			//_, feSearchExists := postDatas["fe_search_"+field.TableField]
			//if feSearchExists {
			//	fieldArr = append(fieldArr, field.TableField)
			//}
			_, formShowExists := postDatas["form_"+field.TableField]
			//if feSearchExists {
			//	formShow = append(formShow, field.TableField)
			//}
			showInPage[field.TableField] = controllers.FieldShowInPageList{
				Show:   showExists,
				Search: search,
				//FeSearch:  feSearchExists,
				FormShow:  formShowExists,
				Formatter: postDatas["formatter_"+field.TableField][0],
			}
		}
		strs, _ := json.Marshal(showInPage)
		model.FieldShowInList = string(strs)
		//if len(fieldArr) != 0 {
		//	model.FeSearchFields = strings.Join(fieldArr, ",")
		//} else {
		//	model.FeSearchFields = "*"
		//}
		model.Formatters = c.Ctx().PostString("formatters", " ")
		_, err := pine.Make(controllers.ServiceXorm).(*xorm.Engine).Table(&tables.DocumentModel{}).Where("id = ?", mid).Update(model)
		if err != nil {
			helper.Ajax("更新失败:"+err.Error(), 1, c.Ctx())
		} else {
			helper.Ajax("更新字段显示列表成功", 0, c.Ctx())
		}
		return
	}

	_ = json.Unmarshal([]byte(model.FieldShowInList), &showInPage)

	c.Ctx().Render().ViewData("shows", showInPage)
	c.Ctx().Render().ViewData("formShow", showInPage)
	c.Ctx().Render().ViewData("fields", fields)
	c.Ctx().Render().ViewData("l", len(fields))
	c.Ctx().Render().ViewData("codeEditorHeight", len(fields)*49)
	c.Ctx().Render().ViewData("mid", mid)
	c.Ctx().Render().ViewData("model", model)
	c.Ctx().Render().HTML("backend/model_field_show_list_edit.html")
}

func (c *DocumentController) ModelList() {
	page, _ := c.Ctx().GetInt64("page")
	rows, _ := c.Ctx().GetInt64("rows")
	list, total := models.NewDocumentModel().GetList(page, rows)
	helper.Ajax(pine.H{"rows": list, "total": total}, 0, c.Ctx())
}

func (c *DocumentController) ModelAdd() {
	list, _ := models.NewDocumentModelFieldModel().GetList(1, 1000)
	if c.Ctx().IsPost() {
		var data ModelForm
		if err := c.Ctx().BindJSON(&data); err != nil {
			helper.Ajax("表单参数错误: "+err.Error(), 1, c.Ctx())
			return
		}
		data.intID, _ = strconv.ParseInt(data.ID, 10, 64)
		for _, v := range data.FieldType {
			ty, _ := strconv.ParseInt(v, 10, 64)
			data.fieldType = append(data.fieldType, ty)
		}

		//查找重复记录
		exists, err := c.Ctx().Value("orm").(*xorm.Engine).Where("`name`=? or `table`=?", data.Name, data.Table).Exist(&tables.DocumentModel{})
		if exists {
			helper.Ajax("模型名称或者数据表已经存在", 1, c.Ctx())
			return
		}

		// 判断后续字段名称是否一致
		var m = map[string]struct{}{}
		for _, v := range data.FieldName {
			if _, ok := m[v]; ok {
				helper.Ajax("表单名称重复: "+v, 1, c.Ctx())
				return
			} else {
				m[v] = struct{}{}
			}
		}
		m = map[string]struct{}{}
		for _, v := range data.FieldField {
			if _, ok := m[v]; ok {
				helper.Ajax("字段名重复: "+v, 1, c.Ctx())
				return
			} else {
				m[v] = struct{}{}
			}
		}

		var enabled = 0
		if data.Enabled == "on" {
			enabled = 1
		}
		_, err = c.Ctx().Value("orm").(*xorm.Engine).Transaction(func(session *xorm.Session) (i interface{}, err error) {
			dm := &tables.DocumentModel{
				Name:        data.Name,
				Table:       data.Table,
				Enabled:     enabled,
				ModelType:   1,
				FeTplIndex:  helper.EasyUiIDToFilePath(data.FeTplIndex),
				FeTplList:   helper.EasyUiIDToFilePath(data.FeTplList),
				FeTplDetail: helper.EasyUiIDToFilePath(data.FeTplDetail),
			}
			affected, err := session.Insert(dm)
			if affected < 1 {
				if err == nil {
					err = errors.New("保存模型数据失败")
				}
				return nil, err
			}

			// 查找h
			var fieldHtmlsMap = map[int64]*tables.DocumentModelField{}
			for _, field := range list {
				fieldHtmlsMap[field.Id] = field
			}

			var fields []tables.DocumentModelDsl
			for k, name := range data.FieldName {
				listorder, _ := strconv.Atoi(data.FieldSort[k])
				f := tables.DocumentModelDsl{
					Mid:          dm.Id,
					FormName:     name,
					TableField:   data.FieldField[k],
					ListOrder:    int64(listorder),
					FieldType:    data.fieldType[k],
					Datasource:   data.FieldDataSource[k],
					RequiredTips: data.FieldRequiredTips[k],
					Validator:    data.FieldValidator[k],
					Default:      data.FieldDefault[k],
				}
				f.Html = fieldHtmlsMap[data.fieldType[k]].Html
				if strings.HasPrefix(f.Datasource, "[") || strings.HasPrefix(f.Datasource, "{") {
					var dataSourceJson interface{}
					if err := json.Unmarshal([]byte(f.Datasource), &dataSourceJson); err != nil {
						fmt.Println("数据源格式错误", f.Datasource, err)
						return nil, err
					}
				}
				fields = append(fields, f)
			}
			rest, err := session.Insert(fields)
			if rest < int64(len(fields)) {
				if err == nil {
					err = errors.New("批量添加模型字段失败")
				}
				return nil, err
			}
			return true, nil
		})
		if err != nil {
			pine.Logger().Error("添加模型失败", err)
			helper.Ajax("添加模型失败", 1, c.Ctx())
			return
		}
		helper.Ajax("添加模型成功", 0, c.Ctx())
		return
	}
	currentPos := models.NewMenuModel().CurrentPos(64)

	// 传递默认字段给modeladd

	fields := models.NewDocumentFieldDslModel().GetList(0)
	c.Ctx().Render().ViewData("fields", fields)
	c.Ctx().Render().ViewData("fieldslen", len(fields))

	c.Ctx().Render().ViewData("list", list)
	c.Ctx().Render().ViewData("title", currentPos)
	listJson, _ := json.Marshal(list)
	c.Ctx().Render().ViewData("listJson", string(listJson))
	c.Ctx().Render().HTML("backend/model_add.html")
}

func (c *DocumentController) ModelEdit(iCache cache.AbstractCache) {
	list, _ := models.NewDocumentModelFieldModel().GetList(1, 1000)
	if c.Ctx().IsPost() {
		var data ModelForm
		if err := c.Ctx().BindJSON(&data); err != nil {
			helper.Ajax("表单参数错误: "+err.Error(), 1, c.Ctx())
			return
		}

		data.intID, _ = strconv.ParseInt(data.ID, 10, 64)
		for _, v := range data.FieldType {
			ty, _ := strconv.ParseInt(v, 10, 64)
			data.fieldType = append(data.fieldType, ty)
		}
		// 先看模型
		document := models.NewDocumentModel().GetByID(data.intID)
		if document == nil {
			helper.Ajax("模型不存在", 1, c.Ctx())
			return
		}
		//查找重复记录
		exists, err := c.Ctx().Value("orm").(*xorm.Engine).Where("(`name`=? or `table`=?) and id <> ?", data.Name, data.Table, data.ID).Exist(&tables.DocumentModel{})
		if exists {
			helper.Ajax("模型名称或表名已经存在", 1, c.Ctx())
			return
		}

		// 判断后续字段名称是否一致
		var m = map[string]struct{}{}
		for _, v := range data.FieldName {
			if _, ok := m[v]; ok {
				helper.Ajax("表单名称重复: "+v, 1, c.Ctx())
				return
			} else {
				m[v] = struct{}{}
			}
		}
		m = map[string]struct{}{}
		for _, v := range data.FieldField {
			if _, ok := m[v]; ok {
				helper.Ajax("字段名重复: "+v, 1, c.Ctx())
				return
			} else {
				m[v] = struct{}{}
			}
		}

		var enabled = 0
		if data.Enabled == "on" {
			enabled = 1
		}

		_, err = c.Ctx().Value("orm").(*xorm.Engine).Transaction(func(session *xorm.Session) (i interface{}, err error) {
			document.Name = data.Name
			document.Table = data.Table
			document.Enabled = enabled
			document.FeTplIndex = helper.EasyUiIDToFilePath(data.FeTplIndex)
			document.FeTplList = helper.EasyUiIDToFilePath(data.FeTplList)
			document.FeTplDetail = helper.EasyUiIDToFilePath(data.FeTplDetail)
			//document.FieldShowInList = ""
			document.Execed = 0
			_, err = session.ID(document.Id).AllCols().Update(document)
			if err != nil {
				return nil, err
			}
			// 先删除所有字段
			if af, _ := session.Where("mid=?", document.Id).Delete(&tables.DocumentModelDsl{}); af == 0 {
				return nil, errors.New("删除表字段失败")
			}
			var fieldHtmlsMap = map[int64]*tables.DocumentModelField{}
			for _, field := range list {
				fieldHtmlsMap[field.Id] = field
			}

			var fields []tables.DocumentModelDsl
			for k, name := range data.FieldName {
				listorder, _ := strconv.Atoi(data.FieldSort[k])
				f := tables.DocumentModelDsl{
					Mid:          document.Id,
					FormName:     name,
					TableField:   data.FieldField[k],
					FieldType:    data.fieldType[k],
					ListOrder:    int64(listorder),
					Datasource:   data.FieldDataSource[k],
					RequiredTips: data.FieldRequiredTips[k],
					Validator:    data.FieldValidator[k],
					Default:      data.FieldDefault[k],
				}
				f.Html = fieldHtmlsMap[data.fieldType[k]].Html
				if strings.HasPrefix(f.Datasource, "[") || strings.HasPrefix(f.Datasource, "{") {
					var dataSourceJson interface{}
					if err := json.Unmarshal([]byte(f.Datasource), &dataSourceJson); err != nil {
						return nil, err
					}
				}
				if data.FieldRequired[k] == "on" {
					f.Required = 1
				}
				fields = append(fields, f)
			}
			rest, err := session.Insert(fields)
			if rest < int64(len(fields)) {
				if err == nil {
					err = errors.New("批量添加模型字段失败")
				}
				pine.Logger().Error("修改模型", err)
				return nil, err
			}
			return true, nil
		})
		if err != nil {
			helper.Ajax("更新模型失败:"+err.Error(), 1, c.Ctx())
			return
		}
		iCache.Delete(fmt.Sprintf(controllers.CacheDocumentModelPrefix, data.intID))
		helper.Ajax("更新模型成功", 0, c.Ctx())
		return
	}
	mid, err := c.Ctx().GetInt64("mid")
	if err != nil || mid == 0 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	currentPos := models.NewMenuModel().CurrentPos(86)
	// 查找模型信息
	document := models.NewDocumentModel().GetByID(mid)
	if document == nil {
		helper.Ajax("模型不存在或已删除", 1, c.Ctx())
		return
	}
	fields := models.NewDocumentFieldDslModel().GetList(mid)
	c.Ctx().Render().ViewData("fields", fields)
	// 读取系统默认字段
	defaultFields := models.NewDocumentFieldDslModel().GetList(0)
	var conntDelFields []string
	for _, v := range defaultFields {
		conntDelFields = append(conntDelFields, v.TableField)
	}
	c.Ctx().Render().ViewData("dfs", conntDelFields)
	c.Ctx().Render().ViewData("fieldslen", len(fields))
	c.Ctx().Render().ViewData("document", document)
	c.Ctx().Render().ViewData("list", list)
	c.Ctx().Render().ViewData("title", currentPos)
	listJson, _ := json.Marshal(list)

	c.Ctx().Render().ViewData("listJson", string(listJson))
	c.Ctx().Render().HTML("backend/model_edit.html")
}

func (c *DocumentController) ModelDelete() {
	modelID, _ := c.Ctx().GetInt64("id")
	//if modelID < 1 {
	//	helper.Ajax("模型参数错误", 1, c.Ctx())
	//	return
	//}
	model := models.NewDocumentModel()
	dm := model.GetByID(modelID)
	if dm == nil {
		helper.Ajax("模型不存在", 1, c.Ctx())
		return
	}
	if _, err := model.DeleteByID(modelID); err == nil {

		helper.Ajax("删除模型成功", 0, c.Ctx())
	} else {
		helper.Ajax("删除模型失败: "+err.Error(), 1, c.Ctx())
	}
}

var sqlFieldTypeMap = map[string]string{
	"varchar": "varchar(100)",
	"int":     "int(10)",
}

var sqlLite3FieldTypeMap = map[string]string{
	"varchar": "TEXT",
	"int":     "INTEGER",
}

// 生成SQL 传入模型ID
func (c *DocumentController) GenSQL(orm *xorm.Engine) {
	modelID, _ := c.Ctx().GetInt64("mid")
	//if modelID < 1 {
	//	helper.Ajax("模型参数错误", 1, c.Ctx())
	//	return
	//}
	model := models.NewDocumentModel()
	dm := model.GetByID(modelID)
	if dm == nil {
		helper.Ajax("模型不存在", 1, c.Ctx())
		return
	}

	// 如果已经执行过SQL 直接返回一个错误
	if dm.Execed == 1 {
		helper.Ajax("没有任何改动可以执行", 1, c.Ctx())
		return
	}
	//由于执行与SQL显示在同一个控制器内, 所以通过exec区分一下
	exec, _ := c.Ctx().GetBool("exec")
	// 模型字段
	fields := models.NewDocumentFieldDslModel().GetList(modelID)
	// 关联数据
	fieldTypes := models.NewDocumentModelFieldModel().GetMap()
	preg, _ := regexp.Compile("/(.+?)\\?")
	tableSchema := strings.TrimLeft(preg.FindString(orm.DataSourceName()), "/")
	tableSchema = strings.TrimRight(tableSchema, "?")

	var existsFields []map[string]string
	var fieldStrs []string
	querySQL := ""
	tableName := controllers.GetTableName(dm.Table)
	if ok, _ := orm.IsTableExist(tableName); ok {
		querySQL = "ALTER TABLE `" + tableName + "` "
		existsFields, _ = orm.QueryString("select * from information_schema.columns where TABLE_NAME='" + tableName + "' and  table_schema = '" + tableSchema + "'")
		for _, field := range fields {
			var exists bool
			for _, existsField := range existsFields {
				if field.TableField == existsField["COLUMN_NAME"] {
					exists = true
					break
				}
			}
			if !exists {
				colType, ok := sqlFieldTypeMap[fieldTypes[field.FieldType].Type]
				if !ok {
					colType = fieldTypes[field.FieldType].Type
				}
				fieldStrs = append(fieldStrs, fmt.Sprintf("\tADD `%s` %s %s %s %s %s", field.TableField, colType, "", "", "", `COMMENT "`+field.FormName+`"`))
			}
		}
		if len(fieldStrs) > 0 {
			querySQL += "\n" + regexp.MustCompile(" +").ReplaceAllString(strings.Join(fieldStrs, ",\n"), " ")
		} else {
			querySQL = ""
		}
	} else {
		existsFields = append(existsFields, extraFields...)
		querySQL += "CREATE TABLE `" + tableName + "` ( \n"
		querySQL += fmt.Sprintf("\t`%s` %s %s %s %s %s,\n", "id", "int", "NOT NULL", "", "auto_increment", `COMMENT "ID自增字段"`)

		for _, field := range fields {
			colType, ok := sqlFieldTypeMap[fieldTypes[field.FieldType].Type]
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
	}
	querySQL = regexp.MustCompile(" +").ReplaceAllString(querySQL, " ")
	if exec && querySQL != "" {
		_, err := c.Ctx().Value("orm").(*xorm.Engine).Exec(querySQL)
		if err != nil {
			helper.Ajax(err.Error(), 1, c.Ctx())
			return
		}
		af, err := c.Ctx().Value("orm").(*xorm.Engine).ID(modelID).Table(&tables.DocumentModel{}).Update(map[string]interface{}{"execed": 1})
		if af > 0 {
			helper.Ajax("执行SQL成功", 0, c.Ctx())
			return
		}
		helper.Ajax("执行SQL失败", 1, c.Ctx())
	} else {
		helper.Ajax(querySQL, 0, c.Ctx())
	}

}

func (c *DocumentController) GenSQLFromSQLite3(orm *xorm.Engine) {
	modelID, _ := c.Ctx().GetInt64("mid")
	//if modelID < 1 {
	//	helper.Ajax("模型参数错误", 1, c.Ctx())
	//	return
	//}
	model := models.NewDocumentModel()
	dm := model.GetByID(modelID)
	if dm == nil {
		helper.Ajax("模型不存在", 1, c.Ctx())
		return
	}

	// 如果已经执行过SQL 直接返回一个错误
	if dm.Execed == 1 {
		helper.Ajax("没有任何改动可以执行", 1, c.Ctx())
		return
	}
	//由于执行与SQL显示在同一个控制器内, 所以通过exec区分一下
	exec, _ := c.Ctx().GetBool("exec")
	// 模型字段
	fields := models.NewDocumentFieldDslModel().GetList(modelID)
	// 关联数据
	fieldTypes := models.NewDocumentModelFieldModel().GetMap()

	querySQL := ""
	tableName := controllers.GetTableName(dm.Table)
	querySQL += "CREATE TABLE `" + tableName + "` ( \n"
	var createFields []string
	createFields = append(createFields, fmt.Sprintf("\t`%s` %s %s %s %s", "id", "INTEGER", "NOT NULL", "PRIMARY KEY AUTOINCREMENT", ""))

	// 模型字段定义
	for _, field := range fields {
		colType, ok := sqlLite3FieldTypeMap[fieldTypes[field.FieldType].Type]
		if !ok {
			colType = fieldTypes[field.FieldType].Type
		}
		createFields = append(createFields, strings.Trim(fmt.Sprintf("\t`%s` %s %s %s %s", field.TableField, strings.ToUpper(colType), "", "", ""), " "))
	}

	// 额外的字段添加
	for _, f := range extraFields {
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
	}

	querySQL += strings.Join(createFields, ", \n")
	querySQL += "\n);"

	querySQL = regexp.MustCompile(" +").ReplaceAllString(querySQL, " ")
	if exec {
		// 判断表名是否存在
		if querySQL != "" {
			_, err := orm.Transaction(func(sess *xorm.Session) (i interface{}, err error) {
				exist, err := sess.IsTableExist(tableName)
				if err != nil {
					return nil, err
				}
				if exist {
					today := time.Now().Format("20060102150405")
					backTableName := `_` + tableName + `_old_` + today
					// 重命名表
					ret, err := sess.Exec(`ALTER TABLE "` + tableName + `" RENAME TO "` + backTableName + `"`)
					if err != nil {
						return nil, err
					}
					if af, _ := ret.RowsAffected(); af == 0 {
						return nil, errors.New("无法重命名表")
					}
				}
				// 判断表是否存在
				_, err = sess.Exec(querySQL)
				if err != nil {
					return nil, err
				}
				af, _ := sess.ID(modelID).Table(&tables.DocumentModel{}).Update(map[string]interface{}{"execed": 1})
				if af == 0 {
					return nil, errors.New("无法更新模型Execed字段")
				}
				return nil, nil
			})
			if err != nil {
				c.Logger().Error("执行SQL错误", err)
				helper.Ajax("执行SQL失败", 1, c.Ctx())
				return
			}
		}
		helper.Ajax("执行SQL成功, 创建表:"+tableName+"成功", 0, c.Ctx())
	} else {
		helper.Ajax("#SQLite 将原表名备份然后创建一个新的表\n"+querySQL, 0, c.Ctx())
	}
}

// 预览模型表单界面
func (c *DocumentController) PreviewPage() {
	modelID, _ := c.Ctx().GetInt64("mid")
	//if modelID < 1 {
	//	helper.Ajax("模型参数错误", 1, c.Ctx())
	//	return
	//}
	c.Ctx().Render().ViewData("form", template.HTML(buildModelForm(modelID, nil)))
	c.Ctx().Render().ViewData("preview", 1)
	c.Ctx().Render().HTML("backend/model_publish.html")
}
