package taglibs

import (
	"fmt"
	"github.com/CloudyKit/jet"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/common"
	"github.com/xiusin/pinecms/src/config"
	"reflect"
	"strconv"
	"strings"
)

/**
list(typeid, page, pagesize, modelname,  titlelen, orderby, orderway)
*/

func List(args jet.Arguments) reflect.Value {
	if !checkArgType(&args) {
		return defaultArrReturnVal
	}
	defer func() {
		if err := recover(); err != nil {
			pine.Logger().Errorf("ArcList Failed %s", err)
		}
	}()
	catid := getNumber(args.Get(0))
	page := getNumber(args.Get(1))
	pagesize := getNumber(args.Get(2))
	if pagesize == 0 {
		conf, _ := config.SiteConfig()
		ps, _ := strconv.Atoi(conf["SITE_PAGE_SIZE"])
		if ps == 0 {
			pagesize = 15
		} else {
			pagesize = int64(ps)
		}
	}
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pagesize
	tableName := args.Get(3).String()
	if tableName == "" || catid < 1 {
		return defaultArrReturnVal
	}
	titlelen := int(getNumber(args.Get(4)))
	orderBy := args.Get(5).String()
	orderWay := args.Get(6).String()
	if orderWay == "" {
		orderWay = "desc"
	}
	if orderBy == "" {
		orderBy = "listorder"
	}
	if orderBy == "rand" {
		if strings.ToLower(config.DBConfig().Db.DbDriver) == "mysql" {
			orderBy = "RAND()"
		} else {
			orderBy = "RANDOM()"
		}
	} else {
		orderBy = fmt.Sprintf("%s %s", orderBy, orderWay)
	}
	categoryTable := getCategoryTable()
	modelTable := controllers.GetTableName(tableName)
	m := models.NewCategoryModel()
	ids := m.GetNextCategoryOnlyCatids(catid, true)
	sess := getOrmSess(tableName).Where(modelTable+".deleted_time IS NULL").Where(modelTable+".status = 1").In(modelTable+".catid", ids).Limit(int(pagesize), int(offset)).OrderBy(orderBy)
	defer sess.Close()
	sess.Join("LEFT", categoryTable, categoryTable+".catid = "+modelTable+".catid")
	sess.Select(fmt.Sprintf("%s.*, %s.catname as typename", modelTable, categoryTable))
	list, _ := sess.QueryString()
	common.HandleArtListInfo(list, titlelen)
	return reflect.ValueOf(list)
}
