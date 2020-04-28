# 模型

pinecms采用`xorm`作为内置orm引擎. 所以您需要先熟悉`xorm`库的相关用法. 

xrom文档地址: [XORM DOC](http://gobook.io/read/gitea.com/xorm/manual-zh-CN/)


# 模型定义
在目录`src/application/models`可以创建模型:
```go
package models

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type XXModel struct {
	orm *xorm.Engine
    cache cache.AbstractCache
}
func init()  {  // 非必须, 如果需要共享实例,您可以注册
	di.Set(&XXModel{}, func(builder di.AbstractBuilder) (i interface{}, err error) {
		return &XXModel{
			orm:   builder.MustGet("*xorm.Engine").(*xorm.Engine),
			cache: builder.MustGet("cache.AbstractCache").(cache.AbstractCache),
		}, nil
	}, true)
}
func NewXXModel() *XXModel {
	return &XXModel{orm: di.MustGet("*xorm.Engine").(*xorm.Engine)}

    // 也可以直接获取共享实例(这里不确定那种方式更好, todo 记录为性能测试)
    return di.MustGet(&XXModel{}).(*XXModel)

}

func (l *XXModel) GetList(page, limit int64) ([]tables.Advert, int64) {
	offset := (page - 1) * limit
	var list []tables.XX
	var total int64
	var err error
	if total, err = l.orm.Desc("listorder").Limit(int(limit), int(offset)).FindAndCount(&list); err != nil {
		pine.Logger().Error(err)
	}
	return list, total
}
```

一般的每个模型定于都对应一个表定义: 可以到`models/tables`下定义表字段映射关系. 

