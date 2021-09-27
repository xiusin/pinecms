package common

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"math"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/common/helper"
)

const trimChar = " \r\n\t;"

type Process struct {
	*pine.Context
	db       *sqlx.DB
	dbname   string
	lastSQL  string
	formData struct {
		id    string
		page  int
		name  string
		table string
		query string
	}
}

func InitProcess(db *sqlx.DB, ctx *pine.Context) *Process {
	p := &Process{
		db:      db,
		Context: ctx,
	}
	p.SelectVersion()

	p.dbname = p.Session().Get("db.name")

	if p.dbname != "" {
		p.db.Exec("USE " + p.dbname)
	}

	p.formData.query = strings.Trim(p.FormValue("query"), trimChar)
	p.formData.table = strings.Trim(p.FormValue("name"), trimChar)
	p.formData.name = p.formData.table
	p.formData.id = p.FormValue("id")
	p.formData.page = 1

	if match, _ := regexp.MatchString(`^\d+$`, p.formData.table); match {
		p.formData.page, _ = strconv.Atoi(p.formData.table)
		p.formData.table = strings.Trim(p.FormValue("query"), trimChar)
		p.formData.query = ""
	}

	return p
}

func (p *Process) Info() string {
	var html []byte
	if p.dbname != "" {
		tip := T("Database summary") + ": [" + p.dbname + "]"
		html = []byte(p.createDbInfoGrid(tip))
	}
	//  else {
	// html = []byte(p.createResultGrid(""))
	// }
	return string(html)
}

// Showinfo 展示数据表或数据库的基本信息
func (p *Process) Showinfo() string {
	typo := p.PostString("id")
	if typo == "table" || typo == "view" {
		p.formData.id = "table"
		return p.Query()
	} else {
		cmd := p.getCreateCommand(p.formData.id, p.formData.table)
		cmd = p.sanitizeCreateCommand(cmd)

		return string(p.Render("showinfo", pine.H{
			"TYPE":    p.formData.id,
			"NAME":    p.formData.table,
			"COMMAND": cmd,
		}))
	}

}

// Objcreate 创建对象, 如事务,视图,存储过程等
func (p *Process) Objcreate() string {
	typo := "message ui-state-highlight"
	var msg string
	var refresh bool
	objinfo := p.FormValue("objinfo")
	if objinfo != "" {
		pine.Logger().Debug("createObj", objinfo)
		msg = p.createDatabaseObject(objinfo)
		if msg == "" {
			msg = T("The command executed successfully")
			typo = "message ui-state-default"
			refresh = true
		} else {
			typo = "message ui-state-error"
		}
	} else {
		msg = T("Any existing object with the same name should be dropped manually before executing the creation command") +
			"!<br/>" + T("Enter command for object creation")
	}

	return p.displayCreateObjectForm(objinfo, msg, typo, refresh)
}

// displayCreateObjectForm 显示创建对象表单 @ref Objcreate
func (p *Process) displayCreateObjectForm(objInfo, msg, typo string, refresh bool) string {
	if objInfo == "" {
		objInfo = p.getObjectCreateCommand()
	}
	form := "</textarea></td></tr>"

	editorLink := template.HTML("<script type=\"text/javascript\" language=\"javascript\" src=\"/mywebsql/cache?script=editor/codemirror\"></script>")

	editorOptions := template.HTML("parserfile: \"mysql.js\", path: \"/mywebsql/js/editor/\"")

	v := pine.H{
		"ID":             p.formData.id,
		"MESSAGE":        template.HTML(msg),
		"MESSAGE_TYPE":   typo,
		"OBJINFO":        objInfo,
		"EDITOR_LINK":    editorLink,
		"EDITOR_OPTIONS": editorOptions,
		"REFRESH":        "0",
	}
	if refresh {
		v["REFRESH"] = "1"
	}

	return form + string(p.Render("objcreate", v))
}

func (p *Process) createDatabaseObject(info string) string {
	cmd := strings.Trim(info, " \t\r\n;")
	if strings.ToLower(cmd[:6]) != "create" {
		return T("Only create commands are accepted")
	}

	if _, err := p.db.Exec(cmd); err != nil {
		return err.Error()
	}

	ws := p.getWarnings()
	if len(ws) > 0 {
		for _, s := range ws {
			return s
		}
	}
	return ""
}

func (p *Process) getObjectCreateCommand() string {
	templates := map[string]string{
		"0": "templates/table",
		"1": "templates/view",
		"2": "templates/procedure",
		"3": "templates/function",
		"4": "templates/trigger",
		"5": "templates/event",
		"6": "templates/schema",
	}

	return string(p.Render(templates[p.formData.id], nil))
}

// Objlist 切换数据库时触发
func (p *Process) Objlist() string {
	grid := `<div id="objlist">`
	grid += GetDatabaseTreeHTML(p.db, []string{}, p.dbname)
	grid += `</div>`
	return grid
}

func (p *Process) Usermanager() string {
	return "用户管理界面"
}

// Databases 数据库操作管理
func (p *Process) Databases() string {
	dbs, err := GetDbList(p.db)
	if err != nil {
		return p.createErrorGrid("SHOW DATABASES", err)
	}
	byts, _ := json.Marshal(dbs)
	datas := pine.H{"data": pine.H{"objects": template.HTML(byts)}, "objCount": len(dbs), "stats": nil}
	if p.formData.id == "batch" {
		postdata := p.PostData()
		status := map[string]int{"success": 0, "errors": 0}
		databases := postdata["databases[]"]
		pine.Logger().Warning("删除数据库", databases)
		if len(databases) > 0 {
			for _, database := range databases {
				if p.FormValue("dropcmd") == "on" {
					if err := p.dropObject(database, "database"); err == nil {
						status["success"]++
					} else {
						status["errors"]++
					}
				}
			}
			datas["stats"] = pine.H{"drop": status}
			//>' . str_replace('{{ NUM }}', $data['stats']['drop']['success'], __('{{ NUM }} queries successfully executed')) . '
			txt := "<p><span class=\"ui-icon ui-icon-check\"></span>" + strReplace([]string{"{{ NUM }}"}, []string{strconv.Itoa(status["success"])}, T("{{ NUM }} queries successfully executed")) + "</p>"
			if status["errors"] > 0 {
				txt += "<p><span class=\"ui-icon ui-icon-close\"></span>" + strReplace([]string{"{{ NUM }}"}, []string{strconv.Itoa(status["success"])}, T("{{ NUM }} queries failed to execute")) + "</p>"
			}
			datas["statsHtml"] = txt
		}
	}
	return string(p.Render("databases", datas))
}

