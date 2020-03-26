package backend

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"html/template"
	"strconv"

	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type CategoryController struct {
	pine.Controller
}

func (c *CategoryController) RegisterRoute(b pine.IRouterWrapper) {
	//分类相关
	b.ANY("/category/list", "CategoryList")
	b.ANY("/category/category-add", "CategoryAdd")
	b.ANY("/category/category-edit", "CategoryEdit")
	b.ANY("/category/category-select", "CategorySelect")
	b.ANY("/category/category-delete", "CategoryDelete")
	b.ANY("/category/category-order", "CategoryOrder")
}

func (c *CategoryController) CategoryList() {
	if c.Ctx().URLParam("grid") == "treegrid" {
		c.Ctx().Render().JSON(models.NewCategoryModel().GetTree(models.NewCategoryModel().GetAll(), 0))
		return
	}
	menuid, _ := c.Ctx().URLParamInt64("menuid")
	table := helper.Treegrid("category_categorylist_treegrid", "/b/category/list?grid=treegrid", helper.EasyuiOptions{
		"title":     models.NewMenuModel().CurrentPos(menuid),
		"toolbar":   "category_categorylist_treegrid_toolbar",
		"idField":   "catid",
		"treeField": "catname",
	}, helper.EasyuiGridfields{
		"排序":   {"field": "listorder", "width": "10", "align": "center", "formatter": "categoryCategoryListOrderFormatter", "index": "0"},
		"栏目名称": {"field": "catname", "width": "20", "index": "1"},
		"栏目类型": {"field": "type", "width": "10", "formatter": "categoryCategoryListTypeFormatter", "index": "2"},
		"栏目模型": {"field": "model_id", "width": "10", "index": "3"},
		"状态":   {"field": "ismenu", "width": "10", "formatter": "categoryCategoryListStateFormatter", "index": "4"},
		"管理操作": {"field": "catid", "width": "40",  "formatter": "categoryCategoryListOperateFormatter", "index": "5"},
	})
	c.Ctx().Render().ViewData("TreeGrid", template.HTML(table))
	c.Ctx().Render().HTML("backend/category_list.html")
}

//删除分类的判断
func (c *CategoryController) CategoryDelete() {
	id, err := strconv.Atoi(c.Ctx().FormValue("id"))
	if err != nil || id == 0 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	model := models.NewCategoryModel()
	category, _ := model.GetCategory(int64(id))
	if category.Catid <= 0 {
		helper.Ajax("分类不存在或已删除", 1, c.Ctx())
		return
	}
	//判断是否可以删除
	if len(model.GetNextCategory(int64(id))) > 0 {
		helper.Ajax("有下级分类，不可删除", 1, c.Ctx())
		return
	}
	document := models.NewDocumentModel().GetByID(category.ModelId)
	if document.Id <= 0 {
		if models.NewCategoryModel().DeleteById(int64(id)) {
			helper.Ajax("删除分类成功", 0, c.Ctx())
		} else {
			helper.Ajax("删除分类失败", 1, c.Ctx())
		}
		return
	}
	// 查询文档分类
	sql := []interface{}{fmt.Sprintf("SELECT COUNT(*) total FROM `%s` WHERE catid=? and deleted_time IS NULL", controllers.GetTableName(document.Table)), id}
	totals, _ := c.Ctx().Value("orm").(*xorm.Engine).QueryString(sql...)
	var total = "0"
	if len(totals) > 0 {
		total = totals[0]["total"]
	}
	if total == "0" {
		if models.NewCategoryModel().DeleteById(int64(id)) {
			helper.Ajax("删除分类成功", 0, c.Ctx())
		} else {
			helper.Ajax("删除分类失败", 1, c.Ctx())
		}
	} else {
		helper.Ajax("分类下有文章，无法删除", 1, c.Ctx())
	}

}

