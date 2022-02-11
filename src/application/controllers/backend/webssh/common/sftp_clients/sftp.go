package sftp_clients

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/pkg/sftp"
	"log"
	"path"
	"sync"
	"time"
)

const (
	getpwd = "getpwd"
	upload = "upload"
)

type sftp_req struct {
	Type     string `json:"type"`
	FilePath string `json:"filepath"`
	FileName string `json:"filename"`
	FileData string `json:"filedata"`
}

type sftp_resp struct {
	Code int    `json:"code"`
	Type string `json:"type"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type MyClient struct {
	Uid  uint
	Sftp *sftp.Client
}

type clients struct {
	*sync.RWMutex
	C    map[string]*MyClient
}

var Client clients

func init() {
	Client = clients{new(sync.RWMutex),make(map[string]*MyClient)}
}

func (c *MyClient) ReceiveWsMsg(wsConn *websocket.Conn, exitCh chan bool) {
	defer setQuit(exitCh)
	go c.SessionWait(wsConn, exitCh)
	for {
		select {
		case <-exitCh:
			return
		default:
			_, wsData, err := wsConn.ReadMessage()
			if err != nil {
				log.Println(err.Error())
				//logrus.WithError(err).Error("reading webSocket message failed")
				return
			}
			//unmashal bytes into struct
			msgObj := sftp_req{}
			if err := json.Unmarshal(wsData, &msgObj); err != nil {
				log.Println("unmarshal websocket message failed:", string(wsData))
				continue
			}
			msgresp := sftp_resp{}
			switch msgObj.Type {
			case getpwd:
				msgresp.Code = 200
				msgresp.Type = "pwd"
				path, err := c.Sftp.Getwd()
				if err != nil {
					msgresp.Code = 404
					msgresp.Msg = "服务器Path获取失败"
					msgresp.Data = err.Error()
					log.Println("sftp getpwd err:", err.Error())
				}
				msgresp.Data = path
				msg, _ := json.Marshal(msgresp)
				if err := wsConn.WriteMessage(websocket.TextMessage, msg); err != nil {
					log.Println("sftp client getpwd err:", err.Error())
					return
				}
			case upload:
				msgresp.Code = 200
				msgresp.Type = "upload"
				io_data, err := base64.StdEncoding.DecodeString(msgObj.FileData)
				if err != nil {
					msgresp.Code = 401
					msgresp.Msg = "文件解析失败"
					msgresp.Data = err.Error()
					log.Println("sftp base64decode err:", err.Error())
					msg, _ := json.Marshal(msgresp)
					if err := wsConn.WriteMessage(websocket.TextMessage, msg); err != nil {
						log.Println("sftp base64decode send err:", err.Error())
						return
					}
				}

				if err := c.Sftp.MkdirAll(msgObj.FilePath); err != nil {
					msgresp.Code = 402
					msgresp.Msg = "服务器创建目录失败"
					msgresp.Data = err.Error()
					log.Println("sftp mkdir err:", err.Error())
					msg, _ := json.Marshal(msgresp)
					if err := wsConn.WriteMessage(websocket.TextMessage, msg); err != nil {
						log.Println("sftp mkdir send err:", err.Error())
						return
					}
					continue
				}

				file, err := c.Sftp.Create(path.Join(msgObj.FilePath, msgObj.FileName))
				if err != nil {
					msgresp.Code = 403
					msgresp.Msg = "服务器文件创建失败"
					msgresp.Data = err.Error()
					log.Println("sftp create file err:", err.Error())
					msg, _ := json.Marshal(msgresp)
					if err := wsConn.WriteMessage(websocket.TextMessage, msg); err != nil {
						log.Println("sftp create file send err:", err.Error())
						return
					}
					continue
				}

				defer file.Close()

				if _, err := file.Write(io_data); err != nil {
					msgresp.Code = 405
					msgresp.Msg = "服务器文件写入失败"
					msgresp.Data = err.Error()
					log.Println("sftp write file err:", err.Error())
					msg, _ := json.Marshal(msgresp)
					if err := wsConn.WriteMessage(websocket.TextMessage, msg); err != nil {
						log.Println("sftp write file send err:", err.Error())
						return
					}
					continue
				}
				file.Close()
				filepath := path.Join(msgObj.FilePath, msgObj.FileName)
				msgresp.Code = 200
				msgresp.Msg = "OK"
				msgresp.Data = filepath
				//log.Println("file write ok")
				msg, _ := json.Marshal(msgresp)
				if err := wsConn.WriteMessage(websocket.TextMessage, msg); err != nil {
					log.Println("sftp write file send err:", err.Error())
					return
				}
			}
		}
	}
}

func (c *MyClient) SessionWait(wsConn *websocket.Conn, quitChan chan bool) {
	timer := time.NewTicker(time.Second * 30)
	defer timer.Stop()
	defer setQuit(quitChan)
	for {
		select {
		case <-timer.C:
			{
				if err := wsConn.WriteMessage(websocket.TextMessage, []byte("pong")); err != nil {
					log.Println("sftp pong send err :", err.Error())
					return
				}
			}
		case <-quitChan:
			return
		}
	}
}

func setQuit(ch chan bool) {
	ch <- true
}
