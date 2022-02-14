package filemanager

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/xiusin/pine"

	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/controllers/backend/filemanager/tables"
	"github.com/xiusin/pinecms/src/common/storage"
)

type FileManagerController struct {
	backend.BaseController
	User      *tables.FileManagerAccount
	engine    storage.Uploader
	path      string
	urlPrefix string
	disk      string
}

func (c *FileManagerController) Construct() {
	c.Table = &tables.FileManagerAccount{}
	c.BaseController.Construct()
	if !strings.HasSuffix(c.Ctx().Path(), "/me") && !strings.HasSuffix(c.Ctx().Path(), "/login") {
		c.User = &tables.FileManagerAccount{}
		loginId := c.Session().Get("loginId")
		ok, _ := c.Orm.Where("id = ?", loginId).Get(c.User)
		if !ok {
			c.Session().Destroy()
			panic(fmt.Errorf("登陆用户无法匹配"))
		}
		c.engine = GetUserUploader(c.User)
		if c.engine == nil {
			panic(fmt.Errorf("用户配置的存储引擎不存在"))
		}
		c.path, _ = c.Input().GetString("path", "")
		c.disk, _ = c.Input().GetString("disk", "public")
		c.urlPrefix = "/uploads/"
		c.path = strings.TrimPrefix(c.path, c.urlPrefix)
	}
}

func (c *FileManagerController) GetMe() {
	if isLogin := c.Session().Get("isLogin"); isLogin != Logined {
		c.Render().JSON(pine.H{"result": ResResult{Status: "success"}, "isLogin": false})
		return
	}

	c.Render().JSON(pine.H{"result": ResResult{Status: "success"}, "isLogin": true, "nickname": c.Session().Get("nickname")})
}

func (c *FileManagerController) PostLogin() {
	username, _ := c.Input().GetString("username")
	pwd, _ := c.Input().GetString("pwd")

	if ok, _ := c.Orm.Where("username = ?", username).Get(c.Table); !ok {
		ResponseError(c.Ctx(), "userNotExist")
		return
	}

	if c.Table.(*tables.FileManagerAccount).GetMd5Pwd(pwd) != c.Table.(*tables.FileManagerAccount).Password {
		ResponseError(c.Ctx(), "passwordError")
		return
	}

	c.Session().Set("isLogin", Logined)
	c.Session().Set("nickname", c.Table.(*tables.FileManagerAccount).Nickname)
	c.Session().Set("loginId", fmt.Sprintf("%d", c.Table.(*tables.FileManagerAccount).Id))

	c.Render().JSON(pine.H{
		"result":   ResResult{Status: "success", Message: "logined"},
		"isLogin":  true,
		"nickname": c.Table.(*tables.FileManagerAccount).Nickname,
	})
}

func (c *FileManagerController) GetLogout() {
	c.Session().Destroy()
	c.Render().JSON(pine.H{"result": ResResult{Status: "success", Message: "logouted"}})
}

func (c *FileManagerController) GetInitialize() {
	c.Render().ContentType(pine.ContentTypeJSON)

	c.Render().Text(`{
  "result": {
    "status": "success",
    "message": null
  },
  "config": {
    "acl": true,
    "leftDisk": null,
    "leftPath": null,
    "rightDisk": null,
    "rightPath": null,
    "disks": {
      "public": {
        "public": "",
        "driver": "public"
      },
      "private": {
        "private": "",
        "driver": "private"
      }
    }
  }
}`)
}

func (c *FileManagerController) GetTree() {
	l, _ := c.engine.List(c.path)
	dirs, files := c._formatList(l)
	c.Render().JSON(pine.H{"result": ResResult{Status: "success"}, "directories": dirs, "files": files})
}

func (c *FileManagerController) GetContent() {
	c.GetTree()
}

func (c *FileManagerController) PostSelectDisk() {}

