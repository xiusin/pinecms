package backend

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image/png"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/config"

	"github.com/afocus/captcha"
	"github.com/xiusin/pinecms/src/common/helper"

	"github.com/go-xorm/xorm"
)

type PublicController struct {
	pine.Controller
}

func (c *PublicController) RegisterRoute(b pine.IRouterWrapper) {
	b.ANY("/upload", "Upload")
	b.ANY("/fedir-scan", "FeDirScan")
	b.ANY("/attachments", "Attachments")
	b.ANY("/ueditor", "UEditor")
	b.ANY("/verify-code", "VerifyCode")
	b.ANY("/todo", "TODO")
}

func (c *PublicController) FeDirScan() {
	c.Ctx().Render().JSON(helper.ScanDir(path.Join(config.AppConfig().View.FeDirname, config.AppConfig().View.Theme)))
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
	// 判断上传类型
	uploadType := models.IMG_TYPE
	id := c.Ctx().FormValue("id")

	if id != "" && strings.Contains(id, "_FILE_") {
		uploadType = models.FILE_TYPE
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
			// 涂鸦上传
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
		fname = helper.GetRandomString(10) + ".png" // 涂鸦上传
	}

	info := strings.Split(fname, ".")
	ext := strings.ToLower(info[len(info)-1])
	flag := false

	if uploadType != models.FILE_TYPE {
		canUpload := []string{"jpg", "jpeg", "png"}
		for _, v := range canUpload {
			if v == ext {
				flag = true
			}
		}
	} else {
		flag = true
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
		Type:       uploadType,
	}); id > 0 {
		uploadAjax(c.Ctx(), resJson, isEditor)
	} else {
		os.Remove(storageName)
		uploadAjax(c.Ctx(), map[string]interface{}{"state": "保存上传失败", "errcode": "1"}, isEditor)
	}

}

func (c *PublicController) VerifyCode() {
	cpt := captcha.New()
	cpt.SetFont(path.Join(helper.GetRootPath(), "resources/fonts/comic.ttf"))

	img, str := cpt.Create(4, captcha.ALL)

	c.Session().AddFlush("verify", str)

	c.Ctx().Writer().Header().Set("Content-type", "img/png")

	png.Encode(c.Ctx().Writer(), img)
}

func (c *PublicController) UEditor() {
	action := c.Ctx().URLParam("action")

	switch action {
	case "config":
		// 读取上传配置文件
		conf, err := config.SiteConfig()
		if err != nil {
			c.Logger().Error(err)
			return
		}
		uploadFileSize, _ := strconv.Atoi(conf["UPLOAD_MAX_SIZE"])
		if uploadFileSize == 0 {
			c.Logger().Error("请配置上传文件大小")
			return
		}

		imageFileExt := strings.Split(conf["UPLOAD_IMG_TYPES"], "|")
		for k := range imageFileExt {
			imageFileExt[k] = "." + strings.Trim(imageFileExt[k], ".")
		}
		imageFileExtByte, _ := json.Marshal(imageFileExt)
		if len(imageFileExtByte) == 0 {
			imageFileExtByte = []byte("[]")
		}

		acctFileExt := strings.Split(conf["UPLOAD_ACCT_TYPES"], "|")
		for k := range acctFileExt {
			acctFileExt[k] = "." + strings.Trim(acctFileExt[k], ".")
		}
		acctFileExtByte, _ := json.Marshal(acctFileExt)
		if len(acctFileExtByte) == 0 {
			acctFileExtByte = []byte("[]")
		}

		uploadSizeByte := uploadFileSize * 1024 * 1000

		c.Ctx().WriteString(`
{
   "imageActionName": "upload",
   "imageFieldName": "filedata",
   "imageMaxSize": ` + strconv.Itoa(uploadSizeByte) + `,
   "imageAllowFiles": ` + string(imageFileExtByte) + `,
   "imageCompressEnable": true,
   "imageCompressBorder": 1600,
   "imageUrlPrefix": "",
   "scrawlActionName": "upload",
   "scrawlFieldName": "filedata",
   "scrawlMaxSize": ` + strconv.Itoa(uploadSizeByte) + `,
   "scrawlUrlPrefix": "",
   "imageManagerActionName": "attachments-img",
   "imageManagerUrlPrefix": "",
   "imageManagerAllowFiles": ` + strconv.Itoa(uploadSizeByte) + `,
   "fileActionName": "upload", 
   "fileFieldName": "filedata", 
   "filePathFormat": "", 
   "fileUrlPrefix": "", 
   "fileMaxSize": ` + strconv.Itoa(uploadSizeByte) + `,
   "fileAllowFiles": ` + string(acctFileExtByte) + `,
   "fileManagerActionName": "attachments-file",
   "fileManagerListPath": "",
   "fileManagerUrlPrefix": "",
   "fileManagerListSize": 20,
   "fileManagerAllowFiles": ` + string(acctFileExtByte) + `
}
`)
	case "upload":
		c.Upload()
	case "attachments-img":
		c.Attachments(models.IMG_TYPE)
	case "attachments-file":
		c.Attachments(models.FILE_TYPE)
	}
}

func uploadAjax(ctx *pine.Context, uploadData map[string]interface{}, isEditor bool) {
	uploadData["errmsg"] = uploadData["state"]
	ctx.Render().JSON(uploadData)
}

// 读取资源列表
func (c *PublicController) Attachments(attachmentType string) {
	page, _ := c.Ctx().URLParamInt64("page")
	if page < 1 {
		page = 1
	}
	start, _ := c.Ctx().URLParamInt64("start")
	if start < 0 {
		start = 0
	}
	var data []*tables.Attachments
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

func (c *PublicController) TODO(orm *xorm.Engine) {
	id, _ := c.Ctx().PostInt64("id")
	if id > 0 {
		af, _ := orm.ID(id).MustCols("status").Update(&tables.Todo{
			Status: false,
		})
		if af > 0 {
			helper.Ajax("success", 0, c.Ctx())
		} else {
			helper.Ajax("failed", 1, c.Ctx())
		}
		return
	}
	todo := c.Ctx().FormValue("todo")
	if todo == "" {
		return
	}
	userid, _ := strconv.Atoi(c.Session().Get("adminid"))
	todoObj := &tables.Todo{
		Message: todo,
		UserID:  int64(userid),
		Status:  true,
	}
	_, err := orm.InsertOne(todoObj)
	if err != nil {
		panic(err)
	}
	helper.Ajax(todoObj.Id, 0, c.Ctx())
}
