package backend

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/go-playground/validator/v10"
	"xorm.io/xorm"

	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/controllers/middleware/apidoc"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/helper"
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
	Field    string                      // 字段
	Op       string                      // 操作字符	默认为 =
	DataExp  string                      // 数据匹配格式 匹配值=$? 如 LIKE %$?% 默认=$?
	SkipFn   func(any) bool              // 校验某些值不作为筛选条件如： 0， false不筛选状态
	CallBack func(*xorm.Session, ...any) // 替换callback 如果设置, 将绝对忽略Field Op DataExp的设置 匹配值=$?
}

type BaseController struct {
	LastErr error // 最后一次错误对象

	SearchFields   []SearchFieldDsl // 设置可以搜索的字段 接收或匹配params的字段
	BindType       uint             // 表单绑定类型
	KeywordsSearch []SearchFieldDsl // 关键字搜索字段 用于关键字匹配字段
	Table          any              // 传入Table结构体引用
	Entries        any              // 传入Table结构体的切片
	Orm            *xorm.Engine
	P              listParam
	apiEntities    map[string]apidoc.Entity

	ExceptCols []string // 排除数据表字段不读取
	Cols       []string // 仅读取指定表字段

	TableKey       string // 表主键
	TableStructKey string // 表结构体主键字段 主要用于更新逻辑反射数据

	OpBefore func(op int, v any) error // 操作前置
	OpAfter  func(op int, v any) error // 操作后置

	apidoc.Entity
	ApiEntityName string

	SelectOp    func(*xorm.Session)      // select下拉列表条件构建函数, 一般用于过滤指定值的数据
	UniqCheckOp func(int, *xorm.Session) // 唯一性校验构建SQL方法 一般用于删除或关停等校验是否有关联数据使用阻止关闭

	SelectListKV struct { // select 下拉列表kv字段设置 需要设置为struct属性
		Key   string
		Value string
	}

	pine.Controller
}

// Construct 默认初始化数据
func (c *BaseController) Construct() {
	c.TableKey = "id"
	c.TableStructKey = "Id"
	c.Orm = helper.GetORM()
	c.AppId = "admin"

	c.SelectListKV = struct {
		Key   string
		Value string
	}{Key: "Id", Value: "Name"}

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

func (c *BaseController) IsOperate(op int) bool {
	return slices.Contains([]int{OpAdd, OpEdit}, op)
}

func (c *BaseController) BindParse(receivers ...any) (err error) {
	receiver := c.Table
	if len(receivers) > 0 {
		receiver = receivers[0]
	}
	switch c.BindType {
	case BindTypeForm:
		err = c.Ctx().BindForm(receiver)
	default:
		err = c.Ctx().BindJSON(receiver)
	}
	if err != nil {
		return err
	}
	if err = validate.Struct(receiver); err != nil {
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
				helper.Ajax(err.Error(), 1, c.Ctx())
			}
		}

		c.buildQueryCols(query)

		if p.Size == 0 {
			err = query.Find(c.Entries)
		} else {
			count, err = query.Limit(p.Size, (p.Page-1)*p.Size).FindAndCount(c.Entries)
		}
		if err != nil {
			pine.Logger().Error(err.Error())
			helper.Ajax(err.Error(), 1, c.Ctx())
			return
		}
		if c.OpAfter != nil {
			if err := c.OpAfter(OpList, &p); err != nil {
				helper.Ajax(err.Error(), 1, c.Ctx())
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
	if err := parseParam(c.Ctx(), &c.P); err != nil {
		pine.Logger().Warn("解析参数错误", err)
	}
	if len(c.KeywordsSearch) > 0 && c.P.Keywords != "" { // 关键字搜索
		var whereBuilder []string
		var whereLikeBind []any
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
				whereLikeBind = append(whereLikeBind, strings.ReplaceAll(v.DataExp, "$?", c.P.Keywords))
			} else {
				v.CallBack(query, c.P.Keywords)
			}
		}
		wl := len(whereBuilder)
		if wl == len(whereLikeBind) && wl != 0 {
			query.Where(strings.Join(whereBuilder, " OR "), whereLikeBind...)
		}
	}
	if c.SearchFields != nil && len(c.P.Params) > 0 {
		for _, v := range c.SearchFields {
			if v.Field == "" {
				continue
			}

			if val, exists := c.P.Params[strings.ReplaceAll(v.Field, "`", "")]; exists {
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
					case []any:
						v.CallBack(query, val...)
					default:
						v.CallBack(query, val)
					}
				}
			}
		}
	}
	if len(c.P.OrderField) > 0 {
		if c.P.Sort == "desc" {
			query.Desc(c.P.OrderField)
		} else {
			query.Asc(c.P.OrderField)
		}
	}
	return &c.P, nil
}

