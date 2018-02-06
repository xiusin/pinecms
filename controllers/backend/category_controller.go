package backend

import (
	"html/template"
	"iriscms/controllers/backend/helper"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/context"
	"iriscms/models"
	"strconv"
	"iriscms/models/tables"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/mvc"
)

type CategoryController struct {
	Ctx context.Context
	Orm *xorm.Engine
	Session *sessions.Session
}

func (c *CategoryController) BeforeActivation(b mvc.BeforeActivation) {
	//分类相关
	b.Handle("ANY","/category/list", "CategoryList")
	b.Handle("ANY","/category/category-add", "CategoryAdd")
	b.Handle("ANY","/category/category-edit", "CategoryEdit")
	b.Handle("ANY","/category/category-select", "CategorySelect")
	b.Handle("ANY","/category/category-delete", "CategoryDelete")
	b.Handle("ANY","/category/category-order", "CategoryOrder")
}



func (this *CategoryController) CategoryList() {
	if this.Ctx.URLParam("grid") == "treegrid" {
		this.Ctx.JSON(models.NewCategoryModel(this.Orm).GetTree(models.NewCategoryModel(this.Orm).GetAll(),0))
		return
	}
	menuid, _ := this.Ctx.URLParamInt64("menuid")
	table := helper.Treegrid("category_categorylist_treegrid", "/b/category/list?grid=treegrid", helper.EasyuiOptions{
		"title":     models.NewMenuModel(this.Orm).CurrentPos(menuid),
		"toolbar":   "category_categorylist_treegrid_toolbar",
		"idField":   "catid",
		"treeField": "catname",
	}, helper.EasyuiGridfields{
		"排序":   {"field": "listorder", "width": "15", "align": "center", "formatter": "categoryCategoryListOrderFormatter", "index": "0"},
		"栏目名称": {"field": "catname", "width": "130", "index": "1"},
		"栏目类型": {"field": "type", "width": "30", "formatter": "categoryCategoryListTypeFormatter", "index": "2"},
		"描述":   {"field": "description", "width": "80", "index": "3"},
		"状态":   {"field": "ismenu", "width": "20", "formatter": "categoryCategoryListStateFormatter", "index": "4"},
		"管理操作": {"field": "catid", "width": "50", "sortable": "true", "align": "center", "formatter": "categoryCategoryListOperateFormatter", "index": "5"},
	})
	this.Ctx.ViewData("TreeGrid",template.HTML(table))
	this.Ctx.View("backend/category_list.html")
}

//删除分类的判断
func (this *CategoryController) CategoryDelete() {
	id, err := strconv.Atoi(this.Ctx.FormValue("id"))
	if err != nil || id == 0 {
		helper.Ajax("参数错误", 1, this.Ctx)
		return
	}
	//判断是否可以删除
	if len(models.NewCategoryModel(this.Orm).GetNextCategory(int64(id))) > 0 {
		helper.Ajax("有下级分类，不可删除", 1, this.Ctx)
		return
	}
	//todo 相关文章和单页分类
	// code to do something...
	if models.NewCategoryModel(this.Orm).DeleteById(int64(id)) {
		helper.Ajax("删除分类成功", 0, this.Ctx)
	} else {
		helper.Ajax("删除分类失败", 1, this.Ctx)
	}
}