func (p *Process) dropObject(name, typo string) error {
	query := "drop " + typo + " `" + name + "`"
	pine.Logger().Warning("执行删除操作: ", query)
	_, err := p.db.Exec(query)
	return err
}

func (p *Process) getObjectTypes() []string {
	return []string{"tables", "views", "procedures", "functions", "triggers", "events"}
}

func (p *Process) getWarnings() map[int]string {
	var ret = map[int]string{}
	if rows, err := p.db.Queryx("SHOW WARNINGS"); err != nil {
		pine.Logger().Warning("获取警告失败", err)
	} else {
		for rows.Next() {
			results := make(map[string]interface{})
			err = rows.MapScan(results)
			code, _ := strconv.Atoi(string(results["Code"].([]byte)))
			ret[code] = string(results["Message"].([]byte))
		}
	}
	return ret
}

func (p *Process) Infodb() string {
	if p.dbname == "" {
		return string(p.Render("invalid_request", nil))
	}
	query := "show table status where Engine is not null"
	return p.createSimpleGrid(T("Database summary")+": ["+p.dbname+"]", query)
}

func (p *Process) Query() string {
	var querySql string
	if p.formData.id == "table" {
		querySql = p.selectFromTable()
	} else {
		querySql = p.simpleQuery()
	}

	querySql = strings.Trim(querySql, " \r\n")

	if len(querySql) == 0 {
		return p.createErrorGrid(querySql, errors.New("无法生成Query SQL"))
	}

	var html string

	p.loadDbVars()
	// TODO 区分查询是否需要返回执行结果
	if strings.ToLower(querySql[:6]) == "select" {
		queryType := p.getQueryType(querySql)
		pine.Logger().Debug(queryType["can_limit"])
		if queryType["can_limit"] {
			html = p.createResultGrid(querySql)
		} else {
			html = p.createSimpleGrid(T("Query")+": "+querySql, querySql)
		}
	} else {

	}

	return html

}

func (p *Process) loadDbVars() {
	varByts := []byte(p.Session().Get("vars"))
	var vars = map[string]string{}
	json.Unmarshal(varByts, &vars)
	for variable, value := range vars {
		p.db.Exec("SET @" + variable + " = " + value)
	}
}

func (p *Process) setDbVar(variable, value string) {
	varByts := []byte(p.Session().Get("vars"))
	var vars = map[string]string{}
	json.Unmarshal(varByts, &vars)
	vars[variable] = value
	varByts, _ = json.Marshal(vars)
	p.Session().Set("vars", string(varByts))
}

func (p *Process) sanitizeCreateCommand(cmd string) string {
	return regexp.MustCompile(`[\n|\r]?[\n]+`).ReplaceAllString(cmd, "<br/>")
}

func (p *Process) selectFromTable() string {
	recordLimit := p.GetCookie("res-max-count")
	if len(recordLimit) == 0 {
		recordLimit = fmt.Sprintf("%d", MAX_RECORD_TO_DISPLAY)
	}
	recordLimitInt, _ := strconv.Atoi(recordLimit)
	var query string
	if p.formData.page > 1 {
		if p.selectSession("limit") != "true" {
			pine.Logger().Debug("无limit数据")
			return ""
		}
		query = p.selectSession("query")
		table := p.selectSession("table")
		count := p.selectSessionInt("count")
		queryType := p.getQueryType(query)
		if !queryType["result"] || table == "" || count < 1 {
			return ""
		}
		totalPages := int(math.Ceil(float64(count) / float64(recordLimitInt)))
		if totalPages < p.formData.page {
			return ""
		}
		p.Session().Set("select.page", p.n2s(p.formData.page))
		query += fmt.Sprintf(" LIMIT %d OFFSET %d", recordLimitInt, (p.formData.page-1)*recordLimitInt)

	} else {
		keys := []string{"table", "can_limit", "limit", "page", "count"}
		for _, key := range keys {
			p.Session().Remove("select." + key)
		}

		query = "SELECT * FROM `" + p.formData.table + "`"

		p.Session().Set("select.query", query)
		p.Session().Set("select.table", p.formData.table)

		var count int

		if err := p.db.Get(&count, "SELECT COUNT(*) FROM `"+p.formData.table+"`"); err != nil {
			pine.Logger().Warning("获取总数失败", err)
		}

		p.Session().Set("select.count", p.n2s(count))
		p.Session().Set("select.page", "1")
		p.Session().Set("select.can_limit", "true")

		if count > recordLimitInt {
			p.Session().Set("select.limit", "true")
			query += " LIMIT " + recordLimit
		}
	}
	return query
}

func (p *Process) n2s(num int) string {
	return strconv.Itoa(num)
}

func (p *Process) simpleQuery() string {
	query := p.formData.query
	if query == "" {
		query = p.selectSession("query")
	}
	if query == "" {
		return ""
	}

	queryType := p.getQueryType(query)

	if !queryType["result"] || !queryType["can_limit"] {
		return query
	}

	if p.selectSession("can_limit") == "true" {
		p.Session().Set("select.can_limit", "true")
	} else {
		p.Session().Set("select.can_limit", "false")
	}

	if p.formData.id == "sort" {
		field := p.formData.name
		if field != "" {
			query = p.sortQuery(query, field)
		}
	}

	return query
}

