package backend

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers/middleware/apidoc"
	"github.com/xiusin/pinecms/src/common/helper"
	"xorm.io/xorm"
)

var validate = validator.New()

const (
	BindTypeJson = iota
	BindTypeForm

	OpList = iota
	OpAdd
	OpEdit
	OpDel
	OpInfo
)

type SearchFieldDsl struct {
	Field    string                              // 字段
	Op       string                              // 操作字符	默认为 =
	DataExp  string                              // 数据匹配格式 匹配值=$? 如 LIKE %$?% 默认=$?
	SkipFn   func(interface{}) bool              // 校验某些值不作为筛选条件如： 0， false不筛选状态
	CallBack func(*xorm.Session, ...interface{}) // 替换callback 如果设置, 将绝对忽略Field Op DataExp的设置 匹配值=$?
}

type BaseController struct {
	LastErr error // 最后一次错误对象

	SearchFields   []SearchFieldDsl // 设置可以搜索的字段 接收或匹配params的字段
	BindType       uint             // 表单绑定类型
	KeywordsSearch []SearchFieldDsl // 关键字搜索字段 用于关键字匹配字段
	Table          interface{}      // 传入Table结构体引用
	Entries        interface{}      // 传入Table结构体的切片
	Orm            *xorm.Engine
	p              listParam
	apiEntities    map[string]apidoc.Entity

	TableKey       string // 表主键
	TableStructKey string // 表结构体主键字段 主要用于更新逻辑反射数据

	OpBefore func(int, interface{}) error // 操作前置
	OpAfter  func(int, interface{}) error // 操作后置

	apidoc.Entity
	ApiEntityName string

	pine.Controller
}

//Construct 默认初始化数据
func (c *BaseController) Construct() {
	c.TableKey = "id"
	c.TableStructKey = "Id"
	c.Orm = helper.GetORM()
	c.AppId = "admin"
	if c.apiEntities == nil {
		c.apiEntities = map[string]apidoc.Entity{ // 内置一个模板, 配合ApiEntityName使用
			"__inner": {}, // 重写时不要附加此参数
			"list":    {Title: "%s列表", Desc: "获取系统未删除的%s列表"},
			"add":     {Title: "新增%s", Desc: "新增一个新的%s"},
			"edit":    {Title: "编辑%s", Desc: "修改给定ID的%s信息"},
			"delete":  {Title: "删除%s", Desc: "删除指定%s"},
			"info":    {Title: "%s详情", Desc: "获取指定%s的详情信息"},
		}
	}
	c.setApiEntity()
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
		return errs[0]
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
		if p.Export {
			p.Size = 0
		}

		if c.OpBefore != nil {
			if err := c.OpBefore(OpList, query); err != nil {
				helper.Ajax("获取列表异常: "+err.Error(), 1, c.Ctx())
			}
		}

		if p.Size == 0 {
			err = query.Find(c.Entries)
		} else {
			count, err = query.Limit(p.Size, (p.Page-1)*p.Size).FindAndCount(c.Entries)
		}
		if err != nil {
			pine.Logger().Error(err)
			helper.Ajax("获取列表异常: "+err.Error(), 1, c.Ctx())
			return
		}
		if c.OpAfter != nil {
			if err := c.OpAfter(OpList, &p); err != nil {
				helper.Ajax("获取列表异常: "+err.Error(), 1, c.Ctx())
			}
		}

		if c.Entries == nil {
			c.Entries = []struct{}{}
		}

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

func (c *BaseController) buildParamsForQuery(query *xorm.Session) (*listParam, error) {
	if err := parseParam(c.Ctx(), &c.p); err != nil {
		pine.Logger().Warning("解析参数错误", err)
	}
	if len(c.KeywordsSearch) > 0 && c.p.Keywords != "" { // 关键字搜索
		var whereBuilder []string
		var whereLikeBind []interface{}
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
				whereBuilder = append(whereBuilder, fmt.Sprintf("%s %s ?", v.Field, v.Op))
				whereLikeBind = append(whereLikeBind, strings.ReplaceAll(v.DataExp, "$?", c.p.Keywords))
			} else {
				v.CallBack(query, c.p.Keywords)
			}
		}
		wl := len(whereBuilder)
		if wl == len(whereLikeBind) && wl != 0 {
			query.Where(strings.Join(whereBuilder, " OR "), whereLikeBind...)
		}
	}
	if c.SearchFields != nil && len(c.p.Params) > 0 {
		for _, v := range c.SearchFields {
			if v.Field == "" {
				continue
			}

			if val, exists := c.p.Params[strings.ReplaceAll(v.Field, "`", "")]; exists {
				if v.SkipFn != nil && v.SkipFn(val) {
					continue
				}
				switch val := val.(type) {
				case string:
					if len(val) == 0 {
						continue
					}
				case bool:
					if !val {
						continue
					}
				}
				if v.CallBack == nil {
					if v.Op == "" {
						v.Op = "="
					}
					if v.DataExp == "" {
						v.DataExp = "$?"
					}
					if v.Op == "LIKE" {
						val = strings.ReplaceAll(v.DataExp, "$?", val.(string))
					}
					query.Where(fmt.Sprintf("%s %s ?", v.Field, v.Op), val)
				} else {
					switch val := val.(type) {
					case []interface{}:
						v.CallBack(query, val...)
					default:
						v.CallBack(query, val)
					}
				}
			}
		}
	}
	if len(c.p.OrderField) > 0 {
		if c.p.Sort == "desc" {
			query.Desc(c.p.OrderField)
		} else {
			query.Asc(c.p.OrderField)
		}
	}
	return &c.p, nil
}

