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
sons 是否附带直属子分类数据, 可用于下级数据的迭代, 可取代channel的sun功能
active 是否附带当前活动状态逻辑, 一般用于下级菜单激活上级菜单的高亮样式
topid 记录当前所处页面的tid
{{channellist = channelartlist(typeid, row, topid, sons, active)}}
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
		orm =  getCategoryOrm().Where("parentid = 0")
	} else if strings.Contains(_typeid, ",") {
		orm =  getCategoryOrm().In("catid", strings.Split(_typeid, ","))
	} else {
		orm =  getCategoryOrm().Where("parentid = ?", _typeid)
	}

	_topid := getNumber(args.Get(2))	// 当前页面的ID

	var cats []tables.Category

	orm.Limit(_row).Select("catid, parentid, catname, type, model_id, description, thumb, url").Asc("listorder").Find(&cats)

	if len(cats) == 0 {
		return defaultArrReturnVal
	}
	m := models.NewCategoryModel()

	withSons := args.Get(3).Bool()
	withActive := args.Get(4).Bool()

	for k,v := range cats {
		if v.Type != 2 {
			if withSons {
				var sons tables.Category
				count, _ := getCategoryOrm().Where("parentid = ?", v.Catid).Count(&sons)
				if count > 0 {
					cats[k].HasSon = true
				}
			}
			if withActive {
				cat1s := m.GetPosArr(v.Catid)
				cats[k].Url = fmt.Sprintf("/%s/", m.GetUrlPrefixWithCategoryArr(cat1s))
				for _, v:= range cat1s {
					if v.Catid == _topid {
						cats[k].Active = true
						break
					}
				}
			}
			cats[k].Url = fmt.Sprintf("/%s/", m.GetUrlPrefixWithCategoryArr(m.GetPosArr(v.Catid)))
		}
	}
	return reflect.ValueOf(cats)
}
