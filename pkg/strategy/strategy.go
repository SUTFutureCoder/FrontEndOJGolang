package strategy

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/consts"
	"FrontEndOJGolang/pkg/e"
	"FrontEndOJGolang/pkg/websocket"
	"encoding/json"
	"log"
)

var strategy map[string]func(v ...interface{})

func Setup() {
	strategy = make(map[string]func(v ...interface{}), 64)
	strategy[consts.HelloWorld] = func(v ...interface{}) {
		log.Println(v)
		log.Println(consts.HelloWorld)
	}
	strategy[consts.SayHello] = func(v ...interface{}) {
		log.Println(v)
		log.Println(consts.SayHello)
	}
	strategy[consts.JudgerResultCallBack] = func(v ...interface{}) {
		if len(v) == 0 {
			return
		}
		// step1 parse user info
		var labSubmit models.LabSubmit
		if str, ok := v[0].(string); ok {
			err := json.Unmarshal([]byte(str), &labSubmit)
			if err != nil {
				log.Println(err.Error())
				return
			}
			websocket.SendToUser(labSubmit.CreatorId, e.SUCCESS, labSubmit)
		}
	}
}

func ExecStrategy(s string, d string) {
	if _, ok := strategy[s]; !ok {
		log.Printf("strategy not exist, target strategy:%s", s)
		return
	}
	strategy[s](d)
}
