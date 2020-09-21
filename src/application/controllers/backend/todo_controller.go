package backend

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type TodoController struct {
	BaseController
}

func (c *TodoController) Construct() {
	c.BindType = "form"
	c.SearchFields = map[string]searchFieldDsl{
		"id":          {Op: "="},
		"userid":      {Op: "="},
		"status":      {Op: "="},
		"set_status":  {Op: "="},
		"enum_status": {Op: "="},
		"city_id":     {Op: "="},
		"createtime":  {Op: "range"},
		"date":        {Op: "range"},
		"time":        {Op: "="},
		"year":        {Op: "="},
	}

	c.Orm = pine.Make(controllers.ServiceXorm).(*xorm.Engine)
	c.Table = &tables.Todo{}
	c.Entries = &[]tables.Todo{}
}
