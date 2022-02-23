package backend

import (
	"errors"
	"fmt"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"xorm.io/xorm"
)

type TableController struct {
	BaseController
}

func (c *TableController) Construct() {
	c.Group = "字段管理"
	c.KeywordsSearch = []SearchFieldDsl{
		{Field: "name", Op: "LIKE", DataExp: "%$?%"},
	}
	c.SearchFields = []SearchFieldDsl{
		{Field: "status"},
	}
	c.Table = &tables.DocumentModelDsl{}
	c.Entries = &[]tables.DocumentModelDsl{}
	c.ApiEntityName = "字段"
	c.OpBefore = c.before
	c.BaseController.Construct()
}

func (c TableController) GetFields() {
	var fields []tables.DocumentModelField
	err := c.Orm.Table(&tables.DocumentModelField{}).Find(&fields)
	if err != nil {
		c.Logger().Error(err)
	}
	helper.Ajax(fields, 0, c.Ctx())
}

func (c *TableController) before(act int, params interface{}) error {
	if OpList == act {
		params.(*xorm.Session).Unscoped().Where("mid <> ?", 0)
		v, _ := c.Input().GetInt64("mid")
		if v != 0 {
			params.(*xorm.Session).Where("mid = ?", v)
		}
	} else if OpAdd == act || OpEdit == act {
		data := params.(*tables.DocumentModelDsl)
		sess := c.Orm.Where("table_field = ?", data.TableField).Where("mid = ?", data.Mid)
		if OpEdit == act {
			sess.Where("id <> ?", data.Id)
		}
		exist, err := sess.Exist(&tables.DocumentModelDsl{})
		if err != nil {
			c.Logger().Error(err)
		}
		if exist {
			return errors.New("字段已经存在，请换个字段名称")
		}
	} else if OpDel == act {
		c.Orm.Where(c.TableKey+"=?", params.(*idParams).Id).Get(c.Table)
		exist, _ := c.Orm.Where("mid = ?", 0).Where("table_field = ?", c.Table.(*tables.DocumentModelDsl).TableField).Exist(&tables.DocumentModelDsl{})
		if exist {
			return errors.New("不可删除模型固有字段")
		}
	}
	return nil
}

func (c *TableController) after(act int, params interface{}) {
	if act == OpDel || act == OpAdd || act == OpEdit {
		pine.Logger().Print("操作模型字段, 清除缓存")
		cacher := di.MustGet(controllers.ServiceICache).(cache.AbstractCache)
		if act == OpDel {
			cacher.Delete(fmt.Sprintf(controllers.CacheModelTablePrefix, params.(*idParams).Id))
		} else {
			mid, _ := c.Input().GetInt("mid")
			cacher.Delete(fmt.Sprintf(controllers.CacheModelTablePrefix, mid))
		}
	}
}
