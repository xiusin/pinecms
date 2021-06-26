package backend

import (
	"errors"
	"github.com/xiusin/pinecms/src/application/controllers/middleware/apidoc"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type DictController struct {
	BaseController
}

func (c *DictController) Construct() {
	c.AppId = "admin"
	c.Group = "字典管理"
	c.SubGroup = "字典列表管理"
	c.KeywordsSearch = []KeywordWhere{
		{Field: "key", Op: "LIKE", DataExp: "%$?%"},
		{Field: "name", Op: "LIKE", DataExp: "%$?%"},
	}
	c.SearchFields = map[string]searchFieldDsl{
		"status": {Op: "="},
	}
	c.Table = &tables.Dict{}
	c.Entries = &[]*tables.Dict{}
	c.apiEntities = map[string]apidoc.Entity{
		"list":   {Title: "字典列表", Desc: "查询指定分类的字典列表"},
		"add":    {Title: "新增字典", Desc: "在给定分类下新增一条字典记录"},
		"edit":   {Title: "编辑字典", Desc: "编辑一条字典记录，在指定分类下标识不可重复"},
		"delete": {Title: "删除字典", Desc: "删除指定字典"},
		"info":   {Title: "字典详情", Desc: "获取指定字典的详情信息"},
	}
	c.BaseController.Construct()

	c.OpBefore = c.before
	c.OpAfter = c.after
}

func (c *DictController) before(act int, param interface{}) error {
	switch act {
	case OpList:
		cid, _ := param.(*listParam).Params["cid"]
		if uint(cid.(float64)) < 1 {
			return errors.New("必须选择字典分类")
		}
	case OpAdd, OpEdit:
		p := param.(*tables.Dict)
		id, cid, name := p.Id, p.Cid, p.Name
		sess := c.Orm.Table(c.Table).Where("name = ?", name).Where("cid = ?", cid)
		if OpEdit == act {
			sess = sess.Where("id <> ?", id)
		}
		exist, err := sess.Exist()
		if err != nil {
			return err
		}
		if exist {
			return errors.New("该分类下已经存在此标识")
		}
	}
	return nil
}

func (c *DictController) after(act int, param interface{}) error {
	switch act {
	case OpList:
		entities := c.Entries.(*[]*tables.Dict)
		var ids []uint
		for _, dict := range *entities {
			ids = append(ids, dict.Cid)
		}
		if len(ids) > 0 {
			var cats []tables.DictCategory
			c.Orm.Where("id in ?", ids).Find(&cats)
			var m = map[uint]string{}
			for _, cat := range cats {
				m[cat.Id] = cat.Name
			}
			for _, dict := range *entities {
				dict.CatName, _ = m[dict.Cid]
			}
		}
	}
	return nil
}
