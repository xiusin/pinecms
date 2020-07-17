package backend

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"html/template"
	"strconv"
	"strings"
	"time"
)

type AdController struct {
	pine.Controller
}

func (c *AdController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY("/ad/list", "AdList")
	b.POST("/ad/add", "AdAdd")
	b.ANY("/ad/edit", "AdEdit")
	b.ANY("/ad/order", "AdOrder")
	b.ANY("/ad/delete", "AdDelete")

	b.GET("/ad/space-list", "AdSpaceList")
	b.POST("/ad/space-add", "AdSpaceAdd")
	b.POST("/ad/space-edit", "AdSpaceEdit")
	b.POST("/ad/space-delete", "AdSpaceDelete")
}

func (c *AdController) AdList() {
	page, _ := c.Ctx().GetInt64("page")
	rows, _ := c.Ctx().GetInt64("rows")
	list, total := models.NewAdModel().GetList(page, rows)
	spaces := models.NewAdSpaceModel().All()
	var h = map[int64]string{}
	for _, space := range spaces {
		h[space.Id] = space.Name
	}
	for k, v := range list {
		list[k].SpaceName = h[v.SpaceID]
	}
	helper.Ajax(pine.H{"rows": list, "total": total}, 0, c.Ctx())
}

func (c *AdController) AdAdd() {
	var ad tables.Advert
	ad.Status, _ = c.Ctx().PostBool("status")
	ad.Name = c.Ctx().PostString("name")
	ad.LinkUrl = c.Ctx().PostString("link_url")
	ad.SpaceID, _ = c.Ctx().PostInt64("space_id")
	if ad.SpaceID == 0 {
		helper.Ajax("广告位参数错误", 1, c.Ctx())
		return
	}
	ad.StartTime = c.Ctx().PostString("start_time", time.Now().In(helper.GetLocation()).Format(helper.TimeFormat))
	ad.EndTime = c.Ctx().PostString("end_time", time.Now().In(helper.GetLocation()).Add(365*24*3600*time.Second).Format(helper.TimeFormat))
	ad.Image = c.Ctx().PostString("image")
	listorder, _ := c.Ctx().PostInt("listorder")
	ad.ListOrder = uint(listorder)

	if models.NewAdModel().Add(&ad) > 0 {
		helper.Ajax("广告添加成功", 0, c.Ctx())
	} else {
		helper.Ajax("广告添加失败", 1, c.Ctx())
	}

}

func (c *AdController) AdEdit() {
	if c.Ctx().IsPost() {
		var ad tables.Advert
		ad.Status, _ = c.Ctx().PostBool("status")
		ad.Name = c.Ctx().PostString("name")
		ad.LinkUrl = c.Ctx().PostString("link_url")
		ad.Id, _ = c.Ctx().PostInt64("id")
		if ad.Id < 1 {
			helper.Ajax("参数错误", 1, c.Ctx())
			return
		}
		ad.SpaceID, _ = c.Ctx().PostInt64("space_id")
		if ad.SpaceID < 1 {
			helper.Ajax("广告位参数错误", 1, c.Ctx())
			return
		}
		// 默认为当前时间向后推迟到一年
		ad.StartTime = c.Ctx().PostString("start_time", time.Now().In(helper.GetLocation()).Format(helper.TimeFormat))
		ad.EndTime = c.Ctx().PostString("end_time", time.Now().In(helper.GetLocation()).Add(365*24*3600*time.Second).Format(helper.TimeFormat))
		ad.Image = c.Ctx().PostString("image")
		listorder, _ := c.Ctx().PostInt("listorder")
		ad.ListOrder = uint(listorder)

		if models.NewAdModel().Update(&ad) {
			helper.Ajax("广告修改成功", 0, c.Ctx())
		} else {
			helper.Ajax("广告修改失败", 1, c.Ctx())
		}
		return
	}
	id, _ := c.Ctx().GetInt64("id")
	ad := models.NewAdModel().Get(id)
	imageUploader := template.HTML(helper.SiginUpload("image", ad.Image, false, "广告图", "", ""))
	c.ViewData("imageUploader", imageUploader)
	c.ViewData("ad", ad)
	c.ViewData("adspaces", models.NewAdSpaceModel().All())
	c.Ctx().Render().HTML("backend/ad_edit.html")
}

