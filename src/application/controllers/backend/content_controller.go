package backend

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/xiusin/pinecms/src/application/models"

	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type ContentController struct {
	BaseController
}

func (c *ContentController) Construct() {
	c.Group = "内容管理"
	c.KeywordsSearch = []SearchFieldDsl{
		{Field: "value", Op: "LIKE", DataExp: "%$?%"},
		{Field: "name", Op: "LIKE", DataExp: "%$?%"},
	}
	c.SearchFields = []SearchFieldDsl{
		{Op: "=", Field: "status"},
	}
	c.Entries = &[]*tables.DocumentModel{}
	c.ApiEntityName = "内容"
	c.BaseController.Construct()
}

func (c *ContentController) PostList() {
	catid, _ := c.Input().GetInt64("cid")
	var category tables.Category
	_, _ = c.Orm.Where("id = ?", catid).Get(&category)
	if category.Catid == 0 {
		helper.Ajax("栏目不存在或已删除", 1, c.Ctx())
		return
	}
	if category.ModelId < 1 {
		helper.Ajax("栏目模型不存在", 1, c.Ctx())
		return
	}
	var document tables.DocumentModel
	c.Orm.Where("id = ?", category.ModelId).Get(&document)
	if document.Id == 0 {
		helper.Ajax("无法找到关联模型", 1, c.Ctx())
		return
	}
	c.Table = controllers.GetTableName(document.Table) // 设置表名

	query := c.Orm.Table(c.Table)
	if p, err := c.buildParamsForQuery(query); err != nil {
		helper.Ajax("参数错误: "+err.Error(), 1, c.Ctx())
	} else {
		var fields tables.ModelDslFields
		c.Orm.Where("mid = ?", category.ModelId).Find(&fields)
		query.Where("catid = ?", catid).OrderBy("listorder DESC").OrderBy("id DESC")
		query.Cols(fields.GetListFields()...)
		var count int64
		var contents []map[string]interface{}
		if p.Size == 0 {
			err = query.Find(c.Entries)
		} else {
			count, err = query.Limit(p.Size, (p.Page-1)*p.Size).FindAndCount(&contents)
			if err == nil {
				contents, err = query.QueryInterface()
			}
		}
		if err != nil {
			helper.Ajax("错误"+err.Error(), 1, c.Ctx())
			return
		}
		for i, content := range contents {
			for field, value := range content {
				switch value.(type) {
				case []byte:
					content[field] = interface{}(helper.Bytes2String(value.([]byte)))
				}
			}
			contents[i] = content
		}
		helper.Ajax(pine.H{
			"list": contents,
			"pagination": pine.H{
				"page":  p.Page,
				"size":  p.Size,
				"total": count,
			},
		}, 0, c.Ctx())
	}
}

// PostEdit 编辑内容
func (c *ContentController) PostEdit() {
	id, _ := c.Input().GetInt("id")
	mid, _ := c.Input().GetInt("mid")
	catid, _ := c.Input().GetInt("catid")
	fmt.Println(id, mid, catid)
	if mid < 1 || catid < 1 || id < 1 {
		helper.Ajax("缺少关键参数", 1, c.Ctx())
		return
	}

	var document tables.DocumentModel
	c.Orm.Where("id = ?", mid).Get(&document)
	if document.Id == 0 {
		helper.Ajax("模型不存在", 1, c.Ctx())
		return
	}
	c.Table = controllers.GetTableName(document.Table) // 设置表名
	query := c.Orm.Table(c.Table)

	var data = map[string]interface{}{}
	c.Ctx().BindJSON(&data)
	data["updated_time"] = helper.NowDate("Y-m-d H:i:s")
	_, err := query.Where("id = ?", id).Where("mid = ?", mid).Where("catid = ?", catid).AllCols().Update(&data)
	if err == nil {
		helper.Ajax("更新内容成功", 1, c.Ctx())
	} else {
		helper.Ajax("更新内容失败: "+err.Error(), 0, c.Ctx())
	}
}

func (c *ContentController) GetInfo() {
	var mid, _ = c.Ctx().GetInt("mid")
	var id, _ = c.Ctx().GetInt("id")
	var document tables.DocumentModel
	c.Orm.Where("id = ?", mid).Get(&document)
	if document.Id == 0 {
		helper.Ajax("模型不存在", 1, c.Ctx())
		return
	}
	c.Table = controllers.GetTableName(document.Table) // 设置表名
	query := c.Orm.Table(c.Table)
	contents, err := query.Where("id = ?", id).QueryInterface()
	if err != nil {
		helper.Ajax("错误"+err.Error(), 1, c.Ctx())
		return
	}
	for field, value := range contents[0] {
		switch value := value.(type) {
		case []byte:
			contents[0][field] = interface{}(helper.Bytes2String(value))
		}
	}
	helper.Ajax(contents[0], 0, c.Ctx())
}

func (c *ContentController) PostDelete() {
	var ids idParams
	if err := parseParam(c.Ctx(), &ids); err != nil {
		helper.Ajax("参数错误: "+err.Error(), 1, c.Ctx())
		return
	}
	mid, _ := c.Input().GetInt("mid")
	if mid < 1 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	var document tables.DocumentModel
	c.Orm.Where("id = ?", mid).Get(&document)
	if document.Id == 0 {
		helper.Ajax("模型不存在", 1, c.Ctx())
		return
	}
	c.Table = controllers.GetTableName(document.Table)
	idArr := []string{"0"}
	for _, id := range ids.Ids {
		idArr = append(idArr, strconv.Itoa(int(id)))
	}
	ret, err := c.Orm.Exec("DELETE FROM `" + c.Table.(string) + "` WHERE `" + c.TableKey + "` IN (" + strings.Join(idArr, ",") + ")")
	if err != nil {
		helper.Ajax(err.Error(), 1, c.Ctx())
		return
	}
	if rowNum, _ := ret.RowsAffected(); rowNum == 0 {
		helper.Ajax("删除失败", 1, c.Ctx())
		return
	}
	helper.Ajax("删除成功", 0, c.Ctx())
}

func (c *ContentController) GetPage() {
	catid, _ := c.Ctx().GetInt64("id")
	if catid == 0 {
		helper.Ajax("页面错误", 1, c.Ctx())
		return
	}
	pageModel := models.NewPageModel()
	page := pageModel.GetPage(catid)
	if page == nil {
		page = &tables.Page{
			Id: catid,
		}
	}
	helper.Ajax(page, 0, c.Ctx())
}

func (c *ContentController) PostPage() {
	var page tables.Page
	c.Ctx().BindJSON(&page)
	if page.Id == 0 {
		helper.Ajax("分类ID不存在", 1, c.Ctx())
		return
	}
	var ret int64
	page.UpdatedAt = tables.LocalTime(time.Now())
	page.CreatedAt = tables.LocalTime(time.Now())
	if exist, _ := c.Orm.Where("id = ?", page.Id).Exist(&tables.Page{}); exist {
		ret, _ = c.Orm.Where("id = ?", page.Id).Update(&page)
	} else {
		ret, _ = c.Orm.InsertOne(&page)
	}
	if ret > 0 {
		helper.Ajax("更新单页成功", 0, c.Ctx())
	} else {
		helper.Ajax("更新单页失败", 1, c.Ctx())
	}
}
