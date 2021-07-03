package cmd

import (
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"github.com/xiusin/pinecms/src/config"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化CMS配置",
	Run: func(cmd *cobra.Command, args []string) {
		app := tview.NewApplication()
		var db config.DbConfig
		form := tview.NewForm().
			AddInputField("服务器地址", "", 20, nil, func(text string) {
				db.Db.Conf.ServeIp = text
			}).
			AddInputField("端口", "3306", 20, nil, func(text string) {
				db.Db.Conf.Port = text
			}).
			AddInputField("账号", "", 20, nil, func(text string) {
				db.Db.Conf.Username = text
			}).
			AddPasswordField("密码", "", 20, '*', func(text string) {
				db.Db.Conf.Password = text
			}).
			AddInputField("数据库", "", 20, nil, func(text string) {
				db.Db.Conf.Name = text
			}).
			AddInputField("表前缀", "", 20, nil, func(text string) {
				db.Db.DbPrefix = text
			}).
			AddButton("保存", func() {
				db.CreateYaml()
				app.Stop()
			}).
			AddButton("取消", func() {
				app.Stop()
			})

		form.SetBorder(true).SetTitle("PineCms初始化安装").SetTitleAlign(tview.AlignCenter)
		if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
