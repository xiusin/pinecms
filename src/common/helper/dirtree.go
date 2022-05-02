package helper

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

type DirInfo struct {
	Label    string      `json:"label"`
	FullPath string      `json:"full_path"`
	IsDir    bool        `json:"is_dir"`
	Children interface{} `json:"children"`
}

func DirTree(dir string) []DirInfo {
	fileInfos, _ := ioutil.ReadDir(dir)
	var ms []DirInfo
	for _, f := range fileInfos {
		fullPath := filepath.Join(dir, f.Name())
		if f.IsDir() {
			s := DirInfo{
				Label:    f.Name(),
				IsDir:    true,
				FullPath: fullPath,
				Children: DirTree(fullPath),
			}
			ms = append(ms, s)
		} else {
			ext := strings.ToLower(filepath.Ext(f.Name()))
			if ext != ".css" && ext != ".js" && ext != ".jet" && ext != ".html" && ext != ".htm" && ext != ".sh" {
				continue
			}
			ms = append(ms, DirInfo{
				Label:    f.Name(),
				FullPath: fullPath,
				Children: "",
			})
		}
	}
	return ms
}
