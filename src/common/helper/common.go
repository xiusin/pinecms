package helper

//避免循环调用错误,公用非依赖变量以函数方式返回
import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/application/controllers"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"golang.org/x/image/bmp"

	"github.com/kataras/go-mailer"

	"github.com/nfnt/resize"
	"golang.org/x/image/draw"
)

type Node struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Children []Node `json:"children"`
}

//GetRootPath 获取IRIS项目根目录 (即 main.go的所在位置)
func GetRootPath() string {
	pwd, _ := os.Getwd()
	return pwd
}

var location *time.Location

func init()  {
	loc, err := time.LoadLocation("PRC")
	if err != nil {
		pine.Logger().Error(err)
		loc, err = time.LoadLocation("UTC")
	}
	location = loc
}

const TimeFormat = "2006-01-02 15:04:05"

func GetLocation() *time.Location {
	return location
}

func ScanDir(dir string) []Node {
	cacher := pine.Make("cache.AbstractCache").(cache.AbstractCache)
	var nodes []Node
	err := cacher.GetWithUnmarshal(controllers.CacheFeTplList, &nodes)
	if err != nil || len(nodes) == 0 {
		fs, err := ioutil.ReadDir(dir)
		if err != nil {
			panic(fmt.Sprintf("打开%s:%s", dir, err.Error()))
		}
		for _, f := range fs {
			if filepath.Ext(f.Name()) != ".jet" {
				continue
			}
			name := f.Name()
			id := FilePathToEasyUiID(name)
			node := Node{
				Id:       id,
				Name:     name,
				Children: nil,
			}
			nodes = append(nodes, node)
		}
		cacher.SetWithMarshal(controllers.CacheFeTplList, &nodes)
	}

	return nodes
}

func FilePathToEasyUiID(path string) string {
	return strings.Replace(strings.Replace(path, "/", "d__ds__fd", -1), ".", "f_dot_e", 1)
}

func EasyUiIDToFilePath(id string) string {
	return strings.Replace(strings.Replace(id, "d__ds__fd", "/", -1), "f_dot_e", ".", 1)
}

// 获取当前执行函数名 只用于日志记录
func GetCallerFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name() + ":"
}

//Krand 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	//随机种子 (如果不以时间戳作为时间种子, 可能每次生成的随机数每次都相同)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll {
			// random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

//GetMd5 md5加密字符串
func GetMd5(str string) string {
	md := md5.New()
	md.Write([]byte(str))
	return hex.EncodeToString(md.Sum(nil))
}

//Ajax Ajax返回数据给前端
func Ajax(errmsg interface{}, errcode int64, this *pine.Context) {
	// 添加操作日志




	this.Render().JSON(pine.H{"errcode": errcode, "errmsg": errmsg})
}

func Dialog(errmsg interface{}, this *pine.Context) {
	this.Render().Text(fmt.Sprintf(`<div class="easyui-dialog" title="错误提醒" style="width:400px;height:200px;"
    data-options="iconCls:'icon-error',resizable:true">%s</div>`, errmsg))
}

//GetTimeStamp 获取时间戳
func GetTimeStamp() int {
	timestamp := time.Now().In(location).Unix()
	return int(timestamp)
}

//NowDate 当前时间 Y m d H:i:s
func NowDate(str string) string {
	return time.Now().In(location).Format(format(str))
}

func format(str string) string {
	str = strings.Replace(str, "Y", "2006", 1)
	str = strings.Replace(str, "m", "01", 1)
	str = strings.Replace(str, "d", "02", 1)
	str = strings.Replace(str, "H", "13", 1)
	str = strings.Replace(str, "i", "04", 1)
	str = strings.Replace(str, "s", "05", 1)
	return str
}

//FormatTime 时间戳格式化时间
func FormatTime(timestamp int64) string {
	t := time.Unix(timestamp, 0).In(location)
	str := TimeFormat
	return t.Format(str)
}

