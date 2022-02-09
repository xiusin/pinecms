package tables

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/xiusin/pinecms/src/common/helper"
)

type FileManagerAccount struct {
	Id       int64  `json:"id"`
	Nickname string `json:"nickname" xorm:"nickname"`
	Username string `json:"username" xorm:"username"`
	Password string `json:"pwd" xorm:"pwd"`
	Salt     string `json:"-" xorm:"salt"`
	Disk     string `json:"disk" xorm:"disk"`
	Engine   string `json:"engine" xorm:"engine"`
}

func (c *FileManagerAccount) Init() error {
	if c.Username == "" {
		return fmt.Errorf("用户不能为空")
	}
	if c.Salt == "" {
		c.Salt = helper.GetRandomString(4)
	}
	return nil
}

func (c *FileManagerAccount) GetMd5Pwd(pwd string) string {
	hash := md5.New()
	hash.Write([]byte(pwd + c.Salt))
	return hex.EncodeToString(hash.Sum(nil))
}
