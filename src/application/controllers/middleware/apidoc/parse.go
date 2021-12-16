package apidoc

import (
	"github.com/fatih/structs"
	"reflect"
	"strings"
)

func parseInterface(reqParams interface{}) ([]apiParam, []apiReturn) {
	if reqParams != nil {
		s := structs.Fields(reqParams)
		var apiReqParams []apiParam
		var apiReturns []apiReturn
		for _, field := range s {
			apiData := field.Tag("api")
			if len(apiData) == 0 {
				continue
			}
			fieldType := "any"
			switch field.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				fieldType = "number"
			case reflect.Bool:
				fieldType = "bool"
			case reflect.String:
				fieldType = "string"
			case reflect.Struct, reflect.Map:
				fieldType = "object"
			case reflect.Slice, reflect.Array:
				fieldType = "array"
			case reflect.Float32, reflect.Float64:
				fieldType = "float"
			}
			apiDatas := strings.Split(apiData, "|")
			name := field.Tag("json")
			if len(name) == 0 {
				name = field.Name()
			}
			apiReqParam := apiParam{Type: fieldType, Name: name}
			apiReturn := apiReturn{Type: fieldType, Name: name, Params: nil}
			for _, data := range apiDatas {
				kv := strings.Split(strings.ToLower(data), ":")
				v := ""
				if len(kv) > 1 {
					v = kv[1]
				}
				switch kv[0] {
				case "require":
					if strings.ToLower(v) == "true" {
						apiReqParam.Require = true
					}
				case "remark":
					apiReqParam.Desc = v
					apiReturn.Desc = v
				case "main":
					if strings.ToLower(v) == "true" {
						apiReturn.Main = true
					}
				case "default":
					apiReqParam.Default = v
					apiReturn.Default = v
				}
			}
			apiReqParams = append(apiReqParams, apiReqParam)
			apiReturns = append(apiReturns, apiReturn)
		}
		return apiReqParams, apiReturns
	}
	return nil, nil
}
