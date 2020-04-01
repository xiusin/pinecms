package taglibs

import (
	"fmt"
	"github.com/CloudyKit/jet"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
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
		orderBy = "listorder asc"
	}
	offset := getInt(args.Get(1))
	limit := getInt(args.Get(2))
	model := &tables.DocumentModel{}
	var modelID string
	// 先不支持多模型调用
	if len(ids) == 0 {
		modelID = args.Get(4).String()
	} else if ids[0] != "0" && ids[0] != "" {
		catid,_ := strconv.Atoi(ids[0])
		catgory,err := models.NewCategoryModel().GetCategoryFullWithCache(int64(catid))
		if err != nil {
			panic("无法查找分类"+strings.Join(ids, ",")+"的信息:" + err.Error())
		}
		modelID = strconv.Itoa(int(catgory.Model.Id))
	}
	if modelID == ""  || modelID == "0"{
		// 默认一个模型ID
		modelID = "1"
	}
	exists, _ := pine.Make(controllers.ServiceXorm).(*xorm.Engine).Table(model).ID(modelID).Get(model)
	if !exists {
		panic("模型ID" + modelID + "不存在")
	}
	categoryTable := controllers.GetTableName("category")
	modelTable := controllers.GetTableName(model.Table)
	sess := getOrmSess(model.Table).
		Limit(limit, offset).Select(fmt.Sprintf("%s.*, %s.catname, %s.url as caturl, %s.type ", modelTable, categoryTable, categoryTable, categoryTable)).
		Join("LEFT", categoryTable, categoryTable + ".catid = " + modelTable + ".catid").
		Where(modelTable +".deleted_time IS NULL").Where(modelTable +".status = 1").OrderBy(orderBy)

	if ids[0] != "0" && ids[0] != "-1" {
		sess.In(modelTable + ".catid", ids)
	}
	list, err := sess.QueryString()

	if err != nil {
		pine.Logger().Error(sess.LastSQL())
		panic(err)
	}
	// 重写URL
	m := models.NewCategoryModel()
	for i, art := range list {
		catid,_ := strconv.Atoi(art["catid"])
		if art["type"] != "2" {
			list[i]["caturl"] = fmt.Sprintf("/%s/", m.GetUrlPrefix(int64(catid)))
		}
	}
	if list == nil {
		list = []map[string]string{}
	}
	return reflect.ValueOf(list)
}
