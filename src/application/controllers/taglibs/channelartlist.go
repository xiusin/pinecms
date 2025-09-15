package taglibs

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/CloudyKit/jet"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"xorm.io/xorm"
)

/*
*
sons 是否附带直属子分类数据, 可用于下级数据的迭代, 可取代channel的sun功能
active 是否附带当前活动状态逻辑, 一般用于下级菜单激活上级菜单的高亮样式
topid 记录当前所处页面的tid
{{channellist = channelartlist(typeid, row, topid, sons, active)}}
*/
func ChannelArtList(args jet.Arguments) reflect.Value {
	var cats []tables.Category
	cacheKey := "pinecms:tag:channelartlist:" + getTagHash(args)
	err := helper.Cache().Remember(cacheKey, &cats, func() (any, error) {
		if !checkArgType(&args) {
			return &[]tables.Category{}, nil
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
			orm = getCategoryOrm().Where("parentid = 0")
		} else if strings.Contains(_typeid, ",") {
			orm = getCategoryOrm().In("catid", strings.Split(_typeid, ","))
		} else {
			orm = getCategoryOrm().Where("parentid = ?", _typeid)
		}

		_topId := getNumber(args.Get(2)) // 当前页面的ID

		var tempCats []tables.Category

		orm.Limit(_row).Select("id, parentid, catname, type, model_id, description, thumb, url").Asc("listorder").Find(&tempCats)

		if len(tempCats) == 0 {
			return &[]tables.Category{}, nil
		}

		allCategories, err := models.NewCategoryModel().GetAll(true)
		if err != nil {
			return nil, err
		}

		childrenMap := map[int64][]tables.Category{}
		for _, cat := range allCategories {
			childrenMap[cat.Parentid] = append(childrenMap[cat.Parentid], cat)
		}

		m := models.NewCategoryModel()

		withSons := args.Get(3).Bool()
		withActive := args.Get(4).Bool()

		for i, v := range tempCats {
			if v.Type != 2 {
				if withSons {
					if _, ok := childrenMap[v.Catid]; ok {
						tempCats[i].HasSon = true
					}
				}
				posArr, _ := m.GetPosArr(v.Catid)

				if withActive {
					for _, p := range posArr {
						if p.Catid == _topId {
							tempCats[i].Active = true
							break
						}
					}
				}
				tempCats[i].Url = fmt.Sprintf("/%s/", m.GetUrlPrefixWithCategoryArr(posArr))
			}
		}
		return &tempCats, nil
	})
	if err != nil {
		pine.Logger().Error(err)
	}
	return reflect.ValueOf(cats)
}
