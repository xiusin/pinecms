package backend

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pinecms/src/application/controllers/middleware/apidoc"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type AttachmentController struct {
	BaseController
}

func (c *AttachmentController) Construct() {
	c.SearchFields = []SearchFieldDsl{
		{Field: "`type`", Op: "="},
	}
	c.Table = &tables.Attachments{}
	c.Entries = &[]*tables.Attachments{}
	c.Group = "系统配置"
	c.SubGroup = "附件管理"

	c.apiEntities = map[string]apidoc.Entity{
		"list":   {Title: "附件列表", Desc: "查询已上传系统的附件列表"},
		"add":    {Title: "新增配置", Desc: "新增上传附件"},
		"del":    {Title: "删除配置", Desc: "删除一个附件"},
	}
	c.OpBefore = c.before
	c.BaseController.Construct()
}

func (c *AttachmentController) before(act int, params interface{}) error {
	if act == OpList {
		cid := c.Input().GetInt64("classifyId")
		if cid > 0 {
			params.(*xorm.Session).Where("classify_id = ?", cid)
		}
	}
	return nil
}
