package taglibs

import (
	"github.com/CloudyKit/jet"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/config"
	"reflect"
	"runtime/debug"
	"strings"
)

func HotWords(args jet.Arguments) reflect.Value {
	if !checkArgType(&args) {
		return defaultArrReturnVal
	}
	defer func() {
		if err := recover(); err != nil {
			pine.Logger().Error("HotWords Failed", err, string(debug.Stack()))
		}
	}()
	conf,_ := config.SiteConfig()
	words := strings.Split(conf["SITE_HOTWORDS"], ",")
	return reflect.ValueOf(words)
}
