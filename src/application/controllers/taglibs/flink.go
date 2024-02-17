package taglibs

import (
	"reflect"
	"strings"

	"github.com/CloudyKit/jet"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

func Flink(args jet.Arguments) reflect.Value {
	if !checkArgType(&args) {
		return defaultArrReturnVal
	}
	defer func() {
		if err := recover(); err != nil {
			pine.Logger().Warning(err)
		}
	}()
	orm := helper.GetORM()
	sess := orm.Table(&tables.Link{})
	defer sess.Close()
	row := int(args.Get(0).Float())
	if row == 0 {
		row = 10
		sess.Limit(row)
	}

	idParam := args.Get(1).String()
	if len(idParam) != 0 {
		ids := strings.Split(idParam, ",")
		sess.In("id", ids)
	}

	sort := args.Get(2).String()
	if len(sort) != 0 {
		sess.OrderBy(sort)
	} else {
		sess.Desc("id")
	}
	data := []tables.Link{}
	helper.PanicErr(sess.Find(&data))
	return reflect.ValueOf(data)
}
