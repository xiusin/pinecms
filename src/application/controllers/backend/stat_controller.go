package backend

import (
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/common/helper"
)

type StatController struct {
	pine.Controller
}

func (_ *StatController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY("/stat/data", "Data")
}


func (c *StatController) Data() {
	helper.Ajax(pine.H{
		"list": []string{},
	}, 0, c.Ctx())
}

