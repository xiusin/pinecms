package controllers

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/pine"
	"reflect"
	"strconv"
	"strings"

	"github.com/CloudyKit/jet"
)

//{dede:arclist typeid='26' row='15' titlelen='50' orderby='pubdate'}
func ArcList() {

}

// 公共方法
func Channel(args jet.Arguments) reflect.Value {
	// id type row
	catid := int(args.Get(0).Float())
	if catid < 1 {
		panic("channel:id参数不能小于1")
	}
	typeInfo := args.Get(1).String()
	switch typeInfo {
	case "top":

	case "son":

	case "self":

	default:
		panic("channel:type 可选值为top,son,self")
	}

	row := int(args.Get(2).Float())
	if row <= 0 {
		row = 10
	}

	orm := pine.Make("*xorm.Engine").(*xorm.Engine)
	sess := orm.Table(&tables.IriscmsCategory{})

	sess.Limit(row)
	defer sess.Close()
	return reflect.Value{}
}

//{dede:pagelist listsize='1' listitem='index,end,pre,next,pageno'/}
func PageList() {

}

// todo 列表页面用与接收参数，应该是控制器内暴露的方法
// {dede:list pagesize ='5'}
// {dede:list perpage='20'}
func List() {

}

// {dede:tag row='60' sort='new'}
func Tag() {

}

//  {dede:type typeid='26'}
func Type(args jet.Arguments) reflect.Value {
	catid := int(args.Get(0).Float())
	if catid < 0 {
		panic("typeid参数不能小于1")
	}
	orm := pine.Make("*xorm.Engine").(*xorm.Engine)
	var data = &tables.IriscmsCategory{}
	sess := orm.Table(data)
	defer sess.Close()

	exists, _ := sess.ID(catid).Get(data)
	if !exists {
		data = nil
	}
	return reflect.ValueOf(data)

}

func Prenext() {

}

func Global() {

}

// 支持嵌套channel 看怎么实现
func ChannelArtList(args jet.Arguments) reflect.Value {
	catid := args.Get(0)
	var ids []string
	switch catid.Type().String() {
	case "float64":
		ids = append(ids, strconv.Itoa(int(catid.Float())))
	case "string":
		ids = strings.Split(catid.String(), ",")
	default:
		panic("channelartlist:typeid 不支持的类型")
	}
	orm := pine.Make("*xorm.Engine").(*xorm.Engine)
	sess := orm.Table(&tables.IriscmsLink{})
	defer sess.Close()

	sess.Table(&tables.IriscmsCategory{})
	var categories []tables.IriscmsCategory
	err := sess.In("parentid", ids).Where("ismenu = ?", 1).Desc("listorder").Find(&categories)
	if err != nil {
		panic(err)
	}

	// 关联查询数据







	return reflect.ValueOf(map[string]interface{}{
		"Categories":categories,
	})
}

func Flink(args jet.Arguments) reflect.Value {
	orm := pine.Make("*xorm.Engine").(*xorm.Engine)
	sess := orm.Table(&tables.IriscmsLink{})
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
	var data []tables.IriscmsLink
	if err := sess.Find(&data); err != nil {
		panic(err)
	}
	return reflect.ValueOf(data)
}
