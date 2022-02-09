package filemanager

import (
	"encoding/json"
	"fmt"
	"github.com/xiusin/pine"
	"path/filepath"
	"strings"

	"github.com/xiusin/pinecms/src/application/controllers/backend"
	"github.com/xiusin/pinecms/src/application/controllers/backend/filemanager/tables"
	"github.com/xiusin/pinecms/src/common/storage"
)

type FileManagerController struct {
	backend.BaseController
	User   *tables.FileManagerAccount
	engine storage.Uploader
	path   string
}

func (c *FileManagerController) Construct() {
	InitInstall()
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
		c.path = strings.TrimPrefix(c.path, "/uploads/")
	}
}

func (c *FileManagerController) GetMe() {
	if isLogin := c.Session().Get("isLogin"); isLogin != Logined {
		c.Render().JSON(pine.H{
			"result": ResResult{
				Status: "success",
			},
			"isLogin": false,
		})
		return
	}

	c.Render().JSON(pine.H{
		"result": ResResult{
			Status: "success",
		},
		"isLogin":  true,
		"nickname": c.Session().Get("nickname"),
	})
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
		"result": ResResult{
			Status:  "success",
			Message: "logined",
		},
		"isLogin":  true,
		"nickname": c.Table.(*tables.FileManagerAccount).Nickname,
	})
}

func (c *FileManagerController) PostLogout() {
	c.Session().Destroy()
	c.Render().JSON(pine.H{
		"result": ResResult{
			Status:  "success",
			Message: "logouted",
		},
	})
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
	c.Render().JSON(pine.H{
		"result": ResResult{
			Status: "success",
		},
		"directories": dirs,
		"files":       files,
	})
}

func (c *FileManagerController) GetContent() {
	c.GetTree()
}

func (c *FileManagerController) PostFileAuthor() {

}

func (c *FileManagerController) PostSelectDisk() {

}

func (c *FileManagerController) PostDownloadFile() {

}

func (c *FileManagerController) GetDownload() {
	content, err := c.engine.Content(c.path)
	if err != nil {
		ResponseError(c.Ctx(), err.Error())
		return
	}
	_ = c.Render().Bytes(content)
}

func (c *FileManagerController) PostZip() {

}

func (c *FileManagerController) PostUnzip() {

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
	if _, err := c.engine.Upload(strings.TrimPrefix(c.path, "/uploads/"), f); err != nil {
		ResponseError(c.Ctx(), err.Error())
		return
	}
	c.Render().JSON(pine.H{"result": ResResult{Status: "success", Message: "updated"}})
}

func (c *FileManagerController) GetThumbnailsLink() {

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
	file, _ := c.Ctx().FormFile("file")
	f, err := file.Open()
	if err != nil {
		ResponseError(c.Ctx(), err.Error())
		return
	}
	defer f.Close()

	path, err := c.engine.Upload(name, f)
	if err != nil {
		ResponseError(c.Ctx(), err.Error())
		return
	}
	finfo, err := c.engine.Info(path)
	if err != nil {
		ResponseError(c.Ctx(), err.Error())
		return
	}

	c.Render().JSON(pine.H{
		"result": ResResult{
			Status:  "success",
			Message: "fileCreated",
		},
		"file": finfo,
	})
}

func (c *FileManagerController) PostCreateDirectory() {

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
		if err := c.engine.Remove(item.Path); err != nil {
			hasErr = true
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

	if err := c.engine.Rename(strings.TrimPrefix(oldname, "/uploads/"), strings.TrimPrefix(newname, "/uploads/")); err != nil {
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
			ID:        nil,
			Basename:  file.Name,
			Filename:  filepath.Base(file.Name),
			Dirname:   strings.ReplaceAll(filepath.Dir(file.FullPath), "\\", "/"),
			Path:      file.FullPath,
			ParentID:  "",
			Timestamp: file.Ctime.Second() * 1000,
			ACL:       0,
			Size:      int(file.Size),
			Extension: strings.TrimLeft(filepath.Ext(file.Name), "."),
			Props:     FMFileProps{},
			Author:    "",
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
