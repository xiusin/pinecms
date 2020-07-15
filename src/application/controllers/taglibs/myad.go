package taglibs

import (
	"github.com/CloudyKit/jet"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"reflect"
	"time"
)

/**
广告标签
{{yield myad(id="", name="")}}
返回一组广告
*/
func MyAd(args jet.Arguments) reflect.Value {
	if !checkArgType(&args) {
		return defaultSignalVal
	}
	id := int(getNumber(args.Get(0)))
	name := args.Get(1).String()
	now := time.Now().In(helper.GetLocation()).Format(helper.TimeFormat)
	orm := pine.Make(controllers.ServiceXorm).(*xorm.Engine).Where("status = 1").Where("start_time <= ?", now).Where("end_time >= ?", now).Select("id, name, image, link_url")
	if id > 0 {
		orm.ID(id)
	}
	if id == 0 && name != "" {
		orm.Where("name = ?", name)
	}
	var advs = tables.Advert{}
	orm.Get(&advs)
	return reflect.ValueOf(advs)

}
