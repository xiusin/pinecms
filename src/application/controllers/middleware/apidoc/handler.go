package apidoc

import (
	"github.com/xiusin/pine"
)

func getConfig(ctx *pine.Context) {
	_ = ctx.WriteJSON(pine.H{"code": 0, "data": defaultConfig})
}

func getApiData(ctx *pine.Context) {
	appKey := ctx.GetString("appKey", "")
	if appKey == "" {
		_ = ctx.WriteJSON(pine.H{"code": 1, "msg": "参数错误"})
		return
	}
	var entities []apiEntity
	_ = simdbDriver.Open(&apiEntity{}).Where("app_id", "=", appKey).Get().AsEntity(&entities)
	groups := []apiGroup{{Title: "全部", Name: 0}}
	var lists []apiList
	for _, entity := range entities {
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
		listItem.Group = entity.Group.Title
		if len(entity.SubGroup) > 0 {
			listItem.MenuKey = entity.SubGroup
			listItem.Title = entity.SubGroup
		} else {
			listItem.Title = "未分组"
			listItem.MenuKey = "no_group"
		}
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
