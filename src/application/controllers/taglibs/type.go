package taglibs

import (
	"fmt"
	"github.com/CloudyKit/jet"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"reflect"
	"runtime/debug"
)

func Type(args jet.Arguments) reflect.Value {
	fmt.Println("arclist")

	if !checkArgType(&args) {
		return defaultArrReturnVal
	}
	defer func() {
		if err := recover(); err != nil {
			pine.Logger().Error("Type Failed", string(debug.Stack()))
		}
	}()
	catid := getNumber(args.Get(0))
	if catid < 0 {
		panic("typeid参数不能小于1")
	}
	orm := helper.GetORM()
	var data = &tables.Category{}
	sess := orm.Table(data)
	defer sess.Close()
	exists, _ := sess.ID(catid).Get(data)
	if exists && data.Type != 2 {
		data.Url = fmt.Sprintf("/%s/", models.NewCategoryModel().GetUrlPrefix(data.Catid))
	}
	return reflect.ValueOf(data)
}
