package backend

import (
	"github.com/xiusin/pine/pointer"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type AdController struct {
	BaseController
}

func (c *AdController) Construct() {
	c.Table = &tables.Advert{}
	c.Entries = &[]tables.Advert{}

	c.AppId = "admin"
	c.Group = "广告管理"
	c.SubGroup = "广告管理"
	c.ApiEntityName = "广告"
	c.BaseController.Construct()

	c.OpBefore = func(i int, a any) error {
		if c.IsOperate(i) {
			t := a.(*tables.Advert)
			if len(t.DateRange) == 2 {
				t.StartTime = pointer.To(helper.ToTableTime(t.DateRange[0]))
				t.EndTime = pointer.To(helper.ToTableTime(t.DateRange[1]))
			}
		}
		return nil
	}
}
