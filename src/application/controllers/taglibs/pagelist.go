package taglibs

import (
	"fmt"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/config"
	"math"
	"reflect"
	"strconv"

	"github.com/CloudyKit/jet"
)

// PageList {{pagelist(row, .TypeID, .ArtCount, .PageNum, .QP) | unsafe}}
func PageList(args jet.Arguments) reflect.Value {
	var pageStr string
	helper.Cache().Remember("pinecms:tag:pagelist:"+getTagHash(args), &pageStr, func() (any, error) {
		limit := int(getNumber(args.Get(0)))
		tid := getNumber(args.Get(1))
		total := int(getNumber(args.Get(2)))
		page := int(getNumber(args.Get(3)))
		qp, _ := args.Get(4).Interface().(map[string][]string)
		conf, err := config.SiteConfig()
		if err != nil {
			return &pageStr, err
		}
		if limit == 0 {
			limit, _ = strconv.Atoi(conf["SITE_PAGE_SIZE"])
			if limit == 0 {
				limit = 15
			}
		}
		totalPage := int(math.Ceil(float64(total) / float64(limit)))
		if page > totalPage {
			page = totalPage
		}
		if len(qp) == 0 {
			prefix, err := models.NewCategoryModel().GetUrlPrefix(tid)
			if err != nil {
				return &pageStr, err
			}
			pageStr = helper.NewPage("/"+prefix, page, limit, total, nil, false).String()
		} else {
			// todo 地址后期path
			pageStr = helper.NewPage("/search.go", page, limit, total, qp, true).String()
		}
		return &pageStr, nil
	})
	return reflect.ValueOf(pageStr)
}
