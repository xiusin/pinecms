package backend

import (
	"encoding/json"
	"fmt"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"html/template"
	"runtime"
	"strconv"
	"strings"

	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"

	"github.com/shirou/gopsutil/disk"
	"github.com/xiusin/pinecms/src/common/helper"
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
	b.ANY("/index/index1", "Index1")
	b.ANY("/index/menu", "Menu")
	b.ANY("/index/main", "Main")

	//go func() {
	//	// 每10秒采集一次服务器信息
	//	for range time.Tick(10 * time.Second) {
	//		vm, err := mem.VirtualMemory()
	//		if err != nil {
	//			pine.Logger().Error("读取服务器内存信息错误:", err)
	//		} else {
	//			mems := getMems()
	//			mems = append(mems, MemPos{TimePos: time.Now().In(helper.GetLocation()).Format("15:04:05"), Percent: int(vm.UsedPercent)})
	//			memsSaveData, _ := json.Marshal(mems)
	//			pine.Make(controllers.ServiceICache).(cache.AbstractCache).Set(controllers.CacheMemCollect, memsSaveData)
	//		}
	//	}
	//}()
}

func getMems() []MemPos {
	var mems []MemPos
	c := pine.Make(controllers.ServiceICache).(cache.AbstractCache)
	memCollect, _ := c.Get(controllers.CacheMemCollect)
	if memCollect == nil {
		memCollect = []byte{}
	}
	err := json.Unmarshal(memCollect, &mems)
	if err == nil {
		if len(mems) > 10 {
			mems = mems[len(mems)-10:]
		}
	}
	return mems
}


func (c *IndexController) Index(icache cache.AbstractCache) {
	menus := c.GetMenus(icache)
	c.ViewData("menus", menus)
	c.ViewData("username", c.Session().Get("username"))
	c.ViewData("rolename", c.Session().Get("role_name"))
	c.View("backend/index_index.html")
}

func (c *IndexController) Index1() {
	roleid := c.Ctx().Value("roleid")
	if roleid == nil {
		c.Ctx().Redirect("/b/login/index")
		return
	}
	menus := models.NewMenuModel().GetMenu(0, roleid.(int64)) //读取一级菜单
	c.ViewData("menus", menus)
	c.ViewData("username", c.Session().Get("username"))
	c.View("backend/index_index1.html")
}

var us, _ = disk.Usage(helper.GetRootPath())

func (c *IndexController) Main(iCache cache.AbstractCache) {

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

	c.ViewData("SiteSize", siteSize)
	c.ViewData("NumCPU", runtime.NumCPU())
	c.ViewData("GoVersion", "Version "+strings.ToUpper(runtime.Version()))
	c.ViewData("pineVersion", "Version "+pine.Version)
	c.ViewData("Goos", strings.ToUpper(runtime.GOOS))
	c.ViewData("Grountues", runtime.NumGoroutine())

	c.ViewData("Mem", "未获得内存情况")

	c.ViewData("mems", getMems())

	todos ,_ := iCache.Get(controllers.CacheToDo)
	c.ViewData("todos", template.HTML(string(todos)))
	
	c.View("backend/index_main.html")
}

func (c *IndexController) Menu(iCache cache.AbstractCache) {
	meid, _ := c.Ctx().PostInt64("menuid")
	roleid := c.Ctx().Value("roleid")
	if roleid == nil {
		roleid = interface{}(int64(0))
	}
	menus := models.NewMenuModel().GetMenu(meid, roleid.(int64)) //获取menuid内容
	cacheKey := fmt.Sprintf(controllers.CacheAdminMenuByRoleIdAndMenuId, roleid, meid)
	var menujs []map[string]interface{} //要返回json的对象
	var data string
	if meid > 0 {
		dataBytes, _ := iCache.Get(cacheKey)
		data = string(dataBytes)
	} else {
		data = ""
	}
	if data == "" || json.Unmarshal([]byte(data), &menujs) != nil {
		pine.Logger().Debug("目录列表没有走缓存")
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
	c.Render().JSON(menujs)
}

func (c *IndexController) GetMenus(iCache cache.AbstractCache) []map[string]interface{} {
	roleid := c.Ctx().Value("roleid")
	if roleid == nil {
		roleid = interface{}(int64(0))
	}
	menus := models.NewMenuModel().GetMenu(1, roleid.(int64)) //获取menuid内容
	cacheKey := fmt.Sprintf(controllers.CacheAdminMenuByRoleIdAndMenuId, roleid, 1)
	var menujs []map[string]interface{} //要返回json的对象
	var data string
		dataBytes, _ := iCache.Get(cacheKey)
		data = string(dataBytes)
	if data == "" || json.Unmarshal([]byte(data), &menujs) != nil {
		pine.Logger().Debug("目录列表没有走缓存")
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
	return menujs
}

