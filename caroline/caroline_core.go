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
		fmt.Println(output)

	}

}

func runTests(url string, labTestcase *models.LabTestcase, output *interface{}) chromedp.Action {
	task := chromedp.Tasks{
		chromedp.Navigate(url),
	}
	if labTestcase.WaitBefore != 0 {
		task = append(task, chromedp.Sleep(time.Duration(labTestcase.WaitBefore)*time.Microsecond))
	}
	task = append(task, chromedp.EvaluateAsDevTools(strings.ReplaceAll(labTestcase.TestcaseCode, "\n", ""), &output))
	return task
}