func (c *AdController) AdOrder(orm *xorm.Engine) {
	posts := c.Ctx().PostData()
	fmt.Println(posts)
	data := map[int64]int64{}
	for k, v := range posts {
		k = strings.Replace(k, "order[", "", 1)
		k = strings.Replace(k, "]", "", 1)
		s, e := strconv.Atoi(k)
		if e != nil {
			continue
		}
		sort, e := strconv.Atoi(v[0])
		if e != nil {
			continue
		}
		data[int64(s)] = int64(sort)
	}
	var flag int64
	for id, val := range data {
		ad := new(tables.Advert)
		ad.ListOrder = uint(val)
		affected, err := orm.Id(id).Update(ad)
		if err != nil {
			c.Logger().Error("adorder", err)
		}
		if affected > 0 {
			flag++
		}
	}
	if flag > 0 {
		helper.Ajax("排序更新成功", 0, c.Ctx())
	} else {
		helper.Ajax("排序规则没有发生任何改变", 1, c.Ctx())
	}
}

func (c *AdController) AdDelete() {
	id, _ := c.Ctx().GetInt64("id")
	if id == 0 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	if models.NewAdModel().Delete(id) {
		helper.Ajax("广告删除成功", 0, c.Ctx())
	} else {
		helper.Ajax("广告删除失败", 1, c.Ctx())
	}
}

func (c *AdController) AdSpaceList() {
	page, _ := c.Ctx().GetInt64("page")
	rows, _ := c.Ctx().GetInt64("rows")
	list, total := models.NewAdSpaceModel().GetList(page, rows)
	helper.Ajax(pine.H{"rows": list, "total": total}, 0, c.Ctx())

}

func (c *AdController) AdSpaceAdd(orm *xorm.Engine) {

	space := &tables.AdvertSpace{Name: c.Ctx().FormValue("name")}
	if exists, _ := orm.Exist(space); exists {
		helper.Ajax("广告位名称已经存在", 1, c.Ctx())
		return
	}
	if models.NewAdSpaceModel().Add(space) > 0 {
		helper.Ajax("广告位添加成功", 0, c.Ctx())
	} else {
		helper.Ajax("广告位添加失败", 1, c.Ctx())
	}

}

func (c *AdController) AdSpaceEdit(orm *xorm.Engine) {
	id, _ := c.Ctx().PostInt64("id")
	if id <= 0 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	if count, _ := orm.Where("id <> ?", id).Where("name = ?", c.Ctx().FormValue("name")).Count(&tables.AdvertSpace{}); count > 0 {
		helper.Ajax("广告位名称已经存在", 1, c.Ctx())
		return
	}
	space := models.NewAdSpaceModel().Get(id)
	if space == nil {
		helper.Ajax("广告位不存在", 1, c.Ctx())
		return
	}
	space.Name = c.Ctx().FormValue("name")
	if models.NewAdSpaceModel().Update(space) {
		helper.Ajax("广告位更新成功", 0, c.Ctx())
	} else {
		helper.Ajax("广告位更新失败", 1, c.Ctx())
	}

}

func (c *AdController) AdSpaceDelete(orm *xorm.Engine) {
	id, _ := c.Ctx().GetInt64("id")
	if id == 0 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	space := models.NewAdSpaceModel()
	advpos := space.Get(id)
	if advpos == nil {
		helper.Ajax("广告位不存在", 1, c.Ctx())
		return
	}
	count, _ := orm.Where("space_id = ?", advpos.Id).Count(&tables.Advert{})
	if count > 0 {
		helper.Ajax("广告位下还有广告,无法直接删除", 1, c.Ctx())
		return
	} else {
		if space.Delete(advpos.Id) {
			helper.Ajax("广告位删除成功", 0, c.Ctx())
		} else {
			helper.Ajax("广告位删除失败", 1, c.Ctx())
		}
	}
}
