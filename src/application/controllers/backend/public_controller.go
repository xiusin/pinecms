package backend

import (
	"path/filepath"
	"strings"

	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/application/models"
	"github.com/xiusin/pinecms/src/application/models/tables"
	"github.com/xiusin/pinecms/src/common/captcha"
	"github.com/xiusin/pinecms/src/common/helper"
	"github.com/xiusin/pinecms/src/config"
)

type PublicController struct {
	BaseController
}

func (c *PublicController) GetMenu() {
	menus := models.NewMenuModel().GetAll()
	helper.Ajax(pine.H{"menus": menus}, 0, c.Ctx())
}

func (c *PublicController) PostUpload() {
	cfg, _ := config.SiteConfig()
	uploader, uploadDir := getStorageEngine(cfg), helper.NowDate("Ymd")
	mf, err := c.Ctx().MultipartForm()
	if err != nil {
		c.Logger().Error("上传文件失败", err)
		helper.Ajax("打开上传临时文件失败", 1, c.Ctx())
		return
	}

	if fss, ok := mf.File["file"]; !ok {
		helper.Ajax("打开上传临时文件失败", 1, c.Ctx())
	} else {
		fs := fss[0]
		f, err := fs.Open()
		if err != nil {
			c.Logger().Error("上传失败", err)
			helper.Ajax("上传失败:"+err.Error(), 1, c.Ctx())
			return
		}
		defer f.Close()
		//md5hash := md5.New()
		//
		//dat,_ := io.ReadAll(f)
		//_, _ = md5hash.Write(dat)
		//md5Byts := fmt.Sprintf("%x", md5hash.Sum(nil))
		//
		md5Byts := ""
		attach := &tables.Attachments{}
		//c.Orm.Where("md5 = ?", md5Byts).Get(attach)
		resJson := map[string]interface{}{"originalName": fs.Filename, "size": fs.Size, "md5": md5Byts}
		if len(attach.Url) == 0 {
			filename := string(helper.Krand(16, 3)) + strings.ToLower(filepath.Ext(fs.Filename))
			storageName := uploadDir + "/" + filename
			path, err := uploader.Upload(storageName, f)
			if err != nil {
				helper.Ajax(err, 1, c.Ctx())
				return
			}
			attach.Name = filename
			attach.Url = path
		}
		resJson["name"] = attach.Name
		resJson["url"] = attach.Url
		helper.Ajax(resJson, 0, c.Ctx())
	}
}

func (c *PublicController) GetCaptcha() {
	id, base64s, err := captcha.Get("")
	if err != nil {
		helper.Ajax(err, 1, c.Ctx())
		return
	}
	helper.Ajax(pine.H{"captchaId": id, "data": base64s}, 0, c.Ctx())
}

func (c *PublicController) GetPprof() {
	baseHost := string(c.Ctx().URI().Scheme()) + "://" + string(c.Ctx().Host())
	helper.Ajax(baseHost+"/debug/pprof/", 0, c.Ctx())
}

func (c *PublicController) GetStatsviz() {
	baseHost := string(c.Ctx().URI().Scheme()) + "://" + string(c.Ctx().Host())
	helper.Ajax(baseHost+"/debug/statsviz/", 0, c.Ctx())
}
