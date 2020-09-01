package cmd

import (
	"bytes"
	"encoding/json"
)

func FormatEnum(field string, opts []map[string]interface{}, item map[string]interface{}) {
	item["type"] = "tpl"
	vmap := map[string]interface{}{}
	for _, opt := range opts {
		vmap[opt["value"].(string)] = opt["label"]
	}
	enumsInfo, _ := json.Marshal(vmap)
	if len(enumsInfo) == 0 {
		enumsInfo = []byte("[]")
	}
	item["tpl"] = `<%= ` + string(enumsInfo) + `[data.` + field + `] %>`
}

func FormatSet(field string, opts []map[string]interface{}, item map[string]interface{}) {
	item["type"] = "tpl"
	vmap := map[string]interface{}{}
	for _, opt := range opts {
		vmap[opt["value"].(string)] = opt["label"]
	}
	enumsInfo, _ := json.Marshal(vmap)
	if len(enumsInfo) == 0 {
		enumsInfo = []byte("[]")
	}
	item["tpl"] = "<% data." + field + ".split(\",\").forEach(function(item) { %><%= " + string(enumsInfo) + "[item] %></span> <% }) %>"
}

func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}
