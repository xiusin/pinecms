package taglibs

import (
	"fmt"
	"reflect"
	"time"

	"github.com/CloudyKit/jet"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

/*
*
typeid = "son | top | self"
top: 顶级栏目 : parentid 为 0
son: 父id 为 reid的下级分类
self: 同级 父ID为reid的同级栏目

row = "10" 调用数量

channel(typeid, reid, type, row)
*/
func Channel(args jet.Arguments) reflect.Value {
	var arr = []tables.Category{}
	helper.Cache().Remember("pine:tag:channel:"+getTagHash(args), &arr, func() (any, error) {
		if !checkArgType(&args) {
			return &arr, nil
		}
		startTime := time.Now()

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
				return &arr, nil
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
				return &arr, nil
			}
			orm.Where("parentid = ?", _typeid)
		case "self":
			orm.Where("parentid = ?", _reid)
			if _noself == "yes" {
				orm.Where("catid <> ?", _typeid)
			}
		}

		err := orm.Find(&arr)
		if err != nil {
			return &arr, err
		}

		if len(arr) == 0 && _type == "son" && _reid != 0 {
			//如果用子栏目模式，当没有子栏目时显示同级栏目
			err = getCategoryOrm().Limit(_row).Asc("listorder").Where("parentid = ?", _reid).Find(&arr)
			if err != nil {
				return &arr, err
			}
		}
		for k, v := range arr {
			if v.Type != 2 {
				cat1s, err := m.GetPosArr(v.Catid)
				if err != nil {
					return &arr, err
				}
				arr[k].Url = fmt.Sprintf("/%s/", m.GetUrlPrefixWithCategoryArr(cat1s))
			}
		}
		fmt.Println("channel 耗时", time.Now().Sub(startTime))

		return &arr, nil
	})
	return reflect.ValueOf(arr)
}
