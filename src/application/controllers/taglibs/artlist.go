package taglibs

import (
	"fmt"
	"github.com/xiusin/pinecms/src/config"
	"reflect"
	"strconv"
	"strings"

	"github.com/CloudyKit/jet"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

func ArcList(args jet.Arguments) reflect.Value {
	if !checkArgType(&args) {
		return defaultArrReturnVal
	}
	defer func() {
		if err := recover(); err != nil {
			pine.Logger().Errorf("ArcList Failed %s", err)
		}
	}()
	catid := args.Get(0)
	var ids []string
	if isNumber(catid) {
		ids = append(ids, strconv.Itoa(int(getNumber(catid))))
	} else {
		ids = strings.Split(catid.String(), ",")
	}

	orderBy := args.Get(3).String()
	if orderBy == "" {
		orderBy = "listorder asc"
	}
	offset := getNumber(args.Get(1))
	limit := getNumber(args.Get(2))
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
		modelID = "1"// 默认模型
	}
	exists, err := pine.Make(controllers.ServiceXorm).(*xorm.Engine).Table(model).ID(modelID).Get(model)
	if !exists {
		if err != nil {
			panic("模型ID" + modelID + "不存在: " + err.Error())
		} else {
			panic("模型ID" + modelID + "不存在")
		}
	}

	if limit  == 0 {	// 没有设置分页条数, 使用系统默认分页数目
		conf,_ := config.SiteConfig()
		t, _ := strconv.Atoi(conf["SITE_PAGE_SIZE"])
		limit = int64(t)
		if offset == 0 { // 如果没传递offset那么就根据系统分页计算偏移量
			offset = (getNumber(args.Get(5)) - 1) * limit
		}
	}
	categoryTable := controllers.GetTableName("category")
	modelTable := controllers.GetTableName(model.Table)
	sess := getOrmSess(model.Table).Select(fmt.Sprintf("%s.*, %s.catname, %s.url as caturl, %s.type ", modelTable, categoryTable, categoryTable, categoryTable)).
		Join("LEFT", categoryTable, categoryTable + ".catid = " + modelTable + ".catid").
		Where(modelTable +".deleted_time IS NULL").Where(modelTable +".status = 1").OrderBy(orderBy)

	if limit > 0 {
		sess.Limit(int(limit), int(offset))
	}

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

	return reflect.ValueOf(list)
}
