package apidoc

const apiDocKey = "apiEntity"

type apiApp struct { // 记录应用分组. 页面右上角下拉
	Title  string `json:"title"`
	Path   string `json:"path,omitempty"`
	Folder string `json:"folder"`
	Host   string `json:"host"`
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
	configured      bool        //程序内调用 SetApiEntity置为true
	Immutable       bool        `json:"immutable"`      //immutable设置不可变, 在文档存在的情况下不会覆盖已生成的文档, 一般用于确定的文档信息, 接口处提交自定义将此参数设置为不可变
	AppId           string      `json:"app_id"`         // 应用ID
	Group           apiGroup    `json:"group"`          // 分组ID
	SubGroup        string      `json:"sub_group"`      // 子分组
	Title           string      `json:"title"`          // 标题
	Desc            string      `json:"desc,omitempty"` // 描述
	Tag             string      `json:"tag"`            // 标签
	Author          string      `json:"author"`         // 作者,默认继承全局配置
	URL             string      `json:"url"`            // API地址段
	Method          string      `json:"method"`         // 请求方式
	Param           []apiParam  `json:"param"`          // 请求参数
	Query           []apiParam  `json:"query"`
	Return          []apiReturn `json:"return"`            // 返回参数
	Header          []apiHeader `json:"header"`            // header
	Name            string      `json:"name"`              // Api名称
	MenuKey         string      `json:"menu_key"`          // 主菜单名称
	RawParam        string      `json:"raw_param"`         // 原始请求参数
	RawReturn       string      `json:"raw_return"`        // 原始返回参数
	RawQuery        string      `json:"raw_query"`         // 原始query
	QueryDataMethod string      `json:"query_data_method"` // 原始请求数据方法
	OnlyParams      []string    `json:"-"`                 // 只允许部分参数
	ExcludeParams   []string    `json:"-"`                 // 过滤部分参数
	NoParams        bool        `json:"-"`
	//Enable          bool        `json:"enable"`            // 是否启用参数
}

func (e *apiEntity) ID() (jsonField string, value interface{}) {
	value, jsonField = e.MenuKey, "menu_key"
	return
}

func (e *apiEntity) ParamExist(name string) bool {
	for _, param := range e.Param {
		if param.Name == name {
			return true
		}
	}
	return false
}

func (e *apiEntity) QueryExist(name string) bool {
	for _, param := range e.Query {
		if param.Name == name {
			return true
		}
	}
	return false
}

// FilterParams 根据配置过滤无用参数, 可以通过设置: 仅允许, 仅排除, 无.
// 此函数仅嵌入时可以生效, 通过管理端设置是可以直接通过
// 当接口设置为不可变时, 不允许改变
func (e *apiEntity) FilterParams() {
	var filterParams []apiParam
	if len(e.OnlyParams) > 0 {
		for _, param := range e.Param {
			for _, onlyParam := range e.OnlyParams {
				if param.Name == onlyParam {
					filterParams = append(filterParams, param)
					break
				}
			}
		}
	} else if len(e.ExcludeParams) > 0 {
		for _, param := range e.Param {
			var exclude bool
			for _, excludeParam := range e.ExcludeParams {
				if param.Name == excludeParam {
					exclude = true
					break
				}
			}
			if !exclude {
				filterParams = append(filterParams, param)
			}
		}
	}
	if e.NoParams {
		e.Param = []apiParam{}
	}
}

type apiReturn struct {
	Name    string     `json:"name"`
	Desc    string     `json:"desc"`
	Type    string     `json:"type"`
	Default string     `json:"default"`
	Main    bool       `json:"main,omitempty"`
	Params  []apiParam `json:"params"`
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

type DemoParams struct {
	Page int `json:"page" api:"require:true,remark:分页数,default:0"`
}

type DemoResponseParam struct {
	Code    int         `json:"code" api:"require:true,remark:状态码,default:0"`
	Message string      `json:"message" api:"require:true,remark:操作描述"`
	Data    interface{} `json:"data" api:"require:true,remark:业务数据"`
}
