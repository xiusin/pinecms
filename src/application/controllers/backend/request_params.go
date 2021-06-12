package backend

// loginUserParam 登录参数
type loginUserParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type idParams struct {
	Id  int64   `json:"id"`
	Ids []int64 `json:"ids"`
}

type listParam struct {
	Page       int    `json:"page"`     // 分页数
	Size       int    `json:"size"`     // 页码
	OrderField string `json:"order"`    // 排序字段
	Sort       string `json:"sort"`     // 排序规则
	Keywords   string `json:"keyWord"`  // 搜索关键字
	Export     bool   `json:"_isExport"` // 是否导出
}

type responsePageParam struct {
	Page  int64 `json:"page"`
	Size  int64 `json:"size"`
	Total int64 `json:"total"`
}
