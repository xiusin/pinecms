package common

type html struct {}

var Html html

// ArrayToOptions 数组转下拉列表
func (h html) ArrayToOptions(array []string, selected string, defaultText string) string {
	str := ""
	if len(defaultText) > 0 {
		str = "<option value=\"\">"+T(defaultText)+"</option>\n"
	}
	for _, val := range array {
		if selected == val {
			str += "<option selected=\"selected\" value=\""+ val + "\">" +  val + "</option>\n"
		} else {
			str += "<option value=\""+ val + "\">" +  val + "</option>\n"
		}
	}
	return str
}
