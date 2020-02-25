package models

import (
	"github.com/go-xorm/xorm"
	tables "github.com/xiusin/iriscms/src/application/models/tables"
)

type PageModel struct {
	Orm *xorm.Engine
}

func NewPageModel(orm *xorm.Engine) *PageModel {
	return &PageModel{Orm: orm}
}

func (p *PageModel) AddPage(page tables.IriscmsPage) bool {
	res, _ := p.Orm.Insert(&page)
	if res != 0 {
		return true
	}
	return false
}

func (p *PageModel) UpdatePage(page tables.IriscmsPage) bool {
	res, _ := p.Orm.Where("catid=?", page.Catid).Update(&page)
	if res != 0 {
		return true
	}
	return false
}

func (p *PageModel) DelPage(catid int64) bool {
	res, _ := p.Orm.Delete(&tables.IriscmsPage{Catid: catid})
	if res != 0 {
		return true
	}
	return false
}

func (p *PageModel) GetPage(catid int64) tables.IriscmsPage {
	page := tables.IriscmsPage{Catid: catid}
	p.Orm.Get(&page)
	return page
}
