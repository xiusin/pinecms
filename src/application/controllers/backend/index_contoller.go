package backend

import (
	"encoding/json"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"runtime"
	"strconv"
	"strings"

	"github.com/xiusin/iriscms/src/application/controllers"
	"github.com/xiusin/iriscms/src/application/models"

	"github.com/kataras/iris/v12"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/xiusin/iriscms/src/common/helper"
)

type IndexController struct {
	pine.Controller
}

func (c *IndexController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY("/index/index", "Index")
	b.ANY("/index/menu", "Menu")
	b.ANY("/index/main", "Main")
	b.ANY("/index/sessionlife", "Sessionlife")
}

func (c *IndexController) Index() {
	roleid := c.Ctx().Value("roleid")
	if roleid == nil {
		c.Ctx().Redirect("/b/login/index")
		return
	}
	menus := models.NewMenuModel(c.Ctx().Value("orm").(*xorm.Engine)).GetMenu(0, roleid.(int64)) //读取一级菜单
	c.Ctx().Render().ViewData("menus", menus)
	c.Ctx().Render().ViewData("username", c.Session().Get("username"))
	c.Ctx().Render().HTML("backend/index_index.html")
}

func (c *IndexController) Main() {
	//统计尺寸
	us, _ := disk.Usage(helper.GetRootPath())

	vm, _ := mem.VirtualMemory()

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
	c.Ctx().Render().ViewData("IrisVersion", "Version "+iris.Version)
	c.Ctx().Render().ViewData("Goos", strings.ToUpper(runtime.GOOS))
	c.Ctx().Render().ViewData("Grountues", runtime.NumGoroutine())
	if vm != nil {
		c.Ctx().Render().ViewData("Mem", "总内存:"+formatMem(vm.Total)+",已使用:"+formatMem(vm.Used))
	} else {
		c.Ctx().Render().ViewData("Mem", "未获得内存情况")
	}
	c.Ctx().Render().HTML("backend/index_main.html")
}

func (c *IndexController) Menu(iCache cache.ICache) {
	meid, _ := c.Ctx().PostInt64("menuid")
	roleid := c.Ctx().Value("roleid")
	if roleid == nil {
		roleid = interface{}(int64(0))
	}
	menus := models.NewMenuModel(c.Ctx().Value("orm").(*xorm.Engine)).GetMenu(meid, roleid.(int64)) //获取menuid内容
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
		for _, v := range menus {
			menu := models.NewMenuModel(c.Ctx().Value("orm").(*xorm.Engine)).GetMenu(v.Id, roleid.(int64))
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