func (this *CategoryController) CategoryAdd() {
	var parentid int
	var err error
	if this.Ctx.Method() == "POST" {
		parentid, err = strconv.Atoi(this.Ctx.FormValue("parentid"))
		cattype, _ := strconv.Atoi(this.Ctx.FormValue("type"))
		ismenu, _ := strconv.Atoi(this.Ctx.FormValue("ismenu"))
		if err != nil {
			helper.Ajax(err.Error(), 1, this.Ctx)
			return
		}
		category := tables.IriscmsCategory{
			Catname:     this.Ctx.FormValue("catname"),
			Parentid:    int64(parentid),
			Type:        int64(cattype),
			Thumb:       this.Ctx.FormValue("thumb"),
			Url:         this.Ctx.FormValue("url"),
			Description: this.Ctx.FormValue("description"),
			Ismenu:      int64(ismenu),
		}
		if !models.NewCategoryModel(this.Orm).AddCategory(category) {
			helper.Ajax("添加分类失败", 1, this.Ctx)
		} else {
			helper.Ajax("添加分类成功", 0, this.Ctx)
		}
		return
	}
	parentid, err = this.Ctx.URLParamInt("parentid")
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	this.Ctx.ViewData("typelist",[]string{
		0: "栏目",
		1: "页面",
		2: "链接",
	})
	this.Ctx.ViewData("parentid",parentid)
	this.Ctx.View("backend/category_add.html")
}

func (this *CategoryController) CategorySelect() {
	cats := []map[string]interface{}{{
		"id":       0,
		"text":     "作为一级栏目",
		"children": models.NewCategoryModel(this.Orm).GetSelectTree(0),
	}}
	this.Ctx.JSON(cats)
}

func (this *CategoryController) CategoryEdit() {
	id, err := this.Ctx.URLParamInt64("id")
	if err != nil || id == 0 {
		this.Ctx.WriteString("参数错误")
		return
	}
	category, err := models.NewCategoryModel(this.Orm).GetCategory(id)
	if err != nil {
		this.Ctx.WriteString("没有找到指定的分类")
		return
	}
	if this.Ctx.Method() == "POST" {
		parentid, err := strconv.Atoi(this.Ctx.FormValue("parentid"))
		cattype, _ := strconv.Atoi(this.Ctx.FormValue("type"))
		ismenu, _ := strconv.Atoi(this.Ctx.FormValue("ismenu"))
		if err != nil {
			helper.Ajax(err.Error(), 1, this.Ctx)
			return
		}
		if category.Catid == int64(parentid) {
			helper.Ajax("不能设置自己为父级分类", 1, this.Ctx)
			return
		}
		//递归查找是否修改的父类是不是自己的子类
		if category.Parentid != int64(parentid) && models.NewCategoryModel(this.Orm).IsSonCategory(category.Catid, int64(parentid)) {
			helper.Ajax("不能把父级分类设置到子类", 1, this.Ctx)
			return
		}
		category.Ismenu = int64(ismenu)
		category.Catname = this.Ctx.FormValue("catname")
		category.Parentid = int64(parentid)
		category.Type = int64(cattype)
		category.Thumb = this.Ctx.FormValue("thumb")
		category.Url = this.Ctx.FormValue("url")
		category.Description = this.Ctx.FormValue("description")
		if !models.NewCategoryModel(this.Orm).UpdateCategory(category) {
			helper.Ajax("修改分类失败", 1, this.Ctx)
		} else {
			helper.Ajax("修改分类成功", 0, this.Ctx)
		}
		return
	}

	this.Ctx.ViewData("category",category)
	this.Ctx.ViewData("typelist",[]string{
		0: "栏目",
		1: "页面",
		2: "链接",
	})
	this.Ctx.View("backend/category_edit.html")
}

func (this *CategoryController) CategoryOrder() {
	data := this.Ctx.FormValues()
	order, ok := data["order"]
	if !ok {
		helper.Ajax("参数错误", 1, this.Ctx)
		return
	}
	id, ok := data["id"]
	if !ok {
		helper.Ajax("参数错误", 1, this.Ctx)
		return
	}
	orm := models.NewCategoryModel(this.Orm)
	for i := 0; i < len(order); i++ {
		catid, err := strconv.Atoi(id[i])
		if err != nil {
			continue
		}
		orderNum, err := strconv.Atoi(order[i])
		if err != nil {
			continue
		}
		orm.UpdateCategory(tables.IriscmsCategory{Catid: int64(catid), Listorder: int64(orderNum)})
	}
	helper.Ajax("更新栏目成功", 0, this.Ctx)
}
