package backend

import (
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type SettingController struct {
	BaseController
}

func (c *SettingController) Construct() {
	c.Table = &tables.Setting{}
	c.Entries = &[]*tables.Setting{}
	c.BaseController.Construct()

	c.SearchFields = map[string]searchFieldDsl{
		"`group`": {Op: "="},
	}

	c.KeywordsSearch = []KeywordWhere{
		{Field: "form_name", Op: "LIKE", DataExp: "%$?%"},
		{Field: "`key`", Op: "LIKE", DataExp: "%$?%"},
	}

}

// PostGroups 获取新分组
func (c *SettingController) PostGroups() {
	var groups []tables.Setting
	c.Orm.Table(&tables.Setting{}).GroupBy("`group`").Find(&groups)
	var kvs = []KV{}

	for _, group := range groups {
		kvs = append(kvs, KV{Label: group.Group, Value: group.Group})
	}
	helper.Ajax(kvs, 0, c.Ctx())
}
