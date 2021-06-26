package backend

import (
	"github.com/xiusin/pinecms/src/application/controllers/middleware/apidoc"
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
	c.AppId = "admin"
	c.Group = "系统配置"
	c.SubGroup = "附件管理"

	c.apiEntities = map[string]apidoc.Entity{
		"list":   {Title: "附件列表", Desc: "查询已上传系统的附件列表"},
		"add":    {Title: "新增配置", Desc: "新增上传附件"},
		"del":    {Title: "删除配置", Desc: "删除一个附件"},
	}

	c.BaseController.Construct()
}
