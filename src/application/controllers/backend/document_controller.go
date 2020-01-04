package backend

import (
	"encoding/json"
	"errors"
	"github.com/go-xorm/xorm"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/xiusin/iriscms/src/application/models"
	"github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/iriscms/src/common/helper"
	"html/template"
	"strconv"
	"strings"
)

/**
1. 文档模型管理
*/
type DocumentController struct {
	Ctx     iris.Context
	Orm     *xorm.Engine
	Session *sessions.Session
}

type ModelForm struct {
	ID                string   `form:"id" json:"id"`
	intID             int64    // 赋值ID的int类型
	Enabled           string   `form:"enabled" json:"enabled"`
	Name              string   `form:"name" json:"name"`
	Table             string   `form:"table" json:"table"`
	FeTplIndex        string   `form:"tpl_index" json:"tpl_index"`
	FeTplList         string   `form:"tpl_list" json:"tpl_list"`
	FeTplDetail       string   `form:"tpl_detail" json:"tpl_detail"`
	FieldID           []string `form:"field_id" json:"field_id"`
	fieldID           []uint
	FieldDataSource   []string `form:"field_datasource" json:"field_datasource"`
	FieldField        []string `form:"field_field" json:"field_field"`
	FieldHtml         []string `form:"field_html" json:"field_html"`
	FieldName         []string `form:"field_name" json:"field_name"`
	FieldRequired     []string `form:"field_required" json:"field_required"`
	FieldRequiredTips []string `form:"field_required_tips" json:"field_required_tips"`
	FieldValidator    []string `form:"field_validator" json:"field_validator"`
	FieldType         []string `form:"field_type" json:"field_type"`
	fieldType         []int64
}

func (c *DocumentController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("ANY", "/model/list", "ModelList")
	b.Handle("ANY", "/model/add", "ModelAdd")
	b.Handle("ANY", "/model/edit", "ModelEdit")
	b.Handle("ANY", "/model/delete", "ModelDelete")
}

func (c *DocumentController) ModelList() {
	page, _ := c.Ctx.URLParamInt64("page")
	rows, _ := c.Ctx.URLParamInt64("rows")
	if page > 0 {
		list, total := models.NewDocumentModel(c.Orm).GetList(page, rows)
		c.Ctx.JSON(map[string]interface{}{"rows": list, "total": total})
		return
	}

	menuid, _ := c.Ctx.URLParamInt64("menuid")
	table := helper.Datagrid("model_list_datagrid", "/b/model/list", helper.EasyuiOptions{
		"title":   models.NewMenuModel(c.Orm).CurrentPos(menuid),
		"toolbar": "model_list_datagrid_toolbar",
	}, helper.EasyuiGridfields{
		"模型名称":  {"field": "name", "width": "30", "index": "0"},
		"数据表名称": {"field": "table", "width": "30", "index": "1"},
		"启用":    {"field": "enabled", "width": "25", "index": "2", "formatter": "enabledFormatter"},
		"系统模型":  {"field": "is_system", "width": "25", "index": "3", "formatter": "systemFormatter"},
		"操作":    {"field": "id", "width": "25", "index": "4", "formatter": "optFormatter"},
	})
	c.Ctx.ViewData("dataGrid", template.HTML(table))
	c.Ctx.View("backend/model_list.html")
}

