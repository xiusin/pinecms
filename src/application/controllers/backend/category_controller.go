package backend

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"html/template"
	"strconv"

	"github.com/xiusin/iriscms/src/application/models"
	"github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/iriscms/src/common/helper"
)

type CategoryController struct {
	pine.Controller
}

func (c *CategoryController) RegisterRoute(b pine.IRouterWrapper) {
	//分类相关
	b.ANY( "/category/list", "CategoryList")
	b.ANY( "/category/category-add", "CategoryAdd")
	b.ANY( "/category/category-edit", "CategoryEdit")
	b.ANY( "/category/category-select", "CategorySelect")
	b.ANY( "/category/category-delete", "CategoryDelete")
	b.ANY( "/category/category-order", "CategoryOrder")
}

func (c *CategoryController) CategoryList() {
	if c.Ctx().URLParam("grid") == "treegrid" {
		c.Ctx().Render().JSON(models.NewCategoryModel(c.Ctx().Value("orm").(*xorm.Engine)).GetTree(models.NewCategoryModel(c.Ctx().Value("orm").(*xorm.Engine)).GetAll(), 0))
		return
	}
	menuid, _ := c.Ctx().URLParamInt64("menuid")
	table := helper.Treegrid("category_categorylist_treegrid", "/b/category/list?grid=treegrid", helper.EasyuiOptions{
		"title":     models.NewMenuModel(c.Ctx().Value("orm").(*xorm.Engine)).CurrentPos(menuid),
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
	model := models.NewCategoryModel(c.Ctx().Value("orm").(*xorm.Engine))
	category,_ := model.GetCategory(int64(id))
	if category.Catid <= 0 {
		helper.Ajax("分类不存在或已删除", 1, c.Ctx())
		return
	}
	//判断是否可以删除
	if len(model.GetNextCategory(int64(id))) > 0 {
		helper.Ajax("有下级分类，不可删除", 1, c.Ctx())
		return
	}
	// 是否有文章
	document := models.NewDocumentModel(c.Ctx().Value("orm").(*xorm.Engine)).GetByID(category.ModelId)
	if document.Id <=0 {
		helper.Ajax("分类文档模型不存在", 1, c.Ctx())
		return
	}
	// 查询文档分类
	sql := []interface{}{fmt.Sprintf("SELECT COUNT(*) total FROM `iriscms_%s` WHERE catid=? and deleted_time IS NULL", document.Table), id}
	totals, _ := c.Ctx().Value("orm").(*xorm.Engine).QueryString(sql...)
	var total = "0"
	if len(totals) > 0 {
		total = totals[0]["total"]
	}
	if total == "0" {
		if models.NewCategoryModel(c.Ctx().Value("orm").(*xorm.Engine)).DeleteById(int64(id)) {
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
		ismenu, _ := strconv.Atoi(c.Ctx().FormValue("ismenu"))
		if err != nil {
			helper.Ajax(err.Error(), 1, c.Ctx())
			return
		}
		url := c.Ctx().FormValue("url")
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
			Catname:     c.Ctx().FormValue("catname"),
			Parentid:    int64(parentid),
			Type:        int64(cattype),
			ModelId:     int64(ModelID),
			Thumb:       c.Ctx().FormValue("thumb"),
			Url:         url,
			Description: c.Ctx().FormValue("description"),
			Ismenu:      int64(ismenu),
		}
		if !models.NewCategoryModel(c.Ctx().Value("orm").(*xorm.Engine)).AddCategory(category) {
			helper.Ajax("添加分类失败", 1, c.Ctx())
		} else {
			//cacheKey := fmt.Sprintf(controllers.CacheCategoryFormat, parentid)
			//if c.cache.IsExist(cacheKey) {
			//	if c.cache.Delete(cacheKey) != nil {
			//		golog.Error("刷新列表缓存失败")
			//	}
			//}

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
	list, _ := models.NewDocumentModel(c.Ctx().Value("orm").(*xorm.Engine)).GetList(1, 1000)
	c.Ctx().Render().ViewData("models", list)
	c.Ctx().Render().ViewData("parentid", parentid)
	c.Ctx().Render().HTML("backend/category_add.html")
}

func (c *CategoryController) CategorySelect() {
	cats := []map[string]interface{}{{
		"id":       0,
		"text":     "作为一级栏目",
		"children": models.NewCategoryModel(c.Ctx().Value("orm").(*xorm.Engine)).GetSelectTree(0),
	}}
	c.Ctx().Render().JSON(cats)
}

func (c *CategoryController) CategoryEdit() {
	id, err := c.Ctx().URLParamInt64("id")
	if err != nil || id == 0 {
		c.Ctx().WriteString("参数错误")
		return
	}
	category, err := models.NewCategoryModel(c.Ctx().Value("orm").(*xorm.Engine)).GetCategory(id)
	if err != nil {
		c.Ctx().WriteString("没有找到指定的分类")
		return
	}
	if c.Ctx().IsPost() {
		parentid, err := strconv.Atoi(c.Ctx().FormValue("parentid"))
		ismenu, _ := strconv.Atoi(c.Ctx().FormValue("ismenu"))
		if err != nil {
			helper.Ajax(err.Error(), 1, c.Ctx())
			return
		}
		if category.Catid == int64(parentid) {
			helper.Ajax("不能设置自己为父级分类", 1, c.Ctx())
			return
		}
		//递归查找是否修改的父类是不是自己的子类
		if category.Parentid != int64(parentid) && models.NewCategoryModel(c.Ctx().Value("orm").(*xorm.Engine)).IsSonCategory(category.Catid, int64(parentid)) {
			helper.Ajax("不能把父级分类设置到子类", 1, c.Ctx())
			return
		}

		category.Ismenu = int64(ismenu)
		category.Catname = c.Ctx().FormValue("catname")
		category.Parentid = int64(parentid)
		category.Thumb = c.Ctx().FormValue("thumb")
		category.Description = c.Ctx().FormValue("description")

		if !models.NewCategoryModel(c.Ctx().Value("orm").(*xorm.Engine)).UpdateCategory(category) {
			helper.Ajax("修改分类失败", 1, c.Ctx())
		} else {
			helper.Ajax("修改分类成功", 0, c.Ctx())
		}
		return
	}
	// 查询模型
	list, _ := models.NewDocumentModel(c.Ctx().Value("orm").(*xorm.Engine)).GetList(1, 1000)
	c.Ctx().Render().ViewData("models", list)
	c.Ctx().Render().ViewData("model_id", int(category.ModelId))
	c.Ctx().Render().ViewData("category", category)
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
	orm := models.NewCategoryModel(c.Ctx().Value("orm").(*xorm.Engine))
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
	helper.Ajax("更新栏目成功", 0, c.Ctx())
}
