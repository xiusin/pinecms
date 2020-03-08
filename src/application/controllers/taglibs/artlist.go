package taglibs

import (
	"github.com/CloudyKit/jet"
	"reflect"
	"strconv"
	"strings"
)

func ArcList(args jet.Arguments) reflect.Value {
	//typeid,offset,limit,orderby,modelid
	catid := args.Get(0)
	var ids []string
	switch catid.Type().String() {
	case "string":
		ids = strings.Split(catid.String(), ",")
	default:
		ids = append(ids, strconv.Itoa(getInt(catid)))
	}
	orderBy := args.Get(3).String()
	if orderBy == "" {
		orderBy = "listorder desc"
	}
	offset := getInt(args.Get(1))
	limit := getInt(args.Get(2))
	sess := getOrmSess().Limit(limit, offset).Where("deleted_time IS NULL").Where("status = 1").OrderBy(orderBy)
	if ids[0] != "0"  && ids[0] != "-1"{
		sess.In("catid", ids)
	}
	list, err := sess.QueryString()
	if err != nil {
		panic(err)
	}
	if list == nil {
		list = []map[string]string{}
	}
	return reflect.ValueOf(list)
}
