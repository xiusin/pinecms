package backend

import (
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type TestController struct {
	BaseController
}

//todo 通过前置控制器方法处理

func (t *TestController) Construct() {
	t.BindType = "json" // 设置绑定表单类型
	t.Orm = pine.Make(controllers.ServiceXorm).(*xorm.Engine)
	t.Table = &tables.Advert{}
	t.Entries = []tables.Advert{}
}
