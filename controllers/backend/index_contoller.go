package backend

import (
	"strconv"
	"github.com/go-xorm/xorm"
	"runtime"
	"iriscms/models"
	"io/ioutil"
	"strings"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/mvc"
)

type IndexController struct {
	Ctx iris.Context
	Orm *xorm.Engine
	Session *sessions.Session
}


func (c *IndexController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("ANY","/index/index", "Index")
	b.Handle("ANY","/index/menu", "Menu")
	b.Handle("ANY","/index/main", "Main")
	b.Handle("ANY","/index/sessionlife", "Sessionlife")
}


func (this *IndexController) Index() {
	roleid,_ := this.Ctx.Values().GetInt64("roleid")
	if roleid == -1 {
		this.Ctx.Redirect("/b/login/index")
		return
	}
	menus := models.NewMenuModel(this.Orm).GetMenu(0, roleid) //读取一级菜单
	this.Ctx.ViewData("menus", menus)
	this.Ctx.ViewData("username", this.Session.Get("username").(string))
	this.Ctx.View("backend/index_index.html")
}

func (this *IndexController) Main() {
	//统计尺寸
	files, _ := ioutil.ReadDir(".")
	var filesize int64 = 0
	for _, file := range files {
		filesize = filesize + file.Size()
	}
	//要转换的值，fmt方式，切割长度如果为-1则显示最大长度，64是float64
	siteSize := strconv.FormatFloat(float64(filesize) / 1024 / 1024, 'f', 3, 64) + " MB"
	this.Ctx.ViewData("SiteSize", siteSize)
	this.Ctx.ViewData("NumCPU", runtime.NumCPU())
	this.Ctx.ViewData("GoVersion", strings.ToUpper(runtime.Version()))
	this.Ctx.ViewData("IrisVersion", iris.Version)
	this.Ctx.ViewData("Goos", strings.ToUpper(runtime.GOOS))
	this.Ctx.ViewData("Grountues", runtime.NumGoroutine())
	this.Ctx.View("backend/index_main.html")
}

func (this *IndexController) Menu() {
	meid, err := strconv.Atoi(this.Ctx.PostValue("menuid"))
	if err != nil {
		meid = 0
	}
	roleid := this.Ctx.Values().Get("roleid").(int64)
	menus := models.NewMenuModel(this.Orm).GetMenu(int64(meid), roleid) //获取menuid内容
	menujs := []map[string]interface{}{}    //要返回json的对象
	for _, v := range menus {
		menu := models.NewMenuModel(this.Orm).GetMenu(v.Id, roleid)
		if len(menu) == 0 {
			continue
		}
		sonmenu := []map[string]interface{}{}
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
	this.Ctx.JSON(menujs)
}

//维持session不过期
func (this *IndexController) Sessionlife() {
	//维持session防止过期
	this.Ctx.WriteString("1")
}
