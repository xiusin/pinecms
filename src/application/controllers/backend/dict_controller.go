package backend

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type DictController struct {
	BaseController
}

func (c *DictController) Construct() {
	c.Group = "字典管理"
	c.SubGroup = "字典列表管理"
	c.KeywordsSearch = []KeywordWhere{
		{Field: "value", Op: "LIKE", DataExp: "%$?%"},
		{Field: "name", Op: "LIKE", DataExp: "%$?%"},
	}
	c.SearchFields = map[string]searchFieldDsl{
		"status": {Op: "="},
	}
	c.Table = &tables.Dict{}
	c.Entries = &[]*tables.Dict{}
	c.ApiEntityName = "字典属性"

	c.BaseController.Construct()
	c.OpBefore = c.before
	c.OpAfter = c.after
}

func (c *DictController) before(act int, param interface{}) error {
	switch act {
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
			c.Orm.In("id", ids).Find(&cats)
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

// GetSelect 下拉列表
func (c *DictController) GetSelect(cacher cache.AbstractCache) {
	cid, _ := c.Ctx().GetInt64("cid")
	if cid == 0 {
		helper.Ajax("请传入字典分类ID", 1, c.Ctx())
		return
	}
	var kv []tables.KV
	if err := cacher.Remember(fmt.Sprintf(controllers.CacheDictPrefix, cid), &kv, func() ([]byte, error) {
		var dicts []tables.KV
		_ = c.Orm.Where("status = 1").Where("cid = ?", cid).Find(c.Entries)
		m := c.Entries.(*[]*tables.Dict)
		for _, model := range *m {
			dicts = append(dicts, tables.KV{Label: model.Name, Value: model.Value})
		}
		return json.Marshal(dicts)
	}); err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	helper.Ajax(kv, 0, c.Ctx())
}
