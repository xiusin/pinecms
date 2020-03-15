package backend

import (
	"encoding/base64"
	"fmt"
	"github.com/xiusin/pine/cache"
	"image/png"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/config"
	"github.com/xiusin/pine"

	"github.com/afocus/captcha"
	"github.com/xiusin/pinecms/src/common/helper"

	"github.com/go-xorm/xorm"
)

type PublicController struct {
	pine.Controller
}

func (c *PublicController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY( "/upload", "Upload")
	b.ANY( "/fedir-scan", "FeDirScan")
	b.ANY( "/attachments", "Attachments")
	b.ANY( "/ueditor", "UEditor")
	b.ANY( "/verify-code", "VerifyCode")
	b.ANY( "/todos", "TODO")
}

func (c *PublicController) FeDirScan() {
	c.Ctx().Render().JSON(helper.ScanDir(config.AppConfig().View.FeDirname))
}

//上传图片
func (c *PublicController) Upload() {
	isEditor := true
	settingData := c.Ctx().Value(controllers.CacheSetting).(map[string]string)
	mid := c.Ctx().GetString("mid")
	// mid 用户ID
	if mid == "" {
		mid = "public"
	}
	uploader := getStorageEngine(settingData)
	uploadDir := fmt.Sprintf("%s/%s", mid, helper.NowDate("Ymd"))
	file, fs, err := c.Ctx().Files("filedata")
	if err != nil {
		if fileData := c.Ctx().FormValue("filedata"); fileData == "" {
			pine.Logger().Error("上传文件失败", err.Error())
			uploadAjax(c.Ctx(), map[string]interface{}{"state": "打开上传临时文件失败 : " + err.Error(), "errcode": "1",}, isEditor)
			return
		} else {
			dist, err := base64.StdEncoding.DecodeString(fileData)
			if err != nil {
				uploadAjax(c.Ctx(), map[string]interface{}{"state": "解码base64数据失败 : " + err.Error(), "errcode": "1",}, isEditor)
				return
			}
			f, err := ioutil.TempFile("", "tempfile_"+strconv.Itoa(rand.Intn(10000)))
			if err != nil {
				uploadAjax(c.Ctx(), map[string]interface{}{"state": "上传失败 : " + err.Error(), "errcode": "1",}, isEditor)
				return
			}
			_, _ = f.Write(dist)
			_ = f.Close()
			fo, _ := os.Open(f.Name())
			file = multipart.File(fo)
			defer os.Remove(f.Name())
		}
	}
	defer file.Close()
	var fname string
	var size int64
	if fs != nil {
		size = fs.Size
		fname = fs.Filename
	} else {
		fname = helper.GetRandomString(10) + ".png"
	}

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
		uploadAjax(c.Ctx(), map[string]interface{}{"state": "不支持的文件类型", "errcode": "1",}, isEditor)
		return
	}
	filename := string(helper.Krand(10, 3)) + "." + ext
	storageName := uploadDir + "/" + filename
	path, err := uploader.Upload(storageName, file)
	if err != nil {
		uploadAjax(c.Ctx(), map[string]interface{}{"state": "上传失败:" + err.Error(), "errcode": "1"}, isEditor)
		return
	}
	resJson := map[string]interface{}{
		"originalName": fname,     //原始名称
		"name":         filename,  //新文件名称
		"url":          path,      //完整文件名,即从当前配置目录开始的URL
		"size":         size,      //文件大小
		"type":         "." + ext, //文件类型
		"state":        "SUCCESS", //上传状态
		"errmsg":       path,
		"errcode":      "0",
	}
	if id, _ := c.Ctx().Value("orm").(*xorm.Engine).InsertOne(&tables.Attachments{
		Name:       filename,
		Url:        path,
		OriginName: fname,
		Size:       size,
		UploadTime: time.Now(),
		Type:       models.IMG_TYPE,
	}); id > 0 {
		uploadAjax(c.Ctx(), resJson, isEditor)
	} else {
		os.Remove(storageName)
		uploadAjax(c.Ctx(), map[string]interface{}{"state":   "保存上传失败", "errcode": "1"}, isEditor)
	}

}

