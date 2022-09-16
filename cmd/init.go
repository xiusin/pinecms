package cmd

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"github.com/xiusin/logger"
	"github.com/xiusin/pinecms/src/config"
	"os"
	"time"
)

var initFilePath = "resources/pinecms.sql"

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "基于文件" + initFilePath + "初始化安装项目, 存在配置则忽略命令操作",
	Run: func(cmd *cobra.Command, args []string) {
		app := tview.NewApplication()
		var submitted bool
		var db config.DbConf
		if db.Inited() {
			logger.Error("项目已初始化")
			return
		}

		if _, err := os.Stat(initFilePath); os.IsNotExist(err) {
			logger.Error("缺少初始数据库文件" + initFilePath)
			return
		}

		db.Db.Conf.Port = "3306"
		db.Db.DbPrefix = "pinecms_"
		db.Db.DbDriver = "mysql"

		pages := tview.NewPages()
		pages.SetBackgroundColor(tcell.ColorDefault)
		connForm := tview.NewForm().
			AddInputField("服务器地址", db.Db.Conf.ServeIp, 20, nil, func(text string) {
				db.Db.Conf.ServeIp = text
			}).
			AddInputField("端口", db.Db.Conf.Port, 20, nil, func(text string) {
				db.Db.Conf.Port = text
			}).
			AddInputField("账号", db.Db.Conf.Username, 20, nil, func(text string) {
				db.Db.Conf.Username = text
			}).
			AddPasswordField("密码", db.Db.Conf.Password, 20, '*', func(text string) {
				db.Db.Conf.Password = text
			}).
			AddInputField("数据库", db.Db.Conf.Name, 20, nil, func(text string) {
				db.Db.Conf.Name = text
			}).
			AddInputField("表前缀", db.Db.DbPrefix, 20, nil, func(text string) {
				db.Db.DbPrefix = text
			}).
			AddButton("保存", func() {
				defer func() {
					if err := recover(); err != nil {
						submitted = false
						alert(pages, "连接数据库异常: "+err.(error).Error())
					}
				}()
				if !db.Db.Conf.Check() {
					alert(pages, "数据库配置信息必须填写")
				} else if !submitted {
					orm := config.InitDB(&db)
					if err := orm.Ping(); err != nil {
						alert(pages, "连接数据库错误: "+err.Error())
						return
					}
					if _, err := orm.ImportFile(initFilePath); err != nil {
						alert(pages, "导入数据库失败: "+err.Error())
						return
					}
					submitted = true
					err := db.BuildYaml()
					if err != nil {
						submitted = false
						alert(pages, "保存失败, "+err.Error())
						return
					}
					go func() {
						time.Sleep(time.Second * 3)
						app.Stop()
					}()
					alert(pages, "保存成功, 3秒后关闭")
				}
			}).
			AddButton("取消", func() {
				app.Stop()
			}).
			SetButtonsAlign(tview.AlignCenter)

		connForm.SetBorder(true).SetTitle("  PineCms Initializer	 ").SetTitleAlign(tview.AlignCenter)
		flex := tview.NewFlex().AddItem(Center(40, 18, connForm), 0, 3, true)
		pages.AddPage("base", flex, true, true)
		if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}
	},
}

func Center(width, height int, p tview.Primitive) tview.Primitive {
	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(p, height, 1, true).
			AddItem(nil, 0, 1, false), width, 1, true).
		AddItem(nil, 0, 1, false)
}

func alert(pages *tview.Pages, message string) *tview.Pages {
	id := "dialog"
	return pages.AddPage(id, tview.NewModal().SetText(message).AddButtons([]string{"确定"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.HidePage(id).RemovePage(id)
		}), false, true)
}
