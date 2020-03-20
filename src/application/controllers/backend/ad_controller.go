package backend

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"html/template"
)

type AdController struct {
	pine.Controller
}

func (c *AdController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY("/ad/list", "AdList")
	b.ANY("/ad/add", "AdAdd")
	b.ANY("/ad/edit", "AdEdit")
	b.ANY("/ad/delete", "AdDelete")

	b.ANY("/ad-space/list", "AdSpaceList")
	b.ANY("/ad-space/add", "AdSpaceAdd")
	b.ANY("/ad-space/edit", "AdSpaceEdit")
	b.ANY("/ad-space/delete", "AdSpaceDelete")
}

func (c *AdController) AdList() {
	page, _ := c.Ctx().URLParamInt64("page")
	rows, _ := c.Ctx().URLParamInt64("rows")
	if page > 0 {
		list, total := models.NewAdModel().GetList(page, rows)
		spaces := models.NewAdSpaceModel().All()
		var h = map[int64]string{}
		for _, space := range spaces {
			h[space.Id] = space.Name
		}
		for k, v := range list {
			list[k].SpaceName = h[v.SpaceID]
		}
		c.Ctx().Render().JSON(map[string]interface{}{"rows": list, "total": total})
		return
	}
	menuid, _ := c.Ctx().URLParamInt64("menuid")
	table := helper.Datagrid("ad_list_datagrid", "/b/ad/list", helper.EasyuiOptions{
		"title":   models.NewMenuModel().CurrentPos(menuid),
		"toolbar": "ad_list_datagrid_toolbar",
	}, helper.EasyuiGridfields{
		"名称":  {"field": "name", "width": "30", "index": "0"},
		"广告位": {"field": "space_name", "width": "50", "index": "1"},
		"图片":  {"field": "image", "width": "30", "index": "2", "formatter": "adListLogoFormatter"},
		"启用":  {"field": "status", "width": "20", "index": "3", "formatter": "adListEnabledFormatter"},
		"操作":  {"field": "id", "index": "4", "formatter": "adListOptFormatter"},
	})
	c.Ctx().Render().ViewData("dataGrid", template.HTML(table))
	c.Ctx().Render().HTML("backend/ad_list.html")
}

func (c *AdController) AdAdd() {
	if c.Ctx().IsPost() {
		var ad tables.Advert
		status := c.Ctx().Request().PostForm.Get("status")
		if status == "" {
			ad.Status = 0
		} else {
			ad.Status = 1
		}
		ad.Name = c.Ctx().PostString("name")
		ad.LinkUrl = c.Ctx().PostString("link_url")
		ad.SpaceID, _ = c.Ctx().PostInt64("space_id")
		if ad.SpaceID == 0 {
			helper.Ajax("广告位参数错误", 1, c.Ctx())
			return
		}
		ad.StartTime = c.Ctx().PostString("start_time")
		ad.EndTime = c.Ctx().PostString("end_time")
		ad.Image = c.Ctx().PostString("image")
		listorder, _ := c.Ctx().PostInt("listorder")
		ad.ListOrder = uint(listorder)

		if models.NewAdModel().Add(&ad) > 0 {
			helper.Ajax("广告添加成功", 0, c.Ctx())
		} else {
			helper.Ajax("广告添加失败", 1, c.Ctx())
		}
		return
	}
	c.ViewData("adspaces", models.NewAdSpaceModel().All())
	imageUploader := template.HTML(helper.SiginUpload("image", "", false, "广告图", "", ""))

	c.ViewData("imageUploader", imageUploader)
	c.Ctx().Render().HTML("backend/ad_add.html")
}

func (c *AdController) AdEdit() {
	if c.Ctx().IsPost() {
		var ad tables.Advert
		status := c.Ctx().Request().PostForm.Get("status")
		if status == "" {
			ad.Status = 0
		} else {
			ad.Status = 1
		}
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
		ad.StartTime = c.Ctx().PostString("start_time")
		ad.EndTime = c.Ctx().PostString("end_time")
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
	page, _ := c.Ctx().URLParamInt64("page")
	rows, _ := c.Ctx().URLParamInt64("rows")
	if page > 0 {
		list, total := models.NewAdSpaceModel().GetList(page, rows)
		c.Ctx().Render().JSON(map[string]interface{}{"rows": list, "total": total})
		return
	}
	menuid, _ := c.Ctx().URLParamInt64("menuid")
	table := helper.Datagrid("adspace_list_datagrid", "/b/ad-space/list", helper.EasyuiOptions{
		"title":   models.NewMenuModel().CurrentPos(menuid),
		"toolbar": "adspace_list_datagrid_toolbar",
	}, helper.EasyuiGridfields{
		"名称": {"field": "name", "width": "30", "index": "0"},
		"操作": {"field": "id", "wdith": "30", "index": "1", "formatter": "adspaceListOptFormatter"},
	})
	c.Ctx().Render().ViewData("dataGrid", template.HTML(table))
	c.Ctx().Render().HTML("backend/adspace_list.html")
}

func (c *AdController) AdSpaceAdd(orm *xorm.Engine) {
	if c.Ctx().IsPost() {
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
		return
	}
	c.Ctx().Render().HTML("backend/adspace_add.html")
}

func (c *AdController) AdSpaceEdit(orm *xorm.Engine) {
	id, _ := c.Ctx().GetInt64("id")
	if c.Ctx().IsPost() {
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
		return
	}
	if c.Ctx().IsPost() {

		return
	}
	adspace := models.NewAdSpaceModel().Get(id)
	c.Ctx().Render().ViewData("adspace", adspace)
	c.Ctx().Render().HTML("backend/adspace_edit.html")
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
