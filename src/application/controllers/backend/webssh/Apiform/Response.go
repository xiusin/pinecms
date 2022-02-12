package Apiform

type Resp struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	Token string      `json:"token"`
}