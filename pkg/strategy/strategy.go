package strategy

import (
	"FrontEndOJGolang/pkg/consts"
	"log"
)

var strategy map[string]func(v ...interface{})

func Setup() {
	strategy = make(map[string]func(v ...interface{}))
	strategy[consts.HelloWorld] = func(v ...interface{}) {
		log.Println(v)
		log.Println(consts.HelloWorld)
	}
	strategy[consts.SayHello] = func(v ...interface{}) {
		log.Println(v)
		log.Println(consts.SayHello)
	}
}

func ExecStrategy(s string) {
	if _, ok := strategy[s]; !ok {
		log.Printf("strategy not exist, target strategy:%s", s)
		return
	}
	strategy[s]()
}

