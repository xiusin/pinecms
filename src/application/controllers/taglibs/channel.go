package taglibs

import (
	"fmt"
	"github.com/CloudyKit/jet"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"reflect"
)

/**
typeid = "son | top | self"
top: 顶级栏目 : parentid 为 0
son: 父id 为 reid的下级分类
self: 同级 父ID为reid的同级栏目

row = "10" 调用数量

channel(typeid, reid, type, row)
*/
func Channel(args jet.Arguments) reflect.Value {
	if !checkArgType(&args) {
		return defaultArrReturnVal
	}
	_typeid := getNumber(args.Get(0))
	_reid := getNumber(args.Get(1))
	_type := args.Get(2).String()
	_row := int(getNumber(args.Get(3)))
	_noself := args.Get(4).String()
	if _row == 0 {
		_row = 10
	}
	m := models.NewCategoryModel()
	if _typeid != 0 {
		cat := m.GetCategory(_typeid)
		if cat == nil {
			return defaultArrReturnVal
		}
		_reid = cat.Parentid
	}
	orm := getCategoryOrm().Limit(_row).Asc("listorder")
	switch _type {
	case "top":
		_reid = 0
		orm.Where("parentid = 0")
	case "son":
		if _typeid == 0 { // 没有设置typeid 返回空
			return defaultArrReturnVal
		}
		orm.Where("parentid = ?", _typeid)
	case "self":
		orm.Where("parentid = ?", _reid)
		if _noself == "yes" {
			orm.Where("catid <> ?", _typeid)
		}
	}
	var arr = []tables.Category{}
	orm.Find(&arr)

	if len(arr) == 0 && _type == "son" && _reid != 0 {
		//如果用子栏目模式，当没有子栏目时显示同级栏目
		getCategoryOrm().Limit(_row).Asc("listorder").Where("parentid = ?", _reid).Find(&arr)
	}
	for k, v := range arr {
		if v.Type != 2 {
			cat1s := m.GetPosArr(v.Catid)
			arr[k].Url = fmt.Sprintf("/%s/", m.GetUrlPrefixWithCategoryArr(cat1s))
		}
	}
	return reflect.ValueOf(arr)
}
