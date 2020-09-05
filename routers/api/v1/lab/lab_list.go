package lab

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RespLabList struct {
	LabList []models.Lab
	Count   int
}

func LabList(c *gin.Context) {
	appGin := app.Gin{
		C: c,
	}
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "20")
	intPage, _ := strconv.Atoi(page)
	intPageSize, _ := strconv.Atoi(pageSize)

	var resp RespLabList
	var err error
	resp.LabList, err = models.GetLabList(intPage, intPageSize)
	resp.Count, err = models.GetLabListCount()
	if err != nil {
		appGin.Response(http.StatusInternalServerError, e.ERROR, err)
		return
	}
	appGin.Response(http.StatusOK, e.SUCCESS, resp)
}
