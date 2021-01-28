package ws

import (
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"FrontEndOJGolang/pkg/strategy"
	ws "FrontEndOJGolang/pkg/websocket"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upGrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}}

func Ws(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	user, err := app.GetUserFromSession(c)
	if err != nil {
		appG.RespErr(e.NOT_LOGINED, "plese login first")
		return
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
	for {
		wsJsonReq := &WsJsonReq{}
		err := c.Conn.ReadJSON(wsJsonReq)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("ws error:%v", err)
				break
			}
			if websocket.IsCloseError(err, websocket.CloseGoingAway) {
				break
			}
			if wsJsonReq.Cmd == "" {
				log.Printf("parse json error:%v", err)
				continue
			}
		}

		// exec
		strategy.ExecStrategy(wsJsonReq.Cmd)
	}
}

func writeData(c *ws.WsClientConn) {
	defer c.Conn.Close()
	for {
		select {
		case message, ok := <- c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
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
