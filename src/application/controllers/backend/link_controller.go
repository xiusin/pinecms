package backend

import (
	"fmt"
	"github.com/xiusin/pine"
	"strconv"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
)

type LinkController struct {
	pine.Controller
}

func (c *LinkController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY("/link/list", "List")
	b.POST("/link/add", "Add")
	b.POST("/link/edit", "Edit")
	b.POST("/link/delete", "Delete")
	b.ANY("/link/order", "Order")
}

func (c *LinkController) List() {
	page, _ := c.Ctx().GetInt64("page")
	rows, _ := c.Ctx().GetInt64("rows")
	list, total := models.NewLinkModel().GetList(page, rows)
	helper.Ajax(map[string]interface{}{"rows": list, "total": total}, 0, c.Ctx())
}

func (c *LinkController) Add() {
	var d tables.Link
	if err := c.Ctx().BindForm(&d); err != nil {
		helper.Ajax("参数错误："+err.Error(), 1, c.Ctx())
		return
	}
	d.Listorder = 30
	d.Addtime = time.Now().In(helper.GetLocation())
	if models.NewLinkModel().Add(&d) > 0 {
		helper.Ajax("添加友链成功", 0, c.Ctx())
	} else {
		helper.Ajax("添加友链失败", 1, c.Ctx())
	}

}

func (c *LinkController) Edit() {
	var d tables.Link
	if err := c.Ctx().BindForm(&d); err != nil || d.Linkid < 1 {
		helper.Ajax(fmt.Sprintf("参数错误:%s", err), 1, c.Ctx())
		return
	}
	d.Addtime = time.Now().In(helper.GetLocation())
	if models.NewLinkModel().Update(&d) {
		helper.Ajax("修改友链成功", 0, c.Ctx())
	} else {
		helper.Ajax("修改友链失败", 1, c.Ctx())
	}
}

func (c *LinkController) Delete() {
	linkId, _ := c.Ctx().GetInt64("id")
	if linkId < 1 {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	if models.NewLinkModel().Delete(linkId) {
		helper.Ajax("删除链接成功", 0, c.Ctx())
	} else {
		helper.Ajax("删除链接失败", 1, c.Ctx())
	}
}

func (c *LinkController) Order() {
	data := c.Ctx().PostData()
	order, ok := data["order"]
	if !ok {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	id, ok := data["id"]
	if !ok {
		helper.Ajax("参数错误", 1, c.Ctx())
		return
	}
	for i := 0; i < len(order); i++ {
		linkid, err := strconv.Atoi(id[i])
		if err != nil {
			continue
		}
		orderNum, err := strconv.Atoi(order[i])
		if err != nil {
			continue
		}
		c.Ctx().Value("orm").(*xorm.Engine).ID(linkid).MustCols("listorder").Update(&tables.Link{Listorder: int64(orderNum)})
	}
	helper.Ajax("更新排序值成功", 0, c.Ctx())
}
