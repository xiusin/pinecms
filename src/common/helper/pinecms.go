package helper

import (
	"fmt"
)

func DetailUrl(tid, aid int, prefix ...string) string {
	return fmt.Sprintf("/news/%d-%d.html", tid, aid)
}

func ListUrl(tid int, prefix ...string) string {
	return fmt.Sprintf("/list/%d-1.html", tid)
}

func PageUrl(tid int, prefix ...string) string {
	return fmt.Sprintf("/page/%d.html", tid)
}