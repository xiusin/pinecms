package taglibs

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/CloudyKit/jet"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/config"
)

/**
list(typeid, page, pagesize, modelname,  titlelen, orderby, orderway)
*/

func List(args jet.Arguments) reflect.Value {
	if !checkArgType(&args) {
		return defaultArrReturnVal
	}
	var list = []map[string]string{}
	helper.Cache().Remember("pinecms:tag:list:"+getTagHash(args), &list, func() (any, error) {
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
			return &list, nil
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
		var isRand bool
		if orderBy == "rand" {
			isRand = true
			if strings.ToLower(config.DB().Db.DbDriver) == "mysql" {
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
		sess := getOrmSess(tableName)
		defer sess.Close()

		b := builder.Select("a.*", "c.catname as typename").From(modelTable, "a").
			LeftJoin(categoryTable, "c", "c.id = a.catid").
			Where(builder.IsNull{"a.deleted_time"}).
			Where(builder.Eq{"a.status": 1}).
			In("a.catid", ids).
			Limit(int(pagesize), int(offset))
		if isRand {
			b.OrderBy(orderBy)
		} else {
			b.OrderBy("a." + orderBy)
		}

		var err error
		list, err = sess.QueryString(b)
		if err != nil {
			pine.Logger().Error(sess.LastSQL())
			return &list, err
		}
		helper.HandleArtListInfo(list, titlelen)
		return &list, nil
	})
	return reflect.ValueOf(list)
}
