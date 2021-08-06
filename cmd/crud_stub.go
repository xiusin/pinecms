package cmd

const (
	controllerDir = "src/application/controllers/backend/"
	tableDir      = "src/application/models/tables/"
	feDir         = "src/cool/modules/base/views/"
	theme         = "vim"
	controllerTpl = `package backend
import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type [ctrl] struct {
	BaseController
}

func (c *[ctrl]) Construct() {
    c.SearchFields = map[string]searchFieldDsl{
[searchFieldDsl]
	}
	c.Group = "[table]管理"
  	c.ApiEntityName = "[table]"
	
	c.Table = &tables.[table]{}
	c.Entries = &[]tables.[table]{}
}`
	tableTpl = `package tables

[struct]
`

	indexVueTpl = ``

	serviceTsTpl = ``
)
