package common

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xiusin/pine"
	"github.com/xiusin/pinecms/src/common/helper"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Process struct {
	db    *sqlx.DB
	ctx   *pine.Context
	query string
}

func InitProcess(db *sqlx.DB, ctx *pine.Context, query string) *Process {
	p := &Process{db, ctx, query}
	p.SelectVersion()
	return p
}

func (p *Process) Info() string {
	var html []byte
	dbname := p.ctx.Session().Get("db.name")
	if dbname != "" {
		tip := T("Database summary") + ": [" + dbname + "]"
		html = []byte(p.createDbInfoGrid(tip))
	} else {
		html = []byte(p.createResultGrid())
	}
	return string(html)
}

// Showinfo
func (p *Process) Showinfo() string {
	typo := p.ctx.PostString("id")
	if typo == "table" || typo == "view" {
		id, page := "table", strings.Trim(p.ctx.FormValue("name")," \r\n\t;")
		querySql := ""
		if  id == "table" {
			querySql = p.selectFromTable(id)
		}

		fmt.Println(page, querySql)
	//Session::del('select');
	//	include('query.php');
	} else {

	}
	//type: showinfo
//id: table
//name:sys_config
//query:
	return "show info"
}

func (p *Process) selectFromTable(id , pageNum string) string  {
	record_limit := p.ctx.GetCookie("res-max-count")
	if len(record_limit) == 0 {
		record_limit = fmt.Sprintf("%d", MAX_RECORD_TO_DISPLAY)
	}

	if len(pageNum) == 0 {
		limit_applied :=  p.ctx.GetString("select.limit")
		query :=  p.ctx.GetString("select.query")
		table :=  p.ctx.GetString("select.table")
		count :=  p.ctx.GetString("select.count")

		queryType := p.getQueryType(query)
		if !queryType["result"] || table == "" {
			return ""
		}

		total_pages :=
	}


	return ""
}

func (p *Process) simpleQuery()  {

}

func (p *Process)getQueryType(query string) map[string]bool {
	types := map[string]bool {"result": false, "can_limit": false, "has_limit": false, "update": false}
	q := strings.ToLower(query[:7])

	if q == "explain" || q[:6] == "select" || q[:4] == "desc" || q[:4] == "show" || q[:4] == "help"  {
		types["result"] = true

		if q[:6] == "select" {
			types["can_limit"] = true
		}
	}
	//define('LIMIT_REGEXP', '/(.*)[\s]+(limit[\s]+[\d]+([\s]*(,|offset)[\s]*[\d]+)?)$/is');
	//define('SORT_REGEXP', '/(.*)[\s]+(ORDER[\s]+BY[\s]+([a-zA-z0-9\._]+|`.*`|\'.*\'|".*")\s*(ASC|DESC)?(\s*\,\s*([a-zA-z0-9\._]+|`.*`|\'.*\'|".*")\s*(ASC|DESC)?)*)$/is');
	match ,err := regexp.MatchString("(.*)[\\s]+(limit[\\s]+[\\d]+([\\s]*(,|offset)[\\s]*[\\d]+)?)$", query)
	if err != nil {
		pine.Logger().Error("匹配sql", err)
	}
	if match && strings.Contains(strings.ToLower(query), "limit") {
		types["has_limit"] = true
	}
	return types
}

func (p *Process) exec(module string) string {
	val := reflect.ValueOf(&p).Elem().MethodByName(helper.UcFirst(module))
	if !val.IsValid() {
		return p.createErrorGrid("")
	}
	return val.Call([]reflect.Value{})[0].String()
}

func (p *Process) QueryVariables() []Variable {
	query := "SHOW VARIABLES"

	var variables []Variable
	p.db.Select(&variables, query)
	return variables
}

func (p *Process) createErrorGrid(query string) string {
	return "错误表格，需要渲染页面"
}

func (p *Process) getTemplateSQL(templateName string) string {
	content, err := GetPlush().Exec("template/"+templateName, nil)
	if err != nil {
		return ""
	}
	return string(content)
}

func (p *Process) createResultGrid() string {
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

	if dbname := p.ctx.Session().Get("db.name"); dbname == "" {
		v["JS"] = `parent.$("#main-menu").find(".db").hide();`
	}
	content, _ := GetPlush().Exec("infoserver.php", v)

	grid += string(content)

	return grid
}

func (p *Process) createDbInfoGrid(message string) string {
	grid := "<div id='results'>"
	grid += "<div class='message ui-state-default'>" + message + "<span style='float:right'>" + T("Quick Search") +
		"&nbsp;<input type=\"text\" id=\"quick-info-search\" maxlength=\"50\" /></div>"

	grid += "<table cellspacing='0' width='100%' border='0' class='results' id='infoTable'><thead>\n"

	//fields := p.getFieldInfo()
	dbname := p.ctx.Session().Get("db.name")
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

func (p *Process) getFieldInfo() []Column {
	dbname := p.ctx.Session().Get("db.name")
	tablename := strings.Trim(p.ctx.PostValue("name"), " \r\n\t;")

	query := "SELECT * FROM information_schema.columns WHERE table_schema = '" + dbname + "' AND table_name = '" + tablename + "'"

	var columns []Column
	p.db.Select(&columns, query)

	return columns
}

func (p *Process) SelectVersion() {
	var variables []Variable
	p.db.Select(&variables, "SHOW VARIABLES LIKE 'version%'")
	if len(variables) > 0 {
		for _, variable := range variables {
			if variable.VariableName == "version" {
				p.ctx.Session().Set("db.version", strings.Split(variable.Value, ".")[0])
				p.ctx.Session().Set("db.version_full", variable.Value)
			} else if variable.VariableName == "version_comment" {
				p.ctx.Session().Set("db.version_comment", variable.Value)
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
