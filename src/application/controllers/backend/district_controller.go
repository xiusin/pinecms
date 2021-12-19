package backend

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"net/http"
)

type DistrictController struct {
	BaseController
}

func (c *DistrictController) Construct() {
	c.Table = &tables.District{}
	c.Entries = &[]tables.District{}
	c.Group = "字典管理"
	c.ApiEntityName = "地区"
	c.BaseController.Construct()
}

func (c *DistrictController) getSubDistrictId() []int64 {
	var p []int64
	for _, entity := range *c.Entries.(*[]tables.District) {
		if len(entity.Code) != 0 {
			p = append(p, entity.Id)
		}
	}
	return p
}

func (c *DistrictController) PostList()  {
	query := c.Orm.Table(c.Table)
	if p, err := c.buildParamsForQuery(query); err != nil {
		helper.Ajax("参数错误: "+err.Error(), 1, c.Ctx())
		return
	} else {
		p.Size = 5
		var count int64
		var err error
		count, err = query.Where("parent_id = 0").Limit(p.Size, (p.Page-1)*p.Size).FindAndCount(c.Entries)
		if err != nil {
			helper.Ajax("获取列表异常: "+err.Error(), 1, c.Ctx())
			return
		}
		ids := c.getSubDistrictId()
		// 分页ID
		var allIds = ids

		for len(ids) > 0 {
			sess := c.Orm.NewSession()
			var entities []tables.District
			sess.In("parent_id", ids).Find(&entities)
			c.Entries = &entities
			ids = c.getSubDistrictId()
			allIds = append(allIds, ids...)
		}
		var entities = []tables.District{}
		c.Orm.In("id", allIds).Find(&entities)
		helper.Ajax(pine.H{
			"list": entities,
			"pagination": pine.H{
				"page":  p.Page,
				"size":  p.Size,
				"total": count,
			},
		}, 0, c.Ctx())
	}
}

// PostImport 导入外部数据库
func (c *DistrictController) PostImport() {
	dbUrl := "https://raw.githubusercontent.com/eduosi/district/master/district-full.sql"
	resp, err := http.Get(dbUrl)
	if err != nil {
		helper.Ajax(err.Error(), 1, c.Ctx())
		return
	}
	defer resp.Body.Close()
	c.Orm.DropTables(&tables.District{}) //先删除table
	if _, err = c.Orm.Import(resp.Body); err != nil {
		helper.Ajax(err.Error(), 1, c.Ctx())
		return
	}
	c.Orm.Exec("rename table district to "+ controllers.GetTableName("district")  + ";")
	helper.Ajax("导入数据库成功", 0, c.Ctx())
}
