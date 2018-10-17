package models

import (
	"github.com/go-xorm/xorm"
	"iriscms/application/models/tables"
)

type PageModel struct {
	Orm *xorm.Engine
}

func NewPageModel(orm *xorm.Engine) *PageModel {
	return &PageModel{Orm: orm}
}

func (this *PageModel) AddPage(page tables.IriscmsPage) bool {
	res, _ := this.Orm.Insert(&page)
	if res != 0 {
		return true
	}
	return false
}

func (this *PageModel) UpdatePage(page tables.IriscmsPage) bool {
	res, _ := this.Orm.Where("catid=?", page.Catid).Update(&page)
	if res != 0 {
		return true
	}
	return false
}

func (this *PageModel) DelPage(catid int64) bool {
	res, _ := this.Orm.Delete(&tables.IriscmsPage{Catid: catid})
	if res != 0 {
		return true
	}
	return false
}

func (this *PageModel) GetPage(catid int64) tables.IriscmsPage {
	page := tables.IriscmsPage{Catid: catid}
	this.Orm.Get(&page)
	return page
}
