package lab_submit

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

func Submit(c *gin.Context) {

	appG := app.Gin{
		C: c,
	}

	labIdStr, _ := c.GetPostForm("lab_id")
	labId, _ := strconv.Atoi(labIdStr)
	submitData, _ := c.GetPostForm("submit_data")
	submitResult := ""
	creator := "Chell01"
	createTime := time.Now().UnixNano() / 1e6

	stmt, err := models.DB.Prepare("INSERT INTO lab_submit (lab_id, submit_data, submit_result, creator, create_time) VALUES (?,?,?,?,?)")
	_, err = stmt.Exec(
		&labId,
		&submitData,
		&submitResult,
		&creator,
		&createTime,
	)
	defer stmt.Close()

	if err != nil {
		log.Printf("[ERROR] lab submit error[%v]\n", err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}
