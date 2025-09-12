package taglibs

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/CloudyKit/jet"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/common/helper"
)

func LikeArticle(args jet.Arguments) reflect.Value {
	var list = []map[string]string{}
	helper.Cache().Remember("pinecms:tag:likearticle:"+getTagHash(args), &list, func() (any, error) {
		if !checkArgType(&args) {
			return &list, nil
		}
		var kws = []string{args.Get(1).String(), args.Get(2).String(), args.Get(3).String()}
		limit := int(getNumber(args.Get(0)))
		catid := getNumber(args.Get(4))
		titlelen := int(getNumber(args.Get(6)))
		if limit < 1 {
			limit = 10
		}
		m := models.NewCategoryModel()
		category, err := m.GetCategoryFByIdForBE(catid)
		if err != nil {
			pine.Logger().Error(err)
			return &list, err
		}
		modelTable := controllers.GetTableName(category.Model.Table)
		sess := getOrmSess(category.Model.Table)
		defer sess.Close()

		b := builder.Select("a.*", "c.catname as typename").From(modelTable, "a").
			LeftJoin(getCategoryTable(), "c", "c.id = a.catid").
			Where(builder.Neq{"a.id": getNumber(args.Get(5))})

		var keywordConds []builder.Cond
		for _, kw := range kws {
			splitKeywords := strings.Split(kw, ",")
			for _, keyword := range splitKeywords {
				if keyword == "" {
					continue
				}
				trimmedKeyword := "%" + strings.Trim(keyword, "") + "%"
				keywordConds = append(keywordConds, builder.Like{"a.keywords", trimmedKeyword})
				keywordConds = append(keywordConds, builder.Like{"a.title", trimmedKeyword})
				keywordConds = append(keywordConds, builder.Like{"a.tags", trimmedKeyword})
			}
		}
		if len(keywordConds) > 0 {
			b.Where(builder.Or(keywordConds...))
		}

		list, err = sess.QueryString(b.Limit(limit).Desc("a.id"))
		if err != nil {
			pine.Logger().Error(err)
			return &list, err
		}
		helper.HandleArtListInfo(list, titlelen)
		return &list, nil
	})
	return reflect.ValueOf(list)
}
