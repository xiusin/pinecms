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
		"FTP_USER_PWD":    "Hh2EptLZAN2KrbXd",
	})


	t.Log(uploader.Mkdir("goftp"))
	t.Log(uploader.Mkdir("goftp2/subftpdir"))

	tmp, _ := 	ioutil.TempFile("", "")
	t.Log(io.WriteString(tmp, "hello world"))
	tmp.Close()
	tmp,_ = os.Open(tmp.Name()) // TODO 临时文件必须先保存一下才能使用
	t.Log("delete", uploader.Remove("goftp/index.html"))
	t.Log(uploader.Upload("goftp/index.html", tmp))
	byts, err := uploader.Content("goftp/index.html")
	if err != nil {
		t.Log(err)
	}
	t.Log("byts", string(byts))

	t.Log(uploader.List("113123123/111111/22222"))

}
