package main

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/setting"
	"FrontEndOJGolang/pkg/strategy"
	"FrontEndOJGolang/pkg/websocket"
	"FrontEndOJGolang/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	setting.Setup()
	setting.Check()
	models.Setup()
	go websocket.NewWsHub().Setup()
	strategy.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := &http.Server{
		Addr:         endPoint,
		Handler:      routersInit(),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	// 自检
	if models.DB == nil {
		log.Fatalln("[FATAL] Database setup failed")
		return
	}

	log.Printf("[SUCCESS] Project Caroline Started 🎂")

	server.ListenAndServe()
}
