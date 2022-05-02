package backend

import (
	"crypto/md5"
	"fmt"
	"io"
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
	// 基于用户过滤菜单
	roleId := c.Ctx().Value("roleid")
	role := &tables.AdminRole{}
	// 获取对于权限配置的菜单
	c.Orm.In("id", roleId).Cols("menu_ids").Get(role)
	menus := models.NewMenuModel().GetAll(role.MenuIdList)
	var perms = map[string]struct{}{}
	for _, menu := range menus {
		permArr := strings.Split(menu.Perms, ",")
		for _, s := range permArr {
			if len(s) > 0 {
				perms[s] = struct{}{}
			}
		}
	}
	permsValues := []string{}
	for i := range perms {
		permsValues = append(permsValues, i)
	}
	helper.Ajax(pine.H{"menus": menus, "perms": permsValues}, 0, c.Ctx())
}

func (c *PublicController) PostUpload() {
	cfg, _ := config.SiteConfig()
	uploader, uploadDir := getStorageEngine(cfg), helper.NowDate("20060102")
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
		md5hash := md5.New()

		io.Copy(md5hash, f) // 不能使用readAll 会读空buffer
		f.Seek(0, 0)        // 读取文件内容后指针会保留在最后一位, 需要seek到首行

		md5sum := fmt.Sprintf("%x", md5hash.Sum(nil))
		attach := &tables.Attachments{}
		c.Orm.Where("md5 = ?", md5sum).Get(attach)
		resJson := map[string]interface{}{"originalName": fs.Filename, "size": fs.Size, "md5": md5sum}
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

func (c *PublicController) GetApidoc() {
	baseHost := string(c.Ctx().URI().Scheme()) + "://" + string(c.Ctx().Host())
	helper.Ajax(baseHost+"/debug/statsviz/", 0, c.Ctx())
}
