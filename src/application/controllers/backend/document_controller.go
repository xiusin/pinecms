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
	ID                uint     `form:"id"`
	Enabled           string   `form:"enabled"`
	Name              string   `form:"name"`
	Table             string   `form:"table"`
	FeTplIndex        string   `form:"tpl_index"`
	FeTplList         string   `form:"tpl_list"`
	FeTplDetail       string   `form:"tpl_detail"`
	FieldID           []uint   `form:"field_id"`
	FieldDataSource   []string `form:"field_datasource"`
	FieldField        []string `form:"field_field"`
	FieldHtml         []string `form:"field_html"`
	FieldName         []string `form:"field_name"`
	FieldRequired     []string `form:"field_required"`
	FieldRequiredTips []string `form:"field_required_tips"`
	FieldValidator    []string `form:"field_validator"`
	FieldType         []int    `form:"field_type"`
}

func (c *DocumentController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("ANY", "/model/list", "ModelList")
	b.Handle("ANY", "/model/add", "ModelAdd")
	b.Handle("ANY", "/model/edit", "ModelEdit")
	b.Handle("POST", "/model/delete", "ModelDelete")
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
		"模型名称": {"field": "name", "width": "30", "index": "0"},
		"模型ID": {"field": "table", "width": "30", "index": "1"},
		"启用":   {"field": "enabled", "width": "25", "index": "2", "formatter": "enabledFormatter"},
		"操作":   {"field": "id", "width": "25", "index": "3", "formatter": "optFormatter"},
	})
	c.Ctx.ViewData("dataGrid", template.HTML(table))
	c.Ctx.View("backend/model_list.html")
}

func (c *DocumentController) ModelAdd() {
	if c.Ctx.Method() == "POST" {
		var data ModelForm
		if err := c.Ctx.ReadForm(&data); err != nil {
			helper.Ajax(err.Error(), 1, c.Ctx)
			return
		}
		var enabled = 0
		if data.Enabled == "on" {
			enabled = 1
		}
		_, err := c.Orm.Transaction(func(session *xorm.Session) (i interface{}, err error) {
			// 组建数据
			documentModel := &tables.IriscmsDocumentModel{
				Name:        data.Name,
				Enabled:     enabled,
				IsSystem:    0,
				ModelType:   0,
				FeTplIndex:  data.FeTplIndex,
				FeTplList:   data.FeTplList,
				FeTplDetail: data.FeTplDetail,
			}
			id, err := session.InsertOne(documentModel)
			if id < 1 {
				if err == nil {
					err = errors.New("保存模型数据失败")
				}
				return nil, err
			}

			var fields []tables.IriscmsDocumentModelDsl

			for k, name := range data.FieldName {
				f := tables.IriscmsDocumentModelDsl{
					Mid:      id,
					FormName: name,
				}
				if len(data.FieldHtml) >= k + 1 {
					f.Html = data.FieldHtml[k]
				}
				if len(data.FieldDataSource) >= k + 1 {
					f.Html = data.FieldDataSource[k]
				}
				if len(data.FieldRequiredTips) >= k + 1 {
					f.Html = data.FieldRequiredTips[k]
				}
				if len(data.FieldValidator) >= k + 1 {
					f.Html = data.FieldValidator[k]
				}
				if len(data.FieldRequired) >= k + 1 && data.FieldRequired[k] == "on" {
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
		helper.Ajax(&data, 0, c.Ctx)
		return
	}

	currentPos := models.NewMenuModel(c.Orm).CurrentPos(64)
	// 查找数据库模型字段
	list, _ := models.NewDocumentModelFieldModel(c.Orm).GetList(1, 1000)
	c.Ctx.ViewData("list", list)
	c.Ctx.ViewData("title", currentPos)
	b, _ := json.Marshal(list)
	c.Ctx.ViewData("listJson", string(b))
	c.Ctx.View("backend/model_add.html")
}

func (c *DocumentController) ModelEdit() {
	c.Ctx.View("backend/model_edit.html")
}

func (c *DocumentController) ModelDelete() {
	c.Ctx.WriteString("hello model111")
}
