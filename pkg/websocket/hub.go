package websocket

import (
	"FrontEndOJGolang/pkg/app"
	"github.com/gorilla/websocket"
)

// 用户下单个连接
type WsClientConn struct {
	User app.UserSession
	Hub  *WsHub
	Conn *websocket.Conn
	Send chan []byte
}

// 用户下多个连接
type WsClient struct {
	User  app.UserSession
	Conns []*WsClientConn
}

// 消息总线
type WsHub struct {
	// 用户集合总线
	clients map[uint64]*WsClient

	// 单用户消息通道
	ClientMsg chan *ClientMsg
	// 广播消息通道
	Broadcast chan []byte

	// 注册
	Register chan *WsClientConn
	// 注销
	Unregister chan *WsClientConn
}

// 单用户消息体
type ClientMsg struct {
	ClientId uint64
	Msg      []byte
}

var Wshub *WsHub

func NewWsHub() *WsHub {
	Wshub = &WsHub{
		clients:    make(map[uint64]*WsClient),
		ClientMsg:  make(chan *ClientMsg),
		Broadcast:  make(chan []byte),
		Register:   make(chan *WsClientConn),
		Unregister: make(chan *WsClientConn),
	}
	return Wshub
}

func (h *WsHub) Setup() {
	for {
		select {
		case clientConn := <-h.Register:
			if _, ok := h.clients[clientConn.User.Id]; !ok {
				// init clients
				wsClient := &WsClient{
					User:  clientConn.User,
					Conns: make([]*WsClientConn, 0),
				}
				h.clients[clientConn.User.Id] = wsClient
			}
			wsClient := h.clients[clientConn.User.Id]
			wsClient.Conns = append(wsClient.Conns, clientConn)
		case clientConn := <-h.Unregister:
			if _, ok := h.clients[clientConn.User.Id]; ok {
				// find conn
				wsClient := h.clients[clientConn.User.Id]
				h.removeClientFromWsClient(wsClient, clientConn)
			}
		case message := <-h.Broadcast:
			// 广播
			for _, clients := range h.clients {
				for _, client := range clients.Conns {
					select {
					case client.Send <- message:
					default:
						h.removeClientFromWsClient(clients, client)
					}
				}
			}
		case clientMsg := <-h.ClientMsg:
			// 对用户
			if _, ok := h.clients[clientMsg.ClientId]; ok {
				for _, client := range h.clients[clientMsg.ClientId].Conns {
					select {
					case client.Send <- clientMsg.Msg:
					default:
						h.removeClientFromWsClient(h.clients[clientMsg.ClientId], client)
					}
				}
			}
		}
	}
}

func (h *WsHub) removeClientFromWsClient(wsClient *WsClient, targetConn *WsClientConn) {
	for i, wsClientConn := range wsClient.Conns {
		if wsClientConn == targetConn {
			close(wsClientConn.Send)
			wsClientConn = nil
			wsClient.Conns[i] = wsClient.Conns[len(wsClient.Conns)-1]
			wsClient.Conns = wsClient.Conns[:len(wsClient.Conns)-1]
			// remove key if empty
			if len(wsClient.Conns) == 0 {
				delete(h.clients, wsClient.User.Id)
			}
			break
		}
	}
}


/**
 * SEND SAMPLE
	// send to one client of one User
	c.Send <- []byte("yooooo")

	// send to all connections of one User
	clientMsg := &ws.ClientMsg{
		ClientId: USERID,
		Msg:      []byte("client"),
	}
	ws.Wshub.ClientMsg <- clientMsg

	// brocast to all Users
	ws.Wshub.Broadcast <- []byte("broadcast")
 */
