package backend

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/application/controllers"
	"regexp"
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
	b.GET("/category/list", "CategoryList")
	b.POST("/category/category-add", "CategoryAdd")
	b.ANY("/category/category-edit", "CategoryEdit")
	b.ANY("/category/category-select", "CategorySelect")
	b.ANY("/category/category-delete", "CategoryDelete")
	b.ANY("/category/category-order", "CategoryOrder")
}

func (c *CategoryController) CategoryList() {
	list := models.NewCategoryModel().GetTree(models.NewCategoryModel().GetAll(false), 0)
	helper.Ajax(list, 0, c.Ctx())
}

//删除分类的判断
func (c *CategoryController) CategoryDelete() {
	id, err := strconv.Atoi(c.Ctx().FormValue("id"))
	if err != nil || id == 0 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	model := models.NewCategoryModel()
	category := model.GetCategory(int64(id))
	if category == nil {
		helper.Ajax("分类不存在或已删除", 1, c.Ctx())
		return
	}
	//判断是否可以删除
	if len(model.GetNextCategory(int64(id))) > 0 {
		helper.Ajax("有下级分类，不可删除", 1, c.Ctx())
		return
	}
	document := models.NewDocumentModel().GetByID(category.ModelId)
	if document == nil || document.Id <= 0 {
		if models.NewCategoryModel().DeleteById(int64(id)) {
			if category.Type == 1 {
				// 删除单页
				models.NewPageModel().DelPage(int64(id))
			}
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

	parentid, _ = strconv.Atoi(c.Ctx().FormValue("parentid"))
	ModelID, _ = strconv.Atoi(c.Ctx().FormValue("model_id"))
	cattype, _ := strconv.Atoi(c.Ctx().FormValue("type"))
	ismenu, _ := c.Ctx().PostBool("ismenu")
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
	category := tables.Category{
		Catname:     c.Ctx().FormValue("catname"),
		Parentid:    int64(parentid),
		Type:        int64(cattype),
		ModelId:     int64(ModelID),
		Thumb:       c.Ctx().FormValue("thumb"),
		Dir:         c.Ctx().FormValue("dir"),
		Url:         url,
		ListTpl:     helper.EasyUiIDToFilePath(c.Ctx().FormValue("list_tpl")),
		DetailTpl:   helper.EasyUiIDToFilePath(c.Ctx().FormValue("detail_tpl")),
		Keywords:    c.Ctx().FormValue("keywords"),
		Description: c.Ctx().FormValue("description"),
	}

	if ismenu {
		category.Ismenu = 1
	}

	// 检查dir的合法性
	if category.Dir != "" && !regexp.MustCompile("^[A-Za-z0-9_-]+$").MatchString(category.Dir) {
		helper.Ajax("静态目录参数错误", 1, c.Ctx())
		return
	}
	m := models.NewCategoryModel()

	if category.Parentid != 0 {
		parentCat := m.GetCategory(category.Parentid)
		if category.Topid > 0 {
			category.Topid = parentCat.Topid
		} else {
			category.Topid = category.Parentid
		}
	}

	if !m.AddCategory(category) {
		helper.Ajax("添加分类失败", 1, c.Ctx())
	} else {
		helper.Ajax("添加分类成功", 0, c.Ctx())
	}

}

func (c *CategoryController) CategorySelect() {
	cats := []map[string]interface{}{{
		"value":    0,
		"label":    "作为一级栏目",
		"children": models.NewCategoryModel().GetSelectTree(0),
	}}
	helper.Ajax(cats, 0, c.Ctx())
}

func (c *CategoryController) CategoryEdit(icache cache.AbstractCache) {
	id, err := c.Ctx().GetInt64("id")
	if err != nil || id == 0 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	category := models.NewCategoryModel().GetCategory(id)
	if category == nil {
		helper.Ajax("没有找到指定的分类", 1, c.Ctx())
		return
	}
	parentid, _ := strconv.Atoi(c.Ctx().FormValue("parentid"))
	ismenu, _ := c.Ctx().PostBool(c.Ctx().FormValue("ismenu"))
	if category.Catid == int64(parentid) {
		helper.Ajax("不能设置自己为父级分类", 1, c.Ctx())
		return
	}
	//递归查找是否修改的父类是不是自己的子类
	if category.Parentid != int64(parentid) && models.NewCategoryModel().IsSonCategory(category.Catid, int64(parentid)) {
		helper.Ajax("不能把父级分类设置到子类", 1, c.Ctx())
		return
	}
	if ismenu {
		category.Ismenu = 1
	}
	category.Catname = c.Ctx().FormValue("catname")
	category.Parentid = int64(parentid)
	category.Thumb = c.Ctx().FormValue("thumb")
	category.Description = c.Ctx().FormValue("description")
	category.Keywords = c.Ctx().FormValue("keywords")
	category.Dir = c.Ctx().FormValue("dir")
	category.Url = c.Ctx().FormValue("url")
	category.ListTpl = helper.EasyUiIDToFilePath(c.Ctx().FormValue("list_tpl"))
	category.DetailTpl = helper.EasyUiIDToFilePath(c.Ctx().FormValue("detail_tpl"))

	// 检查dir的合法性
	if category.Dir != "" && !regexp.MustCompile("^[A-Za-z0-9_-]+$").MatchString(category.Dir) {
		helper.Ajax("静态目录参数错误", 1, c.Ctx())
		return
	}

	if !models.NewCategoryModel().UpdateCategory(category) {
		helper.Ajax("修改分类失败", 1, c.Ctx())
	} else {
		icache.Delete(fmt.Sprintf(controllers.CacheCategoryInfoPrefix, id))
		helper.Ajax("修改分类成功", 0, c.Ctx())
	}
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
		orm.UpdateCategory(&tables.Category{Catid: int64(catid), Listorder: int64(orderNum)})
	}
	helper.Ajax("更新栏目成功", 0, c.Ctx())
}
