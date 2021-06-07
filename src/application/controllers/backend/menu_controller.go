package backend

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type MenuController struct {
	BaseController
}

func (c *MenuController) Construct() {
	c.KeywordsSearch = []KeywordWhere{
		{Field: "name", Op: "LIKE", DataExp: "%$?%"},
	}
	c.Orm = pine.Make(controllers.ServiceXorm).(*xorm.Engine)
	c.Table = &tables.Menu{}
	c.Entries = &[]tables.Menu{}
}

func (c *MenuController) PostList() {
	c.Orm.Find(c.Entries)
	helper.Ajax(c.Entries, 0, c.Ctx())

}
