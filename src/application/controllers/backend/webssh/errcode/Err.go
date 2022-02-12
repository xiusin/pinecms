package errcode

const (
	C_nil_err      = 200 //成功
	C_from_err     = 100 //表单错误
	C_phone_err    = 101 //手机号错误
	C_code_err     = 102 //验证码错误
	S_send_err     = 300 //发送验证码错误
	S_auth_err     = 301 //Token权限校验失败
	S_auth_fmt_err = 302 //Header Token格式错误
	S_Verify_err   = 303 //验证码校验失败
	S_Db_err       = 304
)
