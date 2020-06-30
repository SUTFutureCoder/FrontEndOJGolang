package caroline

import (
	"FrontEndOJGolang/models"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"strings"
	"time"
)

func ExecCaroline(testChamber string, testcases []models.LabTestcase) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	for _, testcase := range testcases {

		var output interface{}
		if err := chromedp.Run(ctx, runTests(testChamber, &testcase, &output)); err != nil {
			fmt.Println(err)
		}
		if output == testcase.Output {
			fmt.Println("PASS")
		} else {
			fmt.Println("FAIL")
		}
		fmt.Println(testcase.Output)
		fmt.Println("###########")
		fmt.Println(output)
		fmt.Println("-----------")

	}

}

func runTests(url string, labTestcase *models.LabTestcase, output *interface{}) chromedp.Action {
	task := chromedp.Tasks{
		chromedp.Navigate(url),
	}
	if labTestcase.WaitBefore != 0 {
		var temp *interface{}
		// 在sleep之前执行一下，需要注意两次执行代码一样，但结果不同，为了保持核心代码和数据表整洁
		task = append(task, chromedp.EvaluateAsDevTools(strings.ReplaceAll(labTestcase.TestcaseCode, "\n", ""), &temp))
		task = append(task, chromedp.Sleep(time.Duration(labTestcase.WaitBefore)*time.Millisecond))
		fmt.Println(labTestcase.WaitBefore)
	}
	task = append(task, chromedp.EvaluateAsDevTools(strings.ReplaceAll(labTestcase.TestcaseCode, "\n", ""), &output))
	return task
}
