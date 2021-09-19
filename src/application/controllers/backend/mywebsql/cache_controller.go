package mywebsql

import (
	"bytes"
	"fmt"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/common/helper"
	"os"
	"regexp"
	"strings"
)

func Cache(ctx *pine.Context) {
	themePath := ctx.Session().Get("theme_path")
	var byts bytes.Buffer
	script := ctx.GetString("script")
	css := ctx.GetString("css")
	if script != "" {
		scriptPath := helper.GetRootPath("mywebsql/js")
		scripts := strings.Split(script, ",")
		ctx.Response.Header.Set("mime-type", "text/javascript")
		ctx.Response.Header.Set("content-type", "text/javascript")
		r := regexp.MustCompile("^(\\w+/){0,2}\\w+$")
		for _, s := range scripts {
			if !r.MatchString(s) {
				continue
			}
			fullPath := scriptPath + "/" + s + ".js"
			if data, err := os.ReadFile(fullPath); err == nil {
				byts.Write(data)
				byts.WriteByte('\n')
				byts.WriteByte('\n')
			} else {
				fmt.Println("无法找到文件", fullPath)
			}
		}

	} else if css != "" {
		styles := strings.Split(css, ",")
		ctx.Response.Header.Set("mime-type", "text/css")
		ctx.Response.Header.Set("content-type", "text/css")
		r := regexp.MustCompile("^(\\w+/){0,2}\\w+$")
		for _, s := range styles {
			if !r.MatchString(s) {
				continue
			}
			fullPath := helper.GetRootPath("mywebsql/themes/_base/" + s + ".css")
			if data, err := os.ReadFile(fullPath); err == nil {
				byts.Write(data)
				byts.WriteByte('\n')
				byts.WriteByte('\n')
			}else {
				fmt.Println("无法找到文件", fullPath)
			}
			fullPath = helper.GetRootPath("mywebsql/themes/"+themePath+"/" + s + ".css")
			if data, err := os.ReadFile(fullPath); err == nil {
				byts.Write(data)
				byts.WriteByte('\n')
				byts.WriteByte('\n')
			}else {
				fmt.Println("无法找到文件", fullPath)
			}
		}
	}
	ctx.Write(byts.Bytes())
}