func (p *Process) sortQuery(query, field string) string {
	query = strings.Trim(query, " \r\n\t")
	sortType := p.selectSession("sort")
	//limit := ""
	if sortType == "" {
		sortType = "DESC"
	}
	// 匹配LIMIT语句
	matches := regexp.MustCompile(LIMIT_REGEXP).FindStringSubmatch(query)
	// 匹配sort语句
	matches = regexp.MustCompile(SORT_REGEXP).FindStringSubmatch(query)
	pine.Logger().Debug("matches", matches)

	p.Session().Set("select.sortcol", field)
	p.Session().Set("select.sort", sortType)

	//query += " ORDER BY " + field + " " + sortType + " " + limit

	pine.Logger().Debug("SORT 最终SQL", query)

	return query
}

func (p *Process) getQueryType(query string) map[string]bool {
	types := map[string]bool{"result": false, "can_limit": false, "has_limit": false, "update": false}
	q := strings.ToLower(query[:7])

	if q == "explain" || q[:6] == "select" || q[:4] == "desc" || q[:4] == "show" || q[:4] == "help" {
		types["result"] = true

		if q[:6] == "select" {
			types["can_limit"] = true
		}
	}
	match, err := regexp.MatchString(LIMIT_REGEXP, query)
	if err != nil {
		pine.Logger().Error("匹配sql错误", err)
	}
	if match && strings.Contains(strings.ToLower(query), "limit") {
		types["has_limit"] = true
	}
	return types
}

func (p *Process) exec(module string) (html string) {
	defer func() {
		if err := recover(); err != nil {
			pine.Logger().Error(err)
			html = p.createErrorGrid("解析执行方法失败", fmt.Errorf("%s", err))
		}
	}()
	val := reflect.ValueOf(&p).Elem().MethodByName(helper.UcFirst(module))
	if !val.IsValid() {
		html = p.createErrorGrid("无法解析处理句柄", nil)
	} else {
		html = val.Call([]reflect.Value{})[0].String()
	}
	return
}

func (p *Process) QueryVariables() []Variable {
	query := "SHOW VARIABLES"
	var variables []Variable
	if err := p.db.Select(&variables, query); err != nil {
		pine.Logger().Warning("获取变量失败", err)
	}
	return variables
}

func (p *Process) createErrorGrid(query string, err error, params ...int) string {
	if query == "" {
		query = p.Session().Get("select.query")
	}

	sessionKeys := []string{"result", "pkey", "ukey", "mkey", "unique_table"}
	for _, v := range sessionKeys {
		p.Session().Remove("select." + v)
	}

	numQueries, affectedRows := 0, -1
	if len(params) > 0 {
		numQueries = params[0]
	}
	if len(params) > 1 {
		affectedRows = params[1]
	}

	grid := "<div id='results'>\n"
	if numQueries > 0 {
		grid += "<div class=\"message ui-state-default\">"
		var msg string
		if numQueries == 1 {
			msg = T("1 query successfully executed")
		} else {
			msg = strings.ReplaceAll(T("{{NUM}} queries successfully executed"), "{{NUM}}", strconv.Itoa(numQueries))
		}

		msg += "<br/><br/>" + strings.ReplaceAll(T("{{NUM}} record(s) were affected"), "{{NUM}}", strconv.Itoa(affectedRows))
		grid += msg + "</div>"
	}

	match := regexp.MustCompile("/[\\n|\\r]?[\\n]+/")

	formattedQuery := match.ReplaceAllString(query, "<br>")

	grid += "<div class=\"message ui-state-error\">" + T("Error occurred while executing the query") +
		":</div><div class=\"message ui-state-highlight\">" + err.Error() +
		"</div><div class=\"sql-text ui-state-error\">" + formattedQuery + "</div>"

	grid += "</div>"
	grid += "<script type=\"text/javascript\" language='javascript'> parent.transferResultMessage(-1, '&nbsp;', '" + T("Error occurred while executing the query") + "');\n"
	grid += "parent.addCmdHistory(\"" + strings.ReplaceAll(p.Session().Get("select.query"), "\r\n", "<br/>") + "\");\n"
	grid += "parent.resetFrame();\n"
	grid += "</script>\n"

	return grid
}

func (p *Process) getTemplateSQL(templateName string) string {
	content, err := GetPlush().Exec("template/"+templateName, nil)
	if err != nil {
		return ""
	}
	return string(content)
}

func (p *Process) Infoserver() string {
	grid := ""
	variables := p.QueryVariables()
	if len(variables) == 0 {
		return ""
	}

	v := pine.H{"JS": "", "SERVER_NAME": "MySQL"}

	for _, variable := range variables {
		switch variable.VariableName {
		case "version":
			v["SERVER_VERSION"] = variable.Value
		case "version_comment":
			v["SERVER_COMMENT"] = variable.Value
		case "character_set_server":
			v["SERVER_CHARSET"] = variable.Value
		case "character_set_client":
			v["CLIENT_CHARSET"] = variable.Value
		case "character_set_database":
			v["DATABASE_CHARSET"] = variable.Value
		case "character_set_results":
			v["RESULT_CHARSET"] = variable.Value
		}
		v[variable.VariableName] = variable.Value
	}

	if p.dbname == "" {
		v["JS"] = `parent.$("#main-menu").find(".db").hide();`
	}
	pine.Logger().Debug("V", v)
	grid += string(p.Render("infoserver", v))

	return grid
}

