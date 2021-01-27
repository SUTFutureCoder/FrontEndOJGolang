package app

import (
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) resp(errCode int, data interface{}) {
	// 二次包装
	g.C.JSON(http.StatusOK, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
}

func (g *Gin) RespSucc(data interface{}) {
	g.resp(e.SUCCESS, data)
}

func (g *Gin) RespErr(errCode int, data interface{}) {
	g.resp(errCode, data)
}
