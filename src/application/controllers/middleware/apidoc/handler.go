package apidoc

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/common/helper"
	"strings"
)

func getConfig(ctx *pine.Context) {
	_ = ctx.WriteJSON(pine.H{"code": 0, "data": defaultConfig})
}

func getApiData(ctx *pine.Context) {
	appKey, _ := ctx.GetString("appKey", "")
	if appKey == "" {
		_ = ctx.WriteJSON(pine.H{"code": 1, "msg": "参数错误"})
		return
	}
	var entities []apiEntity
	_ = simdbDriver.Open(&apiEntity{}).Where("app_id", "=", appKey).Get().AsEntity(&entities)
	groups := []apiGroup{{Title: "全部", Name: 0}}
	var lists = []apiList{}
	for _, entity := range entities {
		if len(entity.SubGroup) == 0 {
			ctx.Logger().Warning(entity.Title + "没有设置SubGroup")
			continue
		}
		var flag = true
		for _, group := range groups {
			if group.Title == entity.Group.Title {
				flag = false
				break
			}
		}
		if flag {
			groups = append(groups, entity.Group)
		}
		var listItem apiList
		var idx = -1
		for i, list := range lists {
			if list.Title == entity.SubGroup {
				listItem = list
				idx = i
			}
		}

		if len(entity.Param) == 0 {
			entity.Param = []apiParam{}
		}

		if len(entity.Query) == 0 {
			entity.Query = []apiParam{}
		}

		listItem.Group = entity.Group.Title
		listItem.MenuKey = entity.SubGroup
		listItem.Title = entity.SubGroup
		listItem.Children = append(listItem.Children, entity)

		if idx > -1 {
			lists[idx] = listItem
		} else {
			lists = append(lists, listItem)
		}
	}

	_ = ctx.WriteJSON(pine.H{"code": 0, "data": apiData{
		Groups: groups,
		List:   lists,
	}})
}

func saveApiData(ctx *pine.Context) {
	typo, _ := ctx.GetString("type")
	switch typo {
	case "request":
		saveRequestData(ctx)
	default:
		helper.Ajax("暂不支持的类型", 1, ctx)
	}
}

func saveRequestData(ctx *pine.Context) {
	var data []apiParam
	var retData []apiReturn
	var entity apiEntity
	if err := ctx.BindJSON(&data); err != nil {
		if err = ctx.BindJSON(&retData); err != nil {
			helper.Ajax(err.Error(), 1, ctx)
			return
		}
	}
	entity.MenuKey, _ = ctx.GetString("menu_key")
	err := simdbDriver.Open(&entity).Where("menu_key", "=", entity.MenuKey).First().AsEntity(&entity)
	if err != nil {
		helper.Ajax(err.Error(), 1, ctx)
		return
	}
	entity.Immutable = true

	subType, _ := ctx.GetString("sub_type")

	switch strings.ToLower(subType) {
	case "query":
		entity.Query = data
	case "param":
		entity.Param = data
	case "return":
		entity.Return = retData
	}

	err = simdbDriver.Open(&entity).Where("menu_key", "=", entity.MenuKey).Update(&entity)
	if err != nil {
		helper.Ajax(err.Error(), 1, ctx)
		return
	}

	helper.Ajax("更新成功", 0, ctx)
}

// 导出Api数据
func exportApiData(ctx *pine.Context) {

}