func (p *Process) createSimpleGrid(message string, query string) string {
	pine.Logger().Debug("createSimpleGrid")
	grid := "<div id='results'>"
	grid += "<div class='message ui-state-default'>" + message + "<span style='float:right'>" + T("Quick Search") +
		"&nbsp;<input type=\"text\" id=\"quick-info-search\" maxlength=\"50\" /></div>"
	grid += "<table cellspacing='0' width='100%' border='0' class='results' id='infoTable'><thead>\n"

	grid += "<tr id='fhead'><th class='th index'><div>#</div></th>\n"

	// 遍历数据
	rows, err := p.db.Query(query)
	if err != nil {
		pine.Logger().Warning("查询异常", query, err)
		return p.createErrorGrid(query, err)
	}

	fields, _ := rows.Columns()
	fieldTypes, _ := rows.ColumnTypes()

	for k, fn := range fields {
		cls, dsrt := "th", "text"
		fieldType := fieldTypes[k]
		if exist, _ := helper.InArray(fieldType.DatabaseTypeName(), []string{"DECIMAL", "INT", "BIGINT", "TINYINT", "FLOAT", "DOUBLE"}); exist {
			cls = "th_numeric"
			dsrt = "numeric"
		}
		grid += "<th nowrap=\"nowrap\" class='" + cls + "' data-sort='" + dsrt + "'><div>" + fn + "</div></th>\n"
	}

	grid += "</tr></thead><tbody>\n"

	datas, err := p.row2arrMap(rows)
	if err != nil {
		pine.Logger().Warning("转换数据类型异常", query, err)
		return p.createErrorGrid(query, err)
	}

	for j, table := range datas {
		grid += `<tr id="rc` + strconv.Itoa(j) + `" class="row"><td class="tj">` + strconv.Itoa(j+1) + `</td>`

		for i, field := range fields {
			class := "tl"
			if j == 0 {
				pine.Logger().Print(field, fieldTypes[i].ScanType().Name(), fieldTypes[i].DatabaseTypeName())
			}
			rs := table[field]
			data := "&nbsp;"
			if rs == nil {
				class = "tnl"
				data = "NULL"
			}

			switch fieldTypes[i].DatabaseTypeName() {
			case "VARCHAR", "CHAR", "TEXT":
				class += " text"
				if rs != nil && len(rs.([]byte)) != 0 {
					data = string(rs.([]byte))
				}
			case "INT", "BIGINT":
				class = "tr"
				if rs != nil && len(rs.([]byte)) != 0 {
					data = string(rs.([]byte))
				}
			case "TIMESTAMP", "DATETIME":
				if rs != nil {
					data = rs.(time.Time).Format(helper.TimeFormat)
				}
			case "binary", "blob": // blob

			}

			// TODO 确认blob类型以及数字类型. 后续碰到补充
			grid += "<td nowrap=\"nowrap\" id=\"r" + p.n2s(j) + "f" + p.n2s(i) + "\" class=\"" + class + "\">" + data + "</td>"
		}
		grid += "</tr>\n"
	}

	grid += "</tbody></table>"
	grid += "</div>"

	grid += "<script type=\"text/javascript\" language=\"javascript\">\n"
	grid += "parent.transferInfoMessage();\n"
	grid += "parent.resetFrame();\n"
	grid += "</script>"
	return grid
}

func (p *Process) createResultGrid(query string) string {
	sessionKeys := []string{"pkey", "ukey", "mkey", "unique_table"}
	for _, v := range sessionKeys {
		p.Session().Remove("select." + v)
	}

	var grid = ""

	recordLimit := p.GetCookie("res-max-count")
	if len(recordLimit) == 0 {
		recordLimit = fmt.Sprintf("%d", MAX_RECORD_TO_DISPLAY)
	}
	recordLimitInt, _ := strconv.Atoi(recordLimit)

	grid += "<div id='results'>"
	grid += "<table cellspacing=\"0\" width='100%' border=\"0\" class='results' id=\"dataTable\"><thead>\n"

	f := p.getFieldInfo()

	if p.Session().Get("select.can_limit") == "true" && len(f) > 0 {
		p.Session().Set("select.unique_table", f[0].TableName)
	}

	grid += "<tr id=\"fhead\">"
	grid += "<th class=\"th tch\"><div>#</div></th>"

	ed := p.Session().Get("select.can_limit") == "true" && p.Session().Get("select.unique_table") != ""

	if ed {
		grid += "<th class=\"th_nosort tch\"><div><input class=\"check-all\" type=\"checkbox\" onclick=\"resultSelectAll()\" title=\"" + T("Select/unselect All records") + "\" /></div></th>"
	}

	var pkey, ukey, mkey []string
	fieldNames := ""
	fieldInfo, _ := json.Marshal(&f)

	for i, column := range f {
		cls := "th"
		if column.Numeric {
			cls = "th_numeric"
		}
		grid += "<th nowrap=\"nowrap\" class='" + cls + "'><div>"
		if column.PKey {
			pkey = append(pkey, column.ColumnName)
			grid += "<span class='pk' title='" + T("Primary key column") + "'>&nbsp;</span>"
		}
		if column.UKey {
			ukey = append(ukey, column.ColumnName)
			grid += "<span class='uk' title='" + T("Unique key column") + "'>&nbsp;</span>"
		}
		if column.MKey && !column.Blob {
			mkey = append(mkey, column.ColumnName)
		}
		grid += column.ColumnName
		// 排序应用
		if p.Session().Get("select.sortcol") == strconv.Itoa(i+1) {
			if p.Session().Get("select.sort") == "DESC" {
				grid += "&nbsp;&#x25BE;"
			} else {
				grid += "&nbsp;&#x25B4"
			}
		}
		grid += "</div></th>"
		fieldNames += "'" + strings.ReplaceAll(column.ColumnName, "'", "\\'") + "',"
	}

	grid += "</tr></thead><tbody>\n"

	// 遍历数据
	rows, err := p.db.Query(query)
	if err != nil {
		pine.Logger().Warning("查询异常", query, err)
		return p.createErrorGrid(query, err)
	}

	datas, err := p.row2arrMap(rows)
	if err != nil {
		pine.Logger().Warning("转换数据类型异常", query, err)
		return p.createErrorGrid(query, err)
	}

	for j, r := range datas {
		grid += "<tr class=\"row\">"
		grid += "<td class=\"tj\">" + strconv.Itoa(j+1) + "</td>"
		if ed {
			grid += "<td class=\"tch\"><input type=\"checkbox\" /></td>"
		}
		for _, column := range f {
			rs, _ := r[column.ColumnName]
			class := "tl"
			if column.Numeric {
				class = "tr"
			}
			if rs == nil {
				class = "tnl"
			}
			if ed {
				class += " edit"
			}

			data := ""
			if !column.Blob {
				if rs == nil {
					data = "NULL"
				} else {
					if v, ok := rs.([]byte); ok && len(v) != 0 {
						data = string(rs.([]byte))
					} else if v, ok := rs.(time.Time); ok {
						data = v.Format(helper.TimeFormat) // TODO 根据时区返回时间
					} else {
						pine.Logger().Debug("字段"+column.ColumnName+"类型", reflect.ValueOf(rs).Type().String())
						data = fmt.Sprintf("%s", rs)
					}
				}
			} else {
				data = p.getBlobDisplay(rs, column, j, ed)
			}

			grid += "<td nowrap=\"nowrap\" class=\"" + class + "\">" + data + "</td>"
		}
		grid += "</tr>\n"
	}

	numRows := len(datas)

	grid += "</tbody></table></div>"

	editTableName := p.Session().Get("select.unique_table")

	gridTitle := T("Query Results")
	if editTableName != "" {
		gridTitle = strings.Replace(T("Data for {{TABLE}}"), "{{TABLE}}", gridTitle, 1)
	}

	grid += "<div id=\"title\">" + gridTitle + "</div>"

	var total_records, current_page, total_pages int
	var message string

	if p.Session().Get("select.can_limit") == "true" {
		if p.Session().Get("select.limit") == "true" {
			total_records = p.selectSessionInt("count")
			total_pages = int(math.Ceil(float64(total_records) / float64(recordLimitInt)))
			current_page = p.selectSessionInt("page")
			from := (current_page-1)*recordLimitInt + 1
			to := from + numRows - 1
			message = "<div class='numrec'>" + strReplace([]string{"{{START}}", "{{END}}"}, []string{p.n2s(from), p.n2s(to)}, T("Showing records {{START}} - {{END}}")) + "</div>"
		} else {
			total_records = numRows
			total_pages = 1
			current_page = 1
			if recordLimitInt > 0 && total_records > recordLimitInt {
				message = "<div class='numrec'>" + strReplace([]string{"{{MAX}}"}, []string{p.n2s(recordLimitInt)}, T("Showing first {{MAX}} records only")) + "!</div>"
			}
		}
	} else {
		total_records = numRows
		total_pages = 1
		current_page = 1
		message = ""
	}

	js := "<script type=\"text/javascript\" language=\"javascript\">\n"
	if len(pkey) > 0 {
		js += "parent.editKey = " + jsonEncode(&pkey) + ";\n"
	} else if len(ukey) > 0 {
		js += "parent.editKey = " + jsonEncode(&ukey) + ";\n"
	} else {
		js += "parent.editKey = [];\n"
	}

	js += "parent.editTableName = \"" + editTableName + "\";\n"
	js += "parent.fieldInfo = " + string(fieldInfo) + ";\n"
	js += "parent.queryID = '" + helper.GetMd5(p.selectSession(query)) + "';\n"
	js += "parent.totalRecords = " + p.n2s(total_records) + ";\n"
	js += "parent.totalPages = " + p.n2s(total_pages) + ";\n"
	js += "parent.currentPage = " + p.n2s(current_page) + ";\n"
	if p.selectSession("table") != "" {
		js += "parent.queryType = \"table\";\n"
	} else {
		js += "parent.queryType = \"query\";\n"
	}
	js += "parent.transferResultGrid(" + p.n2s(numRows) + ", '0', \"" + message + "\");\n"
	js += "parent.addCmdHistory(\"" + strings.ReplaceAll(p.selectSession("query"), "\\n\\r", "<br/>") + "\", 1);\n"
	js += "parent.resetFrame();\n"
	js += "</script>\n"

	grid += js

	return grid

}

