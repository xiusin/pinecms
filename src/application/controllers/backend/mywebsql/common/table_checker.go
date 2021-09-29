package common

import (
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/common/helper"
)

// TableChecker 表检查
type TableChecker struct {
	*sqlx.DB
	Db         string
	Tables     []string
	Opteration string
	Options    map[string]interface{}
	sql        string
	rows       *sqlx.Rows
}

func NewTableChecker(DB *sqlx.DB) *TableChecker {
	return &TableChecker{DB: DB, Options: map[string]interface{}{}}
}

func (t *TableChecker) SetTables(tables []string) {
	t.Tables = tables
}

func (t *TableChecker) SetOperation(op string) {
	t.Opteration = op
}

func (t *TableChecker) SetOptions(options map[string]interface{}) {
	t.Options = options
}

func (t *TableChecker) GetSql() string {
	return t.sql
}

func (t *TableChecker) Runcheck() error {
	cmd := t.Opteration
	if t.Options["skiplog"].(bool) {
		cmd += " NO_WRITE_TO_BINLOG"
	}
	cmd += " tables "

	for _, table := range t.Tables {
		cmd += "`" + table + "`,"
	}

	cmd = strings.TrimRight(cmd, ",")

	if t.Opteration == "check" {
		cmd += " " + t.checkOptions()
	} else if t.Opteration == "repair" {
		cmd += " " + t.repairOptions()
	}

	t.sql = cmd
	var err error
	if t.rows, err = t.Queryx(t.sql); err != nil {
		pine.Logger().Warning("执行语句失败: "+t.sql, err)
	}
	return err
}

func (t *TableChecker) repairOptions() string {
	str := ""
	if ok, _ := helper.InArray("quick", t.Options["repairtype"].([]string)); ok {
		str += " QUICK"
	} else if ok, _ := helper.InArray("extended", t.Options["repairtype"].([]string)); ok {
		str += " EXTENDED"
	} else if ok, _ := helper.InArray("usefrm", t.Options["repairtype"].([]string)); ok {
		str += " USE_FRM"
	}
	return str
}

func (t *TableChecker) checkOptions() string {
	str := ""
	ct, _ := t.Options["checktype"].(string)
	switch ct {
	case "quick":
		str += " QUICK"
	case "extended":
		str += " EXTENDED"
	case "fast":
		str += " FAST"
	case "meduin", "medium":
		str += " MEDIUM"
	case "changed":
		str += " CHANGED"
	}
	return str
}

func (t *TableChecker) GetResults() map[string]map[string]string {
	ret := map[string]map[string]string{}
	if t.rows != nil {
		defer t.rows.Close() // 关闭句柄
		for t.rows.Next() {
			results := make(map[string]interface{})
			t.rows.MapScan(results)
			msg := ""
			if results["Msg_text"] != nil {
				msg = string(results["Msg_text"].([]byte))
			}
			switch string(results["Msg_type"].([]byte)) {
			case "Error":
				ret[string(results["Table"].([]byte))] = map[string]string{"type": "error", "msg": msg}
			case "note":
				ret[string(results["Table"].([]byte))] = map[string]string{"type": "note", "msg": msg}
			case "status":
				ret[string(results["Table"].([]byte))] = map[string]string{"type": "success", "msg": msg}
			}
		}
	}

	return ret
}
