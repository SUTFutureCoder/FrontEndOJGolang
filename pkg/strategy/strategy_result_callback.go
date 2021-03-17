package startegy

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/e"
	"FrontEndOJGolang/pkg/websocket"
	"encoding/json"
	"log"
)

type result struct {
	strategyDto
}

func (r *result) execute(v ...interface{}) {
	// step1 parse user info
	var labSubmit models.LabSubmit
	if str, ok := v[0].(string); ok {
		err := json.Unmarshal([]byte(str), &labSubmit)
		if err != nil {
			log.Println(err.Error())
			return
		}
		r.Data = labSubmit
		websocket.SendToUser(labSubmit.CreatorId, e.SUCCESS, *r)
	}
}
