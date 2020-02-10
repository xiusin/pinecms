package backend

import (
	"html/template"
	"strconv"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/xiusin/iriscms/src/application/models"
	"github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/iriscms/src/common/cache"
	"github.com/xiusin/iriscms/src/common/helper"
)

type CategoryController struct {
	Ctx   iris.Context
	Orm   *xorm.Engine
	cache cache.ICache

	Session *sessions.Session
}

func (c *CategoryController) BeforeActivation(b mvc.BeforeActivation) {
	//分类相关
	b.Handle("ANY", "/category/list", "CategoryList")
	b.Handle("ANY", "/category/category-add", "CategoryAdd")
	b.Handle("ANY", "/category/category-edit", "CategoryEdit")
	b.Handle("ANY", "/category/category-select", "CategorySelect")
	b.Handle("ANY", "/category/category-delete", "CategoryDelete")
	b.Handle("ANY", "/category/category-order", "CategoryOrder")
}

func (this *CategoryController) CategoryList() {
	if this.Ctx.URLParam("grid") == "treegrid" {
		this.Ctx.JSON(models.NewCategoryModel(this.Orm).GetTree(models.NewCategoryModel(this.Orm).GetAll(), 0))
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
		"栏目名称": {"field": "catname", "width": "80", "index": "1"},
		"栏目类型": {"field": "type", "width": "30", "formatter": "categoryCategoryListTypeFormatter", "index": "2"},
		"栏目模型": {"field": "model_id", "width": "30", "index": "3"},
		"描述":   {"field": "description", "width": "80", "index": "4"},
		"状态":   {"field": "ismenu", "width": "20", "formatter": "categoryCategoryListStateFormatter", "index": "5"},
		"管理操作": {"field": "catid", "align": "center", "formatter": "categoryCategoryListOperateFormatter", "index": "6"},
	})
	this.Ctx.ViewData("TreeGrid", template.HTML(table))
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
	var ModelID int
	var err error
	if this.Ctx.Method() == "POST" {
		parentid, err = strconv.Atoi(this.Ctx.FormValue("parentid"))
		ModelID, err = strconv.Atoi(this.Ctx.FormValue("model_id"))
		cattype, _ := strconv.Atoi(this.Ctx.FormValue("type"))
		ismenu, _ := strconv.Atoi(this.Ctx.FormValue("ismenu"))
		if err != nil {
			helper.Ajax(err.Error(), 1, this.Ctx)
			return
		}
		url := this.Ctx.FormValue("url")
		switch cattype {
		case 0:
			url = ""
		case 1:
			url = ""
			ModelID = 0
		case 2:
			ModelID = 0
		}

		category := tables.IriscmsCategory{
			Catname:     this.Ctx.FormValue("catname"),
			Parentid:    int64(parentid),
			Type:        int64(cattype),
			ModelId:     int64(ModelID),
			Thumb:       this.Ctx.FormValue("thumb"),
			Url:         url,
			Description: this.Ctx.FormValue("description"),
			Ismenu:      int64(ismenu),
		}
		if !models.NewCategoryModel(this.Orm).AddCategory(category) {
			helper.Ajax("添加分类失败", 1, this.Ctx)
		} else {
			//cacheKey := fmt.Sprintf(controllers.CacheCategoryFormat, parentid)
			//if this.cache.IsExist(cacheKey) {
			//	if this.cache.Delete(cacheKey) != nil {
			//		golog.Error("刷新列表缓存失败")
			//	}
			//}

			helper.Ajax("添加分类成功", 0, this.Ctx)
		}
		return
	}
	parentid, err = this.Ctx.URLParamInt("parentid")
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	this.Ctx.ViewData("typelist", []string{
		0: "栏目",
		1: "页面",
		2: "链接",
	})

	// 查询模型
	list, _ := models.NewDocumentModel(this.Orm).GetList(1, 1000)
	this.Ctx.ViewData("models", list)
	this.Ctx.ViewData("parentid", parentid)
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
		category.Thumb = this.Ctx.FormValue("thumb")
		category.Description = this.Ctx.FormValue("description")

		if !models.NewCategoryModel(this.Orm).UpdateCategory(category) {
			helper.Ajax("修改分类失败", 1, this.Ctx)
		} else {
			helper.Ajax("修改分类成功", 0, this.Ctx)
		}
		return
	}
	// 查询模型
	list, _ := models.NewDocumentModel(this.Orm).GetList(1, 1000)
	this.Ctx.ViewData("models", list)
	this.Ctx.ViewData("model_id", int(category.ModelId))
	this.Ctx.ViewData("category", category)
	this.Ctx.ViewData("typelist", []string{
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
