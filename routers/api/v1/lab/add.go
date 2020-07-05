package lab

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/app"
	"FrontEndOJGolang/pkg/e"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func AddLab(c *gin.Context) {
	appG := app.Gin{
		C: c,
	}

	labName, _ := c.GetPostForm("lab_name")
	labDesc, _ := c.GetPostForm("lab_desc")
	labSample, _ := c.GetPostForm("lab_sample")
	labTypeStr, _ := c.GetPostForm("lab_type")
	labType, _ := strconv.ParseInt(labTypeStr, 10, 8)

	lab := models.Lab{
		Model: models.Model{
			Creator:    "CaveJohson",
			CreateTime: time.Now().UnixNano() / 1e6,
		},
		LabName:   labName,
		LabDesc:   labDesc,
		LabType:   int8(labType),
		LabSample: labSample,
	}

	err := lab.Insert()
	if err != nil {
		log.Printf("[ERROR] database exec error input[%v] err[%v]", lab, err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