func (c *CategoryController) CategoryAdd() {
	var parentid int
	var ModelID int
	var err error
	if c.Ctx().IsPost() {
		parentid, err = strconv.Atoi(c.Ctx().FormValue("parentid"))
		ModelID, err = strconv.Atoi(c.Ctx().FormValue("model_id"))
		cattype, _ := strconv.Atoi(c.Ctx().FormValue("type"))
		var ismenu int64
		if c.Ctx().FormValue("ismenu") == "on" {
			ismenu = 1
		}
		if err != nil {
			helper.Ajax(err.Error(), 1, c.Ctx())
			return
		}
		manager, add, edit := "", "", ""
		url := c.Ctx().FormValue("url")
		switch cattype {
		case 0:
			url = ""
			if ModelID == 0 {
				add = c.Ctx().FormValue("add_content_router")
				edit = c.Ctx().FormValue("edit_content_router")
				manager = c.Ctx().FormValue("manager_content_router")
			}
		case 1:
			url = ""
			ModelID = 0
		case 2:
			ModelID = 0
		}

		category := tables.Category{
			Catname:              c.Ctx().FormValue("catname"),
			Parentid:             int64(parentid),
			Type:                 int64(cattype),
			ModelId:              int64(ModelID),
			AddContentRouter:     add,
			EditContentRouter:    edit,
			ManagerContentRouter: manager,
			Thumb:                c.Ctx().FormValue("thumb"),
			Dir:                  c.Ctx().FormValue("dir"),
			Url:                  url,
			Description:          c.Ctx().FormValue("description"),
			Ismenu:               ismenu,
		}
		if !models.NewCategoryModel().AddCategory(category) {
			helper.Ajax("添加分类失败", 1, c.Ctx())
		} else {
			helper.Ajax("添加分类成功", 0, c.Ctx())
		}
		return
	}
	parentid, err = c.Ctx().URLParamInt("parentid")
	if err != nil {
		c.Ctx().WriteString(err.Error())
		return
	}
	c.Ctx().Render().ViewData("typelist", []string{0: "栏目", 1: "页面", 2: "链接"})

	// 查询模型
	list, _ := models.NewDocumentModel().GetList(1, 1000)
	c.Ctx().Render().ViewData("models", list)
	c.Ctx().Render().ViewData("parentid", parentid)
	c.Ctx().Render().ViewData("thumbHtml", template.HTML(helper.SiginUpload("thumb", "", false, "", "", "")))
	c.Ctx().Render().HTML("backend/category_add.html")
}

func (c *CategoryController) CategorySelect() {
	cats := []map[string]interface{}{{
		"id":       0,
		"text":     "作为一级栏目",
		"children": models.NewCategoryModel().GetSelectTree(0),
	}}
	c.Ctx().Render().JSON(cats)
}

func (c *CategoryController) CategoryEdit() {
	id, err := c.Ctx().URLParamInt64("id")
	if err != nil || id == 0 {
		c.Ctx().WriteString("参数错误")
		return
	}
	category, err := models.NewCategoryModel().GetCategory(id)
	if err != nil {
		c.Ctx().WriteString("没有找到指定的分类")
		return
	}
	if c.Ctx().IsPost() {
		parentid, err := strconv.Atoi(c.Ctx().FormValue("parentid"))
		var ismenu int64
		if c.Ctx().FormValue("ismenu") == "on" {
			ismenu = 1
		}
		if err != nil {
			helper.Ajax(err.Error(), 1, c.Ctx())
			return
		}
		if category.Catid == int64(parentid) {
			helper.Ajax("不能设置自己为父级分类", 1, c.Ctx())
			return
		}
		//递归查找是否修改的父类是不是自己的子类
		if category.Parentid != int64(parentid) && models.NewCategoryModel().IsSonCategory(category.Catid, int64(parentid)) {
			helper.Ajax("不能把父级分类设置到子类", 1, c.Ctx())
			return
		}

		category.Ismenu = ismenu
		category.Catname = c.Ctx().FormValue("catname")
		category.Parentid = int64(parentid)
		category.Thumb = c.Ctx().FormValue("thumb")
		category.Description = c.Ctx().FormValue("description")
		category.Dir = c.Ctx().FormValue("dir")
		if category.ModelId == 0 && category.Type == 0 {
			category.AddContentRouter = c.Ctx().FormValue("add_content_router")
			category.EditContentRouter = c.Ctx().FormValue("edit_content_router")
			category.ManagerContentRouter = c.Ctx().FormValue("manager_content_router")
		}

		if !models.NewCategoryModel().UpdateCategory(category) {
			helper.Ajax("修改分类失败", 1, c.Ctx())
		} else {
			helper.Ajax("修改分类成功", 0, c.Ctx())
		}
		return
	}
	// 查询模型
	list, _ := models.NewDocumentModel().GetList(1, 1000)
	c.Ctx().Render().ViewData("models", list)
	c.Ctx().Render().ViewData("model_id", int(category.ModelId))
	c.Ctx().Render().ViewData("category", category)
	c.Ctx().Render().ViewData("thumbHtml", template.HTML(helper.SiginUpload("thumb", category.Thumb, false, "", "", "")))
	c.Ctx().Render().ViewData("typelist", []string{0: "栏目", 1: "页面", 2: "链接"})
	c.Ctx().Render().HTML("backend/category_edit.html")
}

func (c *CategoryController) CategoryOrder() {
	data := c.Ctx().PostData()
	order, ok := data["order"]
	if !ok {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	id, ok := data["id"]
	if !ok {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	orm := models.NewCategoryModel()
	for i := 0; i < len(order); i++ {
		catid, err := strconv.Atoi(id[i])
		if err != nil {
			continue
		}
		orderNum, err := strconv.Atoi(order[i])
		if err != nil {
			continue
		}
		orm.UpdateCategory(tables.Category{Catid: int64(catid), Listorder: int64(orderNum)})
	}
	helper.Ajax("更新栏目成功", 0, c.Ctx())
}
