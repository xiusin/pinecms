package backend

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/logger"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
	"net/http"
)

type DistrictController struct {
	BaseController
}

func (c *DistrictController) Construct() {
	c.Table = &tables.District{}
	c.Entries = &[]tables.District{}
	c.AppId = "admin"
	c.Group = "字典管理"
	c.ApiEntityName = "地区"
	c.BaseController.Construct()
}

// PostImport 导入外部数据库
func (c *DistrictController) PostImport() {
	dbUrl := "https://github.com/eduosi/district/blob/master/district-full.sql"
	resp, err := http.Get(dbUrl)
	if err != nil {
		helper.Log2DB(logger.ErrorLevel, c.Ctx(), err)
		helper.Ajax(err.Error(), 1, c.Ctx())
		return
	}
	defer resp.Body.Close()
	c.Orm.Sync2(&tables.District{})
	_, err = c.Orm.Transaction(func(sess *xorm.Session) (interface{}, error) {
		_, err = c.Orm.Where("id > 0").Delete(c.Table)
		if err != nil {
			return nil, err
		}
		return c.Orm.Import(resp.Body)
	})

	if err != nil {
		helper.Log2DB(logger.ErrorLevel, c.Ctx(), err)
		helper.Ajax(err.Error(), 1, c.Ctx())
		return
	}
	helper.Ajax("导入数据库成功", 0, c.Ctx())
}
