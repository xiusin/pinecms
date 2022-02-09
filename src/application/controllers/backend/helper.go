package backend

import (
	"errors"
	"fmt"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/common/storage"
	"reflect"
	"strings"

	"github.com/xiusin/pinecms/src/application/controllers"
)

func getStorageEngine(settingData map[string]string) storage.Uploader {
	engine := settingData["UPLOAD_ENGINE"]
	var uploadEngine storage.Uploader
	uploader, err := di.Get(fmt.Sprintf(controllers.ServiceUploaderEngine, engine))
	if err != nil {
		pine.Logger().Warning("缺少存储驱动, 自动转换为本地存储", err)
		uploadEngine = storage.NewFileUploader(settingData)
	} else {
		uploadEngine = uploader.(storage.Uploader)
	}
	return uploadEngine
}

func strFirstToUpper(str string) string {
	temp := strings.Split(strings.ReplaceAll(str, "_", "-"), "-")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		if y != 0 {
			for i := 0; i < len(vv); i++ {
				if i == 0 {
					vv[i] -= 32
					upperStr += string(vv[i]) // + string(vv[i+1])
				} else {
					upperStr += string(vv[i])
				}
			}
		}
	}
	return temp[0] + upperStr
}

func parseParam(ctx *pine.Context, param interface{}) error {
	if ctx.IsJson() && len(ctx.RequestCtx.PostBody()) > 0 {
		return ctx.BindJSON(param)
	}
	return nil
}

func ArrayCol(arr interface{}, col string) []interface{} {
	val := reflect.ValueOf(arr)
	if val.Kind() != reflect.Slice {
		panic(errors.New("ArrayCol第一个参数必须为切片类型"))
	}
	var cols []interface{}
	for i := 0; i < val.Len(); i++ {
		cols = append(cols, val.Index(i).FieldByName(col).Interface())
	}
	return cols
}

func ArrayColMap(arr interface{}, col string) map[interface{}]interface{} {
	var maps = map[interface{}]interface{}{}
	val := reflect.ValueOf(arr)
	if val.Kind() != reflect.Slice {
		panic(errors.New("ArrayCol第一个参数必须为切片类型"))
	}
	var cols []interface{}
	for i := 0; i < val.Len(); i++ {
		cols = append(cols)
		maps[val.Index(i).FieldByName(col).Interface()] = val.Index(i).Interface()
	}
	return maps
}
