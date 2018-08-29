package api

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/dgrijalva/jwt-go"
	jwt2 "github.com/iris-contrib/middleware/jwt"
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
}

func (c *UserApiController) UserLogin() {
	//生成JwtToken
	userName := c.Ctx.PostValueTrim("username")
	password := c.Ctx.PostValueTrim("password")
	if userName == "" || password == "" {
		c.Ctx.JSON(ReturnApiData{false, "username or password is empty!", nil})
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{	//为了匹配中间件, 在这里使用相同的配置
			"username": userName,
			"password": password,
			"exp": time.Now().Add(time.Hour * time.Duration(1)).Unix(),	//过期时间
			"iat": time.Now().Unix(),	// 当前时间戳
		})
		tokenString, err := token.SignedString([]byte("MySecret"))
		if err != nil {
			c.Ctx.JSON(ReturnApiData{false, err.Error(), nil})
		} else {
			c.Ctx.JSON(ReturnApiData{true, "", struct {
				SignToken  string `json:"sign_token"`
				User map[string]string
			}{
				SignToken:tokenString,
				User: map[string]string{

				},
			}})
		}
	}

}

func (c *UserApiController) UserCenter() {
	user := c.Ctx.Values().Get(jwt2.DefaultContextKey).(*jwt.Token)
	c.Ctx.JSON(ReturnApiData{true, "", user.Claims})
}