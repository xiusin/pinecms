package taglibs

import (
	"github.com/CloudyKit/jet"
	"github.com/xiusin/pine"
	"reflect"
	"strings"
)

func Tags(args jet.Arguments) reflect.Value {
	var tags = []string{}
	helper.Cache().Remember("pinecms:tag:tags:"+getTagHash(args), &tags, func() (any, error) {
		if !checkArgType(&args) {
			return &tags, nil
		}
		tags = strings.Split(args.Get(0).String(), ",")
		return &tags, nil
	})
	return reflect.ValueOf(tags)
}
