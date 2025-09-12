package dede

import (
	"bytes"
	"fmt"
	"github.com/xiusin/pine"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xiusin/pinecms/src/common/helper"
)

var dedeTplCmd = &cobra.Command{
	Use:   "tpl",
	Short: "dede模板转换为pinecms模板",
	Long: `1. 快速转换织梦模板为pinecms模板
2. 只支持标签级转换, 其他的转换完成根据错误提醒自行修复
`,
	Run: func(cmd *cobra.Command, args []string) {
		dedepath, _ := cmd.Flags().GetString("dedepath")
		pinepath, _ := cmd.Flags().GetString("pinepath")
		dirname, _ := cmd.Flags().GetString("dirname")
		force, _ := cmd.Flags().GetBool("force")
		if dedepath == "" || pinepath == "" {
			_ = cmd.Usage()
			return
		}
		dedepath = strings.TrimRight(dedepath, "\\/")
		var theme string
		if dirname != "" {
			theme = dirname
		} else {
			theme = filepath.Base(dedepath)
		}

		fs, err := os.Stat(pinepath)
		if err != nil {
			pine.Logger().Error(err.Error())
			return
		}

		if !fs.IsDir() {
			pine.Logger().Error("您输入的pinepath参数非目录地址")
			return
		}
		themePath := filepath.Join(pinepath, theme)
		if force { // 是否强制删除目录
			err := os.RemoveAll(themePath)
			if err != nil {
				pine.Logger().Error(err.Error())
				return
			}
		}
		err = os.Mkdir(themePath, os.ModePerm)
		if err != nil {
			pine.Logger().Error(err.Error())
			return
		}

		err = filepath.Walk(dedepath, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			ext := strings.ToLower(filepath.Ext(path))
			if ext != ".html" && ext != ".htm" {
				return nil
			}
			// 读取文件内容
			byts, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			relativeFilePath, _ := filepath.Rel(dedepath, path)

			parser := &Parser{
				data: byts,
				dst:  filepath.Join(themePath, strings.Replace(relativeFilePath, ext, ".jet", 1)),
				src:  relativeFilePath,
			}
			if err := parser.Start(); err != nil {
				pine.Logger().Error(err.Error())
			}
			return nil
		})
		if err != nil {
			pine.Logger().Error(err.Error())
			return
		}
		fmt.Printf("%s\n", `

SUCCESS!

1. 导入织梦模板成功, 此操作只能解决80%的繁琐替换过程
2. 测试各个模板数据根据错误修改相应的标签
3. Enjoy! 😃

注意: 导入不保证完全正确,建议进行模型设置(固化字段被设置为text类型)`)
	},
}

func init() {
	Cmd.AddCommand(dedeTplCmd)
	dedeTplCmd.Flags().String("dedepath", "", "织梦模板路径(填入具体主题地址)")
	dedeTplCmd.Flags().String("pinepath", helper.GetRootPath()+"/resources/themes/", "pine的主题路径")
	dedeTplCmd.Flags().Bool("force", false, "是否强制删除pinecms同名主题目录")
	dedeTplCmd.Flags().String("dirname", "", "默认基于pinepath生成目录")
}

type Parser struct {
	data []byte
	dst  string
	src  string
}

func (p *Parser) appendTag() {
}
func (p *Parser) parseInclude() {
	p.data = regexp.MustCompile("{dede:include\\s+filename=[\"|'](.+)?.html?[\"|']\\s*/}").
		ReplaceAll(p.data, []byte(`{{include "$1.jet"}}`))
}

//func (p *Parser) parseTagBlock()  {
//	ts := `{dede:channel type='son' typeid='127''>[field:typename/]</a>{/dede:channel}`
//	regexp.MustCompile("{dede:(.+)?\\s+(.+)?}(.+)?{/dede:$1}").ReplaceAllFunc([]byte(ts), func(i []byte) []byte {
//		fmt.Println(string(i))
//		return []byte{}
//	})
//}

func (p *Parser) parseGlobalSomeField() {
	var fields = map[string]string{ //
		"cfg_powerby":     "site_copyright",
		"cfg_keywords":    "site_keywords",
		"cfg_description": "site_description",
		"cfg_beian":       "site_icp",
		"cfg_webname":     "site_name",
		"cfg_cmsurl":      "site_host",
	}
	p.data = regexp.MustCompile("{dede:global\\.[^\\/]+/}").ReplaceAllFunc(p.data, func(bts []byte) []byte {
		bts = bytes.TrimPrefix(bts, []byte("{dede:global."))
		bts = bytes.TrimSuffix(bts, []byte("/}"))
		fieldInfo := strings.SplitN(string(bts), " ", 2)
		field := fieldInfo[0]
		if val, ok := fields[field]; ok {
			return []byte(`{{global["` + val + `"]}}`)
		}
		return []byte(`{{global["` + field + `"]}}`)
	})
}

