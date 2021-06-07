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
	Page       int    `json:"page"`
	Size       int    `json:"size"`
	OrderField string `json:"order"`
	Sort       string `json:"sort"`
	Keywords   string `json:"keyWord"`
}

type responsePageParam struct {
	Page  int64 `json:"page"`
	Size  int64 `json:"size"`
	Total int64 `json:"total"`
}

type responseListParam struct {
	List       interface{}       `json:"list"`
	Pagination responsePageParam `json:"pagination"`
}