func (c *FileManagerController) GetDownloadFile() {
	c.Ctx().Response.Header.Set("Content-Disposition", "attachment")
	url := c.engine.GetFullUrl(c.path)
	if strings.Contains(url, "?") {
		url += "&fmq=" + DownloadFlag
	} else {
		url += "?fmq=" + DownloadFlag
	}
	c.Render().Text(url)
}

func (c *FileManagerController) GetDownload() {

	if content, err := c.engine.Content(c.path); err != nil {
		ResponseError(c.Ctx(), err.Error())
	} else {
		_ = c.Render().Bytes(content)
	}
}

func (c *FileManagerController) PostUpdateFile() {
	fs, err := c.Ctx().FormFile("file")
	if err != nil {
		ResponseError(c.Ctx(), err.Error())
		return
	}
	f, err := fs.Open()
	if err != nil {
		ResponseError(c.Ctx(), err.Error())
		return
	}
	defer f.Close()
	storageName := filepath.Join(c.path, fs.Filename)
	if _, err := c.engine.Upload(storageName, f); err != nil {
		ResponseError(c.Ctx(), err.Error())
		return
	}

	c.Render().JSON(pine.H{"result": ResResult{Status: "success", Message: "updated"}, "file": &FMFile{
		Size:      int(fs.Size),
		Basename:  fs.Filename,
		Filename:  strings.ReplaceAll(filepath.Base(fs.Filename), filepath.Ext(fs.Filename), ""),
		Dirname:   strings.TrimLeft(strings.TrimPrefix(filepath.Dir(storageName), strings.TrimSuffix(c.urlPrefix, "/")), "/"),
		Path:      strings.TrimLeft(strings.TrimPrefix(strings.ReplaceAll(storageName, "\\", "/"), strings.TrimSuffix(c.urlPrefix, "/")), "/"),
		Timestamp: time.Now().Unix(),
		Extension: strings.TrimLeft(filepath.Ext(fs.Filename), "."),
		Props:     FMFileProps{},
	}})
}

func (c *FileManagerController) GetThumbnailsLink() {
	_ = c.Render().Text(c.engine.GetFullUrl(c.path))
}

func (c *FileManagerController) GetPreview() {

}

func (c *FileManagerController) GetStreamFile() {

}

func (c *FileManagerController) GetUrl() {
	c.Render().JSON(pine.H{"result": ResResult{Status: "success"}, "url": c.engine.GetFullUrl(c.path)})
}

func (c *FileManagerController) PostCreateFile() {
	name, _ := c.Input().GetString("name")
	f, _ := os.CreateTemp("", "")
	defer f.Close()
	storageName := filepath.Join(c.path, name)
	if _, err := c.engine.Upload(storageName, f); err != nil {
		ResponseError(c.Ctx(), err.Error())
		return
	}

	c.Render().JSON(pine.H{"result": ResResult{Status: "success", Message: "fileCreated"}, "file": &FMFile{
		Size:      0,
		Basename:  name,
		Filename:  strings.ReplaceAll(filepath.Base(name), filepath.Ext(name), ""),
		Dirname:   strings.TrimLeft(strings.TrimPrefix(filepath.Dir(storageName), strings.TrimSuffix(c.urlPrefix, "/")), "/"),
		Path:      strings.ReplaceAll(strings.TrimLeft(strings.TrimPrefix(storageName, strings.TrimSuffix(c.urlPrefix, "/")), "/"), "\\", "/"),
		Timestamp: time.Now().Unix(),
		Extension: strings.TrimLeft(filepath.Ext(name), "."),
		Props:     FMFileProps{},
	}})
}

func (c *FileManagerController) PostCreateDirectory() {
	cmp := regexp.MustCompile(`[\\\\/:*?"<>|]`)
	name, _ := c.Input().GetString("name")
	if cmp.MatchString(name) {
		ResponseError(c.Ctx(), "目录创建失败,含有非法字符有\\/:*?\"<>|")
		return
	}
	if err := c.engine.Mkdir(name); err != nil {
		ResponseError(c.Ctx(), err.Error())
		return
	}
	c.Render().JSON(pine.H{"result": ResResult{Status: "success", Message: "目录创建成功"}})

}

