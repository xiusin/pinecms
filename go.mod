module github.com/xiusin/pinecms

go 1.14

require (
	github.com/CloudyKit/jet v2.1.3-0.20180809161101-62edd43e4f88+incompatible
	github.com/afocus/captcha v0.0.0-20191010092841-4bd1f21c8868
	github.com/alecthomas/chroma v0.8.0
	github.com/alexmullins/zip v0.0.0-20180717182244-4affb64b04d0
	github.com/aliyun/aliyun-oss-go-sdk v2.1.0+incompatible
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/betacraft/yaag v1.0.0
	github.com/denisenkom/go-mssqldb v0.0.0-20200428022330-06a60b6afbbc // indirect
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gbrlsnchs/jwt/v3 v3.0.0-rc.2
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-playground/validator/v10 v10.3.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-xorm/xorm v0.7.9
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/gookit/color v1.4.2
	github.com/gorilla/securecookie v1.1.1
	github.com/kataras/go-mailer v0.1.0
	github.com/kataras/tablewriter v0.0.0-20180708051242-e063d29b7c23 // indirect
	github.com/landoop/tableprinter v0.0.0-20200104100433-ae9249991eb1
	github.com/lib/pq v1.7.0 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/microcosm-cc/bluemonday v1.0.2
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/schollz/progressbar v1.0.0
	github.com/spf13/cobra v1.0.0
	github.com/valyala/fasthttp v1.26.0
	github.com/xiusin/logger v0.0.5
	github.com/xiusin/pine v0.0.5
	github.com/xwb1989/sqlparser v0.0.0-20180606152119-120387863bf2
	go.etcd.io/bbolt v1.3.4
	golang.org/x/image v0.0.0-20200119044424-58c23975cae1
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	gopkg.in/yaml.v2 v2.3.0
	xorm.io/builder v0.3.7
	xorm.io/core v0.7.3
)

replace github.com/xiusin/pine => ../pine

replace github.com/xiusin/logger => ../logger
