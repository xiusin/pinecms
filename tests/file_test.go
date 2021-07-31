package tests

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

type DirInfo struct {
	Label    string      `json:"label"`
	Children interface{} `json:"children"`
}

func TestDir(t *testing.T) {
	fmt.Println(dirTree("/Users/xiusin/projects/src/github.com/xiusin/pinecms/resources/themes/default"))
}

func dirTree(dir string) []DirInfo {
	fs, _ := ioutil.ReadDir(dir)
	var ms []DirInfo
	for _, f := range fs {
		if f.IsDir() {
			s := DirInfo{
				Label:    f.Name(),
				Children: dirTree(filepath.Join(dir, f.Name())),
			}
			ms = append(ms, s)
		} else {
			ms = append(ms, DirInfo{
				Label:    f.Name(),
				Children: "",
			})
		}
	}
	return ms
}
