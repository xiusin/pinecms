package backend

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"

	"github.com/xiusin/iriscms/src/application/controllers"
	"github.com/xiusin/iriscms/src/application/models"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/xiusin/iriscms/src/common/helper"
)

type IndexController struct {
	pine.Controller
}

type MemPos struct {
	TimePos string `json:"time_pos"`
	Percent int    `json:"percent"`
}

type CpuPos struct {
	TimePos string    `json:"time_pos"`
	Value   []float64 `json:"percent"`
}

func (c *IndexController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY("/index/index", "Index")
	b.ANY("/index/menu", "Menu")
	b.ANY("/index/main", "Main")
	b.ANY("/index/sessionlife", "Sessionlife")

	// mem cpu collect
	go func() {

		// 每10秒采集一次服务器信息
		for range time.Tick(10 * time.Second) {
			vm, err := mem.VirtualMemory()
			if err != nil {
				pine.Logger().Error("读取服务器内存信息错误:", err)
			} else {
				mems := getMems()
				mems = append(mems, MemPos{TimePos: time.Now().In(helper.GetLocation()).Format("15:04:05"), Percent: int(vm.UsedPercent)})
				memsSaveData, _ := json.Marshal(mems)
				pine.Make("cache.ICache").(cache.ICache).Set("memCollect", memsSaveData)
			}
			cpuInfos, err := cpu.Percent(0, false)
			if err != nil {
				pine.Logger().Error("读取服务器CPU信息错误:", err)
			} else {
				cpus := getCpus()
				cpus = append(cpus, CpuPos{TimePos: time.Now().In(helper.GetLocation()).Format("15:04:05"), Value: cpuInfos})
				memsSaveData, _ := json.Marshal(cpus)
				pine.Make("cache.ICache").(cache.ICache).Set("cpuCollect", memsSaveData)
			}
		}
	}()

}

func getMems() []MemPos {
	var mems []MemPos
	c := pine.Make("cache.ICache").(cache.ICache)
	memCollect, _ := c.Get("memCollect")
	if memCollect == nil {
		memCollect = []byte{}
	}
	_ = json.Unmarshal(memCollect, &mems)
	if len(mems) > 10 {
		mems = mems[len(mems)-10:]
	}
	return mems
}

func getCpus() []CpuPos {
	var cpus []CpuPos
	c := pine.Make("cache.ICache").(cache.ICache)
	memCollect, _ := c.Get("cpuCollect")
	if memCollect == nil {
		memCollect = []byte{}
	}
	_ = json.Unmarshal(memCollect, &cpus)
	if len(cpus) > 10 {
		cpus = cpus[len(cpus)-10:]
	}
	return cpus
}

func (c *IndexController) Index() {
	roleid := c.Ctx().Value("roleid")
	if roleid == nil {
		c.Ctx().Redirect("/b/login/index")
		return
	}
	menus := models.NewMenuModel().GetMenu(0, roleid.(int64)) //读取一级菜单
	c.Ctx().Render().ViewData("menus", menus)
	c.Ctx().Render().ViewData("username", c.Session().Get("username"))
	c.Ctx().Render().HTML("backend/index_index.html")
}

var us, _ = disk.Usage(helper.GetRootPath())

func (c *IndexController) Main() {

	formatMem := func(mem uint64) string {
		fm := map[int64]string{
			1024:                      "K",
			1024 * 1024:               "MB",
			1024 * 1024 * 1024:        "GB",
			1024 * 1024 * 1024 * 1024: "TB",
		}
		for b, s := range fm {
			res := float64(mem) / float64(b)
			if res > float64(1024) || res < 1 {
				continue
			}
			return strconv.FormatFloat(res, 'f', 2, 64) + s
		}
		return strconv.FormatFloat(float64(mem), 'f', 2, 64) + "Bit"
	}

	//要转换的值，fmt方式，切割长度如果为-1则显示最大长度，64是float64
	siteSize := formatMem(us.Total)

	c.Ctx().Render().ViewData("SiteSize", siteSize)
	c.Ctx().Render().ViewData("NumCPU", runtime.NumCPU())
	c.Ctx().Render().ViewData("GoVersion", "Version "+strings.ToUpper(runtime.Version()))
	c.Ctx().Render().ViewData("pineVersion", "Version "+pine.Version)
	c.Ctx().Render().ViewData("Goos", strings.ToUpper(runtime.GOOS))
	c.Ctx().Render().ViewData("Grountues", runtime.NumGoroutine())

	c.Ctx().Render().ViewData("Mem", "未获得内存情况")

	c.Ctx().Render().ViewData("cpus", getCpus())
	c.Ctx().Render().ViewData("mems", getMems())

	c.Ctx().Render().HTML("backend/index_main.html")
}

func (c *IndexController) Menu(iCache cache.ICache) {
	meid, _ := c.Ctx().PostInt64("menuid")
	roleid := c.Ctx().Value("roleid")
	if roleid == nil {
		roleid = interface{}(int64(0))
	}
	menus := models.NewMenuModel().GetMenu(meid, roleid.(int64)) //获取menuid内容
	cacheKey := fmt.Sprintf(controllers.CacheAdminMenuByRoleIdAndMenuId, roleid, meid)
	var menujs []map[string]interface{} //要返回json的对象
	var data string
	//if meid > 0 {
	//	dataBytes, _ := iCache.Get(cacheKey)
	//	data = string(dataBytes)
	//} else {
	data = ""
	//}
	if data == "" || json.Unmarshal([]byte(data), &menujs) != nil {
		for _, v := range menus {
			menu := models.NewMenuModel().GetMenu(v.Id, roleid.(int64))
			if len(menu) == 0 {
				continue
			}
			var sonmenu []map[string]interface{}
			for _, son := range menu {
				sonmenu = append(sonmenu, map[string]interface{}{
					"text": son.Name,
					"id":   son.Id,
					"url":  "/b/" + son.C + "/" + son.A + "?menuid=" + strconv.Itoa(int(son.Id)) + "&" + son.Data,
				})
			}
			menujs = append(menujs, map[string]interface{}{
				"name": v.Name,
				"son":  sonmenu,
			})

		}
		strs, _ := json.Marshal(&menujs)
		if err := iCache.Set(cacheKey, strs); err != nil {
			pine.Logger().Errorf("save cache %s failed: %s", cacheKey, err.Error())
		}
	}
	c.Ctx().Render().JSON(menujs)
}

//维持session不过期
func (c *IndexController) Sessionlife() {

}
