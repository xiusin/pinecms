package backend

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"strings"
)

type TableController struct {
	BaseController
}

func (c *TableController) Construct() {
	c.Group = "字段管理"
	c.KeywordsSearch = []KeywordWhere{
		{Field: "name", Op: "LIKE", DataExp: "%$?%"},
	}
	c.SearchFields = map[string]searchFieldDsl{
		"status": {Op: "="},
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
		if v := c.Input().GetInt64("mid"); v != 0 {
			params.(*xorm.Session).Where("mid = ?", v)
		}
	} else if OpAdd == act || OpEdit == act {
		data := params.(*tables.DocumentModelDsl)
		if strings.HasPrefix(data.Datasource, "[") || strings.HasPrefix(data.Datasource, "{") {
			var dataSourceJson interface{}
			if err := json.Unmarshal([]byte(data.Datasource), &dataSourceJson); err != nil {
				return fmt.Errorf("数据源格式错误： %s", err.Error())
			}
		}
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