//GetImg 根据图片路径生成图片,待优化函数
func GetImg(path, waterPath string) {
	f, err := os.Open(path) //打开文件
	if err != nil {
		fmt.Println("打开文件失败:", err.Error())
		return
	}
	defer f.Close()
	filename := strings.Split(f.Name(), ".")
	if len(filename) != 2 || (filename[1] != "jpg" && filename[1] != "jpeg" && filename[1] != "gif" && filename[1] != "png") {
		return
	}
	var imager image.Image
	if filename[1] == "jpg" {
		imager, err = jpeg.Decode(f)
	} else if filename[1] == "jpeg" {
		imager, err = jpeg.Decode(f)
	} else if filename[1] == "gif" {
		imager, err = gif.Decode(f)
	} else if filename[1] == "png" {
		imager, err = png.Decode(f)
	}
	if err != nil {
		fmt.Println("打开文件失败:", err.Error())
		return
	}

	//获取图片缩略图
	thumbnail := resize.Thumbnail(120, 120, imager, resize.Lanczos3)
	fileThumb, err := os.Create(filename[0] + strconv.Itoa(int(time.Now().Unix())) + "_thmub.jpg")
	if err == nil {
		jpeg.Encode(fileThumb, thumbnail, &jpeg.Options{
			Quality: 80})
		fileThumb.Close()
	}
	rectangler := imager.Bounds()
	//创建画布
	newWidth := 200
	m := image.NewRGBA(image.Rect(0, 0, newWidth, newWidth*rectangler.Dy()/rectangler.Dx()))
	//在画布上绘制图片 m画布 m.bounds画布参数, imager 要参照打开的图片信息 image.Point 图片绘制的其实地址 绘制资源
	draw.Draw(m, m.Bounds(), imager, image.Point{100, 100}, draw.Src)
	//绘制水印图
	//必须是PNG图片
	warter, wterr := os.Open(waterPath)
	if wterr == nil {
		//无错误的时候解码
		watermark, dewaerr := png.Decode(warter)
		if dewaerr == nil {
			//无错误的时候添加水印
			draw.Draw(m, watermark.Bounds().Add(image.Pt(30, 30)), watermark, image.ZP, draw.Over)
		} else {
			fmt.Println("水印图片解码失败")
		}
	} else {
		fmt.Println("水印图片打开失败")
	}
	toimg, _ := os.Create(filename[0] + strconv.Itoa(GetTimeStamp()) + "-120-80.jpg") //创建文件系统
	defer toimg.Close()
	//toimg 保存的名称 要参照的画布，图片选项。默认透明图
	jpeg.Encode(toimg, m, &jpeg.Options{Quality: jpeg.DefaultQuality}) //保存为jpeg图片
}

//MultiUpload 多图上传生成html内容
func MultiUpload(field string, data []string, maxImgNum int, required bool, formName, defaultVal, RequiredTips string) string {
	box := ""
	rid := "MultiUploader_" + strconv.Itoa(rand.Int())
	if RequiredTips == "" {
		RequiredTips = formName + "最少上传一张图片"
	}
	var requiredFunc = ""
	if required {
		requiredFunc = `<script>MultiUploader.push(function(){ 
var flag = false; $("[name='` + field + `']").each(function () { if($(this).val()) { flag = true; } });
if (!flag) { $('#` + rid + `_tip').html("` + RequiredTips + `"); return false; } $('#` + rid + `_tip').html(''); return true; });</script>`
	}
	if len(data) > 0 {
		for _, v := range data {
			box += `<div class="imgbox">
					<input class="imgbox_inputBtn" type="image" onclick="return fromUEImageUploader(this)" style='height: 95px; width:95px;' src="` + v + `" alt="点击上传" onerror='this.src="/assets/backend/static/images/nopic.jpg"' />
					<input type="hidden" value="` + v + `" name="` + field + `" />
					<span style='color:#fff;display:inline-block;width:15px;height:15px;font-size:15px;line-height:15px;text-align:center;background:rgba(0,0,0,0.5);font-weight:normal;cursor:pointer;    position: absolute;left: 72px;top: 10px;'   onclick=''>×</span>
				</div>`
		}
	}

	str := box + `
		<div class="imgbox" onclick="return createHtml(this,'` + field + `', ` + strconv.Itoa(maxImgNum) + `)" style="width: 95px; height: 95px">
			<img style="height: 93px;display: block;border: 1px dashed #888;padding: 30px;" src="/assets/backend/static/images/plus.png" />
		</div>`
	return str + `<div id='` + rid + `_tip' class='errtips'></div>` + requiredFunc
}

//SiginUpload 单图上传
func SiginUpload(field, data string, required bool, formName, defaultVal, RequiredTips string) string {
	rid := "siginUploader_" + strconv.Itoa(rand.Int())
	if RequiredTips == "" {
		RequiredTips = formName + "必须上传"
	}
	var requiredFunc = ""
	if required {
		requiredFunc = `<script>siginUploader.push(function(){ if ($('#` + rid + `').val() == '') {$('#` + rid + `_tip').html("` + RequiredTips + `"); return false; } $('#` + rid + `_tip').html(''); return true; });</script>`
	}
	html := `<input onclick="fromUEImageUploader(this)" class="image_upload_src" type="image" src="` + data + `" onerror='this.src="/assets/backend/static/images/nopic.jpg"' alt="点击上传" style="width:100px;height:100px;display:block;border:1px solid #ddd;padding:2px;float:left;" />
			 <input id='` + rid + `' type="hidden" class="image_upload_val" value="` + data + `" name="` + field + `" />
			 <span title="清空图片内容" class='delImg'  onclick="DelImg(this)">×</span>
<div id='` + rid + `_tip' class='errtips'></div>` + requiredFunc
	return html
}

