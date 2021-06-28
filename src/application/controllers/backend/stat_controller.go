package backend

import (
	"github.com/xiusin/pine"
	"time"
)

var runningStart = time.Now()

type StatController struct {
	pine.Controller
}

func (_ *StatController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY("/stat/data", "Data")
}

func (c *StatController) Data() {
	//runningTime := time.Now().Sub(runningStart) // 运行时长

}
