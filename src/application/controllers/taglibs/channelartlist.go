package taglibs

import (
	"github.com/CloudyKit/jet"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"reflect"
	"strconv"
	"strings"
)

func ChannelArtList(args jet.Arguments) reflect.Value {
	catid := args.Get(0)
	row := getInt(args.Get(1))
	if row == 0 {
		row = 10
	}

	son := args.Get(2).Bool()

	var ids []string
	str := catid.Type().String()
	if strings.HasPrefix(str, "float") {
		ids = append(ids, strconv.Itoa(int(catid.Float())))
	} else if strings.Contains(str, "int") {
		ids = append(ids, strconv.Itoa(int(catid.Int())))
	} else {
		ids = strings.Split(catid.String(), ",")
	}
	sess := pine.Make("*xorm.Engine").(*xorm.Engine).Table(&tables.IriscmsCategory{})
	var categories []tables.IriscmsCategory
	err := sess.In("parentid", ids).Where("ismenu = ?", 1).Limit(int(row)).Desc("listorder").Find(&categories)
	if err != nil {
		panic(err)
	}
	var catids = []int64{}
	for k, v := range categories {
		if v.Type == 0 {
			categories[k].Url = "/list?tid=" + strconv.Itoa(int(v.Catid))
		} else if v.Type == 1 {
			categories[k].Url = "/page?tid=" + strconv.Itoa(int(v.Catid))
		}
		catids = append(catids, v.Catid)
	}

	// 判断是否有下级菜单
	if son {
		sess := pine.Make("*xorm.Engine").(*xorm.Engine).
			Table(&tables.IriscmsCategory{}).
			In("parentid", catids).GroupBy("parentid").Select("parentid,count(*) as total").Where("ismenu = ?", 1)
		rest, _ := sess.QueryString()
		var kvPairs = map[string]string{}
		for _, v := range rest {
			kvPairs[v["parentid"]] = v["total"]
		}
		for k, v := range categories {
			val, exists := kvPairs[strconv.Itoa(int(v.Catid))]
			if exists && len(val) > 0 && val != "0" {
				categories[k].HasSon = true
			}
		}
	}

	return reflect.ValueOf(categories)
}
