package backend

import (
	"errors"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type DepartmentController struct {
	BaseController
}

func (c *DepartmentController) Construct() {
	c.Table = &tables.Department{}
	c.Entries = &[]tables.Department{}
	c.Group = "系统管理"
	c.SubGroup = "部门管理"
	c.ApiEntityName = "部门"
	c.BaseController.Construct()
	c.OpBefore = c.before
}

func (c *DepartmentController) before(act int, params interface{}) error {
	if act == OpEdit || act == OpAdd {
		p := params.(*tables.Department)
		sess := c.Orm.NewSession()
		if act == OpEdit {
			if p.Id < 1 {
				return errors.New("ID不能为空")
			}
			if uint(p.Id) == p.ParentId {
				return errors.New("父级部门不能为自己")
			} else {
				ids := []interface{}{p.Id}
				for len(ids) > 0 {
					var subs []tables.Department
					_ = c.Orm.In("parent_id", ids).Cols("id").Find(&subs)
					ids = ArrayCol(subs, "Id")
					for _, pid := range ids {
						if uint(pid.(int64)) == p.ParentId {
							return errors.New("不可设置自己下级作为父级部门")
						}
					}
				}
			}
			sess.Where("id <> ? and name = ?", p.Id, p.Name)
		} else {
			sess.Where("name = ?", p.Name)
		}
		if exist, _ := sess.Exist(&tables.Department{}); exist {
			return errors.New("部门已存在")
		}
	} else if act == OpDel {
		p := params.(*idParams)
		if exist, _ := c.Orm.In("parent_id", p.Ids).Exist(&tables.Department{}); exist {
			return errors.New("存在下级部门, 不可删除")
		}
	}
	return nil
}
