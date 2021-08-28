package backend

import (
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type AttachmentTypeController struct {
	BaseController
}

func (c *AttachmentTypeController) Construct() {
	c.SearchFields = []SearchFieldDsl{
		 {Op: "=", Field: "type"},
	}
	c.Table = &tables.AttachmentType{}
	c.Entries = &[]tables.AttachmentType{}
	c.Group = "附件分类"
	c.ApiEntityName = "分类"
	c.BaseController.Construct()
}

func (c *AttachmentTypeController) PostList()  {
	c.Orm.Find(c.Entries)
	helper.Ajax(c.Entries, 0, c.Ctx())
}
