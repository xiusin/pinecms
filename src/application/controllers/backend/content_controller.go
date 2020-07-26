package backend

import (
	"encoding/json"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/microcosm-cc/bluemonday"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"html/template"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/xiusin/pinecms/src/application/controllers"

	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/common/helper"
)

type ContentController struct {
	pine.Controller
}

func (c *ContentController) RegisterRoute(b pine.IRouterWrapper) {
	b.GET("/content/aside-category", "AsideCategory")
	b.ANY("/content/news-list", "NewsList")
	b.ANY("/content/news-crud", "NewsModelJson")

	b.ANY("/content/page", "Page")
	b.ANY("/content/add", "AddContent")
	b.ANY("/content/edit", "EditContent")
	b.ANY("/content/delete", "DeleteContent")
	b.ANY("/content/order", "OrderContent")
}

func (c *ContentController) AsideCategory() {
	cats := models.NewCategoryModel().GetContentRightCategoryTree(models.NewCategoryModel().GetAll(false), 0)
	helper.Ajax(cats, 0, c.Ctx())
}

func (c *ContentController) NewsList(orm *xorm.Engine) {
	catid, _ := c.Ctx().GetInt64("catid")
	page, _ := c.Ctx().GetInt64("page")
	rows, _ := c.Ctx().GetInt64("rows")
	catogoryModel := models.NewCategoryModel().GetCategory(catid)
	if catogoryModel == nil {
		helper.Ajax("分类不存在", 1, c.Ctx())
		return
	}
	if catogoryModel.ModelId < 1 {
		helper.Ajax("找不到关联模型", 1, c.Ctx())
		return
	}
	relationDocumentModel := models.NewDocumentModel().GetByID(catogoryModel.ModelId)
	if relationDocumentModel.Id == 0 {
		helper.Ajax("找不到关联模型", 1, c.Ctx())
		return
	}

	// 获取所有字段
	dslFields := models.NewDocumentFieldDslModel().GetList(catogoryModel.ModelId)
	var tMapF = map[string]string{}
	var ff []string
	for _, dsl := range dslFields {
		tMapF[dsl.TableField] = dsl.FormName
		ff = append(ff, dsl.TableField)
	}

	var showInPage = map[string]controllers.FieldShowInPageList{}
	_ = json.Unmarshal([]byte(relationDocumentModel.FieldShowInList), &showInPage)

	var flag bool
	for _, fieldInfo := range showInPage {
		if fieldInfo.Show {
			flag = true
			break
		}
	}
	if !flag {
		helper.Dialog("请配置模型字段显隐属性", c.Ctx())
		return
	}

	querySqlWhere := []string{"catid=?", "deleted_time IS NULL"}
	var whereHolder = []interface{}{catid}
	getData := c.Ctx().GetData()

	for param, values := range getData {
		if (!strings.HasPrefix(param, "search_")) || len(values) == 0 || len(values[0]) == 0 {
			continue
		}
		field := strings.TrimLeft(param, "search_")
		conf, ok := showInPage[field]
		if !ok {
			continue
		}
		if conf.Search == 1 {
			querySqlWhere = append(querySqlWhere, field+"=?")
			whereHolder = append(whereHolder, values[0])
		} else {
			querySqlWhere = append(querySqlWhere, field+" LIKE ?")
			whereHolder = append(whereHolder, "%"+values[0]+"%")
		}
	}

	offset := (page - 1) * rows
	querySql := "SELECT * FROM `%s` WHERE " + strings.Join(querySqlWhere, " AND ") + " ORDER BY listorder DESC, id DESC LIMIT %d,%d"
	sql := []interface{}{fmt.Sprintf(querySql, controllers.GetTableName(relationDocumentModel.Table), offset, rows)}
	sql = append(sql, whereHolder...)

	contents, err := orm.QueryString(sql...)
	if err != nil {
		pine.Logger().Error("请求列表错误", err)
		helper.Ajax("获取文档列表错误", 1, c.Ctx())
		return
	}

	countSql := "SELECT COUNT(*) total FROM `%s` WHERE " + strings.Join(querySqlWhere, " AND ")
	sql = []interface{}{fmt.Sprintf(countSql, controllers.GetTableName(relationDocumentModel.Table))}
	sql = append(sql, whereHolder)

	totals, _ := orm.QueryString(sql...)
	var total = "0"
	if len(totals) > 0 {
		total = totals[0]["total"]
	}
	if contents == nil {
		contents = []map[string]string{}
	}
	c.Ctx().Render().JSON(map[string]interface{}{"rows": contents, "total": total})
}

