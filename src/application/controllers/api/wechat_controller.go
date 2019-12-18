package api

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
	"github.com/silenceper/wechat/user"
	tables "github.com/xiusin/iriscms/src/application/models/tables"
	"github.com/xiusin/iriscms/src/common/cache"
	"github.com/xiusin/iriscms/src/common/helper"
)

type WechatController struct {
	Orm   *xorm.Engine
	Ctx   iris.Context
	cache cache.ICache
}

func (c *WechatController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(iris.MethodGet, "/wechat/endpoint", "EndPointGet")
	b.Handle(iris.MethodPost, "/wechat/endpoint", "EndPointPost")
}

func (c *WechatController) EndPointGet() {
	setting := c.Ctx.Values().Get("setting").(map[string]string)
	conf := c.Ctx.Values().Get("app.config").(map[string]string)
	engine := conf["uploadEngine"]
	wcConfig := &wechat.Config{
		AppID:          setting["WX_APPID"],
		AppSecret:      setting["WX_APPSECRET"],
		Token:          setting["WX_TOKEN"],
		EncodingAESKey: setting["WX_AESKEY"],
	}
	wc := wechat.NewWechat(wcConfig)
	server := wc.GetServer(c.Ctx.Request(), c.Ctx.ResponseWriter())

	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		//回复消息：演示回复用户发送的消息
		id, _ := strconv.Atoi(msg.Content)
		if msg.FromUserName != "" {
			go func() {
				var u tables.IriscmsWechatMember
				// 检查来源
				b, _ := c.Orm.Where("openid=?", msg.FromUserName).Exist(&u)
				if b {
					return
				}
				userService := user.NewUser(wc.Context)
				userInfo, err := userService.GetUserInfo(msg.OpenID)
				if err != nil {
					c.Ctx.Application().Logger().Error("获取用户信息错误", err.Error())
				}
				_, err = c.Orm.InsertOne(&tables.IriscmsWechatMember{
					Time:       time.Now(),
					Nickname:   userInfo.Nickname,
					Headimgurl: userInfo.Headimgurl,
					Sex:        int(userInfo.Sex),
					Openid:     userInfo.OpenID,
					Mpid:       "asdad",
					//SubScribeScene:userInfo.Subscribe,	todo 关注来源
				})
				if err != nil {
					c.Ctx.Application().Logger().Error("保存用户信息失败", err.Error())
				}
			}()
		}

		if id < 1 {
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText("您好,请输入正确的代码编号来获取密码")}
		}
		var content tables.IriscmsContent
		ok, _ := c.Orm.Where("id=?", id).Get(&content)
		if !ok || content.DeletedAt > 0 || content.Status == 0 { //已被删除或不可用
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText("what are you 弄啥类?")}
		}

		if content.PwdType != 1 || content.Money > 0 { //获取密码方式非关注公众号获取 或 收费
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText("资源不可用此方式获取密码")}
		}
		go func() {
			str, err := xml.Marshal(msg)
			if err == nil {
				c.Orm.InsertOne(&tables.IriscmsWechatMessageLog{
					Content: string(str),
					Time:    time.Now(),
				})
			}
		}()
		host := "http://" + c.Ctx.Host()
		articles := make([]*message.Article, 1)
		article := new(message.Article)
		article.Title = content.Title
		article.Description = "您好,资源下载密码为:" + content.SourcePwd + ". 用心提供最有质量的书籍视频,让您学到最有价值的知识. 开通vip下载专属资源."
		if engine == "oss" {
			article.PicURL = content.Thumb
		} else {
			article.PicURL = host + content.Thumb
		}
		article.URL = host + "/#/article/" + strconv.Itoa(int(content.Id)) + "?=fwa" //from_wx_article
		articles[0] = article
		return &message.Reply{MsgType: message.MsgTypeNews, MsgData: message.NewNews(articles)}
	})

	err := server.Serve()
	if err != nil {
		c.Ctx.Application().Logger().Error("验证公众号失败:" + err.Error())
		helper.Ajax("验证失败:"+err.Error(), 1, c.Ctx)
	} else {
		server.Send()
	}
}

// 不知道为什么MVC模式不支持多个RequestMethod注册到同一个方法上
func (c *WechatController) EndPointPost() {
	c.EndPointGet()
}
