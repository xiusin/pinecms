package taglibs

import (
	"fmt"
	"github.com/CloudyKit/jet"
	"github.com/xiusin/pine/cache"
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
	var position []string
	var res string
	var data = struct {
		Arr []tables.Category
		Pos string
	}{}
	key := fmt.Sprintf(controllers.CacheCategoryPosPrefix, tid)
	icache := di.MustGet(controllers.ServiceICache).(cache.AbstractCache)
	err := icache.GetWithUnmarshal(key, &data)
	if err != nil {
		m := models.NewCategoryModel()
		data.Arr = m.GetPosArr(tid)
		for _, cat := range data.Arr {
			if cat.Type != 2 {
				position = append(position, "<a href='"+m.GetUrlPrefix(cat.Catid)+"'>"+cat.Catname+"</a>")
			} else {
				position = append(position, "<a href='"+cat.Url+"'>"+cat.Catname+"</a>")
			}
		}
		if len(data.Arr) > 0 {
			res = strings.Join(position, " > ")
			data.Pos = res
			icache.SetWithMarshal(key, res)
		}
	}
	return data.Pos
}