func (c *BaseController) PostAdd() {
	if err := c.BindParse(); err != nil {
		helper.Ajax(err.Error(), 1, c.Ctx())
		return
	}
	if c.OpBefore != nil {
		if err := c.OpBefore(OpAdd, c.Table); err != nil {
			helper.Ajax(err.Error(), 1, c.Ctx())
			return
		}
	}
	if err := c.add(); err != nil {
		if c.OpAfter != nil {
			if err := c.OpAfter(OpAdd, c.Table); err != nil {
				helper.Ajax(err.Error(), 1, c.Ctx())
				return
			}
		}
		if len(c.Ctx().Response.Body()) == 0 {
			helper.Ajax(c.Table, 0, c.Ctx())
		}
	} else {
		helper.Ajax(err, 1, c.Ctx())
	}
}

func (c *BaseController) add() error {
	_, err := c.Orm.InsertOne(c.Table)
	return err
}

func (c *BaseController) PostUpdate() {
	c.PostEdit() // TODO 需要延迟调用, 否则无法实现重现 ?
}

func (c *BaseController) PostEdit() {
	if err := c.BindParse(); err != nil {
		helper.Ajax(err.Error(), 1, c.Ctx())
		return
	}

	if c.OpBefore != nil {
		if err := c.OpBefore(OpEdit, c.Table); err != nil {
			helper.Ajax(err.Error(), 1, c.Ctx())
			return
		}
	}
	if c.edit() {
		if c.OpAfter != nil {
			if err := c.OpAfter(OpEdit, c.Table); err != nil {
				helper.Ajax(err.Error(), 1, c.Ctx())
				return
			}
		}
		if len(c.Ctx().Response.Body()) == 0 {
			helper.Ajax("修改数据成功", 0, c.Ctx())
		}
	} else if c.LastErr != nil {
		helper.Ajax(c.LastErr.Error(), 1, c.Ctx())
	} else {
		helper.Ajax("修改数据失败", 1, c.Ctx())
	}
}

func (c *BaseController) edit() bool {
	if len(c.TableStructKey) == 0 {
		c.TableStructKey = "Id"
	}
	val := reflect.ValueOf(c.Table).Elem().FieldByName(c.TableStructKey)
	if !val.IsValid() {
		c.Logger().Error("无法匹配字段", c.TableStructKey)
		return false
	}
	result, _ := c.Orm.AllCols().Where(c.TableKey+"=?", val.Interface()).Update(c.Table)
	return result > 0
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
		helper.Ajax(err.Error(), 1, c.Ctx())
		return
	}
	if len(c.Ctx().Response.Body()) == 0 {
		helper.Ajax("删除成功", 0, c.Ctx())
	}
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
		if len(c.Ctx().Response.Body()) == 0 {
			helper.Ajax(c.Table, 0, c.Ctx())
		}
	}
}

func (c *BaseController) setApiEntity() {
	ps := strings.Split(c.Ctx().Path(), "/")
	key := ps[len(ps)-1]

	apiEntity, exist := c.apiEntities[key]
	if len(c.apiEntities) == 0 || !exist {
		return
	}

	if _, exist = c.apiEntities["__inner"]; exist && len(c.ApiEntityName) > 0 {
		apiEntity.Title = fmt.Sprintf(apiEntity.Title, c.ApiEntityName)
		apiEntity.Desc = fmt.Sprintf(apiEntity.Desc, c.ApiEntityName)
	}

	if apiEntity.ApiParam == nil {
		switch key {
		case "list":
			apiEntity.ApiParam = &c.p
		case "edit", "add":
			apiEntity.ApiParam = c.Table
		case "delete":
			apiEntity.ApiParam = &idParams{}
		}
	}
	if len(apiEntity.AppId) == 0 {
		apiEntity.AppId = c.AppId
	}
	if len(apiEntity.Group) == 0 {
		apiEntity.Group = c.Group
	}
	if len(apiEntity.SubGroup) == 0 {
		if len(c.SubGroup) > 0 {
			apiEntity.SubGroup = c.SubGroup
		} else {
			apiEntity.SubGroup = c.Group
		}
	}
	apidoc.SetApiEntity(c.Ctx(), &apiEntity)
}