// 动态json表单
func (c *ContentController) NewsModelJson(orm *xorm.Engine) {
	catid, _ := c.Ctx().GetInt64("catid")
	catogoryModel := models.NewCategoryModel().GetCategory(catid)
	if catogoryModel == nil {
		helper.Ajax("分类不存在", 1, c.Ctx())
		return
	}
	rd := models.NewDocumentModel().GetByID(catogoryModel.ModelId)
	if rd == nil || rd.Id == 0 {
		helper.Ajax("找不到关联模型", 1, c.Ctx())
		return
	}
	var fields []tables.DocumentModelDsl
	orm.Table(new(tables.DocumentModelDsl)).Where("mid = ?", catogoryModel.ModelId).OrderBy("listorder").Find(&fields)	// 按排序查字段
	var forms []FormControl
	fm := models.NewDocumentModelFieldModel().GetMap()
	for _, field := range fields {
		form := FormControl{ // 表单显示
			Type:  fm[field.FieldType].AmisType,
			Name:  field.TableField,
			Label: field.FormName,
		}
		forms = append(forms, form)
	}
	// 构建json
	helper.Ajax(pine.H{
		"type":            "crud",
		"columns":         forms,
		"filterTogglable": true,
		"headerToolbar": []interface{}{
			"filter-toggler",
			map[string]string{
				"type":  "columns-toggler",
				"align": "left",
			},
			map[string]string{
				"type":  "pagination",
				"align": "left",
			},
			map[string]interface{}{
				"type":       "button",
				"align":      "right",
				"actionType": "drawer",
				"label":      "添加",
				"icon":       "fa fa-plus pull-left",
				"size":       "sm",
				"primary":    true,
				"drawer": map[string]interface{}{
					"position": "right",
					"size":     "xl",
					"title":    "发布内容",
					"body": map[string]interface{}{
						"type": "form",
						"mode": "horizontal",
						"api":  "$preset.apis.edit",
						//"$ref": "updateControls",
						"controls": forms,
					},
				},
			},
		},
		"footerToolbar": []string{"statistics", "switch-per-page", "pagination"},
	}, 0, c.Ctx())
}

func (c *ContentController) contentControls()  {

}

func (c *ContentController) Page() {
	catid, _ := c.Ctx().GetInt64("catid")
	if catid == 0 {
		helper.Ajax("页面错误", 1, c.Ctx())
		return
	}
	pageModel := models.NewPageModel()
	page := pageModel.GetPage(catid)
	hasPage := page != nil
	if page == nil {
		page = &tables.Page{}
	}
	var res bool
	if c.Ctx().IsPost() {
		page.Title = c.Ctx().FormValue("title")
		page.Content = c.Ctx().FormValue("content")
		page.Keywords = c.Ctx().FormValue("keywords")
		page.Description = c.Ctx().FormValue("description")
		page.Updatetime = int64(helper.GetTimeStamp())
		if !hasPage {
			page.Catid = catid
			res = pageModel.AddPage(page)
		} else {
			res = pageModel.UpdatePage(page)
		}
		if res {
			helper.Ajax("发布成功", 0, c.Ctx())
		} else {
			helper.Ajax("发布失败", 1, c.Ctx())
		}
		return
	}

	c.Ctx().Render().ViewData("catid", catid)
	c.Ctx().Render().ViewData("info", page)
	c.Ctx().Render().HTML("backend/content_page.html")

}

type customForm map[string]string

func (c customForm) MustCheck() bool {
	var ok bool
	if _, ok = c["catid"]; !ok {
		return false
	}
	if _, ok = c["mid"]; !ok {
		return false
	}
	if _, ok = c["table_name"]; !ok {
		return false
	}
	return true
}

