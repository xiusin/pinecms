package apidoc

type Config struct {
	RoutePrefix   string      `json:"route_prefix"`
	Enable        bool        `json:"enable"`         // 是否启用apidoc
	DataPath      string      `json:"-"`              // 配置数据存储目录
	ResponseParam interface{} `json:"-"`              // 用于反射返回值信息
	Title         string      `json:"title"`          // 标题目录
	Desc          string      `json:"desc"`           // 描述
	Copyright     string      `json:"copyright"`      // 版权
	DefaultAuthor string      `json:"default_author"` // 默认作者
	Apps          []apiApp    `json:"apps"`           // 应用， 例如 前端接口，后端接口
	Groups        []apiGroup  `json:"groups"`         //  应用分组
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

var defaultConfig *Config

func DefaultConfig() *Config {
	return &Config{
		RoutePrefix:   "/apidoc",
		Enable:        true,
		Title:         "PineCMS ApiDoc",
		Desc:          "PineCMS 接口文档",
		Copyright:     "https://github.com/xiusin/pinecms.git",
		DefaultAuthor: "xiusin",
		Debug:         true,
		DataPath:      "docdb",
		ResponseParam: &DemoResponseParam{},
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
				Folder: "admin",
				Host:   "http://localhost:2019",
			},
			{
				Title:  "前端Api",
				Folder: "index",
				Host:   "http://localhost:2019",
			},
		},
	}
}

type Configure func(entity *apiEntity)

func WithImmutable(immutable bool) Configure {
	return func(entity *apiEntity) {
		entity.Immutable = immutable
	}
}

func WithHeaders(headers []apiHeader) Configure {
	return func(entity *apiEntity) {
		entity.Header = append(entity.Header, headers...)
	}
}

// WithOnlyParams 只允许部分参数体现在文档内
func WithOnlyParams(onlyParams []string) Configure {
	return func(entity *apiEntity) {
		entity.OnlyParams = onlyParams
	}
}

// WithExcludeParams 排除部分参数
func WithExcludeParams(excludeParams []string) Configure {
	return func(entity *apiEntity) {
		entity.ExcludeParams = excludeParams
	}
}

// WithNoParams 不带参数
func WithNoParams() Configure {
	return func(entity *apiEntity) {
		entity.NoParams = true
	}
}
