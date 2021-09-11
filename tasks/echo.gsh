package tasks

import (
	"encoding/json"
	"pinecms"
)

/**
脚本样例
 */
func Run(orm *pinecms.DB) (string, error) {
	var data = map[string]interface{}{}
	orm.Table("pinecms_task_info").Get(&data)
	d, _ := json.Marshal(&data)
	return string(d), nil
}
