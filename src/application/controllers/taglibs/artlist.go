package taglibs

import (
	"fmt"
	"github.com/CloudyKit/jet"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/common/helper"
	"reflect"
	"strconv"
	"strings"
)

func ArcList(args jet.Arguments) reflect.Value {
	defer func() {
		if err := recover(); err != nil {
			pine.Logger().Error("ArcList Failed", err)
		}
	}()
	catid := args.Get(0)
	var ids []string
	switch catid.Type().String() {
	case "string":
		ids = strings.Split(catid.String(), ",")
	default:
		ids = append(ids, strconv.Itoa(getInt(catid)))
	}
	orderBy := args.Get(3).String()
	if orderBy == "" {
		orderBy = "listorder desc"
	}
	offset := getInt(args.Get(1))
	limit := getInt(args.Get(2))

	modelTable := controllers.GetTableName("articles")

	categoryTable := controllers.GetTableName("category")

	sess := getOrmSess().
		Limit(limit, offset).Select(fmt.Sprintf("%s.*, %s.catname, %s.url as caturl, %s.type ", modelTable, categoryTable, categoryTable, categoryTable)).
		Join("LEFT", categoryTable, categoryTable + ".catid = " + modelTable + ".catid").
		Where(modelTable +".deleted_time IS NULL").Where(modelTable +".status = 1").OrderBy(orderBy)
	if ids[0] != "0" && ids[0] != "-1" {
		sess.In(modelTable + ".catid", ids)
	}
	list, err := sess.QueryString()


	if err != nil {
		panic(err)
	}
	// 重写URL
	for i, art := range list {
		catid,_ := strconv.Atoi(art["catid"])
		if art["type"] == "0" {
			list[i]["caturl"] = helper.ListUrl(catid)
		} else if art["type"] == "1" {
			list[i]["caturl"]  = helper.PageUrl(catid)
		}
	}
	return reflect.ValueOf(list)
}
