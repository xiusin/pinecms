package models

import (
	"github.com/go-xorm/xorm"
	tables "github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/pine/di"
)

type PageModel struct {
	orm *xorm.Engine
}

func NewPageModel() *PageModel {
	return &PageModel{orm: di.MustGet("*xorm.Engine").(*xorm.Engine)}
}

func (p *PageModel) AddPage(page tables.IriscmsPage) bool {
	res, _ := p.orm.Insert(&page)
	if res != 0 {
		return true
	}
	return false
}

func (p *PageModel) UpdatePage(page tables.IriscmsPage) bool {
	res, _ := p.orm.Where("catid=?", page.Catid).Update(&page)
	if res != 0 {
		return true
	}
	return false
}

func (p *PageModel) DelPage(catid int64) bool {
	res, _ := p.orm.Delete(&tables.IriscmsPage{Catid: catid})
	if res != 0 {
		return true
	}
	return false
}

func (p *PageModel) GetPage(catid int64) tables.IriscmsPage {
	page := tables.IriscmsPage{Catid: catid}
	p.orm.Get(&page)
	return page
}
