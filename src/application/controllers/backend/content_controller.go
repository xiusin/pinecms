package backend

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/spf13/cast"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/models"

	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/common/search"
)

type ContentController struct {
	BaseController
}

func (c *ContentController) Construct() {
	c.Group = "内容管理"
	c.KeywordsSearch = []SearchFieldDsl{
		{Field: "title", Op: "LIKE", DataExp: "%$?%"},
		{Field: "keyword", Op: "LIKE", DataExp: "%$?%"},
		{Field: "description", Op: "LIKE", DataExp: "%$?%"},
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
		var contents []map[string]any
		if p.Size == 0 {
			err = query.Find(&contents)
		} else {
			count, err = query.Limit(p.Size, (p.Page-1)*p.Size).FindAndCount(&contents)
		}
		if err != nil {
			helper.Ajax(err.Error(), 1, c.Ctx())
			return
		}
		for i, content := range contents {
			for field, value := range content {
				switch value := value.(type) {
				case []byte:
					content[field] = any(helper.Bytes2String(value))
				}
			}
			contents[i] = content
		}
		if contents == nil {
			contents = []map[string]any{}
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

func (c *ContentController) PostAdd() {
	mid, _ := c.Input().GetInt("mid")
	catid, _ := c.Input().GetInt("catid")
	if mid < 1 || catid < 1 {
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

	var data = map[string]any{}
	_ = c.Ctx().BindJSON(&data)
	data["created_time"] = helper.NowDate(helper.TimeFormat)
	data["updated_time"] = helper.NowDate(helper.TimeFormat)

	session := c.Orm.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		helper.Ajax("开启事务失败: "+err.Error(), 1, c.Ctx())
		return
	}

	fields := make([]string, 0, len(data))
	args := make([]any, 0, len(data))
	for k, v := range data {
		fields = append(fields, k)
		args = append(args, v)
	}
	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", c.Table, strings.Join(fields, ","), strings.TrimRight(strings.Repeat("?,", len(data)), ","))
	result, err := session.Exec(append([]any{sql}, args...)...)
	if err != nil {
		session.Rollback()
		helper.Ajax("更新内容失败: "+err.Error(), 1, c.Ctx())
		return
	}
	id, _ := result.LastInsertId()
	data["id"] = id
	eid, err := di.MustGet(controllers.ServiceSearchName).(search.ISearch).Index("document", data)
	if err != nil {
		pine.Logger().Error("数据入search失败", err, "id: ", id)
	}
	_, err = session.Table(controllers.GetTableName("search_m_a_e")).Insert(map[string]any{"mid": mid, "aid": id, "eid": eid})
	if err != nil {
		session.Rollback()
		helper.Ajax("更新内容失败: "+err.Error(), 1, c.Ctx())
		return
	}
	session.Commit()
	helper.Ajax("更新内容成功", 0, c.Ctx())
}

// PostEdit 编辑内容
func (c *ContentController) PostEdit() {
	id, _ := c.Input().GetInt("id")
	mid, _ := c.Input().GetInt("mid")
	catid, _ := c.Input().GetInt("catid")
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

	var data = map[string]any{}
	helper.PanicErr(c.Ctx().BindJSON(&data))
	if len(cast.ToString(data["deleted_time"])) == 0 {
		delete(data, "deleted_time")
	}
	if len(cast.ToString(data["pubtime"])) == 0 {
		delete(data, "pubtime")
	}
	data["updated_time"] = helper.NowDate(helper.TimeFormat)

	session := c.Orm.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		helper.Ajax("开启事务失败: "+err.Error(), 1, c.Ctx())
		return
	}

	_, err = session.Table(c.Table).Where("id = ?", id).Where("mid = ?", mid).Where("catid = ?", catid).AllCols().Update(&data)
	if err != nil {
		session.Rollback()
		helper.Ajax("更新内容失败: "+err.Error(), 1, c.Ctx())
		return
	}

	engine := di.MustGet(controllers.ServiceSearchName).(search.ISearch)
	var s tables.SearchMAE
	ok, _ := session.Table(controllers.GetTableName("search_m_a_e")).Where("mid = ? and aid = ?", mid, id).Get(&s)
	if ok {
		err = engine.Update("document", s.Eid, data)
	} else {
		var index string
		index, err = engine.Index("document", data)
		if err == nil {
			_, err = session.Table(controllers.GetTableName("search_m_a_e")).Insert(map[string]any{"mid": mid, "aid": id, "eid": index})
		}
	}

	if err != nil {
		session.Rollback()
		pine.Logger().Error("保存数据到search失败", err)
		helper.Ajax("更新内容失败: "+err.Error(), 1, c.Ctx())
		return
	}
	session.Commit()
	helper.Ajax("更新内容成功", 0, c.Ctx())
}

func (c *ContentController) GetInfo() {
	var mid, _ = c.Ctx().Input().GetInt("mid")
	var id, _ = c.Ctx().Input().GetInt("id")
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
			contents[0][field] = any(helper.Bytes2String(value))
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

	session := c.Orm.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		helper.Ajax("开启事务失败: "+err.Error(), 1, c.Ctx())
		return
	}

	ret, err := session.Table(c.Table).In(c.TableKey, ids.Ids).Delete(c.Table)
	if err != nil {
		session.Rollback()
		helper.Ajax(err.Error(), 1, c.Ctx())
		return
	}
	if ret == 0 {
		session.Rollback()
		helper.Ajax("删除失败", 1, c.Ctx())
		return
	}
	engine := di.MustGet(controllers.ServiceSearchName).(search.ISearch)
	var s []tables.SearchMAE
	session.Table(controllers.GetTableName("search_m_a_e")).Where("mid = ?", mid).In("aid", ids.Ids).Find(&s)
	for _, v := range s {
		engine.Delete("document", v.Eid)
		_, err = session.Table(controllers.GetTableName("search_m_a_e")).Where("mid = ?", mid).In("aid", ids.Ids).Delete()
		if err != nil {
			session.Rollback()
			helper.Ajax("删除失败: "+err.Error(), 1, c.Ctx())
			return
		}
	}
	session.Commit()
	helper.Ajax("删除成功", 0, c.Ctx())
}

func (c *ContentController) GetPage() {
	catid, _ := c.Ctx().Input().GetInt64("id")
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
		var document = map[string]any{}
		structs.FillMap(&page, document)
		var s tables.SearchMAE
		engine := di.MustGet(controllers.ServiceSearchName).(search.ISearch)
		ok, _ := c.Orm.Table(controllers.GetTableName("search_m_a_e")).Where("mid = ? and aid = ?", 0, page.Id).Get(&s)
		if ok {
			engine.Update("pine_cms_system_page", s.Eid, document)
		} else {
			eid, _ := engine.Index("pine_cms_system_page", document)
			c.Orm.Table(controllers.GetTableName("search_m_a_e")).Insert(map[string]any{"aid": page.Id, "eid": eid})
		}
		helper.Ajax("更新单页成功", 0, c.Ctx())
	} else {
		helper.Ajax("更新单页失败", 1, c.Ctx())
	}
}