func FileUpload(field, data string, required bool, formName, defaultVal, RequiredTips string) string {
	rid := "fileUploader_" + strconv.Itoa(rand.Int())
	if RequiredTips == "" {
		RequiredTips = formName + "必须上传"
	}
	jsJSON := "{}"
	if len(data) > 0 {
		jsJSON = data
	}
	var requiredFunc = `<script>buildFileLists(document.getElementById('`+rid+`_button'), `+jsJSON+`);</script>`

	if required {
		requiredFunc = `<script> fileUploader.push(function(){ if ($('#` + rid + `').val() == '') {$('#` + rid + `_tip').html("` + RequiredTips + `"); return false; } $('#` + rid + `_tip').html(''); return true; });
</script>`
	}
	html := `<input name="` + field + `" id="`+rid+`" type="hidden" value='` + data +
		`'/><button class="btn btn-default" type="button" id="`+rid+`_button" onclick="fromUEFileUploader(this, 0)" title="会自动过滤重复文件">上传或选择文件</button><div class='easy-uploader'><ul class="list"></ul></div><div id='` + rid + `_tip' class='errtips'></div>
` + requiredFunc
	return html
}

//Tags 标签
func Tags(field, data string, required bool, formName, defaultVal, RequiredTips string) string {
	rid := "tags_" + strconv.Itoa(rand.Int())
	if RequiredTips == "" {
		RequiredTips = formName + "必须上传"
	}
	var requiredFunc = `<script>$('#` + rid + `').tagsInput({ 'height': '100px',  'width': '446px','interactive': true, 'defaultText': '添加标签', onChange: function (tag) {console.log(arguments);},'placeholderColor': '#666666'});`
	if required {
		requiredFunc += `tagger.push(function(){ if ($('#` + rid + `').val() == '') {$('#` + rid + `_tip').html("` + RequiredTips + `"); return false; } $('#` + rid + `_tip').html(''); return true; });`
	}
	html := `<input type="text" id="` + rid + `" name="` + field + `" value="` + data + `" /><div id='` + rid + `_tip' class='errtips'></div>` + requiredFunc + "</script>"
	return html
}

//Password 生成密码
func Password(password, encrypt string) string {
	return GetMd5(GetMd5(password) + encrypt)
}

//IsFalse 检测字段是否为 空 0 nil
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

//IsError 检测是否有Error
func IsError(args ...error) bool {
	for _, v := range args {
		if v != nil {
			return true
		}
	}
	return false
}

type EmailOpt struct {
	Title string
	UrlOrMessage string
	Address []string
}

/**
1. 配置多个邮箱发送
*/
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

func todayFilename(path string) string {
	today := time.Now().Format("2006-01-02")
	return path + "/" + today + ".log"
}

func NewOrmLogFile(path string) *os.File {
	f, err := os.OpenFile(path+"orm.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

/*
* 图片裁剪
* 入参:
* 规则:如果精度为0则精度保持不变
*https://www.cnblogs.com/cqvoip/p/8078882.html
* 返回:error
 */
func clip(in io.Reader, out io.Writer, x0, y0, x1, y1, quality int) error {
	origin, fm, err := image.Decode(in)
	if err != nil {
		return err
	}

	switch fm {
	case "jpeg":
		img := origin.(*image.YCbCr)
		subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.YCbCr)
		return jpeg.Encode(out, subImg, &jpeg.Options{quality})
	case "png":
		switch origin.(type) {
		case *image.NRGBA:
			img := origin.(*image.NRGBA)
			subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.NRGBA)
			return png.Encode(out, subImg)
		case *image.RGBA:
			img := origin.(*image.RGBA)
			subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
			return png.Encode(out, subImg)
		}
	case "gif":
		img := origin.(*image.Paletted)
		subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.Paletted)
		return gif.Encode(out, subImg, &gif.Options{})
	case "bmp":
		img := origin.(*image.RGBA)
		subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
		return bmp.Encode(out, subImg)
	default:
		return errors.New("ERROR FORMAT")
	}
	return nil
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

