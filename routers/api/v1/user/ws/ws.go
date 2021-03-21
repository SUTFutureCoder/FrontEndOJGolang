package ws

import (
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"FrontEndOJGolang/pkg/setting"
	"FrontEndOJGolang/pkg/strategy"
	ws "FrontEndOJGolang/pkg/websocket"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func Ws(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	user := app.GetUserFromSession(appG)
	if user.Id == 0 {
		return
	}

	handleWsConn(c, &appG, user)
}

// user id start from 10000
var JudgerId uint64 = 0

func WsJudger(c *gin.Context) {
	appG := app.Gin {
		C: c,
	}
	// session token可直接注册
	if setting.SessionSetting.Token != c.GetHeader("session_token") {
		appG.RespErr(e.ERROR, "websocket session token error")
		return
	}
	JudgerId++
	session := app.UserSession {
		Id: JudgerId,
		Name: "JUDGER",
	}
	handleWsConn(c, &appG, session)
}

func handleWsConn(c *gin.Context, appG *app.Gin, user app.UserSession) {
	upGrader := websocket.Upgrader{
		// cross origin domain
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		// 处理 Sec-WebSocket-Protocol Header
		Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")},
	}
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		appG.RespErr(e.ERROR, "websocket upgrade error")
		return
	}

	// init single conn
	client := &ws.WsClientConn{
		User: user,
		Hub:  ws.Wshub,
		Conn: conn,
		Send: make(chan []byte, 10240),
	}
	client.Hub.Register <- client

	go writeData(client)
	go readData(client)
}

type WsJsonReq struct {
	Cmd  string `json:"cmd"`
	Data string `json:"data"`
}

func readData(c *ws.WsClientConn) {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	s := &startegy.Strategy {
		Context: c,
	}
	for {
		wsJsonReq := &WsJsonReq{}
		err := c.Conn.ReadJSON(wsJsonReq)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("ws error:%v", err)
				break
			}
			if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				break
			}
			if wsJsonReq.Cmd == "" {
				log.Printf("parse json error:%v", err)
				continue
			}
		}

		// exec
		s.ExecStrategy(wsJsonReq.Cmd, wsJsonReq.Data)
	}
}

func writeData(c *ws.WsClientConn) {
	defer c.Conn.Close()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				c.Send = nil
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}
