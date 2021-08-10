package backend

import (
	"errors"
	"fmt"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"regexp"
)

type CategoryController struct {
	BaseController
	sql string
}

func (c *CategoryController) Construct() {
	c.Table = &tables.Category{}
	c.Entries = &[]*tables.Category{}
	c.Group = "内容管理"
	c.SubGroup = "分类管理"
	c.ApiEntityName = "分类"
	c.OpBefore = c.before
	c.sql = "SELECT COUNT(*) total FROM `%s` WHERE id=? and deleted_time IS NULL"
	c.BaseController.Construct()
	c.TableStructKey = "Catid"
}

func (c *CategoryController) before(act int, params interface{}) error {
	if act == OpDel {
		ids := params.(*idParams)
		if len(ids.Ids) > 1 {
			return errors.New("分类不支持批量删除")
		}
		ok, _ := c.Orm.In("parent_id", ids.Ids).Exist(c.Table)
		if ok {
			return errors.New("有下级分类，不可删除")
		}
		cat := models.NewCategoryModel().GetCategory(ids.Ids[0])
		document := models.NewDocumentModel().GetByID(cat.ModelId)
		if document == nil || document.Id <= 0 {
			models.NewCategoryModel().DeleteById(cat.Catid)
			if cat.Type == 1 {
				models.NewPageModel().DelPage(cat.Catid)
			}
			return nil
		}
		sql := []interface{}{fmt.Sprintf(c.sql, controllers.GetTableName(document.Table)), cat.Catid}
		totals, _ := c.Orm.QueryString(sql...)
		var total = "0"
		if len(totals) > 0 {
			total = totals[0]["total"]
		}
		if total != "0" {
			return errors.New("分类下有文章，无法删除")
		}
	} else if act == OpAdd {
		cat := params.(*tables.Category)
		switch cat.Type {
		case 0:
			cat.Url = ""
		case 1:
			cat.Url = ""
			cat.ModelId = 0
		case 2:
			cat.ModelId = 0
		}
		if cat.Parentid != 0 {
			parentCat := models.NewCategoryModel().GetCategory(cat.Parentid)
			if cat.Topid > 0 {
				cat.Topid = parentCat.Topid
			} else {
				cat.Topid = cat.Parentid
			}
		}
	} else if act == OpEdit {
		cat := params.(*tables.Category)
		if cat.Dir != "" && !regexp.MustCompile("^[A-Za-z0-9_-]+$").MatchString(cat.Dir) {
			return errors.New("静态目录参数错误")
		}
	}
	return nil
}

func (c *CategoryController) GetSelect() {
	_ = c.Orm.OrderBy("listorder").Find(c.Entries)
	m := c.Entries.(*[]*tables.Category)
	var kv []tables.KV
	for _, model := range *m {
		kv = append(kv, tables.KV{
			Label: model.Catname,
			Value: model.Catid,
		})
	}
	helper.Ajax(kv, 0, c.Ctx())
}
