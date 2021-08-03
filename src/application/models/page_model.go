package models

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type PageModel struct {
	orm *xorm.Engine
}

func NewPageModel() *PageModel {
	return &PageModel{orm: di.MustGet("*xorm.Engine").(*xorm.Engine)}
}

func (p *PageModel) AddPage(page *tables.Page) bool {
	res, _ := p.orm.Insert(page)
	if res != 0 {
		return true
	}
	return false
}

func (p *PageModel) UpdatePage(page *tables.Page) bool {
	res, _ := p.orm.Where("catid=?", page.Id).Update(page)
	if res != 0 {
		return true
	}
	return false
}

func (p *PageModel) DelPage(catid int64) bool {
	res, _ := p.orm.Delete(&tables.Page{Id: catid})
	if res != 0 {
		return true
	}
	return false
}

func (p *PageModel) GetPage(catid int64) *tables.Page {
	page := &tables.Page{Id: catid}
	exists, err := p.orm.Get(page)
	if err != nil {
		pine.Logger().Error("获取page信息失败:", err)
	}
	if exists {
		return page
	}
	return nil
}
