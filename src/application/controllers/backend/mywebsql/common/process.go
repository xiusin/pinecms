package common

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/common/helper"
	"math"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Process struct {
	*pine.Context
	db       *sqlx.DB
	dbname   string
	formData struct {
		id    string
		page  int
		name string
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

	p.formData.query = p.FormValue("query")
	p.formData.table = strings.Trim(p.FormValue("name"), " \r\n\t;")
	p.formData.name = p.formData.table
	p.formData.id = p.FormValue("id")
	p.formData.page = 1

	if match, _ := regexp.MatchString(`^\d+$`, p.formData.table); match {
		p.formData.page, _ = strconv.Atoi(p.formData.table)
		p.formData.table = strings.Trim(p.FormValue("query"), " \r\n\t;")
		p.formData.query = ""
	}

	return p
}

func (p *Process) Info() string {
	var html []byte
	if p.dbname != "" {
		tip := T("Database summary") + ": [" + p.dbname + "]"
		html = []byte(p.createDbInfoGrid(tip))
	} else {
		html = []byte(p.createResultGrid(""))
	}
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

func (p *Process) Query() string {
	var querySql string
	if p.formData.id == "table" {
		querySql = p.selectFromTable()
	} else {
		querySql = p.simpleQuery()
	}

	querySql = strings.Trim(querySql, " \r\n")

	pine.Logger().Debug("获取执行SQL", querySql)

	if len(querySql) == 0 {
		return p.createErrorGrid(querySql, errors.New("无法生存SQL语句"))
	}

	var html string

	p.loadDbVars()

	// TODO 区分查询是否需要返回执行结果
	if strings.ToLower(querySql[:6]) == "select" {
		queryType := p.getQueryType(querySql)
		pine.Logger().Debug("exec select sql", queryType)
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

	if !queryType["result"]  || !queryType["can_limit"]{
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

	return ""
}

func (p *Process) sortQuery(query, field string) string  {
	query = strings.Trim(query, " \r\n\t")
	//sort := ""
	sortType := p.selectSession("sort")

	if sortType == "" {
		sortType = "DESC"
	}

	//pine.Logger().Debug("执行SQL", query)
	matchs := regexp.MustCompile(LIMIT_REGEXP).FindStringSubmatch(query)
	//pine.Logger().Debug("matchs", matchs)
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
	//define('LIMIT_REGEXP', '/(.*)[\s]+(limit[\s]+[\d]+([\s]*(,|offset)[\s]*[\d]+)?)$/is');
	//define('SORT_REGEXP', '/(.*)[\s]+(ORDER[\s]+BY[\s]+([a-zA-z0-9\._]+|`.*`|\'.*\'|".*")\s*(ASC|DESC)?(\s*\,\s*([a-zA-z0-9\._]+|`.*`|\'.*\'|".*")\s*(ASC|DESC)?)*)$/is');
	match, err := regexp.MatchString("(.*)[\\s]+(limit[\\s]+[\\d]+([\\s]*(,|offset)[\\s]*[\\d]+)?)$", query)
	if err != nil {
		pine.Logger().Error("匹配sql错误", err)
	}
	if match && strings.Contains(strings.ToLower(query), "limit") {
		types["has_limit"] = true
	}
	return types
}

func (p *Process) exec(module string) string {
	val := reflect.ValueOf(&p).Elem().MethodByName(helper.UcFirst(module))
	if !val.IsValid() {
		return p.createErrorGrid("", nil)
	}
	return val.Call([]reflect.Value{})[0].String()
}

func (p *Process) QueryVariables() []Variable {
	query := "SHOW VARIABLES"

	var variables []Variable
	p.db.Select(&variables, query)
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
			v["SERVER_CHARSET"] = variable.Value
		case "character_set_database":
			v["DATABASE_CHARSET"] = variable.Value
		case "character_set_results":
			v["RESULT_CHARSET"] = variable.Value
		}
		v[variable.VariableName] = variable.Value
	}

	if dbname := p.Session().Get("db.name"); dbname == "" {
		v["JS"] = `parent.$("#main-menu").find(".db").hide();`
	}
	content, _ := GetPlush().Exec("infoserver.php", v)

	grid += string(content)

	return grid
}

func (p *Process) createSimpleGrid(message string, query string) string {
	grid := "<div id='results'>"
	grid += "<div class='message ui-state-default'>" + message + "<span style='float:right'>" + T("Quick Search") +
		"&nbsp;<input type=\"text\" id=\"quick-info-search\" maxlength=\"50\" /></div>"
	grid += "<table cellspacing='0' width='100%' border='0' class='results' id='infoTable'><thead>\n"

	f := p.getFieldInfo()

	grid += "<tr id='fhead'><th class='th index'><div>#</div></th>\n"

	for _, fn := range f {
		cls, dsrt := "th", "text"
		if fn.Numeric {
			cls = "th_numeric"
			dsrt = "numeric"
		}
		grid += "<th nowrap=\"nowrap\" class='" + cls + "' data-sort='" + dsrt + "'><div>" + fn.ColumnName + "</div></th>\n"
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

	for j, table := range datas {
		grid += `<tr id="rc` + strconv.Itoa(j) + `" class="row"><td class="tj">` + strconv.Itoa(j+1) + `</td>`

		for i, fn := range f {
			class := "tl"
			if fn.Numeric {
				class = "tr"
			}
			rs := table[fn.ColumnName]
			if rs == nil {
				class = "tnl"
			}

			if fn.Blob {
				if fn.DataType == "binary" || fn.DataType == "blob" {
					class += " blob"
				} else {
					class += " text"
				}
			}
			data := "&nbsp;"
			if rs == nil {
				data = "NULL"
			} else if rs.(string) != "" {
				data = rs.(string)
			}
			grid += "<td nowrap=\"nowrap\" id=\"r" + p.n2s(j) + "\".\"f" + p.n2s(i) + "\" class=\"" + class + "\">" + data + "</td>"
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
		pine.Logger().Debug("set select.unique_table")
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
					} else {
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
							data = strconv.FormatFloat(v, 'f', 6, 10)
						} else if v, ok := val.(bool); ok {
							data = strconv.FormatBool(v)
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

	//$j = 0;
	//while($r = $db->fetchRow(0, 'num')) {
	//	$i = 0;
	//	print "<tr id=\"rc$j\" class=\"row\">";
	//	print "<td class=\"tj\">".($j+1)."</td>";
	//
	//	foreach($r as $rs) {
	//		$class = ($rs === NULL) ? "tnl" : ($f[$i]->numeric == 1 ? "tr" : "tl");
	//if ($f[$i]->blob)
	//$class .= $f[$i]->type == 'binary' ? ' blob' : ' text';
	//
	//$data = ($rs === NULL) ? "NULL" : (($rs === "") ? "&nbsp;" : htmlspecialchars($rs));
	//
	//print "<td nowrap=\"nowrap\" id=\"r$j"."f$i\" class=\"$class\">$data</td>";
	//$i++;
	//}
	//print "</tr>\n";
	//$j++;
	//}

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
	var createCommand CreateCommand
	p.db.Select(&createCommand, sql)

	if createCommand.CreateTable != "" {
		cmd = createCommand.CreateTable
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
	p.db.Select(&engines, "show engines")
	var ret []string
	for _, engine := range engines {
		if engine.Support != "NO" {
			ret = append(ret, engine.Engine)
		}
	}
	return ret
}

func (p *Process) GetCharsets() []string {
	var charsets []Charset
	p.db.Select(&charsets, "show character set")
	var ret []string
	for _, charset := range charsets {
		ret = append(ret, charset.Charset)
	}
	sort.Strings(ret)
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
	//p.Session().Destory() // 销毁session
	return string(p.Render("logout", nil))
}

// Infovars 服务器变量
func (p *Process) Infovars() string {
	return ""

}

// Search 搜索数据
func (p *Process) Search() string {
	return ""

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
	return ""

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
	extraMsg := ""
	return extraMsg
}

// Processes 进程管理器
func (p *Process) Processes() string {
	return ""

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
