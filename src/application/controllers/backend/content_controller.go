package backend

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type ContentController struct {
	BaseController
}

func (c *ContentController) Construct() {
	c.Group = "内容管理"
	c.KeywordsSearch = []KeywordWhere{
		{Field: "value", Op: "LIKE", DataExp: "%$?%"},
		{Field: "name", Op: "LIKE", DataExp: "%$?%"},
	}
	c.SearchFields = map[string]searchFieldDsl{
		"status": {Op: "="},
	}
	c.Entries = &[]*tables.DocumentModel{}
	c.ApiEntityName = "内容"
	c.BaseController.Construct()
}

func (c *ContentController) PostList() {
	var catid int64 = c.Input().GetInt64("cid")
	var category tables.Category
	c.Orm.Where("id = ?", catid).Get(&category)
	if category.Catid == 0 {
		helper.Ajax("栏目不存在或已删除", 1, c.Ctx())
		return
	}
	if category.ModelId < 1 {
		helper.Ajax("栏目模型不存在", 1, c.Ctx())
		return
	}
	var document tables.DocumentModel
	c.Orm.Where("id = ?", category.ModelId).Get(&document)
	if document.Id == 0 {
		helper.Ajax("无法找到关联模型", 1, c.Ctx())
		return
	}
	c.Table = controllers.GetTableName(document.Table) // 设置表名

	query := c.Orm.Table(c.Table)
	if p, err := c.buildParamsForQuery(query); err != nil {
		helper.Ajax("参数错误: "+err.Error(), 1, c.Ctx())
	} else {
		var fields tables.ModelDslFields
		c.Orm.Where("mid = ?", category.ModelId).Where("list_visible = 1").Find(&fields)

		query.Where("catid = ?", catid).
			//Where("deleted_at IS NULL").
			OrderBy("listorder DESC").
			OrderBy("id DESC")

		query.Cols(fields.GetListFields() ...)

		var count int64
		var contents []map[string]interface{}
		if p.Size == 0 {
			err = query.Find(c.Entries)
		} else {
			count, err = query.Clone().Count()
			if err == nil {
				contents, err = query.Limit(p.Size, (p.Page-1)*p.Size).QueryInterface()
			}
		}
		if err != nil {
			helper.Ajax(err, 1, c.Ctx())
			return
		}

		for i, content := range contents {
			for field, value := range content {
				switch value.(type) {
				case []byte:
					content[field] = interface{}(helper.Bytes2String(value.([]byte)))
				}
			}
			contents[i] = content
		}

		helper.Ajax(pine.H{
			"list": contents,
			"pagination": pine.H{
				"page":  p.Page,
				"size":  p.Size,
				"total": count,
			},
		}, 0, c.Ctx())
	}
}
