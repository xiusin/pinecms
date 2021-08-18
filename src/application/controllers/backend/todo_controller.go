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
		"type":{Op: "="},
		"name":{Op: "="},
		"introduce":{Op: "="},
		"listorder":{Op: "="},
		"status":{Op: "="},
		"put_date":{Op: "="},
		"put_datetime":{Op: "="},
		"start_time":{Op: "="},
		"end_time":{Op: "="},

	}
	c.Group = "Todo管理"
  	c.ApiEntityName = "Todo"
	
	c.Table = &tables.Todo{}
	c.Entries = &[]tables.Todo{}
}