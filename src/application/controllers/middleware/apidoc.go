package middleware

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/xiusin/pine"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
)

const apiDocKey = "apiEntity"

type apiApp struct { // 记录应用分组. 页面右上角下拉
	Title  string `json:"title"`
	Path   string `json:"path,omitempty"`
	Folder string `json:"folder"`
	Items  []struct {
		Title       string `json:"title"`
		Path        string `json:"path"`
		Folder      string `json:"folder"`
		HasPassword bool   `json:"hasPassword,omitempty"`
	} `json:"items,omitempty"`
}

type apiHeader struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Require bool   `json:"require"`
	Desc    string `json:"desc"`
}

type apiPublicResponse struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Type string `json:"type"`
	Main bool   `json:"main,omitempty"`
}

type Config struct {
	Enable        bool       `json:"enable"`
	DataPath      string     `json:"-"`
	Title         string     `json:"title"`
	Desc          string     `json:"desc"`
	Copyright     string     `json:"copyright"`
	DefaultMethod string     `json:"default_method"`
	DefaultAuthor string     `json:"default_author"`
	Apps          []apiApp   `json:"apps"`
	Groups        []apiGroup `json:"groups"`
	Cache         struct {
		Enable bool   `json:"enable"`
		Path   string `json:"path"`
		Reload bool   `json:"reload"`
		Max    int    `json:"max"`
	} `json:"cache"`
	Auth struct {
		Enable    bool   `json:"enable"`
		SecretKey string `json:"secret_key"`
	} `json:"auth"`
	FilterMethod []interface{}       `json:"filter_method"`
	Headers      []apiHeader         `json:"headers"` // 猜测是保存的公共头部
	Parameters   []interface{}       `json:"parameters"`
	Responses    []apiPublicResponse `json:"responses"`
	Docs         struct {
		MenuTitle string `json:"menu_title"`
		Menus     []struct {
			Title string `json:"title"`
			Path  string `json:"path,omitempty"`
			Items []struct {
				Title string `json:"title"`
				Path  string `json:"path"`
			} `json:"items,omitempty"`
		} `json:"menus"`
	} `json:"docs"`
	Crud struct {
		Controller struct {
			Path     string `json:"path"`
			Template string `json:"template"`
		} `json:"controller"`
		Service struct {
			Path     string `json:"path"`
			Template string `json:"template"`
		} `json:"service"`
		Model struct {
			Path          string `json:"path"`
			Template      string `json:"template"`
			DefaultFields []struct {
				Field       string `json:"field"`
				Desc        string `json:"desc"`
				Type        string `json:"type"`
				Length      int    `json:"length"`
				Default     string `json:"default"`
				NotNull     bool   `json:"not_null"`
				MainKey     bool   `json:"main_key"`
				Incremental bool   `json:"incremental"`
				Validate    string `json:"validate"`
				Query       bool   `json:"query"`
				List        bool   `json:"list"`
				Detail      bool   `json:"detail"`
				Add         bool   `json:"add"`
				Edit        bool   `json:"edit"`
			} `json:"default_fields"`
			FieldsTypes []string `json:"fields_types"`
		} `json:"model"`
		Validate struct {
			Path     string `json:"path"`
			Template string `json:"template"`
			Rules    []struct {
				Name    string `json:"name"`
				Rule    string `json:"rule"`
				Message string `json:"message"`
			} `json:"rules"`
		} `json:"validate"`
	} `json:"crud"`
	Debug bool `json:"debug"`
}

type apiGroup struct {
	Title string      `json:"title"`
	Name  interface{} `json:"name"`
}

type apiItem struct {
	Title   string `json:"title"`
	Path    string `json:"path"`
	Type    string `json:"type"`
	MenuKey string `json:"menu_key"`
}

type apiDoc struct {
	Title   string    `json:"title"`
	Path    string    `json:"path,omitempty"`
	Type    string    `json:"type,omitempty"`
	MenuKey string    `json:"menu_key"`
	Items   []apiItem `json:"items,omitempty"`
	Group   string    `json:"group,omitempty"`
}

// apiParam 请求或响应参数
type apiParam struct {
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	Desc         string      `json:"desc"`
	Default      interface{} `json:"default"`
	Require      bool        `json:"require"`
	ChildrenType string      `json:"childrenType"`
	Params       []apiParam  `json:"params,omitempty"`
}

type apiEntity struct {
	configed  bool
	immutable bool
	Title     string      `json:"title"`
	Desc      string      `json:"desc,omitempty"`
	Tag       string      `json:"tag"`
	Author    string      `json:"author"`
	URL       string      `json:"url"`
	Method    string      `json:"method"`
	Param     []apiParam  `json:"param"`
	Return    []apiReturn `json:"return"`
	Header    []apiHeader `json:"header"`
	Name      string      `json:"name"`
	MenuKey   string      `json:"menu_key"`
	group     apiGroup
	subGroup  string
}

type apiReturn struct {
	Name   string     `json:"name"`
	Desc   string     `json:"desc"`
	Type   string     `json:"type"`
	Main   bool       `json:"main,omitempty"`
	Params []apiParam `json:"params,omitempty"`
}

type apiList struct {
	Group    string      `json:"group"`
	Sort     interface{} `json:"sort"`
	Title    string      `json:"title"`
	MenuKey  string      `json:"menu_key"`
	Children []apiEntity `json:"children"`
}

type apiData struct {
	Groups []apiGroup `json:"groups"`
	List   []apiList  `json:"list"`
	Docs   []apiDoc   `json:"docs"`
	Tags   []string   `json:"tags"`
}

