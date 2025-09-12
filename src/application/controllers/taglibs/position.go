package taglibs

import (
	"fmt"
	"github.com/CloudyKit/jet"
	"github.com/xiusin/pine/contracts"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"reflect"
	"strings"
)

/**
 * 标签：{{yield position() }}
 * 作用：获取当前页面面包屑只可适用于列表页和详情页
 */
func Position(args jet.Arguments) reflect.Value {
	if !checkArgType(&args) {
		return defaultArrReturnVal
	}
	typeid := getNumber(args.Get(0))
	return reflect.ValueOf(getCategoryPos(typeid))
}

func getCategoryPos(tid int64) string {
	var pos string
	key := fmt.Sprintf(controllers.CacheCategoryPosPrefix, tid)
	icache := di.MustGet(controllers.ServiceICache).(contracts.Cache)
	err := icache.Get(key, &pos)
	if err != nil {
		m := models.NewCategoryModel()
		arr, err := m.GetPosArr(tid)
		if err != nil {
			pine.Logger().Error(err)
			return ""
		}
		var position []string
		for _, cat := range arr {
			if cat.Type != 2 {
				prefix, err := m.GetUrlPrefix(cat.Catid)
				if err != nil {
					pine.Logger().Error(err)
					return ""
				}
				position = append(position, "<a href='"+prefix+"'>"+cat.Catname+"</a>")
			} else {
				position = append(position, "<a href='"+cat.Url+"'>"+cat.Catname+"</a>")
			}
		}
		if len(arr) > 0 {
			pos = strings.Join(position, " > ")
			_ = icache.Set(key, pos, 0)
		}
	}
	return pos
}