func (c *FileManagerController) PostDelete() {
	var items []DelItem
	byts, _ := json.Marshal(c.Input().Get("items"))
	if err := json.Unmarshal(byts, &items); err != nil {
		ResponseError(c.Ctx(), err.Error())
		return
	}

	hasErr := false
	for _, item := range items {
		if item.Type == "dir" {
			if err := c.engine.Rmdir(item.Path); err != nil {
				c.Logger().Debug(err)
				hasErr = true
			}
		} else {
			if err := c.engine.Remove(item.Path); err != nil {
				c.Logger().Debug(err)
				hasErr = true
			}
		}
	}
	if hasErr {
		ResponseError(c.Ctx(), "部分删除失败")
	} else {
		c.Render().JSON(pine.H{"result": ResResult{Status: "success", Message: "deleted"}})
	}
}

func (c *FileManagerController) PostPaste() {
	typ, _ := c.Input().GetString("type")
	switch typ {
	case "copy":

	}
}

func (c *FileManagerController) PostRename() {
	oldname, _ := c.Input().GetString("oldName")
	newname, _ := c.Input().GetString("newName")

	cmp := regexp.MustCompile(`[\\\\/:*?"<>|]`)

	if cmp.MatchString(newname) {
		ResponseError(c.Ctx(), "重命名失败,含有非法字符有\\/:*?\"<>|")
		return
	}

	if err := c.engine.Rename(strings.TrimPrefix(oldname, c.urlPrefix), strings.TrimPrefix(newname, c.urlPrefix)); err != nil {
		ResponseError(c.Ctx(), err.Error())
		return
	}
	c.Render().JSON(pine.H{"result": ResResult{Status: "success", Message: "renamed"}})
}

func (c *FileManagerController) PostUpload() {
	fs, err := c.Ctx().FormFile("files[]")
	if err != nil {
		ResponseError(c.Ctx(), err.Error())
		return
	}
	storageName := filepath.Join(c.path, fs.Filename)
	if overwrite, _ := c.Input().GetBool("overwrite"); !overwrite {
		if exist, err := c.engine.Exists(storageName); err != nil {
			ResponseError(c.Ctx(), err.Error())
			return
		} else if exist {
			ResponseError(c.Ctx(), "文件已存在")
			return
		}
	}

	f, err := fs.Open()
	if err != nil {
		ResponseError(c.Ctx(), err.Error())
		return
	}
	defer f.Close()
	if _, err := c.engine.Upload(storageName, f); err != nil {
		ResponseError(c.Ctx(), err.Error())
		return
	}
	c.Render().JSON(pine.H{"result": ResResult{Status: "success", Message: "uploaded"}})
}

func (c *FileManagerController) _formatList(fileList []storage.File) (directories []FMFile, files []FMFile) {
	directories, files = []FMFile{}, []FMFile{}
	for _, file := range fileList {
		file.FullPath = strings.ReplaceAll(file.FullPath, "\\", "/")
		f := FMFile{
			Basename:  file.Name,
			Filename:  strings.ReplaceAll(filepath.Base(file.Name), filepath.Ext(file.Name), ""),
			Dirname:   strings.TrimLeft(strings.TrimPrefix(strings.ReplaceAll(filepath.Dir(file.FullPath), "\\", "/"), strings.TrimSuffix(c.urlPrefix, "/")), "/"),
			Path:      strings.TrimLeft(strings.TrimPrefix(strings.ReplaceAll(file.FullPath, "\\", "/"), strings.TrimSuffix(c.urlPrefix, "/")), "/"),
			Timestamp: file.Ctime.Unix(),
			Size:      int(file.Size),
			Extension: strings.TrimLeft(filepath.Ext(file.Name), "."),
			Props:     FMFileProps{},
		}
		if file.IsDir {
			f.Type = "dir"
			directories = append(directories, f)
		} else {
			f.Type = "file"
			files = append(files, f)
		}
	}
	return
}
