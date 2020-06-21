package main

import (
	"FrontEndOJGolang/pkg/setting"
	"FrontEndOJGolang/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	setting.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := &http.Server{
		Addr: endPoint,
		Handler: routersInit(),
		ReadTimeout: readTimeout,
		WriteTimeout:  writeTimeout,
	}

	log.Printf("[SUCCESS] Falcon X Lunched")

	server.ListenAndServe()
}