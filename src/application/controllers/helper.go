package controllers

type FieldShowInPageList struct {
	Show      bool   `json:"show"`
	Formatter string `json:"formatter"`
}
func GetInMap(data map[string]FieldShowInPageList,key string) FieldShowInPageList {
	s, o := data[key]
	if o {
		return s
	} else {
		return FieldShowInPageList{}
	}
}