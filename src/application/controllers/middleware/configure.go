package middleware

import (
	"github.com/fatih/structs"
	"github.com/xiusin/pine"
)

var defaultConfig *Config

// SetApiEntity 设置接口实体信息 immutable设置不可变, 在文档存在的情况下不会覆盖已生成的文档, 一般用于确定的文档信息
// reqParams 参数实体
func SetApiEntity(ctx *pine.Context, title, desc string, reqParams interface{}, configures ...Configure) {
	if reqParams != nil && !structs.IsStruct(reqParams) {
		ctx.Logger().Warning("不支持非struct类型的请求参数")
	} else if defaultConfig.Enable {
		entity, ok := ctx.Value(apiDocKey).(*apiEntity)
		if ok {
			entity.Title = title
			entity.Desc = desc
			entity.configed = true
			for _, configure := range configures {
				configure(entity)
			}
			entity.Param = parseInterface(reqParams)
		}
	}
}

type Configure func(entity *apiEntity)

func WithImmutable(immutable bool) Configure {
	return func(entity *apiEntity) {
		entity.immutable = immutable
	}
}


func WithHeaders(headers []apiHeader) Configure {
	return func(entity *apiEntity) {
		entity.Header = append(entity.Header, headers...)
	}
}
