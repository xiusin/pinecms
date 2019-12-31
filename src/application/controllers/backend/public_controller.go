package backend

import (
	"github.com/xiusin/iriscms/src/config"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/xiusin/iriscms/src/common/helper"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/xiusin/iriscms/src/common/storage"
)

type PublicController struct {
	Ctx     iris.Context
	Orm     *xorm.Engine
	Session *sessions.Session
}

func (c *PublicController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("ANY", "/upload", "Upload")
	b.Handle("ANY", "/fedir-scan", "FeDirScan")
	//b.Handle("ANY", "/verify-code", "VerifyCode")
}

func (c *PublicController) FeDirScan() {
	_ , err := ioutil.ReadDir(config.AppConfig().FeDirname)
	if err != nil {

	}
}

//上传图片
func (c *PublicController) Upload() {
	mid := c.Session.GetString("mid")
	if mid == "" {
		mid = "public"
	}
	var isEditor = false
	if c.Ctx.URLParam("editorid") != "" {
		//百度编辑器的返回内容
		isEditor = true
	}
	uploadDir := "upload"
	conf := c.Ctx.Values().Get("app.config").(map[string]string)
	setting := c.Ctx.Values().Get("setting").(map[string]string)
	engine := conf["uploadEngine"]
	var uploader storage.Uploader
	if engine != "oss" {
		uploader = storage.NewFileUploader(uploadDir)
		uploadDir = uploader.(*storage.FileUploader).BaseDir
	} else {
		uploader = storage.NewOssUploader(setting)
	}
	//生成要保存到目录和名称
	nowTime := helper.NowDate("Ymd")
	uploadDir = uploadDir + "/" + mid + "/" + nowTime
	file, fs, err := c.Ctx.FormFile("filedata")
	if err != nil {
		c.Ctx.Application().Logger().Error("上传文件失败", err.Error())
		uploadAjax(c.Ctx, map[string]string{
			"errmsg":  "打开上传临时文件失败 : " + err.Error(),
			"state":   "打开上传临时文件失败 : " + err.Error(),
			"errcode": "1",
		}, isEditor)
		return
	}
	defer file.Close()
	fname := fs.Filename
	info := strings.Split(fname, ".")
	ext := strings.ToLower(info[len(info)-1])
	canUpload := []string{"jpg", "jpeg", "png"}
	flag := false
	for _, v := range canUpload {
		if v == ext {
			flag = true
		}
	}
	if !flag {
		uploadAjax(c.Ctx, map[string]string{
			"errmsg":  "不支持的文件类型",
			"state":   "不支持的文件类型",
			"errcode": "1",
		}, isEditor)
		return
	}
	filename := string(helper.Krand(10, 3)) + "." + ext
	storageName := uploadDir + "/" + filename
	path, err := uploader.Upload(storageName, file)
	if err != nil {
		uploadAjax(c.Ctx, map[string]string{
			"errmsg":  "上传失败:" + err.Error(),
			"state":   "上传失败:" + err.Error(),
			"errcode": "1",
		}, isEditor)
		return
	}

	resJson := map[string]string{
		"originalName": fname,        //原始名称
		"name":         filename,     //新文件名称
		"url":          path,         //完整文件名,即从当前配置目录开始的URL
		"size":         "",           //文件大小
		"type":         "image/jpeg", //文件类型
		"state":        "上传成功",       //上传状态
		"errmsg":       path,
		"errcode":      "0",
	}
	uploadAjax(c.Ctx, resJson, isEditor)
	return
}

////生成验证码
//func (this *PublicController) VerifyCode() {
//	cpt := captcha.New()
//	fontPath := helper.GetRootPath() + "/resources/fonts/comic.ttf"
//	// 设置字体
//	cpt.SetFont(fontPath)
//	// 返回验证码图像对象以及验证码字符串 后期可以对字符串进行对比 判断验证
//	this.Ctx.ContentType("img/png")
//	img, str := cpt.Create(1, captcha.ALL)
//	this.Session.SetFlash("verify_code", str)
//	png.Encode(this.Ctx.ResponseWriter(), img) //发送图片内容到浏览器
//}

func uploadAjax(ctx iris.Context, uploadData map[string]string, isEditor bool) {
	if !isEditor {
		errmsg, ok := uploadData["errmsg"]
		if !ok {
			helper.Ajax("未知错误", 1, ctx)
			return
		}
		errcode, ok := uploadData["errcode"]
		code, err := strconv.Atoi(errcode)
		if err != nil {
			errmsg = "未知错误"
		}
		helper.Ajax(errmsg, int64(code), ctx)
	} else {
		ctx.JSON(uploadData)
	}
}
