package storage

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFtpUploader(t *testing.T) {
	uploader := NewFtpUploader(map[string]string{
		"FTP_SERVER_URL":  "124.222.103.232",
		"FTP_SERVER_PORT": "21",
		"FTP_USER_NAME":   "test",
		"FTP_USER_PWD":    "",
		"SITE_URL":        "http://localhost:2019/xxx/",
		"FTP_URL_PREFIX":  "", // 如果配置则使用此配置拼接地址, 否则使用系统接口
	})

	t.Log(uploader.Mkdir("goftp"))
	t.Log(uploader.Mkdir("goftp2/subftpdir"))

	tmp, _ := ioutil.TempFile("", "")
	t.Log(io.WriteString(tmp, "hello world"))

	tmp.Close()
	tmp, _ = os.Open(tmp.Name())
	defer func() {
		tmp.Close()
		os.Remove(tmp.Name())
	}()

	t.Log("delete", uploader.Remove("goftp/index.html"))
	t.Log(uploader.Upload("goftp/index.html", tmp))
	byts, err := uploader.Content("goftp/index.html")
	if err != nil {
		t.Log(err)
	}
	t.Log("byts", string(byts))

	t.Log("exist: ")
	t.Log(uploader.Exists("goftp/index.html"))

	t.Log("rename", uploader.Rename("goftp/index.html", "goftp2/index_1.html"))

	t.Log("exist: ")
	t.Log(uploader.Exists("goftp/index.html"))

	t.Log("rmdir", uploader.Rmdir("goftp2"))

	t.Log(uploader.List("113123123/111111/22222"))

}
