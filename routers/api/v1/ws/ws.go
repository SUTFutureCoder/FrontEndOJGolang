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
	//for {
	//	// read from ws
	//	mt, message, err := ws.ReadMessage()
	//	if err != nil {
	//		log.Println(err.Error())
	//		break
	//	}
	//	log.Println(mt)
	//	log.Println(message)
	//	log.Println(userSession)
	//	ws.WriteMessage(mt, []byte("TEST"))
	//}
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
			if wsJsonReq.Cmd == "" {
				log.Printf("parse json error:%v", err)
				continue
			}
		}

		// exec
		strategy.ExecStrategy(wsJsonReq.Cmd)

		c.Send <- []byte("yooooo")

		// all USE Wshub
		clientMsg := &ws.ClientMsg{
			ClientId: 0,
			Msg:      []byte("client"),
		}
		ws.Wshub.ClientMsg <- clientMsg

		// brocast
		ws.Wshub.Broadcast <- []byte("broadcast")
		log.Println(*c)
		log.Println(*wsJsonReq)
	}
}

func writeData(c *ws.WsClientConn) {
	defer c.Conn.Close()
	for {
		select {
		case message, ok := <-c.Send:
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
