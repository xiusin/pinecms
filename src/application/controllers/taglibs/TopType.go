package taglibs

import (
	"fmt"
	"reflect"

	"github.com/CloudyKit/jet"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/common/helper"
)

/**
 * 标签：{{yield toptype() content}}{{field.Catname}}{{end}}
 * 作用：获取当前页面面包屑只可适用于列表页和详情页
 */
func TopType(args jet.Arguments) reflect.Value {
	fmt.Println("arclist")

	if !checkArgType(&args) {
		return defaultArrReturnVal
	}
	typeid := getNumber(args.Get(0))
	cat, err := models.NewCategoryModel().GetCategoryFByIdForBE(typeid)
	helper.PanicErr(err)
	if cat.Topid != 0 {
		cat, err = models.NewCategoryModel().GetCategoryFByIdForBE(cat.Topid)
		helper.PanicErr(err)
	}
	cat.Content = ""
	cat.Page = nil
	return reflect.ValueOf(cat)
}
