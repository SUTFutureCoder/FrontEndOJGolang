package caroline

import (
	"FrontEndOJGolang/models"
	"FrontEndOJGolang/pkg/setting"
	"fmt"
	"io/ioutil"
	"os"
)

/**
 写入本地磁盘 ./test_chamber/{creator}/{submitid}
 */
func WriteSubmitToFile(labSubmit *models.LabSubmit) string {
	testChanberDirName := fmt.Sprintf("%s/test_chamber/%s/%d/", setting.JudgerSetting.TestChamberBaseDir, labSubmit.Creator, labSubmit.ID)
	testChamberFileName := fmt.Sprintf("%sindex.html", testChanberDirName)
	fmt.Println(testChamberFileName)
	os.MkdirAll(testChanberDirName, 0777)
	ioutil.WriteFile(testChamberFileName, []byte(labSubmit.SubmitData), 0777)
	return testChamberFileName
}
