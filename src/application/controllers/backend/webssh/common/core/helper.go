package core

import (
	"github.com/gorilla/websocket"
	"github.com/xiusin/pine"
	"log"
	"time"
)

func JsonError(c *pine.Context, msg interface{}) {
	c.WriteJSON(pine.H{"ok": false, "msg": msg})
}

func HandleError(c *pine.Context, err error) bool {
	if err != nil {
		//logrus.WithError(err).Error("gin context http handler error")
		JsonError(c, err.Error())
		return true
	}
	return false
}

func WshandleError(ws *websocket.Conn, err error) bool {
	if err != nil {
		log.Println("handler ws ERROR:",err.Error())
		dt := time.Now().Add(time.Second)
		if err := ws.WriteControl(websocket.CloseMessage, []byte(err.Error()), dt); err != nil {
			log.Println("websocket writes control message failed:",err.Error())
		}
		return true
	}
	return false
}
