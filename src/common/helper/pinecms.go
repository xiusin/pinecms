package helper

import (
	"fmt"
	"path/filepath"
)

func DetailUrl(tid, aid int, prefix ...string) string {
	if len(prefix) == 0 {
		prefix = append(prefix, "")
	}
	return filepath.Join(prefix[0], fmt.Sprintf("news/%d-%d.html", tid, aid))
}

func ListUrl(tid int, prefix ...string) string {
	if len(prefix) == 0 {
		prefix = append(prefix, "")
	}
	return fmt.Sprintf("/%s/", prefix[0])
}

func PageUrl(tid int, prefix ...string) string {
	if len(prefix) == 0 {
		prefix = append(prefix, "")
	}
	return fmt.Sprintf("/page/%d.html", tid)
}