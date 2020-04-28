package taglibs

import (
	"github.com/CloudyKit/jet"
	"github.com/xiusin/pine"
	"reflect"
	"runtime/debug"
	"strings"
)

func Tags(args jet.Arguments) reflect.Value {
	if !checkArgType(&args) {
		return defaultArrReturnVal
	}
	defer func() {
		if err := recover(); err != nil {
			pine.Logger().Error("HotWords Failed", err, string(debug.Stack()))
		}
	}()
	tags := strings.Split(args.Get(0).String(), ",")
	return reflect.ValueOf(tags)
}
