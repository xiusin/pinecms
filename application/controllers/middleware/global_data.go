package middleware

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"iriscms/application/models/tables"
)

func FrontendGlobalViewData(app *iris.Application) func(ctx context.Context) {
	return func(ctx context.Context) {
		ctx.ViewData("setting", nil)
		//app.Logger().Println("后端使用 groupcache 缓存前端需要的全局数据,在这里也可以添加公共函数和缓存 header")
		ctx.Next()
	}
}

func SetGlobalConfigData(xorm *xorm.Engine) func(ctx context.Context)  {
	//读取配置项
	return func(ctx context.Context) {
		var settings []tables.IriscmsSetting
		err := xorm.Find(&settings)
		if err != nil {
			ctx.Application().Logger().Error("无法读取到配置内容")
			ctx.StopExecution()
			return
		}
		var settingData = map[string] string{}
		if len(settings) != 0 {
			for _, v := range settings {
				settingData[v.Key] = v.Value
			}
		}
		ctx.Values().Set("setting", settingData)
		ctx.Next()
	}


}