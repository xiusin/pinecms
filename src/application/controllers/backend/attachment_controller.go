package backend

import (
	"github.com/xiusin/pinecms/src/application/controllers/middleware/apidoc"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"xorm.io/xorm"
)

type AttachmentController struct {
	BaseController
}

func (c *AttachmentController) Construct() {
	c.SearchFields = []SearchFieldDsl{
		//{Field: "`type`", Op: "="},
	}
	c.Table = &tables.Attachments{}
	c.Entries = &[]*tables.Attachments{}
	c.Group = "系统配置"
	c.SubGroup = "附件管理"

	c.apiEntities = map[string]apidoc.Entity{
		"list": {Title: "附件列表", Desc: "查询已上传系统的附件列表"},
		"add":  {Title: "新增配置", Desc: "新增上传附件"},
		"del":  {Title: "删除配置", Desc: "删除一个附件"},
	}
	c.OpBefore = c.before
	c.BaseController.Construct()
}

func (c *AttachmentController) before(act int, params interface{}) error {
	if act == OpList {
		cid, _ := c.Input().GetInt64("classifyId")
		if cid > 0 {
			params.(*xorm.Session).Where("classify_id = ?", cid)
		}
		params.(*xorm.Session).Desc("id")
	}
	return nil
}

func (c *AttachmentController) PostAdd() {
	if err := c.BindParse(); err != nil {
		helper.Ajax(err.Error(), 1, c.Ctx())
		return
	}

	md5, _ := c.Input().GetString("md5")
	if len(md5) == 0 {
		helper.Ajax("缺少必要的MD5参数", 1, c.Ctx())
		return
	}
	data := &tables.Attachments{}

	if exist, _ := c.Orm.Where("md5 = ?", md5).Get(data); exist {
		helper.Ajax(data, 0, c.Ctx())
		return
	} else if err := c.add(); err == nil  {
		helper.Ajax(c.Table, 0, c.Ctx())
	} else {
		helper.Ajax(err, 1, c.Ctx())
	}
}