func replaceFieldName(field string) string {
	switch field {
	case "picname", "litpic":
		field = "thumb"
	case "pubdate":
		field = "pubtime"
	}
	return field
}

// 解析页面级别的织梦标签
func (p *Parser) parsePageTagField() {
	p.data = regexp.MustCompile("{dede:field[^}]+?/}").ReplaceAllFunc(p.data, func(bts []byte) []byte {
		matched := string(bts)
		matched = strings.TrimPrefix(matched, "{dede:field")
		matched = strings.TrimSuffix(matched, "/}")
		var field string
		if strings.HasPrefix(matched, ".") {
			field = strings.TrimPrefix(matched, ".")
		} else {
			field = strings.Trim(matched, ` `)
		}
		fs := strings.SplitN(field, " ", 2)
		field = strings.TrimPrefix(fs[0], `name=`)
		field = replaceFieldName(strings.Trim(field, `'" `))
		var unsafe string
		if field == "content" || field == "body" {
			unsafe = " | unsafe"
		}
		switch field {
		case "typename": //todo 前端使用标签实现.
			field = `(isset(field) && field!=nil) ? (field["` + field + `"] ? field["` + field + `"] : field["Catname"]) : (isset(.Field["` + field + `"]) ? .Field["` + field + `"] : .Field["Catname"])`
		case "typeurl", "typelink":
			field = `(isset(field) && field!=nil) ? (field["` + field + `"] ? field["` + field + `"] : field["Url"]) : (isset(.Field["` + field + `"]) ? .Field["` + field + `"] : .Field["Url"])`
		case "body":

			field = `(isset(field) && field!=nil) ? (field["` + field + `"] ? field["` + field + `"] :( field["Content"] ?  field["Content"] : field["content"])) : (isset(.Field["` + field + `"]) ? .Field["` + field + `"] : (.Field["Content"] ? .Field["Content"] : .Field["content"]))`
		default:
			strArry := []rune(field)
			if strArry[0] >= 97 && strArry[0] <= 122 {
				strArry[0] -= 32
			}
			field = `(!isset(field)||field==nil) ? (isset(.Field["` + field + `"]) ? .Field["` + field + `"] : .Field["` + string(strArry) + `"]) : field["` + field + `"]`
		}
		if len(fs) > 1 {
			function := strings.Trim(fs[1], " ")
			if strings.HasPrefix(function, "function=") {
				function = strings.TrimPrefix(function, "function=")
				function = strings.Trim(function, `'"`) // 按字符trim,直到不包括cutset
				if !strings.HasPrefix(function, "html2text") {
					if strings.HasPrefix(function, "GetDateTimeMk") {
						function = strings.ReplaceAll(function, "GetDateTimeMk", "format_time")
					}
					return []byte(`{{` + strings.ReplaceAll(function, "@me", field) + unsafe + `}}`)
				}
			}
		}
		return []byte(`{{` + field + unsafe + `}}`)
	})
}

// 解析标签
func (p *Parser) parseDedeBlockTags() {
	// 开始标签视情况替换内容(唯一恶心的地方就是标签类型非强类型)
	p.data = regexp.MustCompile("{dede:(.+?)\\s+(?s:.+?)}").ReplaceAllFunc(p.data, func(i []byte) []byte {
		matched := string(i)
		if strings.Contains(matched, "runphp") ||
			strings.Contains(matched, "include") ||
			strings.Contains(matched, "global") ||
			strings.Contains(matched, "sql=") {
			return i
		}
		matched = strings.TrimPrefix(matched, "{dede:")
		matched = strings.TrimRight(matched, "}/")
		// arclist
		fs := strings.SplitN(matched, " ", 2)

		if len(fs) == 1 {
			fs = append(fs, " ")
		}
		fs[1] = regexp.MustCompile(`\s+`).ReplaceAllString(fs[1], " ")
		var pineTagAttrs []string
		tag, ok := p.getPineTag(fs[0])
		if !ok {
			return i
		}

		if tag == "channel" {
			fs[1] = regexp.MustCompile(`currentstyle="(?s:.+?)"`).ReplaceAllStringFunc(fs[1], func(s string) string {
				pineTagAttrs = append(pineTagAttrs, strings.ReplaceAll(s, "\n", ""))
				return ""
			})
		}
		pineTagAttrs = append(pineTagAttrs, p.parseAttrs(fs[1])...)

		pine.Logger().Info("%s 替换标签内容 \n%s \n↓\n%s\n\n", p.src, string(i), `{{yield `+tag+`(`+strings.Join(pineTagAttrs, ", ")+`) content}}`)
		return []byte(`{{yield ` + tag + `(` + strings.Join(pineTagAttrs, ", ") + `) content}}`)
	})

	// 结尾标签直接替换
	p.data = regexp.MustCompile("{/dede:(type|likeart|arclist|channel|channelartlist|prenext|flink|list)}").ReplaceAll(p.data, []byte("{{end}}"))
}