func (p *Process) selectSessionInt(key string) int {
	d, _ := strconv.Atoi(p.Session().Get("select." + key))
	return d
}

func (p *Process) selectSession(key string) string {
	return p.Session().Get("select." + key)
}

func (p *Process) getBlobDisplay(rs interface{}, info *Column, numRecord int, editable bool) string {
	binary := info.DataType == "binary"
	span := "<span class=\"i\">"
	var length int
	var size string
	if rs == nil {
		length = 0
		size = "0 B"
		span += "NULL"
	} else {
		length := len(rs.([]byte))
		size = formatBytes(int64(length))
		if length == 0 {
			span += "&nbsp;"
		} else {
			if MAX_TEXT_LENGTH_DISPLAY >= length {
				span += string(rs.([]byte))
			} else if binary {
				span += strings.Replace(T("Blob Data [{{SIZE}}]"), "{{SIZE}}", size, 1)
			} else {
				span += strings.Replace(T("Text Data [{{SIZE}}]"), "{{SIZE}}", size, 1)
			}
		}
	}

	extra, btype := "", "text"
	if binary {
		pine.Logger().Debug("处理为binary")
	}

	span += "</span>"

	if binary {
		//$span .= "<span title=\"" . str_replace('{{NUM}}', $length, __('Click to view/edit column data [{{NUM}} Bytes]')). "\" class=\"blob $btype\" $extra>&nbsp;</span>";
		//return $span;
	}

	span += "<span class=\"d\" style=\"display:none\">" + string(rs.([]byte)) + "</span>"

	if !editable && rs != nil && MAX_TEXT_LENGTH_DISPLAY < length {
		extra = `onclick="vwTxt(this, &quot;` + size + `&quot;, '` + btype + `')"`
		span += "<span title=\"" + strings.ReplaceAll("Click to view/edit column data [{{NUM}} Bytes]", "{{NUM}}", strconv.Itoa(length)) + "\" class=\"blob " + btype + "\" " + extra + ">&nbsp;</span>"
	}

	return span
}

