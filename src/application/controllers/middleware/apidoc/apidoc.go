package apidoc

import (
	"fmt"
	"github.com/sonyarouje/simdb"
	"github.com/xiusin/pine"
	"net/http"
	"os"
	"strings"
)

/**
1. 可以自定义字段备注， 比如想要修改部分字段备注
2. 修改接口信息， 为响应字段添加备注
3.
*/
var simdbDriver *simdb.Driver

func New(app *pine.Application, config *Config) pine.Handler {
	if config == nil {
		config = DefaultConfig()
	}
	defaultConfig = config
	err := os.Mkdir(config.DataPath, os.ModePerm)
	simdbDriver, err = simdb.New(config.DataPath)
	if err != nil {
		panic(err)
	}
	app.ANY(config.RoutePrefix+"/:action", func(ctx *pine.Context) {
		if !config.Enable {
			ctx.Next()
			return
		}
		switch ctx.Params().Get("action") {
		case "config":
			getConfig(ctx)
		case "apiData":
			getApiData(ctx)
		case "edit":
			saveApiData(ctx)
		//标注接口字段为不可再修改， go端不可直接配置
		case "reset":

		}
	})

	return func(ctx *pine.Context) {
		method := string(ctx.Method())
		if strings.Contains(ctx.Path(), config.RoutePrefix+"/") || method == http.MethodOptions {
			ctx.Next()
			return
		}
		ps := strings.Split(ctx.Path(), "/")
		key := strings.ToLower(fmt.Sprintf("%s_%s", method, strings.ReplaceAll(ctx.Path(), "/", "_")))
		entity := &apiEntity{URL: ctx.Path(), Method: method, Name: ps[len(ps)-1], Header: defaultConfig.Headers, MenuKey: key}
		ctx.Set(apiDocKey, entity)
		ctx.Next()
		if !entity.configured {
			return
		}

		var existData apiEntity
		exist := simdbDriver.Open(entity).Where("menu_key", "=", key).First().AsEntity(&existData)
		if entity.Immutable == true && exist == nil { // 代码内传入不可变
			return
		}
		if existData.Immutable == true { // 实体内传入不可变
			return
		}
		entity.Author = defaultConfig.DefaultAuthor
		entity.RawQuery = string(ctx.RequestCtx.QueryArgs().QueryString())
		entity.QueryDataMethod = string(ctx.Response.Header.ContentType())
		if strings.Contains(strings.ToLower(entity.QueryDataMethod), "application/json") {
			entity.RawParam = string(ctx.Request.Body())
		} else {
			entity.RawParam = string(ctx.PostBody())
		}
		entity.RawReturn = string(ctx.Response.Body())
		_, entity.Return = parseInterface(defaultConfig.ResponseParam)
		if err := simdbDriver.Upsert(entity); err != nil {
			pine.Logger().Warning("保存接口数据失败", err)
		}
	}
}
