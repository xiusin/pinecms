package cmd

import (
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化CMS配置",
	Run: func(cmd *cobra.Command, args []string) {
		app := tview.NewApplication()
		var ip string
		form := tview.NewForm().
			AddInputField("服务器地址", "", 20, nil, func(text string) {
				ip = text
			}).
			AddInputField("端口", "3306", 20, nil, nil).
			AddInputField("账号", "", 20, nil, nil).
			AddPasswordField("密码", "", 20, '*', nil).
			AddInputField("数据库", "", 20, nil, nil).
			AddButton("保存", func() {}).
			AddButton("取消", func() {
				app.Stop()
			})

		form.SetBorder(true).SetTitle("\\PineCms初始化安装").SetTitleAlign(tview.AlignCenter)
		if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
