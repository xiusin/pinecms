package api

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-xorm/xorm"
	"github.com/google/uuid"
	jwt2 "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iriscms/application/models/tables"
	"iriscms/common/helper"
	"time"
)

type UserApiController struct {
	Orm *xorm.Engine
	Ctx iris.Context
}

type ReturnApiData struct {
	Status bool        `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func (c *UserApiController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodPost, "/user/login", "UserLogin")
	b.Handle(iris.MethodGet, "/user/center", "UserCenter")
	b.Handle(iris.MethodPost, "/user/register", "UserRegister")
}

func (c *UserApiController) UserLogin() {
	//生成JwtToken
	dd := map[string]string{}
	err := c.Ctx.UnmarshalBody(&dd, iris.UnmarshalerFunc(json.Unmarshal)) //todo 解析body字符串
	if err != nil {
		c.Ctx.JSON(ReturnApiData{false, err.Error(), nil})
		return
	}
	if !helper.VerifyVCaptcha(dd["token"]) {
		c.Ctx.JSON(ReturnApiData{false, "验证码失败", nil})
		return
	}
	if dd["account"] == "" || dd["password"] == "" {
		c.Ctx.JSON(ReturnApiData{false, "username or password is empty!", nil})
	} else {
		var user tables.IriscmsMember
		ok, err := c.Orm.Where("account = ? and password = ?", dd["account"], dd["password"]).Get(&user)
		if !ok || err != nil {
			c.Ctx.JSON(ReturnApiData{false, "login failed!", nil})
			return
		}

		claims := jwt.MapClaims{ //为了匹配中间件, 在这里使用相同的配置
			"user": user,
			"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(), //过期时间
			"iat":  time.Now().Unix(),                                   // 当前时间戳
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte("MySecret"))
		c.Ctx.JSON(ReturnApiData{true, "login success!", struct {
			SignToken string `json:"sign_token"`
		}{SignToken: tokenString}})
	}
}

func (c *UserApiController) UserCenter() {
	user, ok := c.Ctx.Values().Get(jwt2.DefaultContextKey).(*jwt.Token)
	if !ok {
		c.Ctx.JSON(ReturnApiData{false, "author error", nil})
		return
	}
	userC, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		c.Ctx.JSON(ReturnApiData{false, "author error", nil})
		return
	}
	c.Ctx.JSON(ReturnApiData{true, "", userC["user"]})
}

func (c *UserApiController) UserLogout() {
	c.Ctx.JSON(ReturnApiData{true, "", nil})
}

func (c *UserApiController) UserRegister() {
	var user tables.IriscmsMember
	account := c.Ctx.FormValue("account")
	email := c.Ctx.FormValue("email")
	lenth, _ := c.Orm.Where("account=? or email=?", account, email).Count()
	if lenth > 0 {
		c.Ctx.JSON(ReturnApiData{false, "account or email are already exists!", nil})
		return
	}
	user.Account = account
	user.Email = email
	user.Password = c.Ctx.FormValue("password")
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Enabled = 1
	user.VerifyToken = uuid.New().String()
	_, err := c.Orm.InsertOne(&user)
	if err != nil {
		c.Ctx.JSON(ReturnApiData{false, "注册失败, " + err.Error() + "!", nil})
	} else {
		c.Ctx.JSON(ReturnApiData{true, "注册成功!", nil})
	}
}