func (p *Process) createDbInfoGrid(message string) string {
	grid := "<div id='results'>"
	grid += "<div class='message ui-state-default'>" + message + "<span style='float:right'>" + T("Quick Search") +
		"&nbsp;<input type=\"text\" id=\"quick-info-search\" maxlength=\"50\" /></div>"

	grid += "<table cellspacing='0' width='100%' border='0' class='results' id='infoTable'><thead>\n"

	//fields := p.getFieldInfo()
	dbname := p.Session().Get("db.name")
	grid += "<tr id='fhead'><th class='th index'><div>#</div></th>\n"

	headers := GetTableInfoHeaders()

	// 数字类型的字段
	numericFields := []string{"Version", "Rows", "Avg_row_length", "Data_length", "Max_data_length", "Index_length", "Auto_increment"}

	for _, header := range headers {
		cls, dsrt := "th", "text"
		if ok, _ := helper.InArray(header, numericFields); ok {
			cls = "th_numeric"
			dsrt = "numeric"
		}
		grid += "<th nowrap=\"nowrap\" class='" + cls + "' data-sort='" + dsrt + "'><div>" + header + "</div></th>\n"
	}

	grid += "</tr></thead><tbody>\n"

	// 遍历数据
	tables := getTables(p.db, dbname)

	// ------------ print data -----------
	for j, table := range tables {
		grid += `<tr id="rc` + strconv.Itoa(j) + `" class="row"><td class="tj">` + strconv.Itoa(j+1) + `</td>`

		vs := reflect.ValueOf(&table).Elem()

		for i := 0; i < vs.NumField(); i++ {
			rs := vs.Field(i).Interface()
			class := "tl"
			data := ""
			if rs == nil {
				class = "tnl"
				data = "NULL"
			} else if strings.Contains(strings.ToLower(vs.Field(i).String()), "int") {
				class = "tr"
			}

			switch vs.Field(i).Kind() {
			case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
				data = fmt.Sprintf("%d", vs.Field(i).Int())
			case reflect.Uint, reflect.Uint16, reflect.Uint8, reflect.Uint32, reflect.Uint64:
				data = fmt.Sprintf("%d", vs.Field(i).Uint())
			case reflect.String:
				data = vs.Field(i).String()
			default:
				sqlField := vs.Field(i)
				// 其他类型格式兼容
				if strings.Contains(vs.Field(i).Type().String(), "*sql.Null") {
					sqlField = sqlField.Elem()
				}

				if !sqlField.IsValid() || sqlField.IsZero() {
					data, class = "NULL", "tnl"
				} else {
					if strings.Contains(sqlField.Type().String(), "sql.Null") {
						val := sqlField.Field(0).Interface()
						if v, ok := val.(time.Time); ok {
							data = v.Format(helper.TimeFormat)
						} else if v, ok := val.(int64); ok {
							data = fmt.Sprintf("%d", v)
						} else if v, ok := val.(int32); ok {
							data = fmt.Sprintf("%d", v)
						} else if v, ok := val.(float64); ok {
							data = strconv.FormatFloat(v, 'f', 6, 64)
						} else if v, ok := val.(bool); ok {
							data = strconv.FormatBool(v)
						} else if v, ok := val.(string); ok {
							data = v
						} else {
							data = fmt.Sprintf("%s", vs.Field(i).Interface())
						}
					} else {
						data = fmt.Sprintf("%s", vs.Field(i).Interface())
					}
				}
			}

			class += " text"
			grid += "<td nowrap=\"nowrap\" id=\"r" + strconv.Itoa(j) + "f" + strconv.Itoa(i) + "\" class=\"" + class + "\">" + data + "</td>\n"
		}
		grid += "</tr>\n"
	}

	grid += "</tbody></table>"
	grid += "</div>"

	grid += "<script type=\"text/javascript\" language=\"javascript\">\n"
	grid += "parent.transferInfoMessage();\n"
	grid += "parent.resetFrame();\n"
	grid += "</script>"
	return grid
}

func (p *Process) getFieldInfo() []*Column {
	query := "SELECT * FROM information_schema.columns WHERE table_schema = '" + p.dbname + "' AND table_name = '" + p.formData.table + "'"
	var columns []*Column
	if err := p.db.Select(&columns, query); err != nil {
		pine.Logger().Warning("获取表字段信息失败", err)
	}
	for _, column := range columns {
		column.Fill()
	}
	return columns
}

func (p *Process) SelectVersion() {
	var variables []Variable
	p.db.Select(&variables, "SHOW VARIABLES LIKE 'version%'")
	if len(variables) > 0 {
		for _, variable := range variables {
			if variable.VariableName == "version" {
				p.Session().Set("db.version", strings.Split(variable.Value, ".")[0])
				p.Session().Set("db.version_full", variable.Value)
			} else if variable.VariableName == "version_comment" {
				p.Session().Set("db.version_comment", variable.Value)
			}
		}
	}
}

// 获取创建命令语句
func (p *Process) getCreateCommand(typo, name string) string {
	sql, cmd := "", ""
	if typo == "trigger" {
		sql = "show triggers where `trigger` = '" + name + "'"
	} else {
		sql = "show create " + typo + " `" + name + "`"
	}
	var createCommand []CreateCommand
	if err := p.db.Select(&createCommand, sql); err != nil {
		pine.Logger().Warning("查询创建语句异常", err)
	}
	p.lastSQL = sql
	if len(createCommand) > 0 && createCommand[0].CreateTable != "" {
		cmd = createCommand[0].CreateTable
	}
	return cmd
}

func (p *Process) GetDropCommand(table string) string {
	return "drop table if exists '" + table + "'"
}

func (p *Process) GetTruncateCommand(table string) string {
	return "truncate table '" + table + "'"
}

func (p *Process) GetEngines() []string {
	var engines []Engine
	if err := p.db.Select(&engines, "show engines"); err != nil {
		pine.Logger().Warning("获取存储引擎失败", err)
	}
	var ret []string
	for _, engine := range engines {
		if engine.Support != "NO" {
			ret = append(ret, engine.Engine)
		}
	}
	return ret
}

func (p *Process) GetCharsets() []string {
	var ret []string
	if rows, err := p.db.Queryx("show character set"); err != nil {
		pine.Logger().Warning("获取字符集失败", err)
	} else {
		for rows.Next() {
			results := make(map[string]interface{})
			rows.MapScan(results)
			ret = append(ret, string(results["Charset"].([]byte)))
		}
		sort.Strings(ret)
	}
	return ret
}

