package backend

import (
	"github.com/afocus/captcha"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/cache"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/common/helper"
	"image/png"
	"path"
)

type PublicController struct {
	BaseController
}

func (c *PublicController) GetMenu() {
	menus := models.NewMenuModel().GetAll() //获取menuid内容
	helper.Ajax(pine.H{"menus": menus}, 0, c.Ctx())
}

func (c *PublicController) PostUpload() {
	//isEditor := true
	//settingData := c.Ctx().Value(controllers.CacheSetting).(map[string]string)
	//mid := c.Ctx().GetString("mid")
	//// mid 用户ID
	//if mid == "" {
	//	mid = "public"
	//}
	//// 判断上传类型
	//uploadType := models.FILE_TYPE
	//t := c.Ctx().FormValue("type")
	//
	//if t != "" && strings.Contains(t, "image") {
	//	uploadType = models.IMG_TYPE
	//}
	//
	//uploader := getStorageEngine(settingData)
	//uploadDir := fmt.Sprintf("%s/%s", mid, helper.NowDate("Ymd"))
	//
	//mf, err := c.Ctx().MultipartForm()
	//if err != nil {
	//	uploadAjax(c.Ctx(), map[string]interface{}{"state": "打开上传临时文件失败 : " + err.Error(), "errcode": "1"}, isEditor)
	//	return
	//}
	//fss, ok := mf.File["file"]
	//if !ok {
	//	uploadAjax(c.Ctx(), map[string]interface{}{"state": "打开上传临时文件失败", "errcode": "1"}, isEditor)
	//	return
	//}
	//fs := fss[0]
	//var fname string
	//var size int64
	//if fs != nil {
	//	size = fs.Size
	//	fname = fs.Filename
	//} else {
	//	fname = helper.GetRandomString(10) + ".png" // 涂鸦上传
	//}
	//
	//info := strings.Split(fname, ".")
	//ext := strings.ToLower(info[len(info)-1])
	//flag := false
	//
	//if uploadType != models.FILE_TYPE {
	//	canUpload := []string{"jpg", "jpeg", "png"}
	//	for _, v := range canUpload {
	//		if v == ext {
	//			flag = true
	//		}
	//	}
	//} else {
	//	flag = true
	//}
	//if !flag {
	//	uploadAjax(c.Ctx(), map[string]interface{}{"state": "不支持的文件类型", "errcode": "1"}, isEditor)
	//	return
	//}
	//filename := string(helper.Krand(10, 3)) + "." + ext
	//storageName := uploadDir + "/" + filename
	//f, err := fs.Open()
	//var path string
	//if err == nil {
	//	defer f.Close()
	//	path, err = uploader.Upload(storageName, f)
	//}
	//if err != nil {
	//	uploadAjax(c.Ctx(), map[string]interface{}{"state": "上传失败:" + err.Error(), "errcode": "1"}, isEditor)
	//	return
	//}
	//resJson := map[string]interface{}{
	//	"originalName": fname,     //原始名称
	//	"name":         filename,  //新文件名称
	//	"url":          path,      //完整文件名,即从当前配置目录开始的URL
	//	"size":         size,      //文件大小
	//	"type":         "." + ext, //文件类型
	//	"state":        "SUCCESS", //上传状态
	//	"errmsg":       path,
	//	"errcode":      "0",
	//	"value":        path, // 给amsi使用
	//}
	//if id, _ := c.Ctx().Value("orm").(*xorm.Engine).InsertOne(&tables.Attachments{
	//	Name:       filename,
	//	Url:        path,
	//	OriginName: fname,
	//	Size:       size,
	//	Type:       uploadType,
	//}); id > 0 {
	//	helper.Ajax(resJson, 0, c.Ctx())
	//} else {
	//	os.Remove(storageName)
	//	uploadAjax(c.Ctx(), map[string]interface{}{"state": "保存上传失败", "errcode": "1"}, isEditor)
	//}
}

func (c *PublicController) GetVerifyCode(cacher cache.AbstractCache) {
	cpt := captcha.New()
	cpt.SetFont(path.Join(helper.GetRootPath(), "resources/fonts/comic.ttf"))
	img, str := cpt.Create(4, captcha.ALL)
	c.Session().AddFlush(controllers.CacheVerifyCode, str)
	c.Ctx().SetContentType("img/png")
	png.Encode(c.Ctx().Response.BodyWriter(), img)
}
