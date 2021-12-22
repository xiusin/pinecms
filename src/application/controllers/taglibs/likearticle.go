package taglibs

import (
	"fmt"
	"github.com/CloudyKit/jet"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/common"
	"reflect"
	"strings"
)

func LikeArticle(args jet.Arguments) reflect.Value {
	defer func() {
		if err := recover(); err != nil {
			pine.Logger().Errorf("likearticle Failed %s", err)
		}
	}()
	if !checkArgType(&args) {
		return defaultArrReturnVal
	}
	var kws = []string{args.Get(1).String(), args.Get(2).String(), args.Get(3).String()}
	limit := int(getNumber(args.Get(0)))
	catid := getNumber(args.Get(4))
	titlelen := int(getNumber(args.Get(6)))
	var keywords []interface{}
	if limit < 1 {
		limit = 10
	}
	var cond []string
	m := models.NewCategoryModel()
	category, _ := m.GetCategoryFByIdForBE(catid)
	modelTable := controllers.GetTableName(category.Model.Table)
	sess := getOrmSess(category.Model.Table).Where(modelTable+".id <> ?", getNumber(args.Get(5)))
	for _, kw := range kws {
		splitKeywords := strings.Split(kw, ",")
		for _, keyword := range splitKeywords {
			if keyword == "" {
				continue
			}
			cond = append(cond, fmt.Sprintf("%s.keywords LIKE ? OR %s.title LIKE ? OR %s.tags LIKE ?", modelTable, modelTable, modelTable))
			keywords = append(keywords, "%"+strings.Trim(keyword, "")+"%", "%"+strings.Trim(keyword, "")+"%", "%"+strings.Trim(keyword, "")+"%")
		}
	}
	categoryTable := getCategoryTable()
	sess.And(strings.Join(cond, " OR "), keywords...)
	sess.Join("LEFT", categoryTable, categoryTable+".catid = "+modelTable+".catid")
	sess.Select(fmt.Sprintf("%s.*, %s.catname as typename", modelTable, categoryTable))
	list, _ := sess.Limit(limit).Desc(modelTable + ".id").QueryString()
	common.HandleArtListInfo(list, titlelen)
	return reflect.ValueOf(list)
}
