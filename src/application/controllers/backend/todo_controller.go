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
	c.Orm = pine.Make(controllers.ServiceXorm).(*xorm.Engine)
	c.Table = &tables.Todo{}
	c.Entries = &[]tables.Todo{}
}