package frontend

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/render/engine/jet"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (c *IndexController) Search(orm *xorm.Engine) {
	keywords := c.Ctx().GetString("keywords", c.Ctx().GetString("q"))
	keywords = c.Ctx().GetString("keyword", keywords)
	page, _ := c.Ctx().GetInt("page", 1)
	channeltype, _ := c.Ctx().GetInt("channeltype", 0)
	typeid, _ := c.Ctx().GetInt("typeid", 0)
	kwType, _ := c.Ctx().GetInt("kwtype", 0)
	orderby := c.Ctx().GetString("orderby")

	if regexp.MustCompile("^[_A-Z0-9a-z]+$").MatchString(orderby) == false {
		orderby = ""
	}
	if page < 1 {
		page = 1
	}

	searchType := c.Ctx().GetString("searchtype")

	if len(keywords) < 2 {
		pine.Logger().Error("关键字太短")
		return
	}

	startTime := c.Ctx().GetString("starttime")
	// 如果是数字
	if regexp.MustCompile("^/d+$").MatchString(startTime) {
		// 换算为天数
		s, _ := strconv.Atoi(startTime)
		startTime = time.Now().In(helper.GetLocation()).AddDate(0, 0, -s).Format(helper.TimeFormat)
	}
	var tableName string
	var channelType int
	m := models.NewCategoryModel()
	if typeid != 0 && channelType == 0 {
		// 查询对应的模型
		cat, err := m.GetCategoryFByIdForBE(int64(typeid))
		if err != nil {
			pine.Logger().Error(err)
			return
		}
		channelType = int(cat.Model.Id)
		tableName = cat.Model.Table
	} else {
		md := models.NewDocumentModel().GetByIDForBE(int64(channeltype))
		if md != nil {
			tableName = md.Table
		} else {
			channelType = 0
		}
	}

	var whereSqls []string
	var whereArgs []interface{}
	if len(startTime) > 0 {
		whereSqls = append(whereSqls, "$$.create_time > ?")
		whereArgs = append(whereArgs, startTime)
	}

	if typeid > 0 {
		soncats := m.GetNextCategoryOnlyCatids(int64(typeid), true)
		var ids []string
		for _, v := range soncats {
			ids = append(ids, fmt.Sprintf("%d", v))
		}
		whereSqls = append(whereSqls, fmt.Sprintf("$$.catid IN (%s)", strings.Join(ids, ",")))
	}

	kwsql := ""
	kws := strings.Split(keywords, " ")
	if len(kws) == 1 {
		kws = strings.Split(keywords, "+")
	}
	var ks []string
	for _, v := range kws {
		if len(v) == 0 {
			continue
		}
		if searchType == "title" {
			ks = append(ks, "$$.title LIKE ?")
			whereArgs = append(whereArgs, "%"+v+"%")
		} else {
			ks = append(ks, "($$.title LIKE ? OR $$.keywords LIKE ?)")
			whereArgs = append(whereArgs, "%"+v+"%")
			whereArgs = append(whereArgs, "%"+v+"%")
		}
	}
	if len(ks) > 0 {
		if kwType == 0 { // 多字段搜索
			kwsql = strings.Join(ks, " OR ")
		} else {
			kwsql = strings.Join(ks, " AND ")
		}
		whereSqls = append(whereSqls, kwsql)
	}
	whereSqls = append(whereSqls, " $$.status = 1 AND $$.deleted_time IS NULL")
	countField := "COUNT(*) AS total"
	selectField := "$$.id,$$.catid, $$.title, $$.description, $$.thumb, $$.pubtime, $1.catname, $1.url as caturl, $1.type"
	var queries []string
	var query string
	var modelInfos []tables.DocumentModel
	if channelType == 0 { // 如果没设置模型, 使用Union联合
		// 读取所有模型表
		orm.Where("enabled = 1").Find(&modelInfos)

		l := len(whereArgs)
		for _, model := range modelInfos {
			if ok, _ := orm.IsTableExist(controllers.GetTableName(model.Table)); ok {
				_q := "SELECT $? FROM $$ LEFT JOIN $1 ON $$.catid = $1.catid WHERE %s"
				_q = fmt.Sprintf(_q, strings.Join(whereSqls, " AND "))
				_q = strings.ReplaceAll(_q, "$$", controllers.GetTableName(model.Table))
				queries = append(queries, _q)
				whereArgs = append(whereArgs, whereArgs[0:l]...)
			}
		}
		query = strings.Join(queries, " UNION ")
	} else {
		query = "SELECT $? FROM $$ LEFT JOIN $1 ON $$.catid = $1.catid WHERE %s"
		query = fmt.Sprintf(query, strings.Join(whereSqls, " AND "))
		query = strings.ReplaceAll(query, "$$", controllers.GetTableName(tableName))
	}
	query = strings.ReplaceAll(query, "$1", controllers.GetTableName("category"))
	// 先统计数据量
	totals, err := orm.QueryString(append([]interface{}{strings.ReplaceAll(query, "$?", countField)}, whereArgs...)...)
	if err != nil {
		pine.Logger().Error(err)
		c.Ctx().Abort(http.StatusInternalServerError, err.Error())
	}
	var total = 0
	if len(totals) > 0 {
		for _, v := range totals {
			_total, _ := strconv.Atoi(v["total"])
			total += _total
		}
	}
	c.ViewData("keywords", keywords)
	var fn = func(pagesize int64) []map[string]string {
		var list []map[string]string
		if total > 0 {
			var err error

			if orderby == "" {
				orderby = "id"
			}
			offset := (page - 1) * int(pagesize)

			if len(queries) > 0 {
				for k, v := range queries {
					v = strings.ReplaceAll(v, "$?", selectField)
					v = strings.ReplaceAll(v, "$$", controllers.GetTableName(modelInfos[k].Table))
					queries[k] = v
				}
				query = strings.Join(queries, " UNION ")
				query = fmt.Sprintf("SELECT * FROM ("+query+") AS t ORDER BY t.%s LIMIT  %d, %d", orderby, offset, pagesize)
			} else {
				query = strings.ReplaceAll(query, "$?", selectField)
				query += fmt.Sprintf(" ORDER BY $$.%s DESC LIMIT %d, %d", orderby, offset, pagesize)
				query = strings.ReplaceAll(query, "$$", controllers.GetTableName(tableName))
			}
			query = strings.ReplaceAll(query, "$1", controllers.GetTableName("category"))
			list, err = orm.QueryString(append([]interface{}{fmt.Sprintf("%s ", query)}, whereArgs...)...)
			if err != nil {
				pine.Logger().Error(err)
			}
			for k, art := range list {
				catid, _ := strconv.Atoi(art["catid"])
				prefix := m.GetUrlPrefix(int64(catid))
				art["caturl"] = fmt.Sprintf("/%s/", prefix)
				id, _ := strconv.Atoi(art["id"])
				list[k]["arcurl"] = fmt.Sprintf("/%s/%d.html", prefix, id)
				list[k]["arturl"] = art["arcurl"]
			}
		}
		return list
	}
	pineJet := pine.Make(controllers.ServiceJetEngine).(*jet.PineJet)
	tpl, err := pineJet.GetTemplate(template("search.jet"))
	if err != nil {
		pine.Logger().Error("读取搜索页面失败", err)
		c.Ctx().Abort(500, err.Error())
		return
	}
	c.Ctx().Render().ContentType(pine.ContentTypeHTML)
	if err := tpl.Execute(c.Ctx().Response.BodyWriter(), viewDataToJetMap(c.Ctx().Render().GetViewData()), struct {
		Field       *tables.Category
		Position    string
		ArtCount    int64
		ModelName   string
		TopCategory *tables.Category
		QP          map[string][]string
		PageNum     int64
		ListFunc    func(int64) []map[string]string
	}{
		ArtCount:  int64(total),
		PageNum:   int64(page),
		ModelName: tableName,
		QP:        c.Ctx().GetData(),
		ListFunc:  fn}); err != nil {
		panic(err)
	}
}
