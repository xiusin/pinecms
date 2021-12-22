package taglibs

import (
	"github.com/CloudyKit/jet"
	"github.com/xiusin/pinecms/src/application/models"
	"reflect"
)

/**
 * 标签：{{yield toptype() content}}{{field.Catname}}{{end}}
 * 作用：获取当前页面面包屑只可适用于列表页和详情页
 */
func TopType(args jet.Arguments) reflect.Value {
	if !checkArgType(&args) {
		return defaultArrReturnVal
	}
	typeid := getNumber(args.Get(0))
	cat, err := models.NewCategoryModel().GetCategoryFByIdForBE(typeid)
	if err != nil {
		panic(err)
	}
	if cat.Topid != 0 {
		cat, err = models.NewCategoryModel().GetCategoryFByIdForBE(cat.Topid)
		if err != nil {
			panic(err)
		}
	}
	cat.Content = ""
	cat.Page = nil
	return reflect.ValueOf(cat)
}