func (c *BaseController) PostAdd() {

	if c.UniqCheckOp != nil {
		sess := c.Orm.NewSession()
		defer sess.Close()

		c.UniqCheckOp(OpAdd, sess)
		if exist, _ := sess.Exist(c.Table); exist {
			helper.Ajax("有重复记录, 新增失败", 1, c.Ctx())
			return
		}
	}

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
		helper.Ajax(err.Error(), 1, c.Ctx())
		return

	}

	if c.OpAfter != nil {
		if err := c.OpAfter(OpAdd, c.Table); err != nil {
			helper.Ajax(err.Error(), 1, c.Ctx())
			return
		}
	}
	helper.Ajax(c.Table, 0, c.Ctx())
}

func (c *BaseController) add() error {
	_, err := c.Orm.InsertOne(c.Table)
	return err
}

func (c *BaseController) buildQueryCols(sess *xorm.Session) {
	if len(c.Cols) > 0 {
		sess.Cols(c.Cols...)
	} else if len(c.ExceptCols) > 0 {
		tableName := c.Orm.TableName(c.Table)
		var fields = []string{}
		if err := helper.Cache().Remember(fmt.Sprintf(controllers.CacheTableNameFields, tableName), &fields, func() (any, error) {
			concat, err := c.Orm.QueryString(`select group_concat(COLUMN_NAME SEPARATOR ',') as fields from information_schema.COLUMNS where table_name = '` + tableName + `'`)
			if err != nil {
				c.Ctx().Logger().Warn("读取数据表"+tableName+"表字段失败", err)
			}
			fields := strings.Split(concat[0]["fields"], ",")
			return &fields, nil
		}, 3600); err != nil {
			c.Ctx().Logger().Warn("缓存表"+tableName+"字段失败", err)
		}

		cols := []string{}

		for _, s := range fields {
			if !slices.Contains(c.ExceptCols, s) {
				cols = append(cols, s)
			}
		}
		sess.Cols(cols...)
	}
}

func (c *BaseController) PostUpdate() {
	c.PostEdit() // TODO 需要延迟调用, 否则无法实现重现 ?
}

func (c *BaseController) PostEdit() {

	if c.UniqCheckOp != nil {
		sess := c.Orm.NewSession()
		defer sess.Close()
		c.UniqCheckOp(OpEdit, sess)
		if exist, _ := sess.Exist(c.Table); exist {
			helper.Ajax("有重复记录, 修改失败", 1, c.Ctx())
			return
		}
	}

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
		c.LastErr = errors.New("无法匹配字段")
		c.Logger().Error("无法匹配字段", c.TableStructKey)
		return false
	}
	_, err := c.Orm.AllCols().Where(c.TableKey+"=?", val.Interface()).Update(c.Table)
	if err != nil {
		c.LastErr = err
		c.Logger().Warn("更新错误", err.Error())
	}
	return true
}

func (c *BaseController) PostDelete() {
	var ids idParams
	if err := parseParam(c.Ctx(), &ids); err != nil {
		helper.Ajax("参数错误: "+err.Error(), 1, c.Ctx())
		return
	}
	_, err := c.Orm.Transaction(func(session *xorm.Session) (any, error) {
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
	id, _ := c.Ctx().Input().GetInt("id", 0)
	if id < 1 {
		helper.Ajax("id参数范围错误", 1, c.Ctx())
		return
	}
	if len(c.TableKey) == 0 {
		c.TableKey = "id"
	}
	query := c.Orm.Where(c.TableKey+"=?", id)

	c.buildQueryCols(query)

	exist, err := query.Get(c.Table)
	if err != nil {
		helper.Ajax(err.Error(), 1, c.Ctx())
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
			apiEntity.ApiParam = &c.P
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

func (c *BaseController) GetSelect() {
	var kv = []tables.KV{}

	if c.Entries != nil {
		sess := c.Orm.NewSession()
		defer sess.Close()
		if c.SelectOp != nil {
			c.SelectOp(sess)
		}

		sess.Find(c.Entries)
		// 反射类型
		val := reflect.ValueOf(c.Entries)
		length := val.Elem().Len()
		for i := 0; i < length; i++ {
			item := val.Elem().Index(i)
			kv = append(kv, tables.KV{
				Value: item.FieldByName(c.SelectListKV.Key).Interface(),
				Label: item.FieldByName(c.SelectListKV.Value).String(),
			})
		}
	}

	helper.Ajax(kv, 0, c.Ctx())
}
