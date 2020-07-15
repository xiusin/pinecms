package taglibs

import (
	"github.com/CloudyKit/jet"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"reflect"
	"runtime/debug"
	"strings"
)

func Flink(args jet.Arguments) reflect.Value {
	if !checkArgType(&args) {
		return defaultArrReturnVal
	}
	defer func() {
		if err := recover(); err != nil {
			pine.Logger().Error("Flink Failed", string(debug.Stack()))
		}
	}()
	orm := pine.Make("*xorm.Engine").(*xorm.Engine)
	sess := orm.Table(&tables.Link{})
	defer sess.Close()
	row := int(args.Get(0).Float())
	if row == 0 {
		row = 10
		sess.Limit(row)
	}

	idparam := args.Get(1).String()
	if len(idparam) != 0 {
		ids := strings.Split(idparam, ",")
		sess.In("linkid", ids)
	}

	sort := args.Get(2).String()
	if len(sort) != 0 {
		sess.OrderBy(sort)
	} else {
		sess.Desc("linkid")
	}
	data := []tables.Link{}
	if err := sess.Find(&data); err != nil {
		panic(err)
	}
	return reflect.ValueOf(data)
}
