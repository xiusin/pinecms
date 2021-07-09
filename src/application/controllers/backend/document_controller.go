package backend

import (
	"errors"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type DocumentController struct {
	BaseController
}

func (c *DocumentController) Construct() {
	c.Group = "系统管理"
	c.SubGroup = "模型管理"
	c.KeywordsSearch = []KeywordWhere{
		{Field: "value", Op: "LIKE", DataExp: "%$?%"},
		{Field: "name", Op: "LIKE", DataExp: "%$?%"},
	}
	c.SearchFields = map[string]searchFieldDsl{
		"status": {Op: "="},
	}
	c.Table = &tables.DocumentModel{}
	c.Entries = &[]*tables.DocumentModel{}
	c.ApiEntityName = "模型"

	c.OpBefore = c.before
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

func (c *DocumentController) PostSelect() {
	c.Orm.Find(c.Entries)
	models := c.Entries.(*[]*tables.DocumentModel)
	var kv []tables.KV
	for _, model := range *models {
		kv = append(kv, tables.KV{
			Label: model.Name,
			Value: model.Id,
		})
	}
	helper.Ajax(kv, 0, c.Ctx())
}
