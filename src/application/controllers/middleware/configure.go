package middleware

import (
	"github.com/fatih/structs"
	"github.com/xiusin/pine"
)

var defaultConfig *Config

type Entity struct {
	Title    string
	Desc     string
	ApiParam interface{}
	Group    string
	SubGroup string
}

// SetApiEntity 设置接口实体信息 immutable设置不可变, 在文档存在的情况下不会覆盖已生成的文档, 一般用于确定的文档信息
// reqParams 参数实体
func SetApiEntity(ctx *pine.Context, entity *Entity, configures ...Configure) {
	if entity.ApiParam != nil && !structs.IsStruct(entity.ApiParam) {
		ctx.Logger().Warning("不支持非struct类型的请求参数")
	} else if defaultConfig.Enable {
		e, ok := ctx.Value(apiDocKey).(*apiEntity)
		if ok {
			e.Title = entity.Title
			e.Desc = entity.Desc
			e.configed = true
			withGroup(entity.Group)
			e.subGroup = entity.SubGroup
			for _, configure := range configures {
				configure(e)
			}
			e.Param = parseInterface(entity.ApiParam)
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

func withGroup(keyOrName string) Configure {
	return func(entity *apiEntity) {
		for _, group := range defaultConfig.Groups {
			if group.Name == keyOrName || group.Title == keyOrName {
				entity.group = group
			}
		}
	}
}
