package helper

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
	"unsafe"

	"github.com/xiusin/pine/di"

	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"xorm.io/xorm"

	"github.com/kataras/go-mailer"
)

const TimeFormat = "2006-01-02 15:04:05"

var location *time.Location

func init() {
	location = time.FixedZone("CST", 8*3600)
	rand.Seed(time.Now().UnixNano())
}

func GetLocation() *time.Location {
	return location
}

// GetRootPath 获取项目根目录 (即 main.go的所在位置)
func GetRootPath(relPath ...string) string {
	pwd, _ := os.Getwd()
	if len(relPath) > 0 {
		pwd = filepath.Join(pwd, relPath[0])
	}
	return pwd
}

// GetCallerFuncName 获取当前执行函数名 只用于日志记录
func GetCallerFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	return runtime.FuncForPC(pc[0]).Name() + ":"
}

// Krand 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	for i := 0; i < size; i++ {
		if isAll {
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

// GetMd5 md5加密字符串
func GetMd5(str string) string {
	md := md5.New()
	md.Write([]byte(str))
	return hex.EncodeToString(md.Sum(nil))
}

// Ajax Ajax返回数据给前端
func Ajax(msg interface{}, errcode int64, this *pine.Context) {
	if errcode == 0 {
		errcode = 1000
	}
	// 添加操作日志
	data := pine.H{"code": errcode}
	if errcode != 1000 {
		switch msg.(type) {
		case error:
			data["message"] = msg.(error).Error()
		default:
			data["message"] = msg
		}
	} else {
		switch msg.(type) {
		case string:
			data["message"] = msg
		default:
			data["data"] = msg
		}
		data["data"] = msg
	}
	_ = this.Render().JSON(data)
}

// GetTimeStamp 获取时间戳
func GetTimeStamp() int {
	timestamp := time.Now().In(location).Unix()
	return int(timestamp)
}

// NowDate 当前时间 Y m d H:i:s
func NowDate(str string) string {
	return time.Now().In(location).Format(str)
}

// Password 生成密码
func Password(password, encrypt string) string {
	return GetMd5(GetMd5(password) + encrypt)
}

// IsFalse 检测字段是否为 空 0 nil
func IsFalse(args ...interface{}) bool {
	for _, v := range args {
		switch v.(type) {
		case string:
			if v != "" {
				return false
			}
		case int, int64, int8, int32:
			if v != 0 {
				return false
			}
		case bool:
			if !v.(bool) {
				return false
			}
		default:
			return true
		}
	}
	return true
}

type EmailOpt struct {
	Title        string
	UrlOrMessage string
	Address      []string
}

func SendEmail(opt *EmailOpt, conf map[string]string) error {
	port, err := strconv.Atoi(conf["EMAIL_PORT"])
	if err != nil {
		port = 25
	}
	mailService := mailer.New(mailer.Config{
		Host:      conf["EMAIL_SMTP"],
		Username:  conf["EMAIL_USER"],
		Password:  conf["EMAIL_PWD"],
		Port:      port,
		FromAlias: conf["EMAIL_SEND_NAME"],
	})

	str := ""
	if strings.HasPrefix(opt.UrlOrMessage, "http") {
		str = `<a href="` + opt.UrlOrMessage + `" class="a-link" target="_blank">` + opt.UrlOrMessage + `</a>
            <br/>如果链接点击无效，请将链接复制到您的浏览器中继续访问。`
	} else {
		str = opt.UrlOrMessage
	}
	content := `
<html>
<head>
<meta http-equiv="content-type" content="text/html;charset=utf-8">
</head>
<body>
<style>
    .main{width:655px;margin:0px auto;border: 2px solid #0382db;background-color:#f8fcff;}
    .main-box{ padding:15px;}
    .mail-title{ font-size:15px; color:#ffffff; background-color:#0382db; font-weight:bold; padding:0px 0px 0px 15px;height:80px;line-height:80px;}
    .fontstyle{ color:#be0303; font-weight:bold;}
    .a-link{ color:#0078ff; font-weight:bold;}
    .csdn{ text-align:right; color:#000000; font-weight: bold;}
    .csdn-color{ color:#000000; font-weight: bold; padding-top:20px;}
    .logo-right{float:right;display:inline-block;padding-right:15px;}
</style>
<div class="main">
    <div class="mail-title">` + opt.Title + `</div>
    <div class="main-box">
        <p>` + opt.Title + `<br/>
            ` + str + `</p>
        <p class="csdn csdn-color">by PineCMS</p>
    </div>
</div>
</body>
</html>
`
	return mailService.Send(opt.Title, content, opt.Address...)
}

func NewOrmLogFile(path string) *os.File {
	f, err := os.OpenFile(filepath.Join(path, "orm.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetORM() *xorm.Engine {
	return pine.Make(controllers.ServiceXorm).(*xorm.Engine)
}

func ToInterfaces(values interface{}) []interface{} {
	v := reflect.ValueOf(values)
	if v.Kind() != reflect.Slice {
		return nil
	}
	var is []interface{}
	for i := 0; i < v.Len(); i++ {
		is = append(is, v.Index(i).Interface())
	}
	return is
}

func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}

func UcFirst(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArr := []rune(str)
	if strArr[0] >= 97 && strArr[0] <= 122 {
		strArr[0] -= 32
	}
	return string(strArr)
}

func Bytes2String(b []byte) *string {
	return (*string)(unsafe.Pointer(&b))
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// Inject 注入依赖
func Inject(key interface{}, v interface{}, single ...bool) {
	if len(single) == 0 {
		single = append(single, true)
	}
	if vi, ok := v.(di.BuildHandler); ok {
		di.Set(key, vi, single[0])
	} else {
		di.Set(key, func(builder di.AbstractBuilder) (i interface{}, e error) {
			return v, nil
		}, single[0])
	}

}

func GetUrlPrefix(catid int64) string {
	getUrlPrefix := di.MustGet(controllers.ServiceCatUrlPrefixFunc).(func(int64) string)
	return getUrlPrefix(catid)
}

// 处理文章列表信息数据. 补全一些cms生成
func HandleArtListInfo(list []map[string]string, titlelen int) {
	for i, art := range list {
		catid, _ := strconv.Atoi(art["catid"])
		prefix := GetUrlPrefix(int64(catid))
		if art["type"] != "2" {
			art["caturl"] = fmt.Sprintf("/%s/", prefix)
			art["typeurl"] = art["caturl"]
		}
		id, _ := strconv.Atoi(art["id"])
		art["arcurl"] = fmt.Sprintf("/%s/%d.html", prefix, id)
		art["arturl"] = art["arcurl"]
		art["click"] = art["visit_count"]
		art["fulltitle"] = art["title"]
		if titlelen > 0 {
			if utf8.RuneCountInString(art["title"]) > titlelen {
				titleRune := []rune(art["title"])
				art["title"] = string(titleRune[:titlelen])
			}
		}
		list[i] = art
	}
}

// PanicErr 抛出异常
func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}
