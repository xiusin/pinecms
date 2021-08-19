package crud

// public/assets/js/require-table.js:403
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
	topCode = append(topCode, `let _`+field+` =`+string(enumsInfo)+`;`)
	item["tpl"] = "<%=formatterEnum(data." + field + ", _" + field + ")%>"
	//item["tpl"] = `<%= ` + string(enumsInfo) + `[data.` + field + `] %>`
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
	topCode = append(topCode, `let _`+field+` =`+string(enumsInfo)+`;`)
	item["tpl"] = "<%=formatterSet(data." + field + ", _" + field + ")%>"
}

// JSONMarshal 不转义字符串编码
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}
