package wechat

import (
	"github.com/silenceper/wechat/v2/officialaccount/menu"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

const CacheKeyWechatMaterialCount = "pinecms.wechat.material.count.%s"

const CacheKeyWechatMaterialListKeys = "pinecms.wechat.material.list.key"

const CacheKeyWechatUserTags = "pinecms.wechat.user.tags.%s"

const CacheKeyWechatMenu = "pinecms.wechat.menu.%s"

const CacheTimeSecs = 30 * 24 * 3600

// Plugin TODO 回复插件
type Plugin struct {
}

// WechatMsg 微信自动回复混合消息结构
type WechatMsg struct {
	Title string `json:"title"`
	// 小程序
	AppID        string `json:"appid"`
	PagePath     string `json:"pagePath"`
	ThumbMediaID string `json:"thumb_media_id"`

	// 音乐 视频
	Description string `json:"description"`
	MusicURL    string `json:"music_url"`
	HQMusicURL  string `json:"hq_music_url"`

	MediaID string `json:"media_id"`

	KfAccount string `json:"kf_account"`

	Articles []*message.Article
}

// MaterialUploadForm 素材上传表单结构
type MaterialUploadForm struct {
	Appid        string `json:"appid"`
	MediaID      string `json:"mediaId"`
	FileName     string `json:"fileName"`
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
	MediaType    string `json:"mediaType"`
}

// menuParam 菜单编辑参数
type menuParam struct {
	Appid string `json:"appid"`
	Menu  struct {
		Button []*menu.Button `json:"button"`
		MenuID int64          `json:"menuid"`
	} `json:"menu"`
}

// replyMsg 回复消息结构体
type replyMsg struct {
	AppId        string `json:"appid"`
	OpenId       string `json:"openid"`
	ReplyContent string `json:"replyContent"`
	ReplyType    string `json:"replyType"`
}

// UserTags 用户标签结构
type UserTags struct {
	Appid   string   `json:"appid"`
	Id      int64    `json:"id"`
	Name    string   `json:"name"`
	Openids []string `json:"openids"`
	Action  string   `json:"action"`
}
