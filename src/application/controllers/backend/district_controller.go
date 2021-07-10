package backend

import (
	"github.com/xiusin/pinecms/src/application/models/tables"
	"net/http"
)

type DistrictController struct {
	BaseController
}

func (c *DistrictController) Construct() {
	c.Table = &tables.District{}
	c.Entries = &[]*tables.District{}
	c.AppId = "admin"
	c.Group = "字典管理"
	c.SubGroup = "地区管理"
	c.ApiEntityName = "地区"
	c.BaseController.Construct()
}

// PostImport 导入外部数据库
func (c *DistrictController) PostImport() {
	dbUrl := "https://github.com/eduosi/district/blob/master/district-full.sql"
	resp, err := http.Get(dbUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	_, err = c.Orm.Import(resp.Body)
	if err != nil {
		panic(err)
	}
}
