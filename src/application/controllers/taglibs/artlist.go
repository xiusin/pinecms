package taglibs

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/config"
	"xorm.io/builder"

	"github.com/CloudyKit/jet"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

/**
    <iterm>row:返回文档列表总数</iterm>
    <iterm>typeid:栏目ID,在列表模板和档案模板中一般不需要指定，在首页模板中允许用","分开表示多个栏目</iterm>
    <iterm>getall:在没有指定这属性的情况下,在栏目页、文章页模板,不会获取以","分开的多个栏目的下级子类</iterm>
    <iterm>titlelen:标题长度 等同于titlelength</iterm>
    <iterm>infolen:表示内容简介长度 等同于infolength</iterm>
    <iterm>imgwidth:缩略图宽度</iterm>
    <iterm>imgheight:缩略图高度</iterm>
    <iterm>listtype: 栏目类型 image含有缩略图 commend推荐</iterm>
    <iterm>orderby:文档排序方式</iterm>
    <iterm>keyword:含有指定关键字的文档列表，多个关键字用","分</iterm>
    <iterm>innertext:单条记录样式</iterm>
    <iterm>aid:指定文档ID</iterm>
    <iterm>idlist:提取特定文档（文档ID</iterm>
    <iterm>channelid:频道ID</iterm>
    <iterm>limit:（起始ID从0开始）表示限定的记录范围（如：limit='1,2'  表示从ID为1的记录开始，取2条记录</iterm>
    <iterm>flag:自定义属性值：头条[h]推荐[c]图片[p]幻灯[f]滚动[s]跳转[j]图文[a]加粗[b]</iterm>
    <iterm>noflag:同flag，但这里是表示不包含这些属性</iterm>
    <iterm>orderway:值为 desc 或 asc ，指定排序方式是降序还是顺向排序，默认为降序</iterm>
    <iterm>subday:表示在多少天以内的文档</iterm>

artlist(typeid, offset, row, orderby, modelid, page, keyword, flag, noflag, titlelen, getall, subday, limit)
*/

func ArcList(args jet.Arguments) reflect.Value {
	if !checkArgType(&args) {
		return defaultArrReturnVal
	}
	//getall := args.Get(10).Bool()
	orderway := args.Get(11).String()
	subday := int(args.Get(12).Float())
	catid := args.Get(0)
	var ids []string
	if isNumber(catid) {
		ids = append(ids, fmt.Sprintf("%d", int(getNumber(catid))))
	} else {
		ids = strings.Split(catid.String(), ",")
	}
	offset := getNumber(args.Get(1))
	limit := getNumber(args.Get(2))
	if limit == 0 { // 没有设置分页条数, 使用系统默认分页数目
		conf, _ := config.SiteConfig()
		t, _ := strconv.Atoi(conf["SITE_PAGE_SIZE"])
		limit = int64(t)
		if limit == 0 {
			limit = 15
		}
		if offset == 0 { // 如果没传递offset那么就根据系统分页计算偏移量
			offset = (getNumber(args.Get(5)) - 1) * limit
		}
	}

	orderBy := args.Get(3).String()
	if orderway == "" {
		orderway = "desc"
	}
	if orderBy == "" {
		orderBy = "listorder"
	}
	var isRand bool
	if orderBy == "rand" {
		isRand = true
		if strings.ToLower(config.DB().Db.DbDriver) == "mysql" {
			orderBy = "RAND()"
		} else {
			orderBy = "RANDOM()"
		}
	} else {
		orderBy = fmt.Sprintf("%s %s", orderBy, orderway)
	}

	model := &tables.DocumentModel{}
	var modelID int64
	m := models.NewCategoryModel()
	if len(ids) == 0 || ids[0] == "0" { // 没有设置typeid直接查看模型ID
		modelID = getNumber(args.Get(4))
	} else if len(ids) > 1 { // 设置多个id以第一个ID查找对应模型
		catid, _ := strconv.Atoi(ids[0])
		catgory, err := m.GetCategoryFByIdForBE(int64(catid))
		if err != nil {
			panic("无法查找分类" + strings.Join(ids, ",") + "的信息:" + err.Error())
		}
		for _, v := range ids {
			catID, _ := strconv.Atoi(v)
			soncats := m.GetNextCategoryOnlyCatids(int64(catID), false)
			for _, v := range soncats {
				ids = append(ids, fmt.Sprintf("%d", v))
			}
		}
		modelID = catgory.Model.Id
	} else if len(ids) == 1 { // 设置了单个类型并且获取所有下级子类&& ids[0] != "0" && getall
		firstCatID := getNumber(catid)
		soncats := m.GetNextCategoryOnlyCatids(firstCatID, false)
		for _, v := range soncats {
			ids = append(ids, fmt.Sprintf("%d", v))
		}
		// 读取模型ID
		catgory, err := m.GetCategoryFByIdForBE(firstCatID)
		if err != nil {
			panic("无法查找分类" + strings.Join(ids, ",") + "的信息:" + err.Error())
		}
		modelID = catgory.Model.Id
	}
	if modelID == 0 {
		modelID = 1
	}
	exists, _ := helper.GetORM().Table(model).ID(modelID).Get(model)
	if !exists {
		panic(fmt.Sprintf("模型ID%d不存在", modelID))
	}

	categoryTable := controllers.GetTableName("category")

	modelTable := controllers.GetTableName(model.Table)
	sess := getOrmSess(model.Table).
		Join("LEFT", categoryTable, categoryTable+".catid = "+modelTable+".catid").
		Where(modelTable + ".deleted_time IS NULL").Where(modelTable + ".status = 1")
	defer sess.Close()
	if isRand {
		sess.OrderBy(orderBy)
	} else {
		sess.OrderBy(modelTable + "." + orderBy)
	}
	if ids[0] != "0" {
		sess.In(modelTable+".catid", ids)
	}

	if subday > 0 {
		sess.Where(modelTable+".pubtime > ?", time.Now().AddDate(0, 0, -subday).In(helper.GetLocation()).Format("2006-01-02"))
	}

	if keywords := strings.Split(args.Get(6).String(), ","); len(keywords) > 0 {
		var conds []builder.Cond
		for _, v := range keywords {
			if v == "" {
				continue
			}
			conds = append(conds, builder.Expr(fmt.Sprintf("%s.keywords LIKE ?", modelTable), "%"+v+"%"))
		}
		sess.Where(builder.Or(conds...))
	}

	if flags := strings.Split(args.Get(7).String(), ","); len(flags) > 0 {
		var conds []builder.Cond
		for _, v := range flags {
			if v == "" {
				continue
			}
			conds = append(conds, builder.Expr(fmt.Sprintf("%s.flag LIKE ?", modelTable), "%"+v+"%"))
		}
		sess.Where(builder.Or(conds...))
	}

	if noflags := strings.Split(args.Get(8).String(), ","); len(noflags) > 0 {
		var conds []builder.Cond
		for _, v := range noflags {
			if v == "" {
				continue
			}
			conds = append(conds, builder.Expr(fmt.Sprintf("%s.flag NOT LIKE ?", modelTable), "%"+v+"%"))
		}
		sess.Where(builder.Or(conds...))
	}

	titlelen := getNumber(args.Get(9))

	if limit > 0 {
		sess.Limit(int(limit), int(offset))
	}

	sess.Select(fmt.Sprintf("%s.*, %s.catname as typename", modelTable, categoryTable))

	list, err := sess.QueryString()
	if err != nil {
		pine.Logger().Error(sess.LastSQL())
		panic(err)
	}
	// 重写URL
	helper.HandleArtListInfo(list, int(titlelen))
	return reflect.ValueOf(list)
}
