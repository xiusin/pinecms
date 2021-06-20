package backend

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type LinkController struct {
	BaseController
}

func (c *LinkController) Construct() {
	c.KeywordsSearch = []KeywordWhere{
		{Field: "name", Op: "LIKE", DataExp: "%$?%"},
	}
	c.Orm = pine.Make(controllers.ServiceXorm).(*xorm.Engine)
	c.Table = &tables.Link{}
	c.Entries = &[]*tables.Link{}
	c.TableStructKey = "Linkid"
}
