package backend

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xiusin/pine/di"
	"regexp"
	"strings"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type DocumentController struct {
	sqlFieldTypeMap map[string]string
	extraFields     []map[string]string
	BaseController
}

type table struct {
	Columns      []column      `json:"columns"`
	Props        interface{}   `json:"props"`
	UpsetComps   []interface{} `json:"upset_comps"`
	SearchFields []interface{} `json:"search_fields"`
}

type column struct {
	Prop                string      `json:"prop"`                // 绑定字段
	Label               string      `json:"label"`               // 显示内容
	Component           interface{} `json:"component"`           // 自渲染组件
	Dict                []dictItem  `json:"dict"`                // 字典, 一般针对下拉数据
	Width               uint        `json:"width"`               // 固定宽度
	MinWidth            uint        `json:"minWidth"`            // 最小宽度
	Align               string      `json:"align"`               // 对齐
	Sortable            bool        `json:"sortable"`            // 可排序
	SortBy              interface{} `json:"sortBy"`              // 排序字段
	ShowOverflowTooltip bool        `json:"showOverflowTooltip"` // 溢出自动tooltip
}

type dictItem struct {
	Label string      `json:"label"`
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}

func (c *DocumentController) Construct() {
	c.Group = "系统管理"
	c.SubGroup = "模型管理"
	c.KeywordsSearch = []SearchFieldDsl{
		{Field: "value", Op: "LIKE", DataExp: "%$?%"},
		{Field: "name", Op: "LIKE", DataExp: "%$?%"},
	}
	c.SearchFields = []SearchFieldDsl{
		{Field: "status"},
	}
	c.Table = &tables.DocumentModel{}
	c.Entries = &[]*tables.DocumentModel{}
	c.ApiEntityName = "模型"
	c.sqlFieldTypeMap = map[string]string{
		"varchar": "varchar(100)",
		"int":     "int(10)",
	}
	c.extraFields = []map[string]string{
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

	c.OpBefore = c.before
	c.OpAfter = c.after
	c.BaseController.Construct()
}

func (c *DocumentController) before(act int, params interface{}) error {
	if act == OpDel {
		modelID := params.(*idParams).Ids[0]
		if modelID < 1 {
			return errors.New("模型参数错误")
		}
		if modelID == 1 {
			return errors.New("默认模型不可删除")
		}
	}
	return nil
}

func (c *DocumentController) after(act int, params interface{}) error {
	if act == OpAdd {
		var fields tables.ModelDslFields
		c.Orm.Where("mid = 0").Find(&fields)
		// 生成固定类型的字段
		for _, field := range fields {
			field.Id = 0
			field.Mid = c.Table.(*tables.DocumentModel).Id
			t := tables.LocalTime(time.Now())
			field.UpdatedAt = &t
			_, err := c.Orm.InsertOne(field)
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}
	if act == OpDel { // TODO 删除指定字段列表

	}
	return nil
}

func (c *DocumentController) GetSelect() {
	_ = c.Orm.Find(c.Entries)
	m := c.Entries.(*[]*tables.DocumentModel)
	var kv []tables.KV
	for _, model := range *m {
		kv = append(kv, tables.KV{
			Label: model.Name,
			Value: model.Id,
		})
	}
	helper.Ajax(kv, 0, c.Ctx())
}

func (c *DocumentController) GetSql(orm *xorm.Engine) {
	modelID, _ := c.Ctx().GetInt64("mid")
	model := models.NewDocumentModel()
	dm := model.GetByID(modelID)
	if dm == nil {
		helper.Ajax("模型不存在", 1, c.Ctx())
		return
	}
	dm.Execed = false
	// 如果已经执行过SQL 直接返回一个错误
	if dm.Execed {
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
				colType, ok := c.sqlFieldTypeMap[fieldTypes[field.FieldType].Type]
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
		existsFields = append(existsFields, c.extraFields...)
		querySQL += "CREATE TABLE `" + tableName + "` ( \n"
		querySQL += fmt.Sprintf("\t`%s` %s %s %s %s %s,\n", "id", "int", "NOT NULL", "", "auto_increment", `COMMENT "ID自增字段"`)

		for _, field := range fields {
			colType, ok := c.sqlFieldTypeMap[fieldTypes[field.FieldType].Type]
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
		querySQL += "\tPRIMARY KEY (`id`) USING BTREE) \nENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT \"" +
			strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(dm.Remark, "\"", ""), "`", ""), "\n", "\t") + "\";"

	}
	querySQL = regexp.MustCompile(" +").ReplaceAllString(querySQL, " ")
	if exec && querySQL != "" {
		_, err := di.MustGet(&xorm.Engine{}).(*xorm.Engine).Exec(querySQL)
		if err != nil {
			helper.Ajax(err.Error(), 1, c.Ctx())
			return
		}
		ret, _ := orm.ID(modelID).Table(&tables.DocumentModel{}).Update(map[string]interface{}{"execed": 1})
		if ret > 0 {
			helper.Ajax("执行SQL成功", 0, c.Ctx())
			return
		}
		helper.Ajax("执行SQL失败", 1, c.Ctx())
	} else {
		helper.Ajax(querySQL, 0, c.Ctx())
	}
}

func (c *DocumentController) GetTable() {
	mid, err := c.publicLogic()
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}

	// 系统字段定义
	var fieldDefines []*tables.DocumentModelField

	var fieldDefineMap = map[int64]*tables.DocumentModelField{}

	c.Orm.Find(&fieldDefines)

	for _, define := range fieldDefines {
		fieldDefineMap[define.Id] = define
	}

	var fields tables.ModelDslFields
	c.Orm.Where("mid = ?", mid).
		Asc("listorder").Find(&fields)
	// TODO 允许搜索字段构建
	table := table{Props: nil, Columns: []column{}, UpsetComps: []interface{}{}}
	for _, field := range fields {
		listColumn := column{
			Prop:                field.TableField,
			Label:               field.FormName,
			Width:               field.FieldLen,
			Dict:                nil,
			Sortable:            field.Sortable,
			ShowOverflowTooltip: true,
			Align:               "left",
			MinWidth:            80,
		}
		if field.Center {
			listColumn.Align = "center"
		}
		if field.ListWidth > 80 {
			listColumn.MinWidth = field.ListWidth
		}

		if len(field.Datasource) > 0 {

		}
		if len(field.Component) > 0 {
			var component = map[string]interface{}{}
			if err := json.Unmarshal([]byte(field.Component), &component); err == nil {
				listColumn.Component = component
			} else {
				listColumn.Component = field.Component
			}
		} else if len(fieldDefineMap[field.FieldType].ListComp) > 0 {
			listColumn.Component = fieldDefineMap[field.FieldType].ListComp
		}

		table.Columns = append(table.Columns, listColumn)

		var props = map[string]interface{}{}
		_ = json.Unmarshal([]byte(fieldDefineMap[field.FieldType].Props), &props)
		comp := map[string]interface{}{
			"name":  fieldDefineMap[field.FieldType].FormComp,
			"props": props,
		}
		if field.FieldLen == 0 {
			field.FieldLen = 24
		}

		comp = map[string]interface{}{
			"prop":      field.TableField,
			"label":     field.FormName,
			"span":      field.Span,
			"component": comp,
		}

		if field.Required {
			comp["rules"] = map[string]interface{}{
				"required": true,
				"message":  field.RequiredTips,
			}
		}
		table.UpsetComps = append(table.UpsetComps, comp)
	}
	table.SearchFields = fields.GetSearchableFields()
	helper.Ajax(table, 0, c.Ctx())
}

func (c *DocumentController) publicLogic() (int, error) {
	mid, _ := c.Ctx().GetInt("mid")
	if mid < 1 {
		return 0, errors.New("模型ID错误")
	}

	var model tables.DocumentModel
	c.Orm.Where("id = ?", mid).Get(&model)
	if model.Id == 0 {
		return 0, errors.New("模型不存在")
	}

	if model.Enabled == 0 {
		return 0, errors.New("模型被禁用")
	}
	return mid, nil
}
