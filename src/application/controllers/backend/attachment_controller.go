package backend

import (
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type AttachmentController struct {
	BaseController
}

func (c *AttachmentController) Construct() {
	c.SearchFields = map[string]searchFieldDsl{
		"type": {Op: "="},
	}
	c.Table = &tables.Attachments{}
	c.Entries = &[]*tables.Attachments{}
	c.BaseController.Construct()
}