//添加内容
func (c *ContentController) AddContent() {
	if c.Ctx().IsPost() {
		mid, _ := strconv.Atoi(c.Ctx().FormValue("mid"))
		if mid < 1 {
			helper.Ajax("模型参数错误， 无法确定所属模型", 1, c.Ctx())
			return
		}
		var data = customForm{}
		postData := c.Ctx().PostData()
		for formName, values := range postData {
			if formName == "attrs" {
				data[formName] = strings.Join(values, ",")
			} else {
				data[formName] = values[0]
			}
		}
		data["catid"] = c.Ctx().GetString("catid")
		if !data.MustCheck() {
			helper.Ajax("缺少必要参数", 1, c.Ctx())
			return
		}

		if _, ok := data["status"]; ok {
			data["status"] = "1"
		} else {
			data["status"] = "0"
		}

		if data["description"] == "" {
			cont := bluemonday.NewPolicy().Sanitize(data["content"])
			if len(cont) > 250 {
				data["description"] = cont[:250]
			} else {
				data["description"] = cont
			}
		}

		data["created_time"] = time.Now().In(helper.GetLocation()).Format(helper.TimeFormat)
		model := models.NewDocumentModel().GetByID(int64(mid))
		var fields []string
		var values []interface{}
		for k, v := range data {
			if k == "table_name" {
				continue
			}
			fields = append(fields, "`"+k+"`")
			values = append(values, v)
		}

		params := append([]interface{}{fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%s)", controllers.GetTableName(model.Table), strings.Join(fields, ","), strings.TrimRight(strings.Repeat("?,", len(values)), ","))}, values...)
		// 先直接入库对应表内
		insertID, err := c.Ctx().Value("orm").(*xorm.Engine).Exec(params...)
		if err != nil {
			glog.Error(err)
			helper.Ajax("添加失败:"+err.Error(), 1, c.Ctx())
			return
		}
		id, _ := insertID.LastInsertId()
		if id > 0 {
			helper.Ajax(id, 0, c.Ctx())
		} else {
			helper.Ajax("添加失败", 1, c.Ctx())
		}
		return
	}
	//根据catid读取出相应的添加模板
	catid, _ := c.Ctx().GetInt64("catid")
	if catid == 0 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	cat := models.NewCategoryModel().GetCategory(catid)
	if cat == nil {
		helper.Ajax("分类不存在", 1, c.Ctx())
		return
	}
	if cat.Catid == 0 {
		helper.Ajax("不存在的分类", 1, c.Ctx())
		return
	}
	c.Ctx().Render().ViewData("category", cat)
	c.Ctx().Render().ViewData("form", template.HTML(buildModelForm(cat.ModelId, nil)))
	c.Ctx().Render().ViewData("submitURL", template.HTML("/b/content/add"))
	c.Ctx().Render().ViewData("preview", 0)
	c.Ctx().Render().HTML("backend/model_publish.html")
}

//修改内容
func (c *ContentController) EditContent(orm *xorm.Engine) {
	//根据catid读取出相应的添加模板
	catid, _ := c.Ctx().GetInt64("catid")
	id, _ := c.Ctx().GetInt64("id")
	if catid < 1 || id < 1 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	catogoryModel := models.NewCategoryModel().GetCategory(catid)
	if catogoryModel == nil {
		helper.Ajax("分类不存在", 1, c.Ctx())
		return
	}
	if catogoryModel.ModelId < 1 {
		helper.Ajax("找不到关联模型", 1, c.Ctx())
		return
	}
	relationDocumentModel := models.NewDocumentModel().GetByID(catogoryModel.ModelId)
	if relationDocumentModel.Id == 0 {
		helper.Ajax("找不到关联模型", 1, c.Ctx())
		return
	}
	sql := []interface{}{fmt.Sprintf("SELECT * FROM `%s` WHERE catid=? and deleted_time IS NULL AND id = ? LIMIT 1", controllers.GetTableName(relationDocumentModel.Table)), catid, id}
	contents, err := orm.QueryString(sql...)
	if err != nil {
		c.Logger().Error(err)
		helper.Ajax("获取文章内容错误", 1, c.Ctx())
		return
	}

	if len(contents) == 0 {
		helper.Ajax("文章不存在或已删除", 1, c.Ctx())
		return
	}
	if c.Ctx().IsPost() {
		var data = customForm{}
		postData := c.Ctx().PostData()
		for formName, values := range postData {
			if formName == "flag" {
				data[formName] = strings.Join(values, ",")
			} else {
				data[formName] = values[0]
			}
		}
		data["catid"] = c.Ctx().GetString("catid")
		if !data.MustCheck() {
			helper.Ajax("缺少必要参数", 1, c.Ctx())
			return
		}
		delete(data, "id")

		if _, ok := data["status"]; ok {
			data["status"] = "1"
		} else {
			data["status"] = "0"
		}

		data["updated_time"] = time.Now().In(helper.GetLocation()).Format(helper.TimeFormat)
		var sets []string
		var values []interface{}

		if data["description"] == "" {
			cont := bluemonday.NewPolicy().Sanitize(data["content"])
			if len(cont) > 250 {
				data["description"] = cont[:250]
			} else {
				data["description"] = cont
			}
		}

		for k, v := range data {
			if k == "table_name" {
				continue
			}
			sets = append(sets, "`"+k+"`= ?")
			values = append(values, v)
		}

		values = append(values, id, catid)
		params := append([]interface{}{fmt.Sprintf("UPDATE `%s` SET %s WHERE id=? and catid=?", controllers.GetTableName(relationDocumentModel.Table), strings.Join(sets, ", "))}, values...)
		insertID, err := c.Ctx().Value("orm").(*xorm.Engine).Exec(params...)
		if err != nil {
			glog.Error(err)
			helper.Ajax("修改失败:"+err.Error(), 1, c.Ctx())
			return
		}
		res, _ := insertID.RowsAffected()
		if res > 0 {
			helper.Ajax("修改成功", 0, c.Ctx())
		} else {
			helper.Ajax("修改失败", 1, c.Ctx())
		}
		return
	}
	c.Ctx().Render().ViewData("form", template.HTML(buildModelForm(catogoryModel.ModelId, contents[0])))
	c.Ctx().Render().ViewData("category", catogoryModel)
	c.Ctx().Render().ViewData("submitURL", template.HTML("/b/content/edit"))
	c.Ctx().Render().ViewData("preview", 0)
	c.Ctx().Render().HTML("backend/model_publish.html")
}

//删除内容
func (c *ContentController) DeleteContent(orm *xorm.Engine) {
	catid, _ := c.Ctx().GetInt64("catid")
	id, _ := c.Ctx().GetInt64("id")
	if catid < 1 || id < 1 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	catogoryModel := models.NewCategoryModel().GetCategory(catid)
	if catogoryModel == nil {
		helper.Ajax("分类不存在", 1, c.Ctx())
		return
	}
	if catogoryModel.ModelId < 1 {
		helper.Ajax("找不到关联模型", 1, c.Ctx())
		return
	}
	relationDocumentModel := models.NewDocumentModel().GetByID(catogoryModel.ModelId)
	if relationDocumentModel.Id == 0 {
		helper.Ajax("找不到关联模型", 1, c.Ctx())
		return
	}
	sqlOrArgs := []interface{}{fmt.Sprintf("UPDATE `%s` SET `deleted_time`='"+time.Now().In(helper.GetLocation()).Format(helper.TimeFormat)+"' WHERE id = ? and catid=?", controllers.GetTableName(relationDocumentModel.Table)), id, catid}
	res, err := orm.Exec(sqlOrArgs...)
	if err != nil {
		c.Logger().Error(err.Error())
		helper.Ajax("删除失败", 1, c.Ctx())
		return
	}
	if ret, _ := res.RowsAffected(); ret > 0 {
		helper.Ajax("删除成功", 0, c.Ctx())
	} else {
		helper.Ajax("删除失败", 1, c.Ctx())
	}
}

//排序内容
func (c *ContentController) OrderContent() {
	data := c.Ctx().PostData()
	var order = map[string]string{}
	for k, v := range data {
		order[strings.ReplaceAll(strings.ReplaceAll(k, "order[", ""), "]", "")] = v[0]
	}
	id, _ := c.Ctx().GetInt64("catid")
	if id < 1 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	catogoryModel := models.NewCategoryModel().GetCategory(id)
	if catogoryModel == nil {
		helper.Ajax("分类不存在", 1, c.Ctx())
		return
	}
	if catogoryModel.ModelId < 1 {
		helper.Ajax("找不到关联模型", 1, c.Ctx())
		return
	}
	relationDocumentModel := models.NewDocumentModel().GetByID(catogoryModel.ModelId)
	if relationDocumentModel.Id == 0 {
		helper.Ajax("找不到关联模型", 1, c.Ctx())
		return
	}
	for artID, orderNum := range order {
		sqlOrArgs := []interface{}{fmt.Sprintf("UPDATE `%s` SET `listorder`=? , updated_time = '"+time.Now().In(helper.GetLocation()).Format(helper.TimeFormat)+"' WHERE id = ? and catid=?", controllers.GetTableName(relationDocumentModel.Table)), orderNum, artID, id}
		if _, err := c.Ctx().Value("orm").(*xorm.Engine).Exec(sqlOrArgs...); err != nil {
			c.Logger().Error(err)
			helper.Ajax("更新文档排序失败", 1, c.Ctx())
			return
		}
	}
	helper.Ajax("更新排序成功", 0, c.Ctx())
}
