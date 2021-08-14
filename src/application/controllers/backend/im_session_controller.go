package backend

import "github.com/xiusin/pine"

type ImSessionController struct {
	pine.Controller
}

func (c *ImSessionController) RegisterRoute(b pine.IRouterWrapper) {
	b.POST("/im/session/page", "SessionPage")
	b.POST("/im/session/unreadCount", "SessionUnreadCount")
	b.POST("/im/message/page", "MessagePage")
}

func (c *ImSessionController) SessionPage() {

}

func (c *ImSessionController) SessionUnreadCount()  {

}

func (c *ImSessionController) MessagePage()  {

}
