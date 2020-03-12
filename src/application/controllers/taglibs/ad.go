package taglibs

import (
	"github.com/CloudyKit/jet"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"reflect"
	"strconv"
	"strings"
	"time"
)

/**
广告标签
{{yield ad(id=3, pos="首页banner图")}}
返回一组广告包括一个
*/
func Ad(args jet.Arguments) reflect.Value {
	id := args.Get(0)
	var ids []string
	switch id.Type().String() {
	case "string":
		if id.String() != "" {
			ids = strings.Split(id.String(), ",")
		}
	default:
		ids = append(ids, strconv.Itoa(getInt(id)))
	}
	order := args.Get(2).String()
	if order == "" {
		order = "listorder desc"
	}
	orm := pine.Make("*xorm.Engine").(*xorm.Engine).OrderBy(order)
	// 获取广告位信息
	var pos = args.Get(1)
	switch pos.Type().String() {
	case "string":
		if len(pos.String()) > 0 {
			advPos  := &tables.AdvertSpace{}
			exists ,_ := pine.Make("*xorm.Engine").(*xorm.Engine).Table(advPos).Where("name = ?", pos.String()).Get(advPos)
			if !exists {
				return reflect.ValueOf([]tables.Advert{})
			}
			orm.Where("space_id = ?", advPos.Id)
		}
	default:
		if pid := getInt(pos); pid > 0 {
			orm.Where("space_id = ?", pid)
		}
	}
	if len(ids) > 0 {
		orm.In("id", ids)
	}
	now := time.Now().In(helper.GetLocation()).Format(helper.TimeFormat)
	orm.Where("status = 1").Where("start_time <= ?", now).Where("end_time >= ?", now)
	var advs = []tables.Advert{}
	if err := orm.Find(&advs); err != nil || len(advs) == 0 {
		return reflect.ValueOf([]tables.Advert{})
	}

	return reflect.ValueOf(advs)
}
