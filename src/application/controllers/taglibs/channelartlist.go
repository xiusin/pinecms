package taglibs

import (
	"fmt"
	"github.com/CloudyKit/jet"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"reflect"
	"strings"
)

/**
{{channellist = channelartlist(typeid, row, topid)}}
 */
func ChannelArtList(args jet.Arguments) reflect.Value {
	if !checkArgType(&args) {
		return defaultArrReturnVal
	}
	_row := int(getNumber(args.Get(1)))
	if _row <= 0 {
		_row = 20
	}
	var _typeid string
	if isNumber(args.Get(0)) {
		_typeid = fmt.Sprintf("%d", getNumber(args.Get(0)))
	} else {
		_typeid = args.Get(0).String()
	}

	var orm *xorm.Session
	if _typeid == "0" || _typeid == "top" {
		orm =  getCategoryOrm().Where("parentid = 0").Where("model_id > 0")
	} else if strings.Contains(_typeid, ",") {
		orm =  getCategoryOrm().In("catid", strings.Split(_typeid, ","))
	} else {
		orm =  getCategoryOrm().Where("parentid = ?", _typeid)
	}

	_topid := getNumber(args.Get(2))	// 当前页面的ID

	var cats []tables.Category

	orm.Limit(_row).Asc("listorder").Find(&cats)
	if len(cats) == 0 {
		return defaultArrReturnVal
	}
	m := models.NewCategoryModel()

	for k,v := range cats {
		if v.Type != 2 {
			cat1s := m.GetPosArr(v.Catid)
			cats[k].Url = fmt.Sprintf("/%s/", m.GetUrlPrefixWithCategoryArr(cat1s))
			for _, v:= range cat1s {
				if v.Catid == _topid {
					cats[k].Active = true
					break
				}
			}
		}
	}
	return reflect.ValueOf(cats)
}
