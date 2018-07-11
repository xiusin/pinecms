package backend

import (
	"image/png"
	"io"
	"os"
	"strconv"
	"strings"

	"iriscms/controllers/backend/helper"

	"github.com/afocus/captcha"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris"
)

type PublicController struct {
	Ctx iris.Context
	Orm *xorm.Engine
	Session *sessions.Session
}


func (c *PublicController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("ANY","/upload", "Upload")
	b.Handle("ANY","/verify-code", "VerifyCode")
}



//上传图片
func (this *PublicController) Upload() {
	mid := this.Session.GetString("mid")
	if mid == "" {
		mid = "public"
	}
	var isEditor = false
	if this.Ctx.URLParam("editorid") != "" {
		//百度编辑器的返回内容
		isEditor = true
	}
	uploadDir := helper.GetRootPath() + "/upload/" + mid
	nowTime := helper.NowDate("Ymd")
	uploadDir = uploadDir + "/" + nowTime
	f, err := os.Open(uploadDir)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(uploadDir, os.ModePerm)
		if err != nil {
			uploadAjax(this.Ctx, map[string]string{
				"errmsg":  "创建文件夹" + uploadDir + "失败,原因：" + err.Error(),
				"errcode": "1",
				"state":   "创建文件夹" + uploadDir + "失败,原因：" + err.Error(),
			}, isEditor)
			return
		}

	}
	f.Close()
	file, fs, err := this.Ctx.FormFile("filedata")
	if err != nil {
		uploadAjax(this.Ctx, map[string]string{
			"errmsg":  "打开上传临时文件失败 : " + err.Error(),
			"state":   "打开上传临时文件失败 : " + err.Error(),
			"errcode": "1",
		}, isEditor)
		return
	}
	defer file.Close()
	fname := fs.Filename
	info := strings.Split(fname, ".")
	ext := strings.ToLower(info[len(info) - 1])
	canUpload := []string{"jpg", "jpeg", "png"}
	flag := false
	for _, v := range canUpload {
		if v == ext {
			flag = true
		}
	}
	if !flag {
		uploadAjax(this.Ctx, map[string]string{
			"errmsg":  "不支持的文件类型",
			"state":   "不支持的文件类型",
			"errcode": "1",
		}, isEditor)
		return
	}
	filename := string(helper.Krand(10, 3)) + "." + ext
	tofile := uploadDir + "/" + filename
	out, err := os.OpenFile(tofile, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		uploadAjax(this.Ctx, map[string]string{
			"errmsg":  "创建上传文件失败 : " + err.Error(),
			"state":   "创建上传文件失败 : " + err.Error(),
			"errcode": "1",
		}, isEditor)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		uploadAjax(this.Ctx, map[string]string{
			"errmsg":  "上传失败:" + err.Error(),
			"state":   "上传失败:" + err.Error(),
			"errcode": "1",
		}, isEditor)
		return
	}

	resJson := map[string]string{
		"originalName": fname, //原始名称
		"name":         filename, //新文件名称
		"url":          "/upload/" + mid + "/" + nowTime + "/" + filename, //完整文件名,即从当前配置目录开始的URL
		"size":         "", //文件大小
		"type":         "image/jpeg", //文件类型
		"state":        "上传成功", //上传状态
		"errmsg":       "/upload/" + mid + "/" + nowTime + "/" + filename,
		"errcode":      "0",
	}
	uploadAjax(this.Ctx, resJson, isEditor)
	return

}

//生成验证码
func (this *PublicController) VerifyCode() {
	cpt := captcha.New()
	fontPath := helper.GetRootPath() + "/resources/fonts/comic.ttf"
	// 设置字体
	cpt.SetFont(fontPath)
	// 返回验证码图像对象以及验证码字符串 后期可以对字符串进行对比 判断验证
	this.Ctx.ContentType("img/png")
	img, str := cpt.Create(1, captcha.ALL)

	this.Session.Set("verify_code", str)
	png.Encode(this.Ctx.ResponseWriter(), img) //发送图片内容到浏览器
}


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
