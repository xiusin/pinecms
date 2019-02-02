package api

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iriscms/application/models"
	"iriscms/application/models/tables"
	"iriscms/common/helper"
	"strconv"
)

type ContentController struct {
	Orm *xorm.Engine
	Ctx iris.Context
}

func (c *ContentController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodGet, "/RngNb/list", "ContentList")
	b.Handle(iris.MethodGet, "/content/info/:id", "ContentInfo")
	b.Handle(iris.MethodPost, "/content/pay", "ContentPay")
}

func (c *ContentController) ContentList() {
	catid, _ := c.Ctx.URLParamInt("catid")
	pageNo, _ := c.Ctx.URLParamInt("pageNo")
	pageSize, _ := c.Ctx.URLParamInt("pageSize")
	six := c.Ctx.URLParam("six")
	sort := c.Ctx.URLParam("sort")
	var arts []tables.IriscmsContent
	var offset = (pageNo - 1) * pageSize

	cats := models.NewCategoryModel(c.Orm).GetNextCategory(int64(catid))
	var catID []interface{}
	catID = append(catID, int64(catid))
	for _, v := range cats {
		catID = append(catID, v.Catid)
	}
	q := c.Orm.Where("deleted_at = 0 and status = 1").In("catid", catID...).Limit(pageSize, offset)
	if sort == "desc" {
		q.Desc(six)
	} else {
		q.Asc(six)
	}
	q.Find(&arts)
	if len(arts) == 0 {
		arts = []tables.IriscmsContent{}
	}
	c.Ctx.JSON(ReturnApiData{Status: true, Msg: "成功", Data: arts})
}

func (c *ContentController) ContentInfo() {
	id, err := c.Ctx.Params().GetInt64("id")
	if err != nil {
		c.Ctx.JSON(ReturnApiData{Status: false, Msg: "资源不存在", Data: nil})
		return
	}
	var content tables.IriscmsContent
	ok, _ := c.Orm.Id(id).Get(&content)
	if !ok {
		c.Ctx.JSON(ReturnApiData{Status: false, Msg: "资源不存在", Data: nil})
		return
	}

	if content.DeletedAt > 0 || content.Status != 1 {
		c.Ctx.JSON(ReturnApiData{Status: false, Msg: "资源不存在", Data: nil})
		return
	}
	c.Ctx.JSON(ReturnApiData{Status: true, Msg: "获取资源成功", Data: content})
}

func (c *ContentController) ContentPay() {
	data := map[string]string{}
	err := c.Ctx.ReadJSON(&data)
	if err != nil {
		c.Ctx.JSON(ReturnApiData{Status: false, Msg: "接口异常", Data: nil})
		return
	}

	var content tables.IriscmsContent
	ok, _ := c.Orm.Id(data["id"]).Get(&content)
	if !ok {
		c.Ctx.JSON(ReturnApiData{Status: false, Msg: "资源不存在", Data: nil})
		return
	}

	if content.DeletedAt > 0 || content.Status != 1 {
		c.Ctx.JSON(ReturnApiData{Status: false, Msg: "资源不存在", Data: nil})
		return
	}

	if content.Money == 0 || content.PwdType == 1 {
		c.Ctx.JSON(ReturnApiData{Status: false, Msg: "资源已免费无需支付,请使用公众号获取资源密码", Data: nil})
		return
	}

	//todo 检查是否存在未支付的订单
	var d = map[string]interface{}{}
	orderId := "U" + "LOGINID" + "A" + strconv.Itoa(int(content.Id)) + "T" + strconv.Itoa(helper.GetTimeStamp())
	fee := strconv.Itoa(int(content.Money))
	r, err := helper.Pay(c.Ctx, data["paytype"], content.Title, orderId, fee)
	if err != nil {
		c.Ctx.JSON(ReturnApiData{Status: false, Msg: "获取支付链接失败" + err.Error()})
	} else {
		err = r.ToJSON(&d)
		fmt.Println(d)
		if err != nil {
			c.Ctx.JSON(ReturnApiData{Status: false, Msg: "获取支付链接失败" + err.Error()})
		} else {
			d["order_id"] = orderId
			d["title"] = content.Title //未支付的订单直接返还给客户
			c.Ctx.JSON(ReturnApiData{Status: true, Msg: "获取支付链接成功", Data: d})
		}
	}
}