func (p *Parser) getPineTag(dedeTag string) (string, bool) {
	switch dedeTag {
	case "type", "likeart", "list", "flink", "channel", "channelartlist", "prenext":
		return dedeTag, true
	case "arclist":
		return "artlist", true
	default:
		return "", false
	}
}

func (p *Parser) parseAttrs(attrsStr string) []string {
	var pineTagAttrs []string
	attrs := strings.Split(attrsStr, " ")
	prevAttr := ""
	for _, attr := range attrs {
		if attr == "" {
			continue
		}
		if !strings.Contains(attr, "=") {
			prevAttr = attr
			continue
		}
		if strings.HasPrefix(attr, "=") {
			attr = prevAttr + attr
			prevAttr = ""
		}
		attrkv := strings.Split(attr, "=")
		k := attrkv[0]
		v := strings.Trim(attrkv[1], `'"`)
		if strings.HasPrefix(k, "att") {
			pineTagAttrs = append(pineTagAttrs, `flag="`+v+`"`)
		} else if strings.Contains(k, "id") {
			if strings.Contains(v, ",") || !regexp.MustCompile("^\\d+$").MatchString(v) {
				pineTagAttrs = append(pineTagAttrs, k+`="`+v+`"`)
			} else {
				pineTagAttrs = append(pineTagAttrs, k+`=`+v)
			}
		} else if strings.Contains(k, "len") || k == "row" {
			pineTagAttrs = append(pineTagAttrs, k+`=`+v)
		} else if k == "limit" {
			if strings.Contains(v, ",") {
				vv := strings.Split(v, ",")
				pineTagAttrs = append(pineTagAttrs, `row=`+vv[1]+``)
				pineTagAttrs = append(pineTagAttrs, `offset=`+vv[0]+``)
			} else {
				pineTagAttrs = append(pineTagAttrs, `offset=`+v)
			}
		} else if k == "orderby" {
			switch v {
			case "hot", "click":
				v = "visit_count"
			case "pubdate":
				v = "pubtime"
			}
			pineTagAttrs = append(pineTagAttrs, k+`="`+v+`"`)
		} else {
			pineTagAttrs = append(pineTagAttrs, k+`="`+v+`"`)
		}
	}
	return pineTagAttrs
}

// 解析标签内部field 如: [field:xxx /]
func (p *Parser) parseFieldsInTagBlock() {
	p.data = regexp.MustCompile("\\[field:\\s*([^\\]]+)?\\s*/\\]").ReplaceAllFunc(p.data, func(bts []byte) []byte {
		bts = bytes.TrimPrefix(bts, []byte("[field:"))
		bts = bytes.TrimSuffix(bts, []byte("/]"))
		fieldInfo := strings.SplitN(string(bts), " ", 2)
		field := replaceFieldName(fieldInfo[0])
		switch field {
		case "typename":
			field = `field["` + field + `"] ? field["` + field + `"] : field["Catname"]`
		case "typeurl", "typelink":
			field = `field["` + field + `"] ? field["` + field + `"] : field["Url"]`
		default:
			field = `field["` + field + `"]`
		}
		if len(fieldInfo) > 1 {
			function := strings.Trim(fieldInfo[1], " ")
			if strings.HasPrefix(function, "function=") {
				function = strings.TrimPrefix(function, "function=")
				function = strings.Trim(function, `'"`) // 按字符trim,直到不包括cutset
				if !strings.HasPrefix(function, "html2text") {
					// todo 貌似这里不生效
					if strings.Contains(function, "GetDateTimeMk") {
						function = strings.ReplaceAll(function, "GetDateTimeMk", "format_time")
					}
					return []byte(`{{` + strings.ReplaceAll(function, "@me", field) + `}}`)
				}
			}
		}
		return []byte(`{{` + field + `}}`)
	})
}

func (p *Parser) Start() error {
	p.parseInclude()
	p.parseFieldsInTagBlock()
	//{dede:global.cfg_webname/}
	p.parseGlobalSomeField()
	p.parsePageTagField()
	//p.parseTagBlock()
	p.parseDedeBlockTags()
	p.data = append([]byte(`{{import "tags.jet"}}
`), p.data...)
	// 生成文件
	err := os.MkdirAll(filepath.Dir(p.dst), os.ModePerm)
	if err != nil {
		return err
	}
	return os.WriteFile(p.dst, p.data, os.ModePerm)
}
