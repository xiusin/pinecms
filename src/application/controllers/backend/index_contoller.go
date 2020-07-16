package backend

import (
	"encoding/json"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/shirou/gopsutil/mem"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"

	"github.com/shirou/gopsutil/disk"
	"github.com/xiusin/pinecms/src/common/helper"
)

type IndexController struct {
	pine.Controller
}

type MemPos struct {
	TimePos int64 `json:"x"`
	Percent int   `json:"y"`
}

func (c *IndexController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY("/index/index", "Index")
	b.ANY("/index/menu", "Menu")
	b.ANY("/index/main", "Main")
	go func() {
		for range time.Tick(2 * time.Second) {
			vm, err := mem.VirtualMemory()
			if err != nil {
				pine.Logger().Error("读取服务器内存信息错误:", err)
			} else {
				mems := getMems()
				mems = append(mems, MemPos{TimePos: time.Now().In(helper.GetLocation()).Unix(), Percent: int(vm.UsedPercent)})
				memsSaveData, _ := json.Marshal(mems)
				pine.Make(controllers.ServiceICache).(cache.AbstractCache).Set(controllers.CacheMemCollect, memsSaveData)
			}
		}
	}()
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

func (c *IndexController) Menu(icache cache.AbstractCache) {
	menus := c.GetMenus(icache)
	helper.Ajax(menus, 1, c.Ctx())
}

func (c *IndexController) Main(orm *xorm.Engine, iCache cache.AbstractCache) {
	var us, _ = disk.Usage(helper.GetRootPath())
	//要转换的值，fmt方式，切割长度如果为-1则显示最大长度，64是float64
	c.ViewData("FullSize", us.Total/1024/1024/1024)
	c.ViewData("usedSize", us.Used/1024/1024/1024)
	c.ViewData("usedPercent", int(us.UsedPercent))
	c.ViewData("NumCPU", runtime.NumCPU())
	c.ViewData("GoVersion", "Version "+strings.ToUpper(runtime.Version()))
	c.ViewData("pineVersion", "Version "+pine.Version)
	c.ViewData("Goos", strings.ToUpper(runtime.GOOS))
	c.ViewData("Grountues", runtime.NumGoroutine())
	memBytes, _ := json.Marshal(getMems())

	data := strings.ReplaceAll(string(memBytes), "\"", "")
	c.ViewData("mems", data)

	var todos []tables.Todo
	orm.Where("status = ?", 1).Find(&todos)
	c.ViewData("todos", todos)
	var statistics = map[string]uint8{}
	_ = iCache.GetWithUnmarshal(controllers.CacheStatistics, &statistics)
	curMonth := time.Now().In(helper.GetLocation()).Format("01")
	// 整合数据
	currentYear, currentMonth, _ := time.Now().Date()

	lastDay := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, time.Now().Location()).AddDate(0, 1, -1).Format("02")
	l, _ := strconv.Atoi(lastDay)
	var statiData = make([]int, l)
	totalVisits := 0
	for k, v := range statistics {
		md := strings.Split(k, "-")
		if md[0] == curMonth /*只统计当月内的数据*/ {
			day, _ := strconv.Atoi(md[1])
			statiData[day] = int(v)
			totalVisits += int(v)
		}
	}
	b, _ := json.Marshal(&statiData)
	bs := string(b)
	c.ViewData("month", curMonth)
	c.ViewData("totalVisits", totalVisits)
	c.ViewData("statiData", bs)
	c.collationFormatVisits(iCache)
	c.View("backend/index_main.html")
}

type visitStruct struct {
	Color   string
	Present int
	Total   int
}

func (c *IndexController) collationFormatVisits(iCache cache.AbstractCache) {
	// 统计访问来源
	referStruct := map[string]visitStruct{}
	var defaultRefers = map[string]int{"baidu": 0, "google": 0, "so": 0, "bing": 0, "sougou": 0, "other": 0}
	var referNames = map[string]string{"baidu": "1.百度", "google": "2.谷歌", "so": "3.360搜索", "bing": "4.必应", "sougou": "5.搜狗", "other": "6.其他"}
	var totalRefer = 0
	var refers = map[string]int{}
	iCache.GetWithUnmarshal(controllers.CacheRefer, &refers)
	if len(refers) == 0 {
		refers = defaultRefers
	} else {
		for _, v := range refers {
			totalRefer += v
		}
		for k, v := range defaultRefers {
			if _, ok := refers[k]; !ok {
				refers[k] = v
			}
		}
	}
	var colors = []string{"green", "cyan", "amethyst", "orange", "red", "hotpink"}
	i := 0
	for k, v := range refers {
		persent := 0
		if totalRefer > 0 {
			persent = v * 100 / totalRefer
		}
		referStruct[referNames[k]] = struct {
			Color   string
			Present int
			Total   int
		}{Color: colors[i], Present: persent, Total: v}
		i++
	}
	c.ViewData("refers", referStruct)
}

func (c *IndexController) GetMenus(iCache cache.AbstractCache) []MenuV2 {
	roleid := c.Ctx().Value("roleid")
	if roleid == nil {
		roleid = interface{}(int64(0))
	}
	menus := models.NewMenuModel().GetMenu(1, roleid.(int64)) //获取menuid内容
	cacheKey := fmt.Sprintf(controllers.CacheAdminMenuByRoleIdAndMenuId, roleid, 1)
	var menujs []MenuV2 //要返回json的对象
	var data string
	dataBytes, _ := iCache.Get(cacheKey)
	data = string(dataBytes)
	data = ""
	if data == "" || json.Unmarshal([]byte(data), &menujs) != nil {
		for _, v := range menus {
			menu := models.NewMenuModel().GetMenu(v.Id, roleid.(int64))
			if len(menu) == 0 {
				continue
			}
			var sonmenu []MenuV2
			for _, son := range menu {
				sonmenu = append(sonmenu, MenuV2{
					Label:           son.Name,
					NodePath:        son.A,
					SideVisible:     true,
					Icon:            "fa fa-home",
					PathToComponent: strFirstToUpper(son.C + "/" + son.A), // 指定要路由到的模块， 请注意横线问题
				})
			}
			menujs = append(menujs, MenuV2{
				Label:       v.Name,
				NodePath:    v.C,
				SideVisible: true,
				Icon:        "fa fa-home",
				Children:    sonmenu,
			})

		}
		strs, _ := json.Marshal(&menujs)
		if err := iCache.Set(cacheKey, strs); err != nil {
			pine.Logger().Errorf("save cache %s failed: %s", cacheKey, err.Error())
		}
	}
	return menujs
}
