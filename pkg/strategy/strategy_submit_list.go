package startegy

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/e"
	"FrontEndOJGolang/pkg/websocket"
	"encoding/json"
	"fmt"
)

type submitList struct {
	strategyDto
}

type submitListData struct {
	LabId uint64 `json:"lab_id"`
	ResultList interface{} `json:"result_list"`
}

type submitListReq struct {
	LabId uint64 `json:"lab_id"`
}

func (s *submitList) execute(v ...interface{}) {
	fmt.Println(s)

	if c, ok := s.Context.(*websocket.WsClientConn); ok {
		if str, ok := v[0].(string); ok {
			req := &submitListReq{}
			err := json.Unmarshal([]byte(str), &req)
			if err != nil {
				fmt.Printf("Unmarshal submitlist req error:%v", err)
			}
			labSubmit := &models.LabSubmit{}
			data, _ := labSubmit.GetUserSubmitsByLabId(c.User.Id, req.LabId)
			s.Data = &submitListData{
				LabId: req.LabId,
				ResultList: data,
			}
			websocket.SendToClient(c, e.SUCCESS, *s)
		}
	}

	fmt.Println(v)
}
