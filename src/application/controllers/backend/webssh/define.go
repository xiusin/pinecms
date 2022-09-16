package webssh

import (
	"github.com/xiusin/pinecms/src/application/controllers/backend/webssh/tables"
	"sync"

	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/common/helper"
)

var once sync.Once

func InitInstall(app *pine.Application, urlPrefix, dir string) {
	once.Do(func() {
		app.Static(urlPrefix, dir, 2)
		orm := helper.GetORM()
		defer func() {
			if err := recover(); err != nil {
				pine.Logger().Warning("初始化安装失败", err)
			}
		}()

		if err := orm.Sync2(&tables.SSHServer{}, &tables.SSHUser{}); err != nil {
			pine.Logger().Warning(err)
		}
	})
}
