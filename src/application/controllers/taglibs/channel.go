package taglibs

import (
	"github.com/CloudyKit/jet"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"reflect"
	"strconv"
)

func Channel(args jet.Arguments) reflect.Value {
	row := getInt(args.Get(2))
	if row <= 0 {
		row = 10
	}
	catid := getInt(args.Get(0))
	typeInfo := args.Get(1).String()
	sess := pine.Make("*xorm.Engine").(*xorm.Engine).Table(&tables.Category{})
	sess.Limit(row)
	switch typeInfo {
	case "son":
		sess.Where("parentid = ?", catid)
	case "self":
		s := &tables.Category{}
		exist,_ := pine.Make("*xorm.Engine").(*xorm.Engine).Table(s).ID(catid).Cols("parentid").Get(s)
		if exist {
			sess.Where("parentid = ?", catid)
		} else {
			return reflect.ValueOf([]tables.Category{})
		}
	}
	var data = []tables.Category{}
	sess.Find(&data)
	for k, v := range data {
		if v.Type == 0 {
			data[k].Url = "/list?tid=" + strconv.Itoa(int(v.Catid))
		} else if v.Type == 1 {
			data[k].Url = "/page?tid=" + strconv.Itoa(int(v.Catid))
		}
	}
	return reflect.ValueOf(data)
}