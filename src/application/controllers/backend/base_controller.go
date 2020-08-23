package backend

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/common/helper"
	"strings"
)

var validate = validator.New()
var NotSupportErr = errors.New("暂不支持的传参类型")
var BindParseErr = errors.New("解析参数错误")

type KeywordWhere struct {
	Field    string                      // 字段
	Op       string                      // 操作字符	默认为 =
	DataExp  string                      // 数据匹配格式 匹配值=$? 如 LIKE %$?% 默认=$?
	CallBack func(*xorm.Session, string) // 替换callback 如果设置, 将绝对忽略Field Op DataExp的设置 匹配值=$?
}

type BaseController struct {
	//SearchFields []string	// 设置可以搜索的字段
	BindType       string
	KeywordsSearch []KeywordWhere // 关键字搜索字段 用于关键字匹配字段
	Table          interface{}    // 传入Table结构体引用
	Entries        interface{}    // 传入Table结构体的切片
	Orm            *xorm.Engine
	pine.Controller
}

func (c *BaseController) BindParse() error {
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
	if err := validate.Struct(c.Table); err != nil {
		return err
	}
	return BindParseErr
}

func (c *BaseController) GetList() {
	query := c.Orm.Table(c.Table)
	c.buildParamsForQuery(query)
	count, err := query.FindAndCount(c.Entries)
	if err != nil {
		c.Logger().Error(err)
		helper.Ajax("读取数据列表错误", 1, c.Ctx())
		return
	}
	helper.Ajax(pine.H{"total": count, "list": c.Entries}, 0, c.Ctx())
}

func (c *BaseController) buildParamsForQuery(query *xorm.Session) {
	page, _ := c.Ctx().GetInt64("page", 1)
	if page < 1 {
		page = 1
	}
	prePage, _ := c.Ctx().GetInt64("prePage", 10)
	if prePage < 1 {
		prePage = 10
	}
	orderBy := c.Ctx().GetString("orderBy", "id")
	orderDir := c.Ctx().GetString("orderDir", "desc")
	keywords := c.Ctx().GetString("keywords", "")
	if len(c.KeywordsSearch) > 0 && keywords != "" {
		for _, v := range c.KeywordsSearch {
			if v.Field == "" && v.CallBack == nil {
				continue
			}
			if v.CallBack == nil {
				if v.Op == "" {
					v.Op = "="
				}
				if v.DataExp == "" {
					v.DataExp = "$?"
				}
				query.Where(
					fmt.Sprintf("%s %s ?", v.Field, v.Op),
					strings.ReplaceAll(v.DataExp, "$?", keywords),
				)
			} else {
				v.CallBack(query, keywords)
			}
		}
	}
	query.OrderBy(fmt.Sprintf("%s %s", orderBy, orderDir))
}

func (c *BaseController) PostAdd() {
	if err := c.BindParse(); err != nil {
		helper.Ajax(err.Error(), 1, c.Ctx())
		return
	}

}

func (c *BaseController) PostEdit() {
	if err := c.BindParse(); err != nil {
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