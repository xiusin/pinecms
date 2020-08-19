package backend

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/common/helper"
	"reflect"
)

var validate = validator.New()
var NotSupportErr = errors.New("暂不支持的传参类型")
var BindParseErr = errors.New("解析参数错误")

type BaseController struct {
	BindType string
	Table    interface{} // 传入Table结构体引用
	Entries  interface{} // 传入Table结构体的切片
	Orm      *xorm.Engine
	pine.Controller
}

func (c *BaseController) bindParse() error {
	var err error
	switch c.BindType {
	case "json", "JSON":
		err = c.Ctx().BindJSON(c.Ctx())
	case "form", "FORM":
		err = c.Ctx().BindForm(c.Ctx())
	default:
		return NotSupportErr
	}
	if err != nil {
		pine.Logger().Error(err)
	}
	// 验证结构体
	if err := validate.Struct(c.Table); err != nil {
		return err
	}
	return BindParseErr
}

func (c *BaseController) GetList() {
	fmt.Println(reflect.ValueOf(c.Table).Elem())
	return
	list := reflect.MakeSlice(reflect.ValueOf(c.Table).Elem().Type(), 0, 0)
	fmt.Println(list)
}

func (c *BaseController) PostAdd() {
	if err := c.bindParse(); err != nil {
		helper.Ajax(err.Error(), 1, c.Ctx())
		return
	}

}

func (c *BaseController) PostEdit() {
	if err := c.bindParse(); err != nil {
		helper.Ajax(err.Error(), 1, c.Ctx())
		return
	}
}

func (c *BaseController) PostOrder() {

}

func (c *BaseController) GetDelete() {
	id, _ := c.Ctx().GetInt("id", 0)
	if id < 1 {
		helper.Ajax("id参数范围错误", 1, c.Ctx())
		return
	}
	count, err := c.Orm.Where("id = ?", id).Delete(c.Table)
	if err != nil {
		pine.Logger().Error(err)
		helper.Ajax("删除异常", 1, c.Ctx())
		return
	}
	if count == 0 {
		helper.Ajax("删除失败", 1, c.Ctx())
		return
	}
	helper.Ajax("删除成功", 0, c.Ctx())
}

func (c *BaseController) GetInfo() {
	id, _ := c.Ctx().GetInt("id", 0)
	if id < 1 {
		helper.Ajax("id参数范围错误", 1, c.Ctx())
		return
	}
	has, err := c.Orm.Where("id = ?", id).Get(c.Table)
	if err != nil || !has {
		helper.Ajax("获取信息失败", 1, c.Ctx())
	} else {
		helper.Ajax(c.Table, 0, c.Ctx())
	}
}