func (c *DocumentController) ModelAdd() {
	list, _ := models.NewDocumentModelFieldModel(c.Orm).GetList(1, 1000)
	if c.Ctx.Method() == "POST" {
		var data ModelForm
		if err := c.Ctx.ReadJSON(&data); err != nil {
			helper.Ajax("表单参数错误: "+err.Error(), 1, c.Ctx)
			return
		}
		data.intID, _ = strconv.ParseInt(data.ID, 10, 64)
		for _, v := range data.FieldType {
			ty, _ := strconv.ParseInt(v, 10, 64)
			data.fieldType = append(data.fieldType, ty)
		}

		//查找重复记录
		exists, err := c.Orm.Where("`name`=? or `table`=?", data.Name, data.Table).Exist(&tables.IriscmsDocumentModel{})
		if exists {
			helper.Ajax("模型名称换货数据表已经存在", 1, c.Ctx)
			return
		}

		// 判断后续字段名称是否一致
		var m = map[string]struct{}{}
		for _, v := range data.FieldName {
			if _, ok := m[v]; ok {
				helper.Ajax("表单名称重复: "+v, 1, c.Ctx)
				return
			} else {
				m[v] = struct{}{}
			}
		}
		m = map[string]struct{}{}
		for _, v := range data.FieldField {
			if _, ok := m[v]; ok {
				helper.Ajax("字段名重复: "+v, 1, c.Ctx)
				return
			} else {
				m[v] = struct{}{}
			}
		}

		var enabled = 0
		if data.Enabled == "on" {
			enabled = 1
		}
		_, err = c.Orm.Transaction(func(session *xorm.Session) (i interface{}, err error) {
			documentModel := &tables.IriscmsDocumentModel{
				Name:        data.Name,
				Table:       data.Table,
				Enabled:     enabled,
				IsSystem:    0,
				ModelType:   0,
				FeTplIndex:  helper.EasyUiIDToFilePath(data.FeTplIndex),
				FeTplList:   helper.EasyUiIDToFilePath(data.FeTplList),
				FeTplDetail: helper.EasyUiIDToFilePath(data.FeTplDetail),
			}
			affected, err := session.Insert(documentModel)
			if affected < 1 {
				if err == nil {
					err = errors.New("保存模型数据失败")
				}
				return nil, err
			}

			// 查找h
			var fieldHtmlsMap = map[int64]*tables.IriscmsDocumentModelField{}
			for _, field := range list {
				fieldHtmlsMap[field.Id] = field
			}

			var fields []tables.IriscmsDocumentModelDsl
			for k, name := range data.FieldName {
				f := tables.IriscmsDocumentModelDsl{
					Mid:          documentModel.Id,
					FormName:     name,
					TableField:   data.FieldField[k],
					FieldType:    data.fieldType[k],
					Datasource:   data.FieldDataSource[k],
					RequiredTips: data.FieldRequiredTips[k],
					Validator:    data.FieldValidator[k],
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
				return nil, err
			}
			return true, nil
		})
		if err != nil {
			golog.Error("添加模型失败", err)
			helper.Ajax("添加模型失败", 1, c.Ctx)
			return
		}
		helper.Ajax("添加模型成功", 0, c.Ctx)
		return
	}
	currentPos := models.NewMenuModel(c.Orm).CurrentPos(64)

	c.Ctx.ViewData("list", list)
	c.Ctx.ViewData("title", currentPos)
	listJson, _ := json.Marshal(list)
	c.Ctx.ViewData("listJson", string(listJson))
	c.Ctx.View("backend/model_add.html")
}

func (c *DocumentController) ModelEdit() {
	list, _ := models.NewDocumentModelFieldModel(c.Orm).GetList(1, 1000)
	if c.Ctx.Method() == "POST" {
		var data ModelForm
		if err := c.Ctx.ReadJSON(&data); err != nil {
			helper.Ajax("表单参数错误: "+err.Error(), 1, c.Ctx)
			return
		}

		data.intID, _ = strconv.ParseInt(data.ID, 10, 64)
		for _, v := range data.FieldType {
			ty, _ := strconv.ParseInt(v, 10, 64)
			data.fieldType = append(data.fieldType, ty)
		}

		// 先看模型
		document := models.NewDocumentModel(c.Orm).GetByID(data.intID)
		if document == nil || document.Id < 1 {
			helper.Ajax("模型不存在", 1, c.Ctx)
			return
		}
		//查找重复记录
		exists, err := c.Orm.Where("(`name`=? or `table`=?) and id <> ?", data.Name, data.Table, data.ID).Exist(&tables.IriscmsDocumentModel{})
		if exists {
			helper.Ajax("模型名称或表名已经存在", 1, c.Ctx)
			return
		}

		// 判断后续字段名称是否一致
		var m = map[string]struct{}{}
		for _, v := range data.FieldName {
			if _, ok := m[v]; ok {
				helper.Ajax("表单名称重复: "+v, 1, c.Ctx)
				return
			} else {
				m[v] = struct{}{}
			}
		}
		m = map[string]struct{}{}
		for _, v := range data.FieldField {
			if _, ok := m[v]; ok {
				helper.Ajax("字段名重复: "+v, 1, c.Ctx)
				return
			} else {
				m[v] = struct{}{}
			}
		}

		var enabled = 0
		if data.Enabled == "on" {
			enabled = 1
		}
		_, err = c.Orm.Transaction(func(session *xorm.Session) (i interface{}, err error) {
			document.Name = data.Name
			document.Table = data.Table
			document.Enabled = enabled
			document.FeTplIndex = helper.EasyUiIDToFilePath(data.FeTplIndex)
			document.FeTplList = helper.EasyUiIDToFilePath(data.FeTplList)
			document.FeTplDetail = helper.EasyUiIDToFilePath(data.FeTplDetail)
			session.Update(document)
			// 先删除所有字段
			if models.NewDocumentFieldDslModel(c.Orm).DeleteByMID(document.Id) == false {
				return nil, errors.New("删除表字段失败")
			}
			var fieldHtmlsMap = map[int64]*tables.IriscmsDocumentModelField{}
			for _, field := range list {
				fieldHtmlsMap[field.Id] = field
			}

			var fields []tables.IriscmsDocumentModelDsl
			for k, name := range data.FieldName {
				f := tables.IriscmsDocumentModelDsl{
					Mid:          document.Id,
					FormName:     name,
					TableField:   data.FieldField[k],
					FieldType:    data.fieldType[k],
					Datasource:   data.FieldDataSource[k],
					RequiredTips: data.FieldRequiredTips[k],
					Validator:    data.FieldValidator[k],
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
				golog.Error("修改模型", err)
				return nil, err
			}
			return true, nil
		})
		if err != nil {
			helper.Ajax("更新模型失败:"+err.Error(), 1, c.Ctx)
			return
		}
		helper.Ajax("更新模型成功", 0, c.Ctx)
		return
	}
	mid, err := c.Ctx.URLParamInt64("mid")
	if err != nil || mid == 0 {
		helper.Ajax("参数错误", 1, c.Ctx)
		return
	}
	currentPos := models.NewMenuModel(c.Orm).CurrentPos(64)
	// 查找模型信息
	document := models.NewDocumentModel(c.Orm).GetByID(mid)
	if document == nil || document.Id < 1 {
		helper.Ajax("模型不存在或已删除", 1, c.Ctx)
		return
	}
	fields := models.NewDocumentFieldDslModel(c.Orm).GetList(mid)
	c.Ctx.ViewData("fields", fields)
	c.Ctx.ViewData("fieldslen", len(fields))
	c.Ctx.ViewData("document", document)
	c.Ctx.ViewData("list", list)
	c.Ctx.ViewData("title", currentPos)
	listJson, _ := json.Marshal(list)
	golog.Error(string(listJson))

	c.Ctx.ViewData("listJson", string(listJson))
	c.Ctx.View("backend/model_edit.html")
}

func (c *DocumentController) ModelDelete() {
	modelID, _ := c.Ctx.URLParamInt64("id")
	if modelID < 1 {
		helper.Ajax("模型参数错误", 1, c.Ctx)
		return
	}
	model := models.NewDocumentModel(c.Orm)
	dm := model.GetByID(modelID)
	if dm == nil {
		helper.Ajax("模型不存在", 1, c.Ctx)
		return
	}
	if _, err := model.DeleteByID(modelID); err == nil {
		helper.Ajax("删除模型成功", 0, c.Ctx)
	} else {
		helper.Ajax("删除模型失败: "+err.Error(), 1, c.Ctx)
	}
}
