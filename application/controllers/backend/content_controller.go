package backend

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"html/template"
	"iriscms/application/models"
	"iriscms/application/models/tables"
	"iriscms/common/helper"
	"strconv"
	"strings"
)

type ContentController struct {
	Ctx     iris.Context
	Orm     *xorm.Engine
	Session *sessions.Session
}

func (c *ContentController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("ANY", "/content/index", "Index")
	b.Handle("ANY", "/content/right", "Right")
	b.Handle("ANY", "/content/public-welcome", "Welcome")
	b.Handle("ANY", "/content/news-list", "NewsList")
	b.Handle("ANY", "/content/page", "Page")
	b.Handle("ANY", "/content/add", "AddContent")
	b.Handle("ANY", "/content/edit", "EditContent")
	b.Handle("ANY", "/content/delete", "DeleteContent")
	b.Handle("ANY", "/content/order", "OrderContent")
}

func (this *ContentController) Index() {
	menuid, _ := this.Ctx.URLParamInt64("menuid")
	this.Ctx.ViewData("currentPos", models.NewMenuModel(this.Orm).CurrentPos(menuid))
	this.Ctx.View("backend/content_index.html")
}

func (this *ContentController) Welcome() {
	this.Ctx.View("backend/content_welcome.html")
}

func (this *ContentController) Right() {
	if this.Ctx.Method() == "POST" {
		cats := models.NewCategoryModel(this.Orm).GetContentRightCategoryTree(models.NewCategoryModel(this.Orm).GetAll(), 0)
		this.Ctx.JSON(cats)
		return
	}
	this.Ctx.View("backend/content_right.html")
}

func (this *ContentController) NewsList() {
	catid, _ := this.Ctx.URLParamInt64("catid")
	page, _ := this.Ctx.URLParamInt64("page")
	rows, _ := this.Ctx.URLParamInt64("rows")
	if page > 0 {
		var contents []tables.IriscmsContent
		total, _ := this.Orm.Where("catid = ? and deleted_at = 0", catid).Limit(int(rows), int((page-1)*rows)).Desc("id").FindAndCount(&contents)
		this.Ctx.JSON(map[string]interface{}{"rows": contents, "total": total})
		return
	}
	table := helper.Datagrid("category_categorylist_datagrid", "/b/content/news-list?grid=datagrid&catid="+strconv.Itoa(int(catid)), helper.EasyuiOptions{
		"toolbar":      "content_newslist_datagrid_toolbar",
		"singleSelect": "true",
	}, helper.EasyuiGridfields{
		"排序":   {"field": "listorder", "width": "15", "formatter": "contentNewsListOrderFormatter", "index": "0"},
		"资源标题": {"field": "title", "width": "130", "index": "1"},
		"下载扣分": {"field": "money", "width": "30", "index": "2"},
		"获密方式": {"field": "pwd_type", "width": "30", "formatter": "getPwdTypeFormatter", "index": "3"},
		"管理操作": {"field": "id", "width": "50", "formatter": "contentNewsListOperateFormatter", "index": "4"},
	})

	this.Ctx.ViewData("DataGrid", template.HTML(table))
	this.Ctx.ViewData("catid", catid)
	this.Ctx.View("backend/content_newslist.html")
}

func (this *ContentController) Page() {
	catid, _ := this.Ctx.URLParamInt64("catid")
	if catid == 0 {
		helper.Ajax("页面错误", 1, this.Ctx)
		return
	}
	pageModel := models.NewPageModel(this.Orm)
	page := pageModel.GetPage(catid)
	var hasPage bool = false
	if page.Title != "" {
		hasPage = true
	}
	var res bool
	if this.Ctx.Method() == "POST" {
		page.Title = this.Ctx.FormValue("title")
		page.Content = this.Ctx.FormValue("content")
		page.Keywords = this.Ctx.FormValue("keywords")
		page.Updatetime = int64(helper.GetTimeStamp())
		if !hasPage {
			page.Catid = catid
			res = pageModel.AddPage(page)
		} else {
			res = pageModel.UpdatePage(page)
		}
		if res {
			helper.Ajax("发布成功", 0, this.Ctx)
		} else {
			helper.Ajax("发布失败", 1, this.Ctx)
		}
		return
	}

	this.Ctx.ViewData("catid", catid)
	this.Ctx.ViewData("info", page)
	this.Ctx.View("backend/content_page.html")

}

