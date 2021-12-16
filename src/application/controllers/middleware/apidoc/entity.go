package apidoc

import (
	"github.com/fatih/structs"
	"github.com/xiusin/pine"
)

type Entity struct {
	Title    string      // 标题
	Desc     string      // 描述
	ApiParam interface{} // 参数
	AppId    string      // 应用id
	Group    string      // 分组名称
	SubGroup string      // 子分组名称
	Tag      []string    // 标签
}

// SetApiEntity 设置接口实体信息
func SetApiEntity(ctx *pine.Context, entity *Entity, configures ...Configure) {
	if entity.ApiParam != nil && !structs.IsStruct(entity.ApiParam) {
		ctx.Logger().Warning("不支持非struct类型的请求参数")
	} else if defaultConfig.Enable {
		e, ok := ctx.Value(apiDocKey).(*apiEntity)
		if ok {
			e.Title = entity.Title
			e.Desc = entity.Desc
			e.configured = true
			e.Group = apiGroup{Title: entity.Group, Name: entity.Group}
			e.SubGroup = entity.SubGroup
			for _, configure := range configures {
				configure(e)
			}
			e.AppId = entity.AppId
			e.Param, _ = parseInterface(entity.ApiParam)
			e.FilterParams()
		}
	}
}
