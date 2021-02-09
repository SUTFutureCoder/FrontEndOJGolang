package websocket

import (
	"FrontEndOJGolang/pkg/app"
	"github.com/gorilla/websocket"
	"sync"
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
	//clients map[uint64]*WsClient
	clients sync.Map

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

// todo solve concurrent issue
var locker sync.RWMutex

func NewWsHub() *WsHub {
	Wshub = &WsHub{
		clients:    sync.Map{},
		ClientMsg:  make(chan *ClientMsg, 128),
		Broadcast:  make(chan []byte, 128),
		Register:   make(chan *WsClientConn, 128),
		Unregister: make(chan *WsClientConn, 128),
	}
	return Wshub
}

func (h *WsHub) Setup() {
	for {
		select {
		case clientConn := <-h.Register:
			wsClientValue, _ := h.clients.LoadOrStore(clientConn.User.Id, &WsClient{
				User:  clientConn.User,
				Conns: make([]*WsClientConn, 0),
			})
			wsClient := wsClientValue.(*WsClient)
			wsClient.Conns = append(wsClient.Conns, clientConn)
		case clientConn := <-h.Unregister:
			if wsClientValue, ok := h.clients.Load(clientConn.User.Id); ok {
				// find conn
				wsClient := wsClientValue.(*WsClient)
				h.removeClientFromWsClient(wsClient, clientConn)
			}
		case message := <-h.Broadcast:
			// 广播
			h.clients.Range(func(key, value interface{}) bool {
				clients := value.(*WsClient)
				for _, client := range clients.Conns {
					select {
					case client.Send <- message:
					default:
						h.removeClientFromWsClient(clients, client)
					}
				}
				return false
			})
		case clientMsg := <-h.ClientMsg:
			// 对用户
			if wsClientValue, ok := h.clients.Load(clientMsg.ClientId); ok {
				wsClient := wsClientValue.(*WsClient)
				for _, client := range wsClient.Conns {
					select {
					case client.Send <- clientMsg.Msg:
					default:
						h.removeClientFromWsClient(wsClient, client)
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
				h.clients.Delete(wsClient.User.Id)
			}
			break
		}
	}
}
