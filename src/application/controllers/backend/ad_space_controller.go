package backend

import (
	"errors"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type AdSpaceController struct {
	BaseController
}

func (c *AdSpaceController) Construct() {
	c.Table = &tables.AdvertSpace{}
	c.Entries = &[]tables.AdvertSpace{}
	c.AppId = "admin"
	c.Group = "广告管理"
	c.SubGroup = "广告位管理"
	c.ApiEntityName = "广告位"
	c.BaseController.Construct()
	c.OpBefore = c.before
}

func (c *AdSpaceController) before(act int, params interface{}) error {
	if act == OpDel {
		ids := params.(*idParams).Ids
		count, _ := c.Orm.In("space_id", ids).Count(&tables.Advert{})
		if count > 0 {
			return errors.New("广告位下还有广告,无法直接删除")
		}
	}
	if act == OpEdit || act == OpAdd {
		t, p := &tables.AdvertSpace{}, params.(*tables.AdvertSpace)
		if act == OpAdd {
			if exists, _ := c.Orm.Where("name = ? or `key` = ?", p.Name, p.Key).Exist(t); exists {
				return errors.New("广告位名称或标识已经存在")
			}
		} else {
			if exists, _ := c.Orm.Where("id <> ? and (name = ? or `key` = ?)", p.Id, p.Name, p.Key).Exist(t); exists {
				return errors.New("广告位名称或标识已经存在")
			}
		}
	}
	return nil
}
