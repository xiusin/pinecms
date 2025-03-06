package backend

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/common/storage"
)

func getStorageEngine(settingData map[string]string) storage.Uploader {
	engine := settingData["UPLOAD_ENGINE"]
	var uploadEngine storage.Uploader
	uploader, err := di.Get(fmt.Sprintf(controllers.ServiceUploaderEngine, engine))
	if err != nil {
		pine.Logger().Warn("缺少存储驱动, 自动转换为本地存储", err)
		uploadEngine = storage.NewFileUploader(settingData)
	} else {
		uploadEngine = uploader.(storage.Uploader)
	}
	return uploadEngine
}

func parseParam(ctx *pine.Context, param any) error {
	if ctx.Input().IsJson() && len(ctx.RequestCtx.PostBody()) > 0 {
		return ctx.BindJSON(param)
	}
	return nil
}

func ArrayCol(arr any, col string) []any {
	val := reflect.ValueOf(arr)
	if val.Kind() != reflect.Slice {
		panic(errors.New("ArrayCol第一个参数必须为切片类型"))
	}
	var cols []any
	for i := range val.Len() {
		cols = append(cols, val.Index(i).FieldByName(col).Interface())
	}
	return cols
}

func ArrayColMap(arr any, col string) map[any]any {
	var maps = map[any]any{}
	val := reflect.ValueOf(arr)
	if val.Kind() != reflect.Slice {
		panic(errors.New("ArrayCol第一个参数必须为切片类型"))
	}
	for i := range val.Len() {
		maps[val.Index(i).FieldByName(col).Interface()] = val.Index(i).Interface()
	}
	return maps
}