func (p *Process) GetCollations() []string {
	var collations []Collation
	p.db.Select(&collations, "show collation")
	var ret []string
	for _, collation := range collations {
		ret = append(ret, collation.Collation)
	}
	sort.Strings(ret)
	return ret
}

func (p *Process) Logout() string {
	p.Session().Destroy() // 销毁session
	return string(p.Render("logout", nil))
}

// Infovars 服务器变量
func (p *Process) Infovars() string {
	return p.createSimpleGrid(T("Server Variables"), "SHOW VARIABLES")
}

// Search 搜索数据
func (p *Process) Search() string {
	return ""

}

// Options 选项配置 用于配置系统内数据
func (p *Process) Options() string {
	pk := p.PostString("p", "ui")

	pagesort := []string{"results", "editing", "misc", "ui"}

	pages := pine.H{
		"results": T("Results"),
		"editing": T("Record Editing"),
		"misc":    T("Miscellaneous"),
		"ui":      T("Interface"),
	}

	if _, exist := pages[pk]; !exist {
		pk = "ui"
	}

	content := string(p.Render("options/"+pk, nil))

	lis := ""
	for _, x := range pagesort {
		y := pages[x]
		if pk == x {
			lis += "<li class=\"current\"><img border=\"0\" align=\"absmiddle\" src='/mywebsql/img/options/o_" + x + ".gif' alt=\"\" />" + y.(string) + "</li>"
		} else {
			lis += "<li><a href=\"#" + x + "\"><img border=\"0\" align=\"absmiddle\" src='/mywebsql/img/options/o_" + x + ".gif' alt=\"\" />" + y.(string) + "</a></li>"
		}
	}

	// <?php
	// foreach($data['pages'] as $x=>$y) {
	// 		if ($data['page'] == $x)
	// 			echo "<li class=\"current\"><img border=\"0\" align=\"absmiddle\" src='img/options/o_$x".".gif' alt=\"\" />$y</li>";
	// 		else
	// 			echo "<li><a href=\"#$x\"><img border=\"0\" align=\"absmiddle\" src='img/options/o_$x".".gif' alt=\"\" />$y</a></li>";
	// 	}
	// ?>

	return string(p.Render("options", pine.H{
		"CONTENT": template.HTML(content),
		"lis":     template.HTML(lis),
		"data": pine.H{
			"pages": pages,
			"page":  pk,
		},
	}))

}

// Queryall
func (p *Process) Queryall() string {
	return ""

}

// Truncate 截断数据表
func (p *Process) Truncate() string {
	return ""

}

// Drop 删除数据表
func (p *Process) Drop() string {
	return ""

}

// Rename 重命名
func (p *Process) Rename() string {
	return ""

}

// Dbrepair 修复表
func (p *Process) Dbrepair() string {
	// TODO tables 必须得数组
	if p.FormValue("optype") != "" && p.FormValue("tables[]") != "" {
		return p.checkTables()
	} else {
		tableStrs := getTables(p.db, p.dbname)
		var tables []string
		for _, v := range tableStrs {
			tables = append(tables, v.Name)
		}
		byts, _ := json.Marshal(&tables)
		extra := ""
		if len(tables) > 0 {
			extra += "$('#db_objects').html('');\n"
			extra += "uiShowObjectList(tables, 'tables', '" + T("Tables") + "');\n"
		}

		return string(p.Render("dbrepair", pine.H{
			"tables":  string(byts),
			"extraJs": template.HTML(extra)}))
	}
}

func (p *Process) checkTables() string {
	typo := p.FormValue("optype")
	options := map[string]interface{}{}

	postdata := p.PostData()
	pine.Logger().Debug("postData", postdata)
	if p.FormValue("skiplog") == "on" {
		options["skiplog"] = true
	} else {
		options["skiplog"] = false
	}
	options["checktype"] = p.PostValue("checktype")
	options["repairtype"] = postdata["repairtype"]
	tables := postdata["tables[]"]
	checker := NewTableChecker(p.db)
	checker.SetOperation(typo)
	checker.SetOptions(options)
	checker.SetTables(tables)
	checker.Runcheck()
	results := checker.GetResults()
	byts, _ := json.Marshal(&results)
	return string(p.Render("dbrepair_results", pine.H{"RESULTS": template.HTML(string(byts))}))
}

// Dbcreate 创建表
func (p *Process) Dbcreate() string {
	p.removeSelectSession()

	name := p.FormValue("name")
	dbSelect := p.FormValue("query")

	sql := "create database `" + name + "`"

	if _, err := p.db.Exec(sql); err != nil {
		return p.createErrorGrid("", err)
	}
	redirect := "0"

	if dbSelect != "" {
		p.Session().Set("db.change", "true")
		p.Session().Set("db.name", name)
		redirect = "1"
	}

	return string(p.Render("dbcreate", pine.H{
		"DB_NAME":  name,
		"SQL":      sql,
		"TIME":     0,
		"REDIRECT": redirect,
	}))

}

// Tableinsert 插入表数据
func (p *Process) Tableinsert() string {
	return ""

}

// Tableupdate 更新表数据
func (p *Process) Tableupdate() string {
	return ""
}

// Showcreate 展示创建语句
func (p *Process) Showcreate() string {
	dels := []string{"result", "pkey", "ukey", "mkey", "unique_table"}
	for _, v := range dels {
		p.Session().Remove("select." + v)
	}

	cmd := p.sanitizeCreateCommand(p.getCreateCommand(p.formData.id, p.formData.name))

	v := pine.H{
		"TYPE":    p.formData.id,
		"NAME":    p.formData.name,
		"COMMAND": template.HTML(cmd),
		"TIME":    0,
		"SQL":     template.HTML(p.lastSQL),
		"MESSAGE": strReplace([]string{"{{TYPE}}", "{{NAME}}"}, []string{p.formData.id, p.formData.name}, T("Create command for {{TYPE}} {{NAME}}")),
	}

	return string(p.Render("showcreate", v))
}

