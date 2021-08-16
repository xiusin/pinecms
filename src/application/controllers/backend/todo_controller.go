package backend
import (
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type TodoController struct {
	BaseController
}

func (c *TodoController) Construct() {
    c.SearchFields = map[string]searchFieldDsl{
		"id":{Op: "="},
		"linktype":{Op: "="},
		"name":{Op: "="},
		"url":{Op: "="},
		"introduce":{Op: "="},
		"listorder":{Op: "="},
		"passed":{Op: "="},
		"addtime":{Op: "range"},

	}
	c.Group = "Todo管理"
  	c.ApiEntityName = "Todo"
	
	c.Table = &tables.Todo{}
	c.Entries = &[]tables.Todo{}
}