package storage

import (
	"fmt"
	"testing"
)

func TestFtpUploader(t *testing.T) {
	uploader := NewFtpUploader(map[string]string{
		"FTP_SERVER_URL":  "124.222.103.232",
		"FTP_SERVER_PORT": "21",
		"FTP_USER_NAME":   "test",
		"FTP_USER_PWD":    "Hh2EptLZAN2KrbXd",
		"FTP_UPLOAD_DIR":  "/www/wwwroot/test",
	})

	fmt.Println(uploader.List(""))
}