func (p *Process) Describe() string {
	if p.dbname == "" || p.formData.name == "" {
		return string(p.Render("invalid_request", nil))
	}

	p.lastSQL = "DESCRIBE `" + p.formData.name + "`"
	return p.createSimpleGrid(T("Table Description")+": ["+p.formData.name+"]", p.lastSQL)
}

func (p *Process) Copy() string {
	if p.formData.name == "" || p.formData.query == "" {
		return p.createErrorGrid("", errors.New("参数错误或不足"))
	}

}

func (p *Process) copyObject() string {

}

// Processes 进程管理器
func (p *Process) Processes() string {
	html := "<link href='/mywebsql/cache?css=theme,default,alerts,results' rel=\"stylesheet\" />\n"

	typo := "message ui-state-highlight"

	msg := T("Select a process and click the button to kill the process")

	prcid := p.FormValue("prcid") // TODO 需要支持传递数组参数

	if prcid != "" {
		//prcids := []string{prcid}
		// TODO 提交杀死进程
	}

	return html + p.displayProcessList(msg, typo)

}

func (p *Process) displayProcessList(msg, typo string) string {
	html := "<input type='hidden' name='q' value='wrkfrm' />"
	html += "<input type='hidden' name='type' value='processes' />"
	html += "<input type='hidden' name='id' value='' />"
	html += "<table border=0 cellspacing=2 cellpadding=2 width='100%'>"
	if len(msg) > 0 {
		html += "<tr><td height=\"25\"><div class=\"" + typo + "\">" + msg + "</div></td></tr>"
	}
	html += "<tr><td colspan=2 valign=top>"

	if rows, err := p.db.Queryx("show full processlist"); err != nil {
		pine.Logger().Warning("获取进程列表失败", err)
		html += T("Failed to get process list")
	} else {
		html += "<table class='results postsort' border=0 cellspacing=1 cellpadding=2 width='100%' id='processes'><tbody>"
		html += "<tr id='fhead'><th></th><th class='th'>" + T("Process ID") + "</th><th class='th'>" + T("Command") + "</th><th class='th'>" + T("Time") +
			"</th><th class='th'>" + T("Info") + "</th></tr>"

		for rows.Next() {
			results := make(map[string]interface{})
			err = rows.MapScan(results)

			id := string(results["Id"].([]byte))
			command := string(results["Command"].([]byte))
			var timed string
			if results["Time"] != nil {
				timed = string(results["Time"].([]byte))
			}
			var info string
			if results["Info"] != nil {
				info = string(results["Info"].([]byte))
			}

			html += "<tr class='row'><td class=\"tch\"><input type=\"checkbox\" name='prcid[]' value='" + id + "' /></td>" +
				"<td class='tl'>" + id + "</td><td class='tl'>" + command + "</td>" +
				"<td class='tl'>" + timed + "</td><td class='tl'>" + info + "</td></tr>"
		}

		html += "</tbody></table>"

		html += "<tr><td colspan=2 align=right><div id=\"popup_buttons\"><input type='submit' id=\"btn_kill\" name='btn_kill' value='" + T("Kill Process") + "' /></div></td></tr>"

		html += "<script type=\"text/javascript\" language='javascript' src=\"/mywebsql/cache?script=common,jquery,ui,query,sorttable,tables\"></script>\n"

		html += `<script type="text/javascript" language='javascript'>
			window.title = "` + T("Process Manager") + `";
			$('#btn_kill').button().click(function() { document.frmquery.submit(); });
			setupTable('processes', {sortable:'inline', highlight:true, selectable:true});
			</script>`
	}
	return html
}

// Help 帮助页面
func (p *Process) Help() string {
	page := p.FormValue("p")
	if page == "" {
		page = "queries"
	}
	pages := map[string]string{
		"queries":  "Executing queries",
		"results":  "Working with results",
		"keyboard": "Keyboard shortcuts",
		"prefs":    "Preferences",
		"misc":     "Miscellaneous",
		"credits":  "Credits",
		"about":    "About",
	}
	if _, ok := pages[page]; !ok {
		page = "queries"
	}
	contents := p.Render("help/"+page, nil)
	return string(p.Render("help", pine.H{"pages": pages, "page": page, "contents": contents}))
}

// Createtbl 创建表
func (p *Process) Createtbl() string {
	action := p.formData.id
	var html string
	if action == "create" || action == "alter" {

	} else {
		engines := Html.ArrayToOptions(p.GetEngines(), "", "Default")
		charsets := Html.ArrayToOptions(p.GetCharsets(), "", "Default")
		collations := Html.ArrayToOptions(p.GetCollations(), "", "Default")
		html = string(p.Render("editable", pine.H{
			"ID":          action,
			"MESSAGE":     "",
			"ROWINFO":     "[]",
			"ALTER_TABLE": "false",
			"TABLE_NAME":  "",
			"ENGINE":      engines,
			"CHARSET":     charsets,
			"COLLATION":   collations,
			"COMMENT":     "",
		}))
	}
	return html
}

// Backup 备份数据
func (p *Process) Backup() string {
	return ""
}

// removeSelectSession 移除选择状态的session数据
func (p *Process) removeSelectSession() {
	sessionKeys := []string{"result", "pkey", "ukey", "mkey", "unique_table"}
	for _, v := range sessionKeys {
		p.Session().Remove("select." + v)
	}
}

// Render 渲染模板
func (p *Process) Render(name string, data pine.H) []byte {
	var byts []byte
	var err error
	if byts, err = GetPlush().Exec(name+".php", data); err != nil {
		pine.Logger().Warning("渲染模板"+name+".php失败", err)
	}
	return byts
}

func (p *Process) row2arrMap(rows *sql.Rows) ([]map[string]interface{}, error) {
	columns, _ := rows.Columns()
	columnLength := len(columns)
	cache := make([]interface{}, columnLength)
	for index := range cache {
		var a interface{}
		cache[index] = &a
	}
	var list []map[string]interface{}
	for rows.Next() {
		_ = rows.Scan(cache...)
		item := make(map[string]interface{})
		for i, data := range cache {
			item[columns[i]] = *data.(*interface{})
		}
		list = append(list, item)
	}
	err := rows.Close()
	return list, err
}
