package helper

import "fmt"

func DetailUrl(tid, aid int) string {
	return fmt.Sprintf("/view/%d-%d.html", tid, aid)
}

func ListUrl(tid int) string {
	return fmt.Sprintf("/list/%d.html", tid)
}