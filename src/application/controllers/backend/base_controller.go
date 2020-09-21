package backend

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/logger"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/common/helper"
)

var validate = validator.New()
var trans, _ = ut.New(zh.New(), en.New()).GetTranslator("zh")
var NotSupportErr = errors.New("暂不支持的传参类型")
var BindParseErr = errors.New("解析参数错误")

type KeywordWhere struct {
	Field    string                      // 字段
	Op       string                      // 操作字符	默认为 =
	DataExp  string                      // 数据匹配格式 匹配值=$? 如 LIKE %$?% 默认=$?
	CallBack func(*xorm.Session, string) // 替换callback 如果设置, 将绝对忽略Field Op DataExp的设置 匹配值=$?
}

type searchFieldDsl struct {
	Op string
}

type BaseController struct {
	SearchFields   map[string]searchFieldDsl // 设置可以搜索的字段
	BindType       string
	KeywordsSearch []KeywordWhere // 关键字搜索字段 用于关键字匹配字段
	Table          interface{}    // 传入Table结构体引用
	Entries        interface{}    // 传入Table结构体的切片
	Orm            *xorm.Engine
	pine.Controller
}

func (c *BaseController) BindParse() (err error) {
	switch c.BindType {
	case "json", "JSON":
		err = c.Ctx().BindJSON(c.Table)
	case "form", "FORM":
		err = c.Ctx().BindForm(c.Table)
	default:
		return NotSupportErr
	}
	if err != nil {
		return err
	}
	if err = validate.Struct(c.Table); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			return errors.New(e.Translate(trans))
		}
	}
	return nil
}

func (c *BaseController) GetList() {
	query := c.Orm.Table(c.Table)
	c.buildParamsForQuery(query)
	count, err := query.FindAndCount(c.Entries)
	if err != nil {
		logger.Error(err)
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
	if len(c.KeywordsSearch) > 0 && keywords != "" { // 关键字搜索
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
	if c.SearchFields != nil {
		for field, dsl := range c.SearchFields { // 其他字段搜索
			val := c.Ctx().GetString(field)
			if val != "" {
				if dsl.Op == "range" { // 范围查询， 组件between and
					ranges := strings.SplitN(val, ",", 2)
					if len(ranges) == 2 {
						query.Where(fmt.Sprintf("%s >= ?", field), ranges[0]).Where(fmt.Sprintf("%s <= ?", field), ranges[1])
					}
				} else {
					query.Where(fmt.Sprintf("%s %s ?", field, dsl.Op), val)
				}

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
	result, _ := c.Orm.InsertOne(c.Table)
	if result > 0 {
		helper.Ajax("新增数据成功", 0, c.Ctx())
	} else {
		helper.Ajax("新增数据失败", 1, c.Ctx())
	}
}

func (c *BaseController) PostEdit() {
	if err := c.BindParse(); err != nil {
		helper.Ajax(err.Error(), 1, c.Ctx())
		return
	}
	result, _ := c.Orm.AllCols().Update(c.Table)
	if result > 0 {
		helper.Ajax("修改数据成功", 0, c.Ctx())
	} else {
		helper.Ajax("修改数据失败", 1, c.Ctx())
	}
}

func (c *BaseController) PostOrder() {

}

func (c *BaseController) PostDelete() {
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
