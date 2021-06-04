package backend

// loginUserParam 登录参数
type loginUserParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type listParam struct {
	Page       int64  `json:"page"`
	Size       int64  `json:"size"`
	OrderField string `json:"order"`
	Sort       string `json:"sort"`
}
