package taglibs

import (
	"github.com/CloudyKit/jet"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"reflect"
	"strconv"
	"strings"
	"time"
	"xorm.io/xorm"
)

/**
广告标签
{{yield adlist(id=3, pos="首页banner图")}}
返回一组广告
*/
func AdList(args jet.Arguments) reflect.Value {
	if !checkArgType(&args) {
		return defaultArrReturnVal
	}
	id := args.Get(0)
	var ids []string
	if isNumber(id) {
		ids = append(ids, strconv.Itoa(int(getNumber(id))))
	} else if id.String() != "" {
		ids = strings.Split(id.String(), ",")
	}
	order := args.Get(2).String()
	if order == "" {
		order = "listorder desc"
	}
	orm := pine.Make(controllers.ServiceXorm).(*xorm.Engine).OrderBy(order)
	// 获取广告位信息
	var pos = args.Get(1)
	switch pos.Type().String() {
	case "string":
		if len(pos.String()) > 0 {
			advPos := &tables.AdvertSpace{}
			exists, _ := pine.Make(controllers.ServiceXorm).(*xorm.Engine).Table(advPos).Where("name = ?", pos.String()).Get(advPos)
			if !exists {
				return reflect.ValueOf([]tables.Advert{})
			}
			orm.Where("space_id = ?", advPos.Id)
		}
	default:
		if pid := getNumber(pos); pid > 0 {
			orm.Where("space_id = ?", pid)
		}
	}
	if len(ids) > 0 {
		orm.In("id", ids)
	}

	now := time.Now().In(helper.GetLocation()).Format(helper.TimeFormat)
	orm.Where("status = 1").Where("start_time <= ?", now).Where("end_time >= ?", now).Select("id, name, image, link_url")
	var advs = []tables.Advert{}
	orm.Find(&advs)
	return reflect.ValueOf(advs)
}
