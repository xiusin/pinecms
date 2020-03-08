module github.com/xiusin/pinecms

go 1.13

require (
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/CloudyKit/jet v2.1.3-0.20180809161101-62edd43e4f88+incompatible
	github.com/afocus/captcha v0.0.0-20191010092841-4bd1f21c8868
	github.com/alexmullins/zip v0.0.0-20180717182244-4affb64b04d0
	github.com/aliyun/aliyun-oss-go-sdk v2.0.4+incompatible
	github.com/fatih/color v1.9.0
	github.com/fsnotify/fsnotify v1.4.7
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/xorm v0.7.9
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/gorilla/securecookie v1.1.1
	github.com/gorilla/websocket v1.4.1
	github.com/hpcloud/tail v1.0.0
	github.com/imroc/req v0.2.4
	github.com/kataras/go-mailer v0.1.0
	github.com/kataras/golog v0.0.10
	github.com/kataras/iris/v12 v12.1.2
	github.com/kataras/tablewriter v0.0.0-20180708051242-e063d29b7c23 // indirect
	github.com/landoop/tableprinter v0.0.0-20180806200924-8bd8c2576d27
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mattn/go-runewidth v0.0.8 // indirect
	github.com/microcosm-cc/bluemonday v1.0.2
	github.com/mitchellh/go-homedir v1.1.0
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/prometheus/common v0.4.0
	github.com/satori/go.uuid v1.2.0
	github.com/shirou/gopsutil v2.19.11+incompatible
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.6.2
	github.com/xiusin/debug v0.0.0-00010101000000-000000000000
	github.com/xiusin/logger v0.0.0-00010101000000-000000000000
	github.com/xiusin/pine v0.0.0-20200301045755-d3e2bba0b14b
	github.com/yanyiwu/gojieba v1.1.1
	golang.org/x/image v0.0.0-20191214001246-9130b4cfad52
	gopkg.in/yaml.v2 v2.2.7
	xorm.io/builder v0.3.6
	xorm.io/core v0.7.3
)

replace github.com/xiusin/pine => ../pine

replace github.com/xiusin/logger => ../logger

replace github.com/xiusin/debug => ../debug