type apiDocApp struct {
	sync.Mutex
	config  *Config
	groups  []apiGroup
	ApiData map[string]apiList
}

func DefaultConfig() *Config {
	return &Config{
		Enable:        true,
		Title:         "PineCMS ApiDoc",
		Desc:          "PineCMS 接口文档",
		Copyright:     "https://github.com/xiusin/pinecms.git",
		DefaultMethod: "GET",
		DefaultAuthor: "xiusin",
		Debug:         true,
		DataPath:      "apidoc",
		//Responses: []apiPublicResponse{ // todo 使用解析
		//	{Name: "code", Desc: "状态码", Type: "int"},
		//	{Name: "message", Desc: "操作描述", Type: "string"},
		//	{Name: "data", Desc: "业务数据", Type: "object", Main: true},
		//},
		Headers: []apiHeader{
			{
				Name:    "Authorization",
				Type:    "string",
				Require: true,
				Desc:    "登录票据",
			},
		},
		Apps: []apiApp{
			{
				Title:  "后端Api",
				Folder: "backend",
			},
			{
				Title:  "前端Api",
				Folder: "index",
			},
		},
	}
}

type DemoParams struct {
	Page int `json:"page" api:"require:true,remark:分页数,default:0"`
}

type DemoResponseParam struct {
	Code    int         `json:"code" api:"require:true,remark:状态码,default:0"`
	Message string      `json:"message" api:"require:true,remark:操作描述"`
	Data    interface{} `json:"data" api:"require:true,remark:业务数据"`
}

func parseInterface(reqParams interface{}) []apiParam {
	if reqParams != nil {
		s := structs.Fields(reqParams)
		var apiReqParams []apiParam
		for _, field := range s {
			apiData := field.Tag("api")
			if len(apiData) == 0 {
				continue
			}
			fieldType := "any"
			switch field.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				fieldType = "number"
			case reflect.Bool:
				fieldType = "bool"
			case reflect.String:
				fieldType = "string"
			case reflect.Struct:
				fieldType = "object"
			case reflect.Slice, reflect.Array:
				fieldType = "array"
			case reflect.Float32, reflect.Float64:
				fieldType = "float"
			}
			apiDatas := strings.Split(apiData, ",")
			name := field.Tag("json")
			if len(name) == 0 {
				name = field.Name()
			}
			apiReqParam := apiParam{Type: fieldType, Name: name}
			for _, data := range apiDatas {
				kv := strings.Split(strings.ToLower(data), ":")
				v := ""
				if len(kv) > 1 {
					v = kv[1]
				}
				switch kv[0] {
				case "require":
					if strings.ToLower(v) == "true" {
						apiReqParam.Require = true
					}
				case "remark":
					apiReqParam.Desc = v
				case "default":
					apiReqParam.Default = v
				}

			}

			apiReqParams = append(apiReqParams, apiReqParam)

		}
		return apiReqParams
	}
	return nil
}

var apiJsonData []byte

func ApiDoc(config *Config) pine.Handler {
	if config == nil {
		config = DefaultConfig()
	}
	defaultConfig = config
	_ = os.Mkdir(config.DataPath, os.ModePerm)
	fileName := filepath.Join(config.DataPath, "apidoc.json")
	apiJsonData, _ = os.ReadFile(fileName)

	//json.Unmarshal(apiJsonData, )

	return func(ctx *pine.Context) {
		if !config.Enable {
			ctx.Next()
			return
		}
		switch ctx.Path() {
		case "/apidoc/config":
			getConfig(ctx)
		case "/apidoc/apiData":
			getApiData(ctx)
		default:
			ps := strings.Split(ctx.Path(), "/")
			method := string(ctx.Method())
			key := strings.ToLower(fmt.Sprintf("%s_%s", method, strings.ReplaceAll(ctx.Path(), "/", "_")))
			entity := &apiEntity{
				URL:     ctx.Path(),
				Method:  method,
				Name:    ps[len(ps)-1],
				Header:  defaultConfig.Headers,
				MenuKey: key,
			}
			ctx.Set(apiDocKey, entity)

			ctx.Next()

			if !entity.configed {
				return
			}

			// 记录响应信息
			if entity.immutable == true {
				// 判断文档是否已经存在 存在则直接返回
				return
			}
			entity.Author = defaultConfig.DefaultAuthor
			// 读取请求参数
			m := structs.Fields(structs.Map(ctx.Response.Body()))

			fmt.Printf("%+v\n", entity)
		}
	}
}

func getConfig(ctx *pine.Context) {
	_ = ctx.WriteJSON(pine.H{"code": 0, "data": defaultConfig})
}

func getApiData(ctx *pine.Context) {
	_ = ctx.WriteJSON(pine.H{"code": 0, "data": apiData{
		Groups: []apiGroup{
			{
				Title: "全部",
				Name:  0,
			},
			{
				Title: "开发文档",
				Name:  "markdown_doc",
			},
		},
		List: []apiList{},
		Docs: []apiDoc{
			{
				Title:   "关于Apidoc",
				Path:    "docs/readme",
				Type:    "md",
				MenuKey: "md_39417",
			},
			{
				Title: "HTTP响应编码",
				Items: []apiItem{
					{
						Title:   "status错误码说明",
						Path:    "docs/HttpStatus",
						Type:    "md",
						MenuKey: "md_64795",
					},
				},
				Group:   "markdown_doc",
				MenuKey: "md_group_47194",
			},
		},
		Tags: []string{"后台", "前台", "测试", "基础"},
	}})
}
