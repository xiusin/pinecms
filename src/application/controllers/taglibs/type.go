package taglibs

import (
	"github.com/CloudyKit/jet"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"reflect"
)

func Type(args jet.Arguments) reflect.Value {
	catid := int(args.Get(0).Float())
	if catid < 0 {
		panic("typeid参数不能小于1")
	}
	orm := pine.Make("*xorm.Engine").(*xorm.Engine)
	var data = &tables.Category{}
	sess := orm.Table(data)
	defer sess.Close()

	exists, _ := sess.ID(catid).Get(data)
	if !exists {
		data = nil
	}
	return reflect.ValueOf(data)
}