//添加内容
func (this *ContentController) AddContent() {
	if this.Ctx.Method() == "POST" {
		var content tables.IriscmsContent
		if err := this.Ctx.ReadForm(&content); err != nil {
			helper.Ajax("添加失败"+err.Error(), 1, this.Ctx)
			return
		}
		content.Id = 0
		content.Status = 1
		content.UpdatedAt = int64(helper.GetTimeStamp())
		content.CreatedAt = content.UpdatedAt
		id, err := this.Orm.InsertOne(&content)
		if err != nil {
			helper.Ajax("添加失败:getid:"+strconv.Itoa(int(id))+":"+err.Error(), 1, this.Ctx)
			return
		}
		if id == 0 {
			helper.Ajax("添加失败", 1, this.Ctx)
			return
		}
		helper.Ajax("添加成功", 0, this.Ctx)
		return
	}
	//根据catid读取出相应的添加模板
	catid, _ := this.Ctx.URLParamInt64("catid")
	if catid == 0 {
		helper.Ajax("参数错误", 1, this.Ctx)
		return
	}
	cat, err := models.NewCategoryModel(this.Orm).GetCategory(catid)
	if err != nil {
		helper.Ajax("读取数据错误:"+err.Error(), 1, this.Ctx)
		return
	}
	if cat.Catid == 0 {
		helper.Ajax("不存在的分类", 1, this.Ctx)
		return
	}
	this.Ctx.ViewData("category", cat)
	this.Ctx.View("backend/" + models.NewCategoryModel(this.Orm).GetCategoryTplPrefix(cat.Parentid, cat.TplPrefix) + "add.html")
}

//修改内容
func (this *ContentController) EditContent() {
	//根据catid读取出相应的添加模板
	catid, _ := this.Ctx.URLParamInt64("catid")
	id, _ := this.Ctx.URLParamInt64("id")
	if catid == 0 || id == 0 {
		helper.Ajax("参数错误", 1, this.Ctx)
		return
	}
	var content = tables.IriscmsContent{Id: id}
	ok, _ := this.Orm.Get(&content)
	if !ok {
		helper.Ajax("无法获取内容", 1, this.Ctx)
		return
	}
	if this.Ctx.Method() == "POST" {
		if err := this.Ctx.ReadForm(&content); err != nil {
			helper.Ajax("添加失败:readform:"+err.Error(), 1, this.Ctx)
			return
		}
		content.UpdatedAt = int64(helper.GetTimeStamp())
		res, err := this.Orm.Where("id=? and catid=?", id, catid).Update(&content)
		if err != nil {
			helper.Ajax("修改失败:"+err.Error(), 1, this.Ctx)
			return
		}
		if res == 0 {
			helper.Ajax("修改失败", 1, this.Ctx)
			return
		}
		helper.Ajax("修改成功", 0, this.Ctx)
		return
	}

	cat, err := models.NewCategoryModel(this.Orm).GetCategory(catid)
	if err != nil {
		helper.Ajax("读取数据错误:"+err.Error(), 1, this.Ctx)
		return
	}
	if cat.Catid == 0 {
		helper.Ajax("不存在的分类", 1, this.Ctx)
		return
	}

	this.Ctx.ViewData("content", content)
	this.Ctx.ViewData("category", cat)
	this.Ctx.View("backend/" + cat.TplPrefix + "edit.html")
}

//删除内容
func (this *ContentController) DeleteContent() {
	id := this.Ctx.FormValue("ids")
	if id == "" {
		helper.Ajax("参数错误", 1, this.Ctx)
		return
	}
	ids := strings.Split(id, ",")
	res, _ := this.Orm.In("id", ids).Update(&tables.IriscmsContent{DeletedAt: int64(helper.GetTimeStamp())})
	if res > 0 {
		helper.Ajax("删除成功", 0, this.Ctx)
	} else {
		helper.Ajax("删除失败", 1, this.Ctx)
	}
}

//排序内容
func (this *ContentController) OrderContent() {

}