////生成验证码
func (c *PublicController) VerifyCode() {
	cpt := captcha.New()
	cpt.SetFont(helper.GetRootPath() + "/resources/fonts/comic.ttf")
	img, str := cpt.Create(4, captcha.ALL)
	c.Session().AddFlush("verify_code", str)
	c.Ctx().Writer().Header().Set("Content-type","img/png")
	png.Encode(c.Ctx().Writer(), img)
}

func (c *PublicController) UEditor() {
	action := c.Ctx().URLParam("action")
	switch action {
	case "config":
		c.Ctx().WriteString(`
{
   "imageActionName": "upload",
   "imageFieldName": "filedata",
   "imageMaxSize": 2048000,
   "imageAllowFiles": [".png", ".jpg", ".jpeg", ".gif", ".bmp"],
   "imageCompressEnable": true,
   "imageCompressBorder": 1600,
   "imageInsertAlign": "none",
   "imageUrlPrefix": "",
   "scrawlActionName": "upload",
   "scrawlFieldName": "filedata",
   "scrawlMaxSize": 2048000,
   "scrawlUrlPrefix": "",
   "scrawlInsertAlign": "none",
   "catcherLocalDomain": ["127.0.0.1", "localhost", "img.baidu.com"],
   "catcherActionName": "catchimage",
   "catcherFieldName": "source",
   "catcherPathFormat": "/ueditor/php/upload/image/{yyyy}{mm}{dd}/{time}{rand:6}",
   "catcherUrlPrefix": "",
   "catcherMaxSize": 2048000,
   "catcherAllowFiles": [".png", ".jpg", ".jpeg", ".gif", ".bmp"],
   "imageManagerActionName": "attachments-img",
   "imageManagerUrlPrefix": "",
   "imageManagerInsertAlign": "none",
   "imageManagerAllowFiles": [".png", ".jpg", ".jpeg", ".gif", ".bmp"],
   "fileManagerActionName": "attachments-file",
   "fileManagerListPath": "/ueditor/php/upload/file/",
   "fileManagerUrlPrefix": "",
   "fileManagerListSize": 20,
   "fileManagerAllowFiles": [
       ".png", ".jpg", ".jpeg", ".gif", ".bmp",
       ".flv", ".swf", ".mkv", ".avi", ".rm", ".rmvb", ".mpeg", ".mpg",
       ".ogg", ".ogv", ".mov", ".wmv", ".mp4", ".webm", ".mp3", ".wav", ".mid",
       ".rar", ".zip", ".tar", ".gz", ".7z", ".bz2", ".cab", ".iso",
       ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".pdf", ".txt", ".md", ".xml"
   ]
}
`)
	case "upload":
		c.Upload()
	case "attachments-img":
		c.Ctx().Request().ParseForm()
		c.Ctx().Request().Form.Add("type", models.IMG_TYPE)
		c.Attachments()
	}
}

func uploadAjax(ctx *pine.Context, uploadData map[string]interface{}, isEditor bool) {
	uploadData["errmsg"] = uploadData["state"]
	ctx.Render().JSON(uploadData)
}

// 读取资源列表
func (c *PublicController) Attachments() {
	page, _ := c.Ctx().URLParamInt64("page")
	if page < 1 {
		page = 1
	}
	start, _ := c.Ctx().URLParamInt64("start")
	if start < 0 {
		start = 0
	}
	var data []*tables.Attachments
	attachmentType := c.Ctx().GetString("type", models.IMG_TYPE)
	cnt, _ := c.Ctx().Value("orm").(*xorm.Engine).Limit(30, int(start)).Where("`type` = ?", attachmentType).FindAndCount(&data)
	c.Ctx().Render().JSON(map[string]interface{}{
		"state":   "SUCCESS",
		"list":    data,
		"total":   cnt,
		"start":   start,
		"errmsg":  "读取成功",
		"errcode": "0",
	})
}

// TODO
func (c *PublicController) TODO(icache cache.ICache) {
	todos := c.Ctx().FormValue("todos")
	icache.Set("backend_todos", []byte(todos))
	c.Ctx().Render().JSON("success")
}