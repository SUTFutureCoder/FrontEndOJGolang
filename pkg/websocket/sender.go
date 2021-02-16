package websocket

import (
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"encoding/json"
	"log"
)

type WsResp app.Response

// send to one client of one User
func SendToClient(c *WsClientConn, errCode int, data interface{}) {
	j, err := marshalData(errCode, data)
	if err != nil {
		log.Printf("marshal data when SendToClient error:%v", err)
		return
	}
	c.Send <- j
}

// send to all connections of one User
func SendToUser(userId uint64, errCode int, data interface{}) {
	j, err := marshalData(errCode, data)
	if err != nil {
		log.Printf("marshal data when SendToUser error:%v", err)
		return
	}
	clientMsg := &ClientMsg{
		ClientId: userId,
		Msg:      j,
	}
	Wshub.ClientMsg <- clientMsg
}

// brocast to all Users
func BroadCastToAllUsers(errCode int, data interface{}) {
	j, err := marshalData(errCode, data)
	if err != nil {
		log.Printf("marshal data when BroadCastToAllUsers error:%v", err)
		return
	}
	Wshub.Broadcast <- j
}

func marshalData(errCode int, data interface{}) ([]byte, error) {
	return json.Marshal(WsResp{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
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
