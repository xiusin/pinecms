package backend

import (
	"fmt"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type SettingController struct {
	BaseController
}


func (c *SettingController) Construct() {
	c.Table = &tables.Setting{}
	c.Entries = &[]*tables.Setting{}
	c.BaseController.Construct()
}

// PostGroupList 获取新分组
func (c *LogController) PostGroupList() {
	var groups []string
	c.Orm.GroupBy("group").Cols("group").Find(&groups)
	fmt.Println("groups", groups)
}
