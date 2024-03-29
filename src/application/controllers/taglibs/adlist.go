package taglibs

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/CloudyKit/jet"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
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
	orm := helper.GetORM().OrderBy(order)
	// 获取广告位信息
	var pos = args.Get(1)
	switch pos.Type().String() {
	case "string":
		if len(pos.String()) > 0 {
			advPos := &tables.AdvertSpace{}
			exists, _ := helper.GetORM().Table(advPos).Where("name = ?", pos.String()).Get(advPos)
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

	now := helper.NowDate(helper.TimeFormat)
	orm.Where("status = 1").Where("start_time <= ?", now).Where("end_time >= ?", now).Select("id, name, image, link_url")
	var advs = []tables.Advert{}
	orm.Find(&advs)
	return reflect.ValueOf(advs)
}
