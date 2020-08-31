package cmd

import (
	"encoding/json"
	"html/template"
)

func FormatEnum(field string, enums []string, item map[string]interface{}) {
	item["type"] = "tpl"
	enumsInfo, _ := json.Marshal(enums)
	if len(enumsInfo) == 0 {
		enumsInfo = []byte("[]")
	}
	item["tpl"] = template.HTML(`<%= `+ string(enumsInfo) +`[data.`+ field +`] %>`)
}


func FormatSet(field string, set []string, item map[string]interface{}) {
	item["type"] = "tpl"
	enumsInfo, _ := json.Marshal(set)
	if len(enumsInfo) == 0 {
		enumsInfo = []byte("[]")
	}
	json.
	item["tpl"] = "'<% data."+field+".split(\",\").forEach(function(item) { %><%= " + string(enumsInfo) + "[item] %></span> <% }) %>'"
}

func FormatLink()  {

}

func FormatFile()  {

}

func FormatImage()  {

}

func FormatStatus(field string, set []string, item map[string]interface{})  {
	
}

