package helper

import "fmt"

func DetailUrl(tid, aid int) string {
	return fmt.Sprintf("/news/%d-%d.html", tid, aid)
}

func ListUrl(tid int) string {
	return fmt.Sprintf("/list/%d-1.html", tid)
}

func PageUrl(tid int) string {
	return fmt.Sprintf("/page/%d.html", tid)
}