package backend

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"html/template"
	"iriscms/application/models"
	"iriscms/common/helper"
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
	fmt.Println(rows)
	if page > 0 {
		this.Ctx.JSON(map[string]interface{}{"rows": []string{}, "total": 0})
		return
	}
	table := helper.Treegrid("category_categorylist_treegrid", "/b/content/news-list?grid=datagrid", helper.EasyuiOptions{
		"toolbar":      "content_newslist_datagrid_toolbar",
		"singleSelect": "false",
	}, helper.EasyuiGridfields{
		"排序":   {"field": "listorder", "width": "15", "align": "center", "formatter": "contentNewsListOrderFormatter", "index": "0"},
		"新闻名称": {"field": "title", "width": "130", "index": "1"},
		"管理操作": {"field": "catid", "width": "50", "sortable": "true", "align": "center", "index": "2"},
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
	this.Ctx.View("backend/" + cat.TplPrefix + "add.html")
}

//修改内容
func (this *ContentController) EditContent() {

}

//删除内容
func (this *ContentController) DeleteContent() {

}

//排序内容
func (this *ContentController) OrderContent() {

}
