package controller

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pinecms/src/application/controllers/backend"
)

// EmailController 控制器区域
type EmailController struct {
	backend.BaseController
}

func (c *EmailController) Construct() {
	c.BaseController.Construct()
	c.OpBefore = c.before
}

func (c *EmailController) before(act int, query interface{}) error {
	if act == backend.OpList {
		sess := query.(*xorm.Session)
		status := c.Input().GetInt("status")
		sess.Where("status = ?", status)
		if status == 1 {
			sess.Where("`type` = ?", c.Input().GetInt("type"))
		}
	}
	return nil
}
