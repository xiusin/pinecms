package backend

// loginUserParam 登录参数
type loginUserParam struct {
	Username     string `json:"username" api:"remark:登录账号|require:true"`
	Password     string `json:"password" api:"remark:登录密码|require:true"`
	CaptchaId    string `json:"captchaId" api:"remark:验证码ID|require:true"`
	CaptchaValue string `json:"verifyCode" api:"remark:验证码|require:true"`
}

type idParams struct {
	Id  int64   `json:"id" api:"remark:删除单个记录"`
	Ids []int64 `json:"ids" api:"remark:删除多个记录"`
}

type listParam struct {
	Page       int            `json:"page" api:"remark:分页数|default:1|require:true"`   // 分页数
	Size       int            `json:"size" api:"remark:分页条数|default:10|require:true"` // 页码
	OrderField string         `json:"order" api:"remark:排序字段"`                        // 排序字段
	Sort       string         `json:"sort" api:"remark:排序方法desc=逆序,asc=正序"`           // 排序规则
	Keywords   string         `json:"keyWord" api:"remark:查询关键字"`                     // 搜索关键字
	Export     bool           `json:"_isExport" api:"remark:是否导出"`                    // 是否导出
	Params     map[string]any `json:"params" api:"remark:额外参数用于非配置字段导出"`              // 额外附加参数
	Param      map[string]any `json:"param" api:"remark:cl-filter组件参数"`               // 额外附加参数
	Filters    *FilterGroup   `json:"filters,omitempty" api:"remark:结构化筛选参数"`         // 结构化筛选参数
}

// FilterGroup 定义了一组筛选条件
type FilterGroup struct {
	Op         string `json:"op" api:"remark:逻辑操作符 AND 或 OR"`
	Conditions []any  `json:"conditions" api:"remark:条件,可以是Filter或FilterGroup"`
}

// Filter 定义了单个筛选条件
type Filter struct {
	Field    string `json:"field" api:"remark:筛选字段"`
	Function string `json:"function,omitempty" api:"remark:函数"`
	Op       string `json:"op" api:"remark:筛选操作符"`
	Value    any    `json:"value" api:"remark:筛选值"`
}
