package backend

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
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

const (
	BindTypeJson = "JSON"
	BindTypeForm = "FORM"
	OpList       = iota
	OpAdd
	OpEdit
	OpOrder
	OpDel
	OpInfo
)

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
	BindType       string                    // 表单绑定类型
	KeywordsSearch []KeywordWhere            // 关键字搜索字段 用于关键字匹配字段
	Table          interface{}               // 传入Table结构体引用
	Entries        interface{}               // 传入Table结构体的切片
	Orm            *xorm.Engine
	TableKey       string // 表主键
	TableStructKey string // 表结构体主键字段 主要用于更新逻辑反射数据

	OpBefore func(int, interface{}) error // 操作前置
	OpAfter  func(int, interface{}) error // 操作后置

	pine.Controller
}

func (c *BaseController) BindParse() (err error) {
	switch c.BindType {
	case BindTypeForm:
		err = c.Ctx().BindForm(c.Table)
	default:
		err = c.Ctx().BindJSON(c.Table)
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

func (c *BaseController) PostList() {
	query := c.Orm.Table(c.Table)
	if p, err := c.buildParamsForQuery(query); err != nil {
		helper.Ajax("参数错误: "+err.Error(), 1, c.Ctx())
		return
	} else {
		var count int64
		var err error
		if p.Size == 0 {
			err = query.Limit(p.Size, (p.Page-1)*p.Size).Find(c.Entries)
		} else {
			count, err = query.Limit(p.Size, (p.Page-1)*p.Size).FindAndCount(c.Entries)
		}
		if err != nil {
			logger.Error(err)
			helper.Ajax("获取列表异常: "+err.Error(), 1, c.Ctx())
			return
		}
		if c.OpAfter != nil {
			if err := c.OpAfter(OpList, &p); err != nil {
				helper.Ajax("获取列表异常: "+err.Error(), 1, c.Ctx())
			}
		}
		if p.Size == 0 {
			helper.Ajax(c.Entries, 0, c.Ctx())
		} else {
			helper.Ajax(pine.H{
				"list": c.Entries,
				"pagination": pine.H{
					"page":  p.Page,
					"size":  p.Size,
					"total": count,
				},
			}, 0, c.Ctx())
		}
	}
}

func (c *BaseController) buildParamsForQuery(query *xorm.Session) (*listParam, error) {
	var p listParam
	if err := parseParam(c.Ctx(), &p); err != nil {
		return nil, err
	}
	if len(c.KeywordsSearch) > 0 && p.Keywords != "" { // 关键字搜索
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
					strings.ReplaceAll(v.DataExp, "$?", p.Keywords),
				)
			} else {
				v.CallBack(query, p.Keywords)
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
	if len(p.OrderField) > 0 {
		if p.Sort == "desc" {
			query.Desc(p.OrderField)
		} else {
			query.Asc(p.OrderField)
		}
	}
	return &p, nil
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

func (c *BaseController) add() bool {
	result, _ := c.Orm.InsertOne(c.Table)
	return result > 0
}

func (c *BaseController) PostEdit() {
	if err := c.BindParse(); err != nil {
		helper.Ajax(err.Error(), 1, c.Ctx())
		return
	}
	if c.edit() {
		helper.Ajax("修改数据成功", 0, c.Ctx())
	} else {
		helper.Ajax("修改数据失败", 1, c.Ctx())
	}
}

func (c *BaseController) edit() bool {
	if len(c.TableStructKey) == 0 {
		c.TableStructKey = "Id"
	}
	val := reflect.ValueOf(c.Table).Elem().FieldByName(c.TableStructKey)
	if !val.IsValid()  {
		c.Logger().Error("无法匹配字段", c.TableStructKey)
		return false
	}
	result, _ := c.Orm.AllCols().Where(c.TableKey+"=?", val.Interface()).Update(c.Table)
	return result > 0
}

func (c *BaseController) PostOrder() {

}

func (c *BaseController) PostDelete() {
	var ids idParams
	if err := parseParam(c.Ctx(), &ids); err != nil {
		helper.Ajax("参数错误: "+err.Error(), 1, c.Ctx())
		return
	}
	_, err := c.Orm.Transaction(func(session *xorm.Session) (interface{}, error) {
		if c.OpBefore != nil {
			if err := c.OpBefore(OpDel, &ids); err != nil {
				return nil, err
			}
		}
		_, err := c.Orm.In(c.TableKey, ids.Ids).Delete(c.Table)
		if err != nil {
			return nil, err
		}
		if c.OpAfter != nil {
			if err := c.OpAfter(OpDel, &ids); err != nil {
				return nil, err
			}
		}
		return nil, nil
	})

	if err != nil {
		helper.Ajax("删除异常:"+err.Error(), 1, c.Ctx())
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
	if len(c.TableKey) == 0 {
		c.TableKey = "id"
	}
	exist, err := c.Orm.Where(c.TableKey+"=?", id).Get(c.Table)
	if err != nil {
		helper.Ajax("获取"+strconv.Itoa(id)+"信息失败: "+err.Error(), 1, c.Ctx())
	} else if !exist {
		helper.Ajax("获取详情信息失败", 1, c.Ctx())
	} else {
		if c.OpAfter != nil {
			if err := c.OpAfter(OpInfo, &idParams{Id: int64(id)}); err != nil {
				helper.Ajax(err, 1, c.Ctx())
			}
		}
		helper.Ajax(c.Table, 0, c.Ctx())
	}
}